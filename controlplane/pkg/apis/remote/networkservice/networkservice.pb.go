// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkservice.proto

package networkservice

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	connection "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/connection"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterType((*NetworkServiceRequest)(nil), "remote.networkservice.NetworkServiceRequest")
}

func init() { proto.RegisterFile("networkservice.proto", fileDescriptor_361e8247d5a9945c) }

var fileDescriptor_361e8247d5a9945c = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4b, 0xf4, 0x30,
	0x10, 0xa6, 0xef, 0x8b, 0x0a, 0x59, 0x58, 0x24, 0x6c, 0xa5, 0x14, 0x85, 0xc5, 0xd3, 0x1e, 0x24,
	0x81, 0x7a, 0xd6, 0x83, 0x8b, 0x47, 0x45, 0x2b, 0x78, 0xf0, 0x22, 0x6d, 0x98, 0x6d, 0xc3, 0x36,
	0x99, 0x98, 0xa4, 0xca, 0xfe, 0x23, 0xfd, 0x97, 0xb2, 0x66, 0x3f, 0x5a, 0x28, 0x7b, 0x29, 0x43,
	0x9f, 0x79, 0xbe, 0x26, 0x64, 0xa2, 0xc1, 0x7f, 0xa1, 0x5d, 0x3a, 0xb0, 0x9f, 0x52, 0x00, 0x33,
	0x16, 0x3d, 0xd2, 0xd8, 0x82, 0x42, 0x0f, 0xac, 0x0f, 0xa6, 0xb2, 0x92, 0xbe, 0x6e, 0x4b, 0x26,
	0x50, 0xf1, 0x3e, 0xa4, 0xc0, 0xd5, 0x43, 0xbf, 0x04, 0x6a, 0x6f, 0xb1, 0x31, 0x4d, 0xa1, 0x81,
	0x9b, 0x65, 0xc5, 0x0b, 0x23, 0x1d, 0x0f, 0xea, 0x6b, 0x50, 0x83, 0xf0, 0x12, 0x75, 0x67, 0x0c,
	0x09, 0xd2, 0xc4, 0xf8, 0x95, 0x01, 0xc7, 0x41, 0x19, 0xbf, 0x0a, 0xdf, 0x80, 0x5c, 0xfe, 0x44,
	0x24, 0x7e, 0x0c, 0x4e, 0x2f, 0xc1, 0x29, 0x87, 0x8f, 0x16, 0x9c, 0xa7, 0x37, 0x84, 0xec, 0x75,
	0x92, 0x68, 0x1a, 0xcd, 0x46, 0xd9, 0x05, 0xdb, 0x54, 0xe9, 0x38, 0xcc, 0x77, 0x63, 0xde, 0x21,
	0xd0, 0x67, 0x12, 0x2b, 0x10, 0x75, 0xa1, 0xa5, 0x53, 0xef, 0xc6, 0xc2, 0x02, 0x2c, 0x68, 0x01,
	0x2e, 0xf9, 0x37, 0xfd, 0x3f, 0x1b, 0x65, 0xe7, 0x03, 0x4a, 0x0f, 0xdb, 0xfd, 0x7c, 0xb2, 0xa3,
	0x3e, 0xed, 0x99, 0xd9, 0x77, 0x44, 0xc6, 0xfd, 0xac, 0xf4, 0x95, 0x9c, 0x6c, 0xf3, 0x5e, 0xb1,
	0xc1, 0x33, 0xb3, 0xc1, 0x76, 0xe9, 0xe1, 0x26, 0xf4, 0x96, 0x1c, 0xcd, 0x1b, 0x74, 0x40, 0x0f,
	0xef, 0xa5, 0x67, 0xac, 0x42, 0xac, 0x9a, 0xcd, 0x4b, 0x97, 0xed, 0x82, 0xdd, 0xaf, 0x8f, 0x7b,
	0x77, 0xfa, 0x36, 0xee, 0xc7, 0x28, 0x8f, 0xff, 0x36, 0xae, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x2f, 0xc0, 0x78, 0x76, 0x23, 0x02, 0x00, 0x00,
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
	Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error)
	Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error) {
	out := new(connection.Connection)
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
	Request(context.Context, *NetworkServiceRequest) (*connection.Connection, error)
	Close(context.Context, *connection.Connection) (*empty.Empty, error)
}

// UnimplementedNetworkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (*UnimplementedNetworkServiceServer) Request(ctx context.Context, req *NetworkServiceRequest) (*connection.Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedNetworkServiceServer) Close(ctx context.Context, req *connection.Connection) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
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
