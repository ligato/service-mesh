// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplaneregistrar.proto

package dataplaneregistrar

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
	DataplaneName        string            `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName,proto3" json:"dataplane_name,omitempty"`
	DataplaneSocket      string            `protobuf:"bytes,2,opt,name=dataplane_socket,json=dataplaneSocket,proto3" json:"dataplane_socket,omitempty"`
	DataplaneParameters  map[string]string `protobuf:"bytes,3,rep,name=dataplane_parameters,json=dataplaneParameters,proto3" json:"dataplane_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DataplaneRegistrationRequest) Reset()         { *m = DataplaneRegistrationRequest{} }
func (m *DataplaneRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationRequest) ProtoMessage()    {}
func (*DataplaneRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_83b7c0b92c2471dc, []int{0}
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

func (m *DataplaneRegistrationRequest) GetDataplaneParameters() map[string]string {
	if m != nil {
		return m.DataplaneParameters
	}
	return nil
}

type DataplaneRegistrationReply struct {
	Registred            bool     `protobuf:"varint,1,opt,name=registred,proto3" json:"registred,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationReply) Reset()         { *m = DataplaneRegistrationReply{} }
func (m *DataplaneRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationReply) ProtoMessage()    {}
func (*DataplaneRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_83b7c0b92c2471dc, []int{1}
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

func (m *DataplaneRegistrationReply) GetRegistred() bool {
	if m != nil {
		return m.Registred
	}
	return false
}

func init() {
	proto.RegisterType((*DataplaneRegistrationRequest)(nil), "dataplaneregistrar.DataplaneRegistrationRequest")
	proto.RegisterMapType((map[string]string)(nil), "dataplaneregistrar.DataplaneRegistrationRequest.DataplaneParametersEntry")
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

// DataplaneRegistrationServer is the server API for DataplaneRegistration service.
type DataplaneRegistrationServer interface {
	RequestDataplaneRegistration(context.Context, *DataplaneRegistrationRequest) (*DataplaneRegistrationReply, error)
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

var _DataplaneRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneregistrar.DataplaneRegistration",
	HandlerType: (*DataplaneRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDataplaneRegistration",
			Handler:    _DataplaneRegistration_RequestDataplaneRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataplaneregistrar.proto",
}

func init() {
	proto.RegisterFile("dataplaneregistrar.proto", fileDescriptor_dataplaneregistrar_83b7c0b92c2471dc)
}

var fileDescriptor_dataplaneregistrar_83b7c0b92c2471dc = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x49, 0x2c, 0x49,
	0x2c, 0xc8, 0x49, 0xcc, 0x4b, 0x2d, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0x4a, 0x2c, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc2, 0x94, 0x51, 0xda, 0xc6, 0xc4, 0x25, 0xe3, 0x02, 0x13,
	0x0e, 0x82, 0x0a, 0x97, 0x64, 0xe6, 0xe7, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xa9,
	0x72, 0xf1, 0xc1, 0xb5, 0xc5, 0xe7, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xf1, 0xc2, 0x45, 0xfd, 0x12, 0x73, 0x53, 0x85, 0x34, 0xb9, 0x04, 0x10, 0xca, 0x8a, 0xf3, 0x93,
	0xb3, 0x53, 0x4b, 0x24, 0x98, 0xc0, 0x0a, 0xf9, 0xe1, 0xe2, 0xc1, 0x60, 0x61, 0xa1, 0x1a, 0x2e,
	0x11, 0x84, 0xd2, 0x82, 0xc4, 0xa2, 0xc4, 0xdc, 0xd4, 0x92, 0xd4, 0xa2, 0x62, 0x09, 0x66, 0x05,
	0x66, 0x0d, 0x6e, 0x23, 0x4f, 0x3d, 0x2c, 0xee, 0xc7, 0xe7, 0x42, 0x84, 0x64, 0x00, 0xdc, 0x2c,
	0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20, 0xe1, 0x14, 0x4c, 0x19, 0x29, 0x37, 0x2e, 0x09, 0x5c, 0x1a,
	0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0xa1, 0x1e, 0x04, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb,
	0x12, 0x73, 0x4a, 0x53, 0xa1, 0x7e, 0x81, 0x70, 0xac, 0x98, 0x2c, 0x18, 0x95, 0xac, 0xb8, 0xa4,
	0x70, 0xb8, 0xaa, 0x20, 0xa7, 0x52, 0x48, 0x86, 0x8b, 0x13, 0xea, 0xfa, 0xd4, 0x14, 0xb0, 0x79,
	0x1c, 0x41, 0x08, 0x01, 0xa3, 0x59, 0x8c, 0x5c, 0xa2, 0x58, 0x35, 0x0b, 0x35, 0x30, 0x72, 0xc9,
	0x40, 0xfd, 0x85, 0x5d, 0x81, 0x01, 0xa9, 0xc1, 0x23, 0xa5, 0x47, 0x82, 0x8e, 0x82, 0x9c, 0xca,
	0x24, 0x36, 0x70, 0x62, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x13, 0x71, 0x2c, 0x48,
	0x02, 0x00, 0x00,
}
