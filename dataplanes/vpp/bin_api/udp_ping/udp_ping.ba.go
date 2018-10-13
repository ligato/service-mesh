// Code generated by GoVPP binapi-generator. DO NOT EDIT.
// source: api/udp_ping.api.json

/*
Package udp_ping is a generated VPP binary API of the 'udp_ping' VPP module.

It is generated from this file:
	udp_ping.api.json

It contains these VPP binary API objects:
	4 messages
	2 services
*/
package udp_ping

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Messages */

// UDPPingAddDel represents the VPP binary API message 'udp_ping_add_del'.
// Generated from 'udp_ping.api.json', line 4:
//
//            "udp_ping_add_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "src_ip_address",
//                16
//            ],
//            [
//                "u8",
//                "dst_ip_address",
//                16
//            ],
//            [
//                "u16",
//                "start_src_port"
//            ],
//            [
//                "u16",
//                "end_src_port"
//            ],
//            [
//                "u16",
//                "start_dst_port"
//            ],
//            [
//                "u16",
//                "end_dst_port"
//            ],
//            [
//                "u16",
//                "interval"
//            ],
//            [
//                "u8",
//                "is_ipv4"
//            ],
//            [
//                "u8",
//                "dis"
//            ],
//            [
//                "u8",
//                "fault_det"
//            ],
//            [
//                "u8",
//                "reserve",
//                3
//            ],
//            {
//                "crc": "0x559fcc01"
//            }
//
type UDPPingAddDel struct {
	SrcIPAddress []byte `struc:"[16]byte"`
	DstIPAddress []byte `struc:"[16]byte"`
	StartSrcPort uint16
	EndSrcPort   uint16
	StartDstPort uint16
	EndDstPort   uint16
	Interval     uint16
	IsIPv4       uint8
	Dis          uint8
	FaultDet     uint8
	Reserve      []byte `struc:"[3]byte"`
}

func (*UDPPingAddDel) GetMessageName() string {
	return "udp_ping_add_del"
}
func (*UDPPingAddDel) GetCrcString() string {
	return "559fcc01"
}
func (*UDPPingAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewUDPPingAddDel() api.Message {
	return &UDPPingAddDel{}
}

// UDPPingAddDelReply represents the VPP binary API message 'udp_ping_add_del_reply'.
// Generated from 'udp_ping.api.json', line 69:
//
//            "udp_ping_add_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type UDPPingAddDelReply struct {
	Retval int32
}

func (*UDPPingAddDelReply) GetMessageName() string {
	return "udp_ping_add_del_reply"
}
func (*UDPPingAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UDPPingAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewUDPPingAddDelReply() api.Message {
	return &UDPPingAddDelReply{}
}

// UDPPingExport represents the VPP binary API message 'udp_ping_export'.
// Generated from 'udp_ping.api.json', line 87:
//
//            "udp_ping_export",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "enable"
//            ],
//            {
//                "crc": "0xb142b369"
//            }
//
type UDPPingExport struct {
	Enable uint32
}

func (*UDPPingExport) GetMessageName() string {
	return "udp_ping_export"
}
func (*UDPPingExport) GetCrcString() string {
	return "b142b369"
}
func (*UDPPingExport) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewUDPPingExport() api.Message {
	return &UDPPingExport{}
}

// UDPPingExportReply represents the VPP binary API message 'udp_ping_export_reply'.
// Generated from 'udp_ping.api.json', line 109:
//
//            "udp_ping_export_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type UDPPingExportReply struct {
	Retval int32
}

func (*UDPPingExportReply) GetMessageName() string {
	return "udp_ping_export_reply"
}
func (*UDPPingExportReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UDPPingExportReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewUDPPingExportReply() api.Message {
	return &UDPPingExportReply{}
}

/* Services */

type Services interface {
	UDPPingAddDel(*UDPPingAddDel) (*UDPPingAddDelReply, error)
	UDPPingExport(*UDPPingExport) (*UDPPingExportReply, error)
}

func init() {
	api.RegisterMessage((*UDPPingAddDel)(nil), "udp_ping.UDPPingAddDel")
	api.RegisterMessage((*UDPPingAddDelReply)(nil), "udp_ping.UDPPingAddDelReply")
	api.RegisterMessage((*UDPPingExport)(nil), "udp_ping.UDPPingExport")
	api.RegisterMessage((*UDPPingExportReply)(nil), "udp_ping.UDPPingExportReply")
}
