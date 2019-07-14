// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplane.proto

package dataplane

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import crossconnect "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/crossconnect"
import connection1 "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/local/connection"
import connection "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/connection"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Message sent by dataplane module informing NSM of any changes in its
// operations parameters or constraints
type MechanismUpdate struct {
	RemoteMechanisms     []*connection.Mechanism  `protobuf:"bytes,1,rep,name=remote_mechanisms,json=remoteMechanisms,proto3" json:"remote_mechanisms,omitempty"`
	LocalMechanisms      []*connection1.Mechanism `protobuf:"bytes,2,rep,name=local_mechanisms,json=localMechanisms,proto3" json:"local_mechanisms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MechanismUpdate) Reset()         { *m = MechanismUpdate{} }
func (m *MechanismUpdate) String() string { return proto.CompactTextString(m) }
func (*MechanismUpdate) ProtoMessage()    {}
func (*MechanismUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplane_3cc7501aa2ae2266, []int{0}
}
func (m *MechanismUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MechanismUpdate.Unmarshal(m, b)
}
func (m *MechanismUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MechanismUpdate.Marshal(b, m, deterministic)
}
func (dst *MechanismUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MechanismUpdate.Merge(dst, src)
}
func (m *MechanismUpdate) XXX_Size() int {
	return xxx_messageInfo_MechanismUpdate.Size(m)
}
func (m *MechanismUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MechanismUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MechanismUpdate proto.InternalMessageInfo

func (m *MechanismUpdate) GetRemoteMechanisms() []*connection.Mechanism {
	if m != nil {
		return m.RemoteMechanisms
	}
	return nil
}

func (m *MechanismUpdate) GetLocalMechanisms() []*connection1.Mechanism {
	if m != nil {
		return m.LocalMechanisms
	}
	return nil
}

func init() {
	proto.RegisterType((*MechanismUpdate)(nil), "dataplane.MechanismUpdate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataplaneClient is the client API for Dataplane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneClient interface {
	Request(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*crossconnect.CrossConnect, error)
	Close(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*empty.Empty, error)
	MonitorMechanisms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Dataplane_MonitorMechanismsClient, error)
}

type dataplaneClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneClient(cc *grpc.ClientConn) DataplaneClient {
	return &dataplaneClient{cc}
}

func (c *dataplaneClient) Request(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*crossconnect.CrossConnect, error) {
	out := new(crossconnect.CrossConnect)
	err := c.cc.Invoke(ctx, "/dataplane.Dataplane/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataplaneClient) Close(ctx context.Context, in *crossconnect.CrossConnect, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/dataplane.Dataplane/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataplaneClient) MonitorMechanisms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Dataplane_MonitorMechanismsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dataplane_serviceDesc.Streams[0], "/dataplane.Dataplane/MonitorMechanisms", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataplaneMonitorMechanismsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dataplane_MonitorMechanismsClient interface {
	Recv() (*MechanismUpdate, error)
	grpc.ClientStream
}

type dataplaneMonitorMechanismsClient struct {
	grpc.ClientStream
}

func (x *dataplaneMonitorMechanismsClient) Recv() (*MechanismUpdate, error) {
	m := new(MechanismUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataplaneServer is the server API for Dataplane service.
type DataplaneServer interface {
	Request(context.Context, *crossconnect.CrossConnect) (*crossconnect.CrossConnect, error)
	Close(context.Context, *crossconnect.CrossConnect) (*empty.Empty, error)
	MonitorMechanisms(*empty.Empty, Dataplane_MonitorMechanismsServer) error
}

func RegisterDataplaneServer(s *grpc.Server, srv DataplaneServer) {
	s.RegisterService(&_Dataplane_serviceDesc, srv)
}

func _Dataplane_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(crossconnect.CrossConnect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplane.Dataplane/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneServer).Request(ctx, req.(*crossconnect.CrossConnect))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataplane_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(crossconnect.CrossConnect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplane.Dataplane/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneServer).Close(ctx, req.(*crossconnect.CrossConnect))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataplane_MonitorMechanisms_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataplaneServer).MonitorMechanisms(m, &dataplaneMonitorMechanismsServer{stream})
}

type Dataplane_MonitorMechanismsServer interface {
	Send(*MechanismUpdate) error
	grpc.ServerStream
}

type dataplaneMonitorMechanismsServer struct {
	grpc.ServerStream
}

func (x *dataplaneMonitorMechanismsServer) Send(m *MechanismUpdate) error {
	return x.ServerStream.SendMsg(m)
}

var _Dataplane_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplane.Dataplane",
	HandlerType: (*DataplaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Dataplane_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Dataplane_Close_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorMechanisms",
			Handler:       _Dataplane_MonitorMechanisms_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dataplane.proto",
}

func init() { proto.RegisterFile("dataplane.proto", fileDescriptor_dataplane_3cc7501aa2ae2266) }

var fileDescriptor_dataplane_3cc7501aa2ae2266 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0x95, 0xf7, 0x15, 0xa0, 0x9a, 0xa1, 0x25, 0x03, 0x42, 0x81, 0x01, 0x31, 0x31, 0xd9,
	0xa8, 0x1d, 0x99, 0x50, 0x01, 0x89, 0xa1, 0x4b, 0x25, 0x66, 0xe4, 0xb8, 0xd7, 0xc4, 0xaa, 0xe3,
	0x33, 0xb6, 0x03, 0xea, 0xf7, 0xe1, 0x33, 0xf1, 0x79, 0x50, 0xec, 0xfe, 0x49, 0x51, 0xc9, 0xd4,
	0x25, 0xca, 0xdd, 0xef, 0xee, 0x79, 0x74, 0xe7, 0x23, 0xfd, 0x19, 0xf7, 0xdc, 0x28, 0xae, 0x81,
	0x1a, 0x8b, 0x1e, 0xd3, 0xde, 0x26, 0x91, 0x95, 0x85, 0xf4, 0x65, 0x9d, 0x53, 0x81, 0x15, 0xd3,
	0xe0, 0x3f, 0xd1, 0x2e, 0x1c, 0xd8, 0x0f, 0x29, 0xa0, 0x02, 0x57, 0xee, 0x4b, 0x09, 0xd4, 0xde,
	0xa2, 0x0a, 0xed, 0xcc, 0x2c, 0x0a, 0xc6, 0x8d, 0x74, 0x4c, 0xa1, 0xe0, 0xaa, 0x61, 0x1a, 0x84,
	0x97, 0xa8, 0x5b, 0xbf, 0xd1, 0x34, 0x93, 0x07, 0x72, 0xb2, 0x50, 0xa1, 0x87, 0x4e, 0xab, 0xf9,
	0x81, 0xac, 0x84, 0x45, 0xe7, 0x56, 0xea, 0x3b, 0xc1, 0xca, 0x67, 0xd4, 0xf2, 0x29, 0x50, 0x71,
	0x5d, 0xb0, 0x00, 0xf2, 0x7a, 0xce, 0x8c, 0x5f, 0x1a, 0x70, 0x0c, 0x2a, 0xe3, 0x97, 0xf1, 0x1b,
	0x9b, 0x6e, 0xbe, 0x12, 0xd2, 0x9f, 0x80, 0x28, 0xb9, 0x96, 0xae, 0x7a, 0x35, 0x33, 0xee, 0x21,
	0x7d, 0x21, 0x67, 0x71, 0xac, 0xb7, 0x6a, 0x4d, 0xdc, 0x45, 0x72, 0xfd, 0xff, 0xf6, 0x74, 0x78,
	0x45, 0x23, 0xa1, 0xad, 0x29, 0x37, 0xed, 0xd3, 0x41, 0x84, 0x9b, 0x84, 0x4b, 0x9f, 0xc9, 0x20,
	0xbc, 0x45, 0x5b, 0xe9, 0x5f, 0x50, 0xba, 0xa4, 0x01, 0xec, 0x17, 0xea, 0x07, 0xb6, 0xd5, 0x19,
	0x7e, 0x27, 0xa4, 0xf7, 0xb8, 0x3e, 0x93, 0xf4, 0x81, 0x9c, 0x4c, 0xe1, 0xbd, 0x06, 0xe7, 0xd3,
	0x8c, 0xee, 0x6c, 0x62, 0xdc, 0x04, 0xe3, 0x18, 0x64, 0x1d, 0x2c, 0xbd, 0x27, 0x47, 0x63, 0x85,
	0x0e, 0x3a, 0x05, 0xce, 0x69, 0x81, 0x58, 0xa8, 0xd5, 0xa1, 0xe6, 0xf5, 0x9c, 0x3e, 0x35, 0xab,
	0x6b, 0x16, 0x34, 0x41, 0x2d, 0x3d, 0xda, 0xd6, 0xa8, 0x7f, 0x14, 0x67, 0x19, 0xdd, 0x1e, 0xfc,
	0xaf, 0x4d, 0xdf, 0x25, 0xf9, 0x71, 0xa8, 0x1e, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xc8,
	0x4f, 0xeb, 0x16, 0x03, 0x00, 0x00,
}
