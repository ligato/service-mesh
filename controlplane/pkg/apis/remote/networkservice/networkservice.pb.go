// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkservice.proto

package networkservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	connection "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/connection"
	grpc "google.golang.org/grpc"
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

type NetworkServiceRequest struct {
	Connection           *connection.Connection  `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	MechanismPreferences []*connection.Mechanism `protobuf:"bytes,2,rep,name=mechanism_preferences,json=mechanismPreferences,proto3" json:"mechanism_preferences,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NetworkServiceRequest) Reset()         { *m = NetworkServiceRequest{} }
func (m *NetworkServiceRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceRequest) ProtoMessage()    {}
func (*NetworkServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_361e8247d5a9945c, []int{0}
}

func (m *NetworkServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceRequest.Unmarshal(m, b)
}
func (m *NetworkServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceRequest.Marshal(b, m, deterministic)
}
func (m *NetworkServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceRequest.Merge(m, src)
}
func (m *NetworkServiceRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceRequest.Size(m)
}
func (m *NetworkServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceRequest proto.InternalMessageInfo

func (m *NetworkServiceRequest) GetConnection() *connection.Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *NetworkServiceRequest) GetMechanismPreferences() []*connection.Mechanism {
	if m != nil {
		return m.MechanismPreferences
	}
	return nil
}

type NetworkServiceReply struct {
	Connection           *connection.Connection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	ResponseJWT          string                 `protobuf:"bytes,2,opt,name=ResponseJWT,proto3" json:"ResponseJWT,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *NetworkServiceReply) Reset()         { *m = NetworkServiceReply{} }
func (m *NetworkServiceReply) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceReply) ProtoMessage()    {}
func (*NetworkServiceReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_361e8247d5a9945c, []int{1}
}

func (m *NetworkServiceReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceReply.Unmarshal(m, b)
}
func (m *NetworkServiceReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceReply.Marshal(b, m, deterministic)
}
func (m *NetworkServiceReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceReply.Merge(m, src)
}
func (m *NetworkServiceReply) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceReply.Size(m)
}
func (m *NetworkServiceReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceReply.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceReply proto.InternalMessageInfo

func (m *NetworkServiceReply) GetConnection() *connection.Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *NetworkServiceReply) GetResponseJWT() string {
	if m != nil {
		return m.ResponseJWT
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkServiceRequest)(nil), "remote.networkservice.NetworkServiceRequest")
	proto.RegisterType((*NetworkServiceReply)(nil), "remote.networkservice.NetworkServiceReply")
}

func init() { proto.RegisterFile("networkservice.proto", fileDescriptor_361e8247d5a9945c) }

var fileDescriptor_361e8247d5a9945c = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x25, 0xfd, 0xf1, 0x53, 0xdc, 0x42, 0x91, 0xd8, 0x4a, 0x08, 0x0a, 0xa1, 0xa7, 0x22, 0xb2,
	0x81, 0x7a, 0xd6, 0x83, 0xc5, 0x8b, 0xa0, 0x68, 0x14, 0x04, 0x2f, 0xd2, 0x2e, 0xd3, 0x64, 0x69,
	0x76, 0x67, 0xdd, 0xdd, 0x54, 0xf2, 0x91, 0xfc, 0x00, 0x7e, 0x3f, 0x89, 0xdb, 0x3f, 0x89, 0x04,
	0x11, 0xbc, 0x2c, 0xc3, 0xbc, 0x99, 0x79, 0xef, 0xcd, 0x2c, 0xe9, 0x4b, 0xb0, 0x6f, 0xa8, 0x17,
	0x06, 0xf4, 0x92, 0x33, 0xa0, 0x4a, 0xa3, 0x45, 0x7f, 0xa0, 0x41, 0xa0, 0x05, 0xda, 0x04, 0x43,
	0x9e, 0x72, 0x9b, 0x15, 0x33, 0xca, 0x50, 0xc4, 0x4d, 0x48, 0x80, 0xc9, 0xda, 0x52, 0x0c, 0xa5,
	0xd5, 0x98, 0xab, 0x7c, 0x2a, 0x21, 0x56, 0x8b, 0x34, 0x9e, 0x2a, 0x6e, 0x62, 0x37, 0xbd, 0x02,
	0x25, 0x30, 0xcb, 0x51, 0xd6, 0x42, 0xa7, 0x20, 0x0c, 0x94, 0x2d, 0x15, 0x98, 0x18, 0x84, 0xb2,
	0xa5, 0x7b, 0x1d, 0x32, 0x7c, 0xf7, 0xc8, 0xe0, 0xd6, 0x31, 0x3d, 0x38, 0xa6, 0x04, 0x5e, 0x0b,
	0x30, 0xd6, 0x3f, 0x27, 0x64, 0x3b, 0x27, 0xf0, 0x22, 0x6f, 0xd4, 0x1d, 0x1f, 0xd3, 0x95, 0x95,
	0x1a, 0xc3, 0x64, 0x13, 0x26, 0xb5, 0x06, 0xff, 0x9e, 0x0c, 0x04, 0xb0, 0x6c, 0x2a, 0xb9, 0x11,
	0x2f, 0x4a, 0xc3, 0x1c, 0x34, 0x48, 0x06, 0x26, 0xe8, 0x44, 0xff, 0x46, 0xdd, 0xf1, 0x51, 0xcb,
	0xa4, 0x9b, 0x75, 0x7d, 0xd2, 0xdf, 0xb4, 0xde, 0x6d, 0x3b, 0x87, 0x4b, 0x72, 0xf0, 0x5d, 0xaa,
	0xca, 0xcb, 0xbf, 0x0a, 0x8d, 0x48, 0x37, 0x01, 0xa3, 0x50, 0x1a, 0xb8, 0x7e, 0x7a, 0x0c, 0x3a,
	0x91, 0x37, 0xda, 0x4b, 0xea, 0xa9, 0xf1, 0x87, 0x47, 0x7a, 0x4d, 0x62, 0x9f, 0x91, 0xdd, 0xf5,
	0x9e, 0x4e, 0x69, 0xeb, 0x79, 0x69, 0xeb, 0x56, 0xc3, 0x93, 0x5f, 0x56, 0x57, 0xc6, 0x2e, 0xc8,
	0xff, 0x49, 0x8e, 0x06, 0xfc, 0x9f, 0xdd, 0x84, 0x87, 0x34, 0x45, 0x4c, 0xf3, 0xd5, 0x77, 0x9b,
	0x15, 0x73, 0x7a, 0x55, 0x5d, 0xf8, 0x72, 0xff, 0xb9, 0xd7, 0x64, 0x99, 0xed, 0x7c, 0x55, 0x9c,
	0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xce, 0x77, 0x8c, 0xa8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*NetworkServiceReply, error)
	Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*NetworkServiceReply, error) {
	out := new(NetworkServiceReply)
	err := c.cc.Invoke(ctx, "/remote.networkservice.NetworkService/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/remote.networkservice.NetworkService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	Request(context.Context, *NetworkServiceRequest) (*NetworkServiceReply, error)
	Close(context.Context, *connection.Connection) (*empty.Empty, error)
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.networkservice.NetworkService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Request(ctx, req.(*NetworkServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(connection.Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.networkservice.NetworkService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Close(ctx, req.(*connection.Connection))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.networkservice.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _NetworkService_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _NetworkService_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkservice.proto",
}
