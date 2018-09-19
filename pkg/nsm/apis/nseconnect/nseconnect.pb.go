// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nseconnect.proto

package nseconnect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import netmesh "github.com/ligato/networkservicemesh/pkg/nsm/apis/netmesh"

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

// EndpointConnectionRequest is sent by a NSM to NSE to build a connection.
type EndpointConnectionRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	NetworkServiceName   string   `protobuf:"bytes,2,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointConnectionRequest) Reset()         { *m = EndpointConnectionRequest{} }
func (m *EndpointConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*EndpointConnectionRequest) ProtoMessage()    {}
func (*EndpointConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_d3bee6930b5b72db, []int{0}
}
func (m *EndpointConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointConnectionRequest.Unmarshal(m, b)
}
func (m *EndpointConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointConnectionRequest.Marshal(b, m, deterministic)
}
func (dst *EndpointConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointConnectionRequest.Merge(dst, src)
}
func (m *EndpointConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_EndpointConnectionRequest.Size(m)
}
func (m *EndpointConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointConnectionRequest proto.InternalMessageInfo

func (m *EndpointConnectionRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointConnectionRequest) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

// EndpointConnectionReply is sent back by NSE to NSM with information required for
// dataplane programming.
type EndpointConnectionReply struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	NetworkServiceName   string   `protobuf:"bytes,2,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	LinuxNamespace       string   `protobuf:"bytes,3,opt,name=linux_namespace,json=linuxNamespace,proto3" json:"linux_namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointConnectionReply) Reset()         { *m = EndpointConnectionReply{} }
func (m *EndpointConnectionReply) String() string { return proto.CompactTextString(m) }
func (*EndpointConnectionReply) ProtoMessage()    {}
func (*EndpointConnectionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_d3bee6930b5b72db, []int{1}
}
func (m *EndpointConnectionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointConnectionReply.Unmarshal(m, b)
}
func (m *EndpointConnectionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointConnectionReply.Marshal(b, m, deterministic)
}
func (dst *EndpointConnectionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointConnectionReply.Merge(dst, src)
}
func (m *EndpointConnectionReply) XXX_Size() int {
	return xxx_messageInfo_EndpointConnectionReply.Size(m)
}
func (m *EndpointConnectionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointConnectionReply.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointConnectionReply proto.InternalMessageInfo

func (m *EndpointConnectionReply) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointConnectionReply) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func (m *EndpointConnectionReply) GetLinuxNamespace() string {
	if m != nil {
		return m.LinuxNamespace
	}
	return ""
}

type EndpointAdvertiseRequest struct {
	RequestId            string                          `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	NetworkEndpoint      *netmesh.NetworkServiceEndpoint `protobuf:"bytes,2,opt,name=network_endpoint,json=networkEndpoint,proto3" json:"network_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *EndpointAdvertiseRequest) Reset()         { *m = EndpointAdvertiseRequest{} }
func (m *EndpointAdvertiseRequest) String() string { return proto.CompactTextString(m) }
func (*EndpointAdvertiseRequest) ProtoMessage()    {}
func (*EndpointAdvertiseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_d3bee6930b5b72db, []int{2}
}
func (m *EndpointAdvertiseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointAdvertiseRequest.Unmarshal(m, b)
}
func (m *EndpointAdvertiseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointAdvertiseRequest.Marshal(b, m, deterministic)
}
func (dst *EndpointAdvertiseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointAdvertiseRequest.Merge(dst, src)
}
func (m *EndpointAdvertiseRequest) XXX_Size() int {
	return xxx_messageInfo_EndpointAdvertiseRequest.Size(m)
}
func (m *EndpointAdvertiseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointAdvertiseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointAdvertiseRequest proto.InternalMessageInfo

func (m *EndpointAdvertiseRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointAdvertiseRequest) GetNetworkEndpoint() *netmesh.NetworkServiceEndpoint {
	if m != nil {
		return m.NetworkEndpoint
	}
	return nil
}

type EndpointAdvertiseReply struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Accepted             bool     `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted,omitempty"`
	AdmissionError       string   `protobuf:"bytes,3,opt,name=admission_error,json=admissionError,proto3" json:"admission_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointAdvertiseReply) Reset()         { *m = EndpointAdvertiseReply{} }
func (m *EndpointAdvertiseReply) String() string { return proto.CompactTextString(m) }
func (*EndpointAdvertiseReply) ProtoMessage()    {}
func (*EndpointAdvertiseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_d3bee6930b5b72db, []int{3}
}
func (m *EndpointAdvertiseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointAdvertiseReply.Unmarshal(m, b)
}
func (m *EndpointAdvertiseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointAdvertiseReply.Marshal(b, m, deterministic)
}
func (dst *EndpointAdvertiseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointAdvertiseReply.Merge(dst, src)
}
func (m *EndpointAdvertiseReply) XXX_Size() int {
	return xxx_messageInfo_EndpointAdvertiseReply.Size(m)
}
func (m *EndpointAdvertiseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointAdvertiseReply.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointAdvertiseReply proto.InternalMessageInfo

func (m *EndpointAdvertiseReply) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointAdvertiseReply) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *EndpointAdvertiseReply) GetAdmissionError() string {
	if m != nil {
		return m.AdmissionError
	}
	return ""
}

// Message sent by NSE to NSM informing of a removal previously advertised
// endpoint. NSM will attempt to locate Customer Resource and delete it.
//
type EndpointRemoveRequest struct {
	RequestId            string                          `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	NetworkEndpoint      *netmesh.NetworkServiceEndpoint `protobuf:"bytes,2,opt,name=network_endpoint,json=networkEndpoint,proto3" json:"network_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *EndpointRemoveRequest) Reset()         { *m = EndpointRemoveRequest{} }
func (m *EndpointRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*EndpointRemoveRequest) ProtoMessage()    {}
func (*EndpointRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_d3bee6930b5b72db, []int{4}
}
func (m *EndpointRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointRemoveRequest.Unmarshal(m, b)
}
func (m *EndpointRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointRemoveRequest.Marshal(b, m, deterministic)
}
func (dst *EndpointRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointRemoveRequest.Merge(dst, src)
}
func (m *EndpointRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_EndpointRemoveRequest.Size(m)
}
func (m *EndpointRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointRemoveRequest proto.InternalMessageInfo

func (m *EndpointRemoveRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointRemoveRequest) GetNetworkEndpoint() *netmesh.NetworkServiceEndpoint {
	if m != nil {
		return m.NetworkEndpoint
	}
	return nil
}

type EndpointRemoveReply struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Accepted             bool     `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted,omitempty"`
	AdmissionError       string   `protobuf:"bytes,3,opt,name=admission_error,json=admissionError,proto3" json:"admission_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointRemoveReply) Reset()         { *m = EndpointRemoveReply{} }
func (m *EndpointRemoveReply) String() string { return proto.CompactTextString(m) }
func (*EndpointRemoveReply) ProtoMessage()    {}
func (*EndpointRemoveReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_nseconnect_d3bee6930b5b72db, []int{5}
}
func (m *EndpointRemoveReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointRemoveReply.Unmarshal(m, b)
}
func (m *EndpointRemoveReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointRemoveReply.Marshal(b, m, deterministic)
}
func (dst *EndpointRemoveReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointRemoveReply.Merge(dst, src)
}
func (m *EndpointRemoveReply) XXX_Size() int {
	return xxx_messageInfo_EndpointRemoveReply.Size(m)
}
func (m *EndpointRemoveReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointRemoveReply.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointRemoveReply proto.InternalMessageInfo

func (m *EndpointRemoveReply) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *EndpointRemoveReply) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *EndpointRemoveReply) GetAdmissionError() string {
	if m != nil {
		return m.AdmissionError
	}
	return ""
}

func init() {
	proto.RegisterType((*EndpointConnectionRequest)(nil), "nseconnect.EndpointConnectionRequest")
	proto.RegisterType((*EndpointConnectionReply)(nil), "nseconnect.EndpointConnectionReply")
	proto.RegisterType((*EndpointAdvertiseRequest)(nil), "nseconnect.EndpointAdvertiseRequest")
	proto.RegisterType((*EndpointAdvertiseReply)(nil), "nseconnect.EndpointAdvertiseReply")
	proto.RegisterType((*EndpointRemoveRequest)(nil), "nseconnect.EndpointRemoveRequest")
	proto.RegisterType((*EndpointRemoveReply)(nil), "nseconnect.EndpointRemoveReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointConnectionClient is the client API for EndpointConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointConnectionClient interface {
	RequestEndpointConnection(ctx context.Context, in *EndpointConnectionRequest, opts ...grpc.CallOption) (*EndpointConnectionReply, error)
}

type endpointConnectionClient struct {
	cc *grpc.ClientConn
}

func NewEndpointConnectionClient(cc *grpc.ClientConn) EndpointConnectionClient {
	return &endpointConnectionClient{cc}
}

func (c *endpointConnectionClient) RequestEndpointConnection(ctx context.Context, in *EndpointConnectionRequest, opts ...grpc.CallOption) (*EndpointConnectionReply, error) {
	out := new(EndpointConnectionReply)
	err := c.cc.Invoke(ctx, "/nseconnect.EndpointConnection/RequestEndpointConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointConnectionServer is the server API for EndpointConnection service.
type EndpointConnectionServer interface {
	RequestEndpointConnection(context.Context, *EndpointConnectionRequest) (*EndpointConnectionReply, error)
}

func RegisterEndpointConnectionServer(s *grpc.Server, srv EndpointConnectionServer) {
	s.RegisterService(&_EndpointConnection_serviceDesc, srv)
}

func _EndpointConnection_RequestEndpointConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointConnectionServer).RequestEndpointConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nseconnect.EndpointConnection/RequestEndpointConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointConnectionServer).RequestEndpointConnection(ctx, req.(*EndpointConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointConnection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nseconnect.EndpointConnection",
	HandlerType: (*EndpointConnectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestEndpointConnection",
			Handler:    _EndpointConnection_RequestEndpointConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nseconnect.proto",
}

// EndpointOperationsClient is the client API for EndpointOperations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointOperationsClient interface {
	AdvertiseEndpoint(ctx context.Context, in *EndpointAdvertiseRequest, opts ...grpc.CallOption) (*EndpointAdvertiseReply, error)
	RemoveEndpoint(ctx context.Context, in *EndpointRemoveRequest, opts ...grpc.CallOption) (*EndpointRemoveReply, error)
}

type endpointOperationsClient struct {
	cc *grpc.ClientConn
}

func NewEndpointOperationsClient(cc *grpc.ClientConn) EndpointOperationsClient {
	return &endpointOperationsClient{cc}
}

func (c *endpointOperationsClient) AdvertiseEndpoint(ctx context.Context, in *EndpointAdvertiseRequest, opts ...grpc.CallOption) (*EndpointAdvertiseReply, error) {
	out := new(EndpointAdvertiseReply)
	err := c.cc.Invoke(ctx, "/nseconnect.EndpointOperations/AdvertiseEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointOperationsClient) RemoveEndpoint(ctx context.Context, in *EndpointRemoveRequest, opts ...grpc.CallOption) (*EndpointRemoveReply, error) {
	out := new(EndpointRemoveReply)
	err := c.cc.Invoke(ctx, "/nseconnect.EndpointOperations/RemoveEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointOperationsServer is the server API for EndpointOperations service.
type EndpointOperationsServer interface {
	AdvertiseEndpoint(context.Context, *EndpointAdvertiseRequest) (*EndpointAdvertiseReply, error)
	RemoveEndpoint(context.Context, *EndpointRemoveRequest) (*EndpointRemoveReply, error)
}

func RegisterEndpointOperationsServer(s *grpc.Server, srv EndpointOperationsServer) {
	s.RegisterService(&_EndpointOperations_serviceDesc, srv)
}

func _EndpointOperations_AdvertiseEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointAdvertiseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointOperationsServer).AdvertiseEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nseconnect.EndpointOperations/AdvertiseEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointOperationsServer).AdvertiseEndpoint(ctx, req.(*EndpointAdvertiseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointOperations_RemoveEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointOperationsServer).RemoveEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nseconnect.EndpointOperations/RemoveEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointOperationsServer).RemoveEndpoint(ctx, req.(*EndpointRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointOperations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nseconnect.EndpointOperations",
	HandlerType: (*EndpointOperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdvertiseEndpoint",
			Handler:    _EndpointOperations_AdvertiseEndpoint_Handler,
		},
		{
			MethodName: "RemoveEndpoint",
			Handler:    _EndpointOperations_RemoveEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nseconnect.proto",
}

func init() { proto.RegisterFile("nseconnect.proto", fileDescriptor_nseconnect_d3bee6930b5b72db) }

var fileDescriptor_nseconnect_d3bee6930b5b72db = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x18, 0x25, 0x0a, 0xb2, 0xfb, 0x09, 0xbb, 0xeb, 0xf8, 0x97, 0x0d, 0xc8, 0x6a, 0x54, 0xf4, 0x2a,
	0x91, 0xfa, 0x04, 0x22, 0x45, 0xf4, 0xa2, 0x42, 0xf4, 0x56, 0xc2, 0x34, 0xf9, 0x48, 0x87, 0x26,
	0x33, 0xe3, 0xcc, 0xb4, 0x1a, 0xd0, 0x1b, 0xc1, 0x27, 0xf0, 0xa9, 0x7c, 0x2b, 0xc9, 0x64, 0x26,
	0xb5, 0x35, 0xda, 0x5e, 0x88, 0x5e, 0xb5, 0x39, 0xe7, 0xcc, 0x77, 0x4e, 0x0e, 0xdf, 0x04, 0xce,
	0xb8, 0xc6, 0x42, 0x70, 0x8e, 0x85, 0x49, 0xa4, 0x12, 0x46, 0x10, 0xd8, 0x20, 0xd1, 0x8b, 0x8a,
	0x99, 0xc5, 0x6a, 0x9e, 0x14, 0xa2, 0x49, 0x6b, 0x56, 0x51, 0x23, 0x52, 0x8e, 0xe6, 0x83, 0x50,
	0x4b, 0x8d, 0x6a, 0xcd, 0x0a, 0x6c, 0x50, 0x2f, 0x52, 0xb9, 0xac, 0x52, 0xae, 0x9b, 0x94, 0x4a,
	0xa6, 0x3b, 0xde, 0x82, 0xee, 0xb7, 0x1f, 0x1a, 0xd7, 0x70, 0x3e, 0xe5, 0xa5, 0x14, 0x8c, 0x9b,
	0xe7, 0xfd, 0x6c, 0x26, 0x78, 0x86, 0xef, 0x57, 0xa8, 0x0d, 0xb9, 0x03, 0xa0, 0xfa, 0xbf, 0x39,
	0x2b, 0xc3, 0xe0, 0x6e, 0xf0, 0xf8, 0x38, 0x3b, 0x76, 0xc8, 0xcb, 0x92, 0x3c, 0x81, 0x1b, 0xce,
	0x34, 0x77, 0xae, 0x39, 0xa7, 0x0d, 0x86, 0x97, 0xac, 0x90, 0x38, 0xee, 0x4d, 0x4f, 0xcd, 0x68,
	0x83, 0xf1, 0xb7, 0x00, 0x6e, 0x8f, 0xd9, 0xc9, 0xba, 0xfd, 0xeb, 0x66, 0xe4, 0x11, 0x9c, 0xd6,
	0x8c, 0xaf, 0x3e, 0x5a, 0x9d, 0x96, 0xb4, 0xc0, 0xf0, 0xb2, 0x15, 0x9f, 0x58, 0x78, 0xe6, 0xd1,
	0xf8, 0x6b, 0x00, 0xa1, 0x4f, 0xf5, 0xac, 0x5c, 0xa3, 0x32, 0x4c, 0xe3, 0x81, 0x1d, 0xbc, 0x82,
	0x33, 0x1f, 0x0b, 0xdd, 0x08, 0x1b, 0xe9, 0xea, 0xe4, 0x22, 0xf1, 0x4d, 0xcf, 0xb6, 0xb2, 0x79,
	0xa7, 0xec, 0xd4, 0x1d, 0xf4, 0x40, 0xfc, 0x09, 0x6e, 0x8d, 0xc4, 0x38, 0xa0, 0x9b, 0x08, 0x8e,
	0x68, 0x51, 0xa0, 0x34, 0x58, 0x5a, 0xf3, 0xa3, 0x6c, 0x78, 0xee, 0x5a, 0xa0, 0x65, 0xc3, 0xb4,
	0x66, 0x82, 0xe7, 0xa8, 0x94, 0x50, 0xbe, 0x85, 0x01, 0x9e, 0x76, 0x68, 0xfc, 0x25, 0x80, 0x9b,
	0x43, 0x36, 0x6c, 0xc4, 0xfa, 0x7f, 0x54, 0xd0, 0xc2, 0xf5, 0xdd, 0x0c, 0xff, 0xe8, 0xfd, 0x27,
	0x9f, 0x81, 0xfc, 0xba, 0x9a, 0xa4, 0x82, 0x73, 0x57, 0xc3, 0x08, 0xf9, 0x30, 0xf9, 0xe9, 0x92,
	0xfe, 0xf6, 0x1a, 0x45, 0xf7, 0xf7, 0xc9, 0x64, 0xdd, 0x4e, 0xbe, 0x07, 0x1b, 0xff, 0xd7, 0x12,
	0x15, 0xed, 0x28, 0x4d, 0xde, 0xc1, 0xb5, 0x61, 0x17, 0x3c, 0x4d, 0x1e, 0x8c, 0x0d, 0xdc, 0xdd,
	0xdc, 0x28, 0xde, 0xa3, 0xea, 0x8a, 0x7d, 0x0b, 0x27, 0x7d, 0xcf, 0xc3, 0xec, 0x7b, 0x63, 0xa7,
	0xb6, 0xf6, 0x21, 0xba, 0xf8, 0x93, 0x44, 0xd6, 0xed, 0xfc, 0x8a, 0xfd, 0xb6, 0x3c, 0xfd, 0x11,
	0x00, 0x00, 0xff, 0xff, 0xd8, 0xb0, 0x54, 0x06, 0xc4, 0x04, 0x00, 0x00,
}
