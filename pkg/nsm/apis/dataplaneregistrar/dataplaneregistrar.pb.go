// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplaneregistrar.proto

package dataplaneregistrar

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/ligato/networkservicemesh/pkg/nsm/apis/common"

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
	DataplaneName        string              `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName" json:"dataplane_name,omitempty"`
	DataplaneSocket      string              `protobuf:"bytes,2,opt,name=dataplane_socket,json=dataplaneSocket" json:"dataplane_socket,omitempty"`
	Mechanisms           []*common.Mechanism `protobuf:"bytes,3,rep,name=mechanisms" json:"mechanisms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DataplaneRegistrationRequest) Reset()         { *m = DataplaneRegistrationRequest{} }
func (m *DataplaneRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationRequest) ProtoMessage()    {}
func (*DataplaneRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_f7526e460ee5418a, []int{0}
}
func (m *DataplaneRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationRequest.Unmarshal(m, b)
}
func (m *DataplaneRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationRequest.Marshal(b, m, deterministic)
}
func (dst *DataplaneRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationRequest.Merge(dst, src)
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

func (m *DataplaneRegistrationRequest) GetMechanisms() []*common.Mechanism {
	if m != nil {
		return m.Mechanisms
	}
	return nil
}

type DataplaneRegistrationReply struct {
	Registered           bool     `protobuf:"varint,1,opt,name=registered" json:"registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationReply) Reset()         { *m = DataplaneRegistrationReply{} }
func (m *DataplaneRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationReply) ProtoMessage()    {}
func (*DataplaneRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_f7526e460ee5418a, []int{1}
}
func (m *DataplaneRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationReply.Unmarshal(m, b)
}
func (m *DataplaneRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationReply.Marshal(b, m, deterministic)
}
func (dst *DataplaneRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationReply.Merge(dst, src)
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

func init() {
	proto.RegisterType((*DataplaneRegistrationRequest)(nil), "dataplaneregistrar.DataplaneRegistrationRequest")
	proto.RegisterType((*DataplaneRegistrationReply)(nil), "dataplaneregistrar.DataplaneRegistrationReply")
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

// Server API for DataplaneRegistration service

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

func init() {
	proto.RegisterFile("dataplaneregistrar.proto", fileDescriptor_dataplaneregistrar_f7526e460ee5418a)
}

var fileDescriptor_dataplaneregistrar_f7526e460ee5418a = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xdd, 0x4a, 0xc3, 0x30,
	0x18, 0x25, 0x0e, 0x44, 0x3f, 0x99, 0xd3, 0x80, 0x50, 0xca, 0x90, 0x31, 0x10, 0xea, 0x4d, 0x33,
	0xb7, 0x5b, 0x2f, 0xdd, 0x9d, 0x7a, 0x51, 0x1f, 0x40, 0xb2, 0xee, 0xa3, 0x0d, 0x6d, 0x7e, 0x4c,
	0xb2, 0xc9, 0xee, 0x7c, 0x17, 0x1f, 0xc8, 0x57, 0x12, 0xbb, 0xb6, 0x4e, 0xac, 0xc2, 0xae, 0x42,
	0xce, 0x77, 0xce, 0xc9, 0xc9, 0xf9, 0x20, 0x58, 0x72, 0xcf, 0x4d, 0xc9, 0x15, 0x5a, 0xcc, 0x84,
	0xf3, 0x96, 0xdb, 0xd8, 0x58, 0xed, 0x35, 0xa5, 0xbf, 0x27, 0xe1, 0x3c, 0x13, 0x3e, 0x5f, 0x2d,
	0xe2, 0x54, 0x4b, 0x56, 0x8a, 0x8c, 0x7b, 0xcd, 0x14, 0xfa, 0x57, 0x6d, 0x0b, 0x87, 0x76, 0x2d,
	0x52, 0x94, 0xe8, 0x72, 0x66, 0x8a, 0x8c, 0x29, 0x27, 0x19, 0x37, 0xc2, 0xb1, 0x54, 0x4b, 0xa9,
	0x55, 0x7d, 0x6c, 0xad, 0xc7, 0xef, 0x04, 0x86, 0x77, 0x8d, 0x7b, 0x52, 0xbb, 0x7b, 0xa1, 0x55,
	0x82, 0x2f, 0x2b, 0x74, 0x9e, 0x5e, 0xc1, 0x69, 0xfb, 0xfa, 0xb3, 0xe2, 0x12, 0x03, 0x32, 0x22,
	0xd1, 0x71, 0xd2, 0x6f, 0xd1, 0x47, 0x2e, 0x91, 0x5e, 0xc3, 0xd9, 0x37, 0xcd, 0xe9, 0xb4, 0x40,
	0x1f, 0x1c, 0x54, 0xc4, 0x41, 0x8b, 0x3f, 0x55, 0x30, 0xbd, 0x01, 0x90, 0x98, 0xe6, 0x5c, 0x09,
	0x27, 0x5d, 0xd0, 0x1b, 0xf5, 0xa2, 0x93, 0xe9, 0x79, 0x5c, 0xa7, 0x7a, 0x68, 0x26, 0xc9, 0x0e,
	0x69, 0x7c, 0x0b, 0xe1, 0x1f, 0x21, 0x4d, 0xb9, 0xa1, 0x97, 0x00, 0xdb, 0x5e, 0xd0, 0xe2, 0xb2,
	0x8a, 0x77, 0x94, 0xec, 0x20, 0xd3, 0x0f, 0x02, 0x17, 0x9d, 0x72, 0xfa, 0x46, 0x60, 0x58, 0x7f,
	0xb4, 0x9b, 0x30, 0x89, 0x3b, 0x96, 0xf2, 0x5f, 0x5f, 0x61, 0xbc, 0x87, 0xe2, 0x2b, 0xfc, 0x0c,
	0x06, 0xb5, 0xf4, 0x5e, 0xac, 0x51, 0xa1, 0x73, 0xb4, 0xdf, 0x94, 0x31, 0x97, 0xc6, 0x6f, 0xc2,
	0x9f, 0xd7, 0x88, 0x4c, 0xc8, 0xe2, 0xb0, 0x5a, 0xde, 0xec, 0x33, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0xf4, 0xde, 0xad, 0x33, 0x02, 0x00, 0x00,
}
