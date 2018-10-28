// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplaneregistrar.proto

package dataplaneregistrar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/ligato/networkservicemesh/pkg/nsm/apis/common"
	math "math"
)

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

// DataplaneRegistrationRequest is sent by the dataplane to NSM
// to advertise itself and inform NSM about the location of the dataplane socket
// and its initially supported parameters.
type DataplaneRegistrationRequest struct {
	DataplaneName        string   `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName,proto3" json:"dataplane_name,omitempty"`
	DataplaneSocket      string   `protobuf:"bytes,2,opt,name=dataplane_socket,json=dataplaneSocket,proto3" json:"dataplane_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationRequest) Reset()         { *m = DataplaneRegistrationRequest{} }
func (m *DataplaneRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationRequest) ProtoMessage()    {}
func (*DataplaneRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4c86488a7f7eab, []int{0}
}

func (m *DataplaneRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationRequest.Unmarshal(m, b)
}
func (m *DataplaneRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *DataplaneRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationRequest.Merge(m, src)
}
func (m *DataplaneRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_DataplaneRegistrationRequest.Size(m)
}
func (m *DataplaneRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneRegistrationRequest proto.InternalMessageInfo

func (m *DataplaneRegistrationRequest) GetDataplaneName() string {
	if m != nil {
		return m.DataplaneName
	}
	return ""
}

func (m *DataplaneRegistrationRequest) GetDataplaneSocket() string {
	if m != nil {
		return m.DataplaneSocket
	}
	return ""
}

type DataplaneRegistrationReply struct {
	Registered           bool     `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationReply) Reset()         { *m = DataplaneRegistrationReply{} }
func (m *DataplaneRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationReply) ProtoMessage()    {}
func (*DataplaneRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4c86488a7f7eab, []int{1}
}

func (m *DataplaneRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationReply.Unmarshal(m, b)
}
func (m *DataplaneRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationReply.Marshal(b, m, deterministic)
}
func (m *DataplaneRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationReply.Merge(m, src)
}
func (m *DataplaneRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_DataplaneRegistrationReply.Size(m)
}
func (m *DataplaneRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneRegistrationReply proto.InternalMessageInfo

func (m *DataplaneRegistrationReply) GetRegistered() bool {
	if m != nil {
		return m.Registered
	}
	return false
}

// DataplaneUnRegistrationRequest is sent by the dataplane to NSM
// to remove itself from the list of available dataplanes.
type DataplaneUnRegistrationRequest struct {
	DataplaneName        string   `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName,proto3" json:"dataplane_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneUnRegistrationRequest) Reset()         { *m = DataplaneUnRegistrationRequest{} }
func (m *DataplaneUnRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneUnRegistrationRequest) ProtoMessage()    {}
func (*DataplaneUnRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4c86488a7f7eab, []int{2}
}

func (m *DataplaneUnRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Unmarshal(m, b)
}
func (m *DataplaneUnRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *DataplaneUnRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneUnRegistrationRequest.Merge(m, src)
}
func (m *DataplaneUnRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Size(m)
}
func (m *DataplaneUnRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneUnRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneUnRegistrationRequest proto.InternalMessageInfo

func (m *DataplaneUnRegistrationRequest) GetDataplaneName() string {
	if m != nil {
		return m.DataplaneName
	}
	return ""
}

type DataplaneUnRegistrationReply struct {
	UnRegistered         bool     `protobuf:"varint,1,opt,name=un_registered,json=unRegistered,proto3" json:"un_registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneUnRegistrationReply) Reset()         { *m = DataplaneUnRegistrationReply{} }
func (m *DataplaneUnRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneUnRegistrationReply) ProtoMessage()    {}
func (*DataplaneUnRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4c86488a7f7eab, []int{3}
}

func (m *DataplaneUnRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Unmarshal(m, b)
}
func (m *DataplaneUnRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Marshal(b, m, deterministic)
}
func (m *DataplaneUnRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneUnRegistrationReply.Merge(m, src)
}
func (m *DataplaneUnRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Size(m)
}
func (m *DataplaneUnRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneUnRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneUnRegistrationReply proto.InternalMessageInfo

func (m *DataplaneUnRegistrationReply) GetUnRegistered() bool {
	if m != nil {
		return m.UnRegistered
	}
	return false
}

func init() {
	proto.RegisterType((*DataplaneRegistrationRequest)(nil), "dataplaneregistrar.DataplaneRegistrationRequest")
	proto.RegisterType((*DataplaneRegistrationReply)(nil), "dataplaneregistrar.DataplaneRegistrationReply")
	proto.RegisterType((*DataplaneUnRegistrationRequest)(nil), "dataplaneregistrar.DataplaneUnRegistrationRequest")
	proto.RegisterType((*DataplaneUnRegistrationReply)(nil), "dataplaneregistrar.DataplaneUnRegistrationReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataplaneRegistrationClient is the client API for DataplaneRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneRegistrationClient interface {
	RequestDataplaneRegistration(ctx context.Context, in *DataplaneRegistrationRequest, opts ...grpc.CallOption) (*DataplaneRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the dataplane that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the dataplane needs to start re-registration logic.
	RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (DataplaneRegistration_RequestLivenessClient, error)
}

type dataplaneRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneRegistrationClient(cc *grpc.ClientConn) DataplaneRegistrationClient {
	return &dataplaneRegistrationClient{cc}
}

func (c *dataplaneRegistrationClient) RequestDataplaneRegistration(ctx context.Context, in *DataplaneRegistrationRequest, opts ...grpc.CallOption) (*DataplaneRegistrationReply, error) {
	out := new(DataplaneRegistrationReply)
	err := c.cc.Invoke(ctx, "/dataplaneregistrar.DataplaneRegistration/RequestDataplaneRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataplaneRegistrationClient) RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (DataplaneRegistration_RequestLivenessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataplaneRegistration_serviceDesc.Streams[0], "/dataplaneregistrar.DataplaneRegistration/RequestLiveness", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataplaneRegistrationRequestLivenessClient{stream}
	return x, nil
}

type DataplaneRegistration_RequestLivenessClient interface {
	Send(*common.Empty) error
	Recv() (*common.Empty, error)
	grpc.ClientStream
}

type dataplaneRegistrationRequestLivenessClient struct {
	grpc.ClientStream
}

func (x *dataplaneRegistrationRequestLivenessClient) Send(m *common.Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataplaneRegistrationRequestLivenessClient) Recv() (*common.Empty, error) {
	m := new(common.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataplaneRegistrationServer is the server API for DataplaneRegistration service.
type DataplaneRegistrationServer interface {
	RequestDataplaneRegistration(context.Context, *DataplaneRegistrationRequest) (*DataplaneRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the dataplane that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the dataplane needs to start re-registration logic.
	RequestLiveness(DataplaneRegistration_RequestLivenessServer) error
}

func RegisterDataplaneRegistrationServer(s *grpc.Server, srv DataplaneRegistrationServer) {
	s.RegisterService(&_DataplaneRegistration_serviceDesc, srv)
}

func _DataplaneRegistration_RequestDataplaneRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataplaneRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneRegistrationServer).RequestDataplaneRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplaneregistrar.DataplaneRegistration/RequestDataplaneRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneRegistrationServer).RequestDataplaneRegistration(ctx, req.(*DataplaneRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataplaneRegistration_RequestLiveness_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataplaneRegistrationServer).RequestLiveness(&dataplaneRegistrationRequestLivenessServer{stream})
}

type DataplaneRegistration_RequestLivenessServer interface {
	Send(*common.Empty) error
	Recv() (*common.Empty, error)
	grpc.ServerStream
}

type dataplaneRegistrationRequestLivenessServer struct {
	grpc.ServerStream
}

func (x *dataplaneRegistrationRequestLivenessServer) Send(m *common.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataplaneRegistrationRequestLivenessServer) Recv() (*common.Empty, error) {
	m := new(common.Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataplaneRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneregistrar.DataplaneRegistration",
	HandlerType: (*DataplaneRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDataplaneRegistration",
			Handler:    _DataplaneRegistration_RequestDataplaneRegistration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestLiveness",
			Handler:       _DataplaneRegistration_RequestLiveness_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dataplaneregistrar.proto",
}

// DataplaneUnRegistrationClient is the client API for DataplaneUnRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneUnRegistrationClient interface {
	RequestDataplaneUnRegistration(ctx context.Context, in *DataplaneUnRegistrationRequest, opts ...grpc.CallOption) (*DataplaneUnRegistrationReply, error)
}

type dataplaneUnRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneUnRegistrationClient(cc *grpc.ClientConn) DataplaneUnRegistrationClient {
	return &dataplaneUnRegistrationClient{cc}
}

func (c *dataplaneUnRegistrationClient) RequestDataplaneUnRegistration(ctx context.Context, in *DataplaneUnRegistrationRequest, opts ...grpc.CallOption) (*DataplaneUnRegistrationReply, error) {
	out := new(DataplaneUnRegistrationReply)
	err := c.cc.Invoke(ctx, "/dataplaneregistrar.DataplaneUnRegistration/RequestDataplaneUnRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataplaneUnRegistrationServer is the server API for DataplaneUnRegistration service.
type DataplaneUnRegistrationServer interface {
	RequestDataplaneUnRegistration(context.Context, *DataplaneUnRegistrationRequest) (*DataplaneUnRegistrationReply, error)
}

func RegisterDataplaneUnRegistrationServer(s *grpc.Server, srv DataplaneUnRegistrationServer) {
	s.RegisterService(&_DataplaneUnRegistration_serviceDesc, srv)
}

func _DataplaneUnRegistration_RequestDataplaneUnRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataplaneUnRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneUnRegistrationServer).RequestDataplaneUnRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplaneregistrar.DataplaneUnRegistration/RequestDataplaneUnRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneUnRegistrationServer).RequestDataplaneUnRegistration(ctx, req.(*DataplaneUnRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataplaneUnRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneregistrar.DataplaneUnRegistration",
	HandlerType: (*DataplaneUnRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDataplaneUnRegistration",
			Handler:    _DataplaneUnRegistration_RequestDataplaneUnRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataplaneregistrar.proto",
}

func init() { proto.RegisterFile("dataplaneregistrar.proto", fileDescriptor_7f4c86488a7f7eab) }

var fileDescriptor_7f4c86488a7f7eab = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x64, 0x3d, 0x88, 0x3e, 0xac, 0x95, 0x05, 0xb1, 0x84, 0x52, 0x24, 0x22, 0xd4, 0x4b, 0x52,
	0xda, 0xab, 0x37, 0x2d, 0x5e, 0xc4, 0x43, 0xc4, 0x73, 0xd9, 0xa6, 0x8f, 0x74, 0x69, 0xf6, 0xc3,
	0xdd, 0x4d, 0xa5, 0x37, 0x4f, 0xfe, 0x10, 0xff, 0x90, 0x7f, 0x49, 0x9a, 0xc6, 0x58, 0xd3, 0xb4,
	0x50, 0x4f, 0x21, 0xb3, 0xf3, 0xe6, 0xcd, 0xec, 0x2c, 0xb4, 0x26, 0xcc, 0x31, 0x9d, 0x32, 0x89,
	0x06, 0x13, 0x6e, 0x9d, 0x61, 0x26, 0xd0, 0x46, 0x39, 0x45, 0xe9, 0xe6, 0x89, 0x37, 0x4c, 0xb8,
	0x9b, 0x66, 0xe3, 0x20, 0x56, 0x22, 0x4c, 0x79, 0xc2, 0x9c, 0x0a, 0x25, 0xba, 0x37, 0x65, 0x66,
	0x16, 0xcd, 0x9c, 0xc7, 0x28, 0xd0, 0x4e, 0x43, 0x3d, 0x4b, 0x42, 0x69, 0x45, 0xc8, 0x34, 0xb7,
	0x61, 0xac, 0x84, 0x50, 0xb2, 0xf8, 0xac, 0xa4, 0x7d, 0x0d, 0xed, 0xfb, 0x1f, 0xf1, 0xa8, 0x10,
	0x77, 0x5c, 0xc9, 0x08, 0x5f, 0x33, 0xb4, 0x8e, 0x5e, 0xc3, 0x69, 0xb9, 0x7c, 0x24, 0x99, 0xc0,
	0x16, 0xb9, 0x24, 0xdd, 0xe3, 0xa8, 0x51, 0xa2, 0x4f, 0x4c, 0x20, 0xbd, 0x81, 0xb3, 0x5f, 0x9a,
	0x55, 0xf1, 0x0c, 0x5d, 0xeb, 0x20, 0x27, 0x36, 0x4b, 0xfc, 0x39, 0x87, 0xfd, 0x5b, 0xf0, 0xb6,
	0x6c, 0xd4, 0xe9, 0x82, 0x76, 0x00, 0x56, 0x19, 0xd1, 0xe0, 0x24, 0xdf, 0x75, 0x14, 0xad, 0x21,
	0xfe, 0x03, 0x74, 0xca, 0xe9, 0x17, 0xf9, 0x7f, 0xc7, 0xfe, 0xdd, 0x5a, 0xf0, 0xaa, 0xd0, 0xd2,
	0xc8, 0x15, 0x34, 0x32, 0x39, 0xda, 0xf0, 0x72, 0x92, 0x15, 0xdc, 0x25, 0xd6, 0xff, 0x22, 0x70,
	0x5e, 0x1b, 0x86, 0xbe, 0x13, 0x68, 0x17, 0x8e, 0xea, 0x09, 0xbd, 0xa0, 0xa6, 0xee, 0x5d, 0x55,
	0x78, 0xc1, 0x1e, 0x13, 0xcb, 0x04, 0x03, 0x68, 0x16, 0xa3, 0x8f, 0x7c, 0x8e, 0x12, 0xad, 0xa5,
	0x8d, 0xa0, 0x28, 0x7f, 0x28, 0xb4, 0x5b, 0x78, 0x7f, 0x7f, 0xbb, 0xa4, 0x47, 0xfa, 0x9f, 0x04,
	0x2e, 0xb6, 0xdc, 0x0b, 0xfd, 0x20, 0xd0, 0xa9, 0x66, 0xaa, 0x50, 0xfa, 0x3b, 0x3d, 0xd6, 0x16,
	0xe6, 0xf5, 0xf6, 0x9a, 0xd1, 0xe9, 0x62, 0x7c, 0x98, 0xbf, 0xdd, 0xc1, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5e, 0x22, 0x80, 0xb2, 0x32, 0x03, 0x00, 0x00,
}
