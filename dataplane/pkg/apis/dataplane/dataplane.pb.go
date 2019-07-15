// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplane.proto

package dataplane

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	crossconnect "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/crossconnect"
	connection1 "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/local/connection"
	connection "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/connection"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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
	return fileDescriptor_617387e490a04ffa, []int{0}
}

func (m *MechanismUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MechanismUpdate.Unmarshal(m, b)
}
func (m *MechanismUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MechanismUpdate.Marshal(b, m, deterministic)
}
func (m *MechanismUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MechanismUpdate.Merge(m, src)
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

func init() { proto.RegisterFile("dataplane.proto", fileDescriptor_617387e490a04ffa) }

var fileDescriptor_617387e490a04ffa = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x15, 0x10, 0xa0, 0x86, 0xa1, 0x25, 0x03, 0x42, 0x81, 0x01, 0x31, 0x31, 0xd9, 0xb4,
	0xdd, 0x60, 0x01, 0x15, 0x90, 0x18, 0xba, 0x54, 0x62, 0x46, 0x8e, 0x7b, 0x4d, 0xac, 0x3a, 0x3e,
	0x63, 0x3b, 0x45, 0x7d, 0x1f, 0x9e, 0x89, 0xe7, 0x41, 0xb1, 0xfb, 0x27, 0x45, 0xa5, 0x53, 0x97,
	0x28, 0x77, 0xbf, 0xbb, 0xef, 0xd3, 0x9d, 0x2f, 0x6e, 0x8f, 0x99, 0x63, 0x5a, 0x32, 0x05, 0x44,
	0x1b, 0x74, 0x98, 0xb4, 0x56, 0x89, 0xb4, 0xc8, 0x85, 0x2b, 0xaa, 0x8c, 0x70, 0x2c, 0xa9, 0x02,
	0xf7, 0x85, 0x66, 0x6a, 0xc1, 0xcc, 0x04, 0x87, 0x12, 0x6c, 0xb1, 0x2d, 0xc5, 0x51, 0x39, 0x83,
	0xd2, 0xb7, 0x53, 0x3d, 0xcd, 0x29, 0xd3, 0xc2, 0x52, 0x89, 0x9c, 0xc9, 0x9a, 0x29, 0xe0, 0x4e,
	0xa0, 0x6a, 0xfc, 0x06, 0xd3, 0x54, 0xec, 0xc9, 0xc9, 0x40, 0x89, 0x0e, 0x76, 0x5a, 0x4d, 0xf6,
	0x64, 0xc5, 0x0d, 0x5a, 0xbb, 0x50, 0xdf, 0x08, 0x16, 0x3e, 0xf7, 0x0d, 0x9f, 0x1c, 0x25, 0x53,
	0x39, 0xf5, 0x20, 0xab, 0x26, 0x8f, 0xb3, 0x2e, 0xe9, 0x93, 0x2e, 0xd5, 0x6e, 0xae, 0xc1, 0x52,
	0x28, 0xb5, 0x9b, 0x87, 0x6f, 0xe8, 0xbd, 0xf9, 0x8e, 0xe2, 0xf6, 0x10, 0x78, 0xc1, 0x94, 0xb0,
	0xe5, 0xbb, 0x1e, 0x33, 0x07, 0xc9, 0x5b, 0x7c, 0x16, 0xa6, 0xfb, 0x28, 0x97, 0xc4, 0x5e, 0x44,
	0xd7, 0x87, 0xb7, 0xa7, 0xbd, 0x2b, 0x12, 0x08, 0x69, 0x0c, 0xbb, 0x6a, 0x1f, 0x75, 0x02, 0x5c,
	0x25, 0x6c, 0xf2, 0x1a, 0x77, 0xfc, 0x93, 0x34, 0x95, 0x0e, 0xbc, 0xd2, 0x25, 0xf1, 0x60, 0xbb,
	0x50, 0xdb, 0xb3, 0xb5, 0x4e, 0xef, 0x27, 0x8a, 0x5b, 0xcf, 0xcb, 0x6b, 0x49, 0x9e, 0xe2, 0x93,
	0x11, 0x7c, 0x56, 0x60, 0x5d, 0x92, 0x92, 0x8d, 0x85, 0x0c, 0xea, 0x60, 0x10, 0x82, 0x74, 0x07,
	0x4b, 0x1e, 0xe2, 0xa3, 0x81, 0x44, 0x0b, 0x3b, 0x05, 0xce, 0x49, 0x8e, 0x98, 0xcb, 0xc5, 0xbd,
	0x66, 0xd5, 0x84, 0xbc, 0xd4, 0xab, 0xab, 0x17, 0x34, 0x44, 0x25, 0x1c, 0x9a, 0xc6, 0xa8, 0xff,
	0x14, 0xa7, 0x29, 0x59, 0xdf, 0xfd, 0x9f, 0x4d, 0xdf, 0x45, 0xd9, 0xb1, 0xaf, 0xee, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x94, 0xd4, 0x39, 0x24, 0x1d, 0x03, 0x00, 0x00,
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

// UnimplementedDataplaneServer can be embedded to have forward compatible implementations.
type UnimplementedDataplaneServer struct {
}

func (*UnimplementedDataplaneServer) Request(ctx context.Context, req *crossconnect.CrossConnect) (*crossconnect.CrossConnect, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedDataplaneServer) Close(ctx context.Context, req *crossconnect.CrossConnect) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (*UnimplementedDataplaneServer) MonitorMechanisms(req *empty.Empty, srv Dataplane_MonitorMechanismsServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorMechanisms not implemented")
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
