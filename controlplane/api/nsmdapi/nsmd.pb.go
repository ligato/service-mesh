// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nsmd.proto

package nsmdapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// ConnectionRequest is sent by a NSM client to build a connection with NSM.
type ClientConnectionRequest struct {
	Workspace            string   `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientConnectionRequest) Reset()         { *m = ClientConnectionRequest{} }
func (m *ClientConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*ClientConnectionRequest) ProtoMessage()    {}
func (*ClientConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_084cb5dcc765b124, []int{0}
}

func (m *ClientConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConnectionRequest.Unmarshal(m, b)
}
func (m *ClientConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConnectionRequest.Marshal(b, m, deterministic)
}
func (m *ClientConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConnectionRequest.Merge(m, src)
}
func (m *ClientConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_ClientConnectionRequest.Size(m)
}
func (m *ClientConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConnectionRequest proto.InternalMessageInfo

func (m *ClientConnectionRequest) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

// ClientConnectionReply is sent back by NSM as a reply to ClientConnectionRequest
// accepted true will indicate that the connection is accepted, otherwise false
// indicates that connection was refused and admission_error will provide details
// why connection was refused.
type ClientConnectionReply struct {
	Workspace            string   `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
	HostBasedir          string   `protobuf:"bytes,2,opt,name=hostBasedir,proto3" json:"hostBasedir,omitempty"`
	ClientBaseDir        string   `protobuf:"bytes,3,opt,name=clientBaseDir,proto3" json:"clientBaseDir,omitempty"`
	NsmServerSocket      string   `protobuf:"bytes,4,opt,name=nsmServerSocket,proto3" json:"nsmServerSocket,omitempty"`
	NsmClientSocket      string   `protobuf:"bytes,5,opt,name=nsmClientSocket,proto3" json:"nsmClientSocket,omitempty"`
	NsmWorkspaceToken    string   `protobuf:"bytes,6,opt,name=nsmWorkspaceToken,proto3" json:"nsmWorkspaceToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientConnectionReply) Reset()         { *m = ClientConnectionReply{} }
func (m *ClientConnectionReply) String() string { return proto.CompactTextString(m) }
func (*ClientConnectionReply) ProtoMessage()    {}
func (*ClientConnectionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_084cb5dcc765b124, []int{1}
}

func (m *ClientConnectionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConnectionReply.Unmarshal(m, b)
}
func (m *ClientConnectionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConnectionReply.Marshal(b, m, deterministic)
}
func (m *ClientConnectionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConnectionReply.Merge(m, src)
}
func (m *ClientConnectionReply) XXX_Size() int {
	return xxx_messageInfo_ClientConnectionReply.Size(m)
}
func (m *ClientConnectionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConnectionReply.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConnectionReply proto.InternalMessageInfo

func (m *ClientConnectionReply) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *ClientConnectionReply) GetHostBasedir() string {
	if m != nil {
		return m.HostBasedir
	}
	return ""
}

func (m *ClientConnectionReply) GetClientBaseDir() string {
	if m != nil {
		return m.ClientBaseDir
	}
	return ""
}

func (m *ClientConnectionReply) GetNsmServerSocket() string {
	if m != nil {
		return m.NsmServerSocket
	}
	return ""
}

func (m *ClientConnectionReply) GetNsmClientSocket() string {
	if m != nil {
		return m.NsmClientSocket
	}
	return ""
}

func (m *ClientConnectionReply) GetNsmWorkspaceToken() string {
	if m != nil {
		return m.NsmWorkspaceToken
	}
	return ""
}

// DeleteConnectionRequest is sent by a nsm-k8s to delete connection with a client.
type DeleteConnectionRequest struct {
	Workspace            string   `protobuf:"bytes,1,opt,name=workspace,proto3" json:"workspace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConnectionRequest) Reset()         { *m = DeleteConnectionRequest{} }
func (m *DeleteConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectionRequest) ProtoMessage()    {}
func (*DeleteConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_084cb5dcc765b124, []int{2}
}

func (m *DeleteConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectionRequest.Unmarshal(m, b)
}
func (m *DeleteConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectionRequest.Merge(m, src)
}
func (m *DeleteConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectionRequest.Size(m)
}
func (m *DeleteConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectionRequest proto.InternalMessageInfo

func (m *DeleteConnectionRequest) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

type DeleteConnectionReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConnectionReply) Reset()         { *m = DeleteConnectionReply{} }
func (m *DeleteConnectionReply) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectionReply) ProtoMessage()    {}
func (*DeleteConnectionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_084cb5dcc765b124, []int{3}
}

func (m *DeleteConnectionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectionReply.Unmarshal(m, b)
}
func (m *DeleteConnectionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectionReply.Marshal(b, m, deterministic)
}
func (m *DeleteConnectionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectionReply.Merge(m, src)
}
func (m *DeleteConnectionReply) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectionReply.Size(m)
}
func (m *DeleteConnectionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectionReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectionReply proto.InternalMessageInfo

type EnumConnectionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumConnectionRequest) Reset()         { *m = EnumConnectionRequest{} }
func (m *EnumConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*EnumConnectionRequest) ProtoMessage()    {}
func (*EnumConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_084cb5dcc765b124, []int{4}
}

func (m *EnumConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumConnectionRequest.Unmarshal(m, b)
}
func (m *EnumConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumConnectionRequest.Marshal(b, m, deterministic)
}
func (m *EnumConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumConnectionRequest.Merge(m, src)
}
func (m *EnumConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_EnumConnectionRequest.Size(m)
}
func (m *EnumConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnumConnectionRequest proto.InternalMessageInfo

type EnumConnectionReply struct {
	Workspace            []string `protobuf:"bytes,1,rep,name=workspace,proto3" json:"workspace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumConnectionReply) Reset()         { *m = EnumConnectionReply{} }
func (m *EnumConnectionReply) String() string { return proto.CompactTextString(m) }
func (*EnumConnectionReply) ProtoMessage()    {}
func (*EnumConnectionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_084cb5dcc765b124, []int{5}
}

func (m *EnumConnectionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumConnectionReply.Unmarshal(m, b)
}
func (m *EnumConnectionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumConnectionReply.Marshal(b, m, deterministic)
}
func (m *EnumConnectionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumConnectionReply.Merge(m, src)
}
func (m *EnumConnectionReply) XXX_Size() int {
	return xxx_messageInfo_EnumConnectionReply.Size(m)
}
func (m *EnumConnectionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumConnectionReply.DiscardUnknown(m)
}

var xxx_messageInfo_EnumConnectionReply proto.InternalMessageInfo

func (m *EnumConnectionReply) GetWorkspace() []string {
	if m != nil {
		return m.Workspace
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientConnectionRequest)(nil), "nsmdapi.ClientConnectionRequest")
	proto.RegisterType((*ClientConnectionReply)(nil), "nsmdapi.ClientConnectionReply")
	proto.RegisterType((*DeleteConnectionRequest)(nil), "nsmdapi.DeleteConnectionRequest")
	proto.RegisterType((*DeleteConnectionReply)(nil), "nsmdapi.DeleteConnectionReply")
	proto.RegisterType((*EnumConnectionRequest)(nil), "nsmdapi.EnumConnectionRequest")
	proto.RegisterType((*EnumConnectionReply)(nil), "nsmdapi.EnumConnectionReply")
}

func init() { proto.RegisterFile("nsmd.proto", fileDescriptor_084cb5dcc765b124) }

var fileDescriptor_084cb5dcc765b124 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x25, 0x6d, 0xbf, 0x7e, 0xf4, 0x8a, 0x8a, 0x23, 0x35, 0xa1, 0x14, 0x09, 0xc1, 0x45, 0x17,
	0x92, 0x85, 0x5d, 0xb8, 0xb7, 0x71, 0x69, 0x17, 0x8d, 0x20, 0xe8, 0x2a, 0xa6, 0x17, 0x0c, 0x49,
	0x66, 0xe2, 0xcc, 0x54, 0xe9, 0x53, 0xf8, 0xc8, 0xca, 0x8c, 0xa3, 0xcd, 0x6f, 0xc1, 0x5d, 0x38,
	0x7f, 0x39, 0x39, 0xb9, 0x00, 0x54, 0xe4, 0x6b, 0xbf, 0xe0, 0x4c, 0x32, 0xf2, 0x5f, 0x3d, 0x47,
	0x45, 0xe2, 0x5d, 0x83, 0xbd, 0xc8, 0x12, 0xa4, 0x72, 0xc1, 0x28, 0xc5, 0x58, 0x26, 0x8c, 0xae,
	0xf0, 0x75, 0x83, 0x42, 0x92, 0x29, 0x8c, 0xde, 0x19, 0x4f, 0x45, 0x11, 0xc5, 0xe8, 0x58, 0xae,
	0x35, 0x1b, 0xad, 0x76, 0x80, 0xf7, 0x69, 0xc1, 0xb8, 0xe9, 0x2c, 0xb2, 0xed, 0x7e, 0x1f, 0x71,
	0xe1, 0xe0, 0x85, 0x09, 0x79, 0x13, 0x09, 0x5c, 0x27, 0xdc, 0xe9, 0x69, 0xbe, 0x0c, 0x91, 0x0b,
	0x38, 0x8c, 0x75, 0xb0, 0x02, 0x82, 0x84, 0x3b, 0x7d, 0xad, 0xa9, 0x82, 0x64, 0x06, 0xc7, 0x54,
	0xe4, 0x21, 0xf2, 0x37, 0xe4, 0x21, 0x8b, 0x53, 0x94, 0xce, 0x40, 0xeb, 0xea, 0xb0, 0x51, 0x7e,
	0x77, 0x35, 0xca, 0x7f, 0xbf, 0xca, 0x32, 0x4c, 0x2e, 0xe1, 0x84, 0x8a, 0xfc, 0xe1, 0xa7, 0xeb,
	0x3d, 0x4b, 0x91, 0x3a, 0x43, 0xad, 0x6d, 0x12, 0x6a, 0xba, 0x00, 0x33, 0x94, 0xf8, 0xd7, 0xe9,
	0x6c, 0x18, 0x37, 0x8d, 0x45, 0xb6, 0x55, 0xc4, 0x2d, 0xdd, 0xe4, 0x8d, 0x3c, 0x6f, 0x0e, 0xa7,
	0x75, 0xa2, 0x65, 0xe9, 0x7e, 0xe5, 0x35, 0x57, 0x1f, 0x3d, 0x18, 0x2c, 0xc3, 0xbb, 0x80, 0x3c,
	0x81, 0x6d, 0x82, 0xea, 0x3f, 0x8c, 0xb8, 0xbe, 0x39, 0x04, 0xbf, 0xe3, 0x0a, 0x26, 0xe7, 0x7b,
	0x14, 0xaa, 0xc3, 0x12, 0x8e, 0xaa, 0xd5, 0xc8, 0xce, 0xd1, 0xfa, 0x31, 0x93, 0x69, 0x27, 0xaf,
	0xf2, 0x1e, 0xe1, 0xcc, 0x8c, 0xd3, 0xdd, 0xb5, 0x63, 0xf6, 0x52, 0xd7, 0xd6, 0x7d, 0x9f, 0x87,
	0xfa, 0xf8, 0xe7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xd3, 0xb5, 0x0e, 0x0a, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NSMDClient is the client API for NSMD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NSMDClient interface {
	RequestClientConnection(ctx context.Context, in *ClientConnectionRequest, opts ...grpc.CallOption) (*ClientConnectionReply, error)
	EnumConnection(ctx context.Context, in *EnumConnectionRequest, opts ...grpc.CallOption) (*EnumConnectionReply, error)
	DeleteClientConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionReply, error)
}

type nSMDClient struct {
	cc grpc.ClientConnInterface
}

func NewNSMDClient(cc grpc.ClientConnInterface) NSMDClient {
	return &nSMDClient{cc}
}

func (c *nSMDClient) RequestClientConnection(ctx context.Context, in *ClientConnectionRequest, opts ...grpc.CallOption) (*ClientConnectionReply, error) {
	out := new(ClientConnectionReply)
	err := c.cc.Invoke(ctx, "/nsmdapi.NSMD/RequestClientConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSMDClient) EnumConnection(ctx context.Context, in *EnumConnectionRequest, opts ...grpc.CallOption) (*EnumConnectionReply, error) {
	out := new(EnumConnectionReply)
	err := c.cc.Invoke(ctx, "/nsmdapi.NSMD/EnumConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSMDClient) DeleteClientConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionReply, error) {
	out := new(DeleteConnectionReply)
	err := c.cc.Invoke(ctx, "/nsmdapi.NSMD/DeleteClientConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NSMDServer is the server API for NSMD service.
type NSMDServer interface {
	RequestClientConnection(context.Context, *ClientConnectionRequest) (*ClientConnectionReply, error)
	EnumConnection(context.Context, *EnumConnectionRequest) (*EnumConnectionReply, error)
	DeleteClientConnection(context.Context, *DeleteConnectionRequest) (*DeleteConnectionReply, error)
}

// UnimplementedNSMDServer can be embedded to have forward compatible implementations.
type UnimplementedNSMDServer struct {
}

func (*UnimplementedNSMDServer) RequestClientConnection(ctx context.Context, req *ClientConnectionRequest) (*ClientConnectionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestClientConnection not implemented")
}
func (*UnimplementedNSMDServer) EnumConnection(ctx context.Context, req *EnumConnectionRequest) (*EnumConnectionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnumConnection not implemented")
}
func (*UnimplementedNSMDServer) DeleteClientConnection(ctx context.Context, req *DeleteConnectionRequest) (*DeleteConnectionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClientConnection not implemented")
}

func RegisterNSMDServer(s *grpc.Server, srv NSMDServer) {
	s.RegisterService(&_NSMD_serviceDesc, srv)
}

func _NSMD_RequestClientConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSMDServer).RequestClientConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsmdapi.NSMD/RequestClientConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSMDServer).RequestClientConnection(ctx, req.(*ClientConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSMD_EnumConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnumConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSMDServer).EnumConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsmdapi.NSMD/EnumConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSMDServer).EnumConnection(ctx, req.(*EnumConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSMD_DeleteClientConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSMDServer).DeleteClientConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsmdapi.NSMD/DeleteClientConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSMDServer).DeleteClientConnection(ctx, req.(*DeleteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NSMD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nsmdapi.NSMD",
	HandlerType: (*NSMDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestClientConnection",
			Handler:    _NSMD_RequestClientConnection_Handler,
		},
		{
			MethodName: "EnumConnection",
			Handler:    _NSMD_EnumConnection_Handler,
		},
		{
			MethodName: "DeleteClientConnection",
			Handler:    _NSMD_DeleteClientConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nsmd.proto",
}
