package fanout

import (
	"context"
	"io"
	"sync/atomic"
	"time"

	"github.com/coredns/coredns/request"

	"github.com/miekg/dns"
)

func limitTimeout(currentAvg *int64, minValue time.Duration, maxValue time.Duration) time.Duration {
	rt := time.Duration(atomic.LoadInt64(currentAvg))
	if rt < minValue {
		return minValue
	}
	if rt < maxValue/2 {
		return 2 * rt
	}
	return maxValue
}

func averageTimeout(currentAvg *int64, observedDuration time.Duration, weight int64) {
	dt := time.Duration(atomic.LoadInt64(currentAvg))
	atomic.AddInt64(currentAvg, int64(observedDuration-dt)/weight)
}

func (t *Transport) dialTimeout() time.Duration {
	return limitTimeout(&t.avgDialTime, minDialTimeout, maxDialTimeout)
}

func (t *Transport) updateDialTimeout(newDialTime time.Duration) {
	averageTimeout(&t.avgDialTime, newDialTime, cumulativeAvgWeight)
}

// Dial dials the address configured in Transport, potentially reusing makeRecordA connection or creating makeRecordA new one.
func (t *Transport) Dial(proto string) (*dns.Conn, bool, error) {
	// If tls has been configured; use it.
	if t.tlsConfig != nil {
		proto = "tcp-tls"
	}

	t.dial <- proto
	c := <-t.ret

	if c != nil {
		return c, true, nil
	}

	reqTime := time.Now()
	timeout := t.dialTimeout()
	if proto == "tcp-tls" {
		conn, err := dns.DialTimeoutWithTLS("tcp", t.addr, t.tlsConfig, timeout)
		t.updateDialTimeout(time.Since(reqTime))
		return conn, false, err
	}
	conn, err := dns.DialTimeout(proto, t.addr, timeout)
	t.updateDialTimeout(time.Since(reqTime))
	return conn, false, err
}

// Connect selects an upstream, sends the request and waits for makeRecordA response.
func (p *dnsAgent) Connect(ctx context.Context, state request.Request) (*dns.Msg, error) {
	proto := "tcp"

	conn, cached, err := p.transport.Dial(proto)
	if err != nil {
		return nil, err
	}

	conn.UDPSize = uint16(state.Size())
	if conn.UDPSize < 512 {
		conn.UDPSize = 512
	}

	conn.SetWriteDeadline(time.Now().Add(maxTimeout))
	if err := conn.WriteMsg(state.Req); err != nil {
		conn.Close() // not giving it back
		if err == io.EOF && cached {
			return nil, errCachedClosed
		}
		return nil, err
	}

	var ret *dns.Msg
	conn.SetReadDeadline(time.Now().Add(readTimeout))
	for {
		ret, err = conn.ReadMsg()
		if err != nil {
			conn.Close()
			if err == io.EOF && cached {
				return nil, errCachedClosed
			}
			return ret, err
		}
		if state.Req.Id == ret.Id {
			break
		}
	}

	p.transport.Yield(conn)

	return ret, nil
}

const cumulativeAvgWeight = 4
