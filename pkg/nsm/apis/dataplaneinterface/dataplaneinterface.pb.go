// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplaneinterface.proto

package dataplaneinterface

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

// Message sent by dataplane module informing NSM of any changes in its
// operations parameters or constraints
type DataplaneUpdate struct {
	RemoteMechanism      []*common.RemoteMechanism `protobuf:"bytes,1,rep,name=remote_mechanism,json=remoteMechanism,proto3" json:"remote_mechanism,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DataplaneUpdate) Reset()         { *m = DataplaneUpdate{} }
func (m *DataplaneUpdate) String() string { return proto.CompactTextString(m) }
func (*DataplaneUpdate) ProtoMessage()    {}
func (*DataplaneUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneinterface_327f578c9150d455, []int{0}
}
func (m *DataplaneUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneUpdate.Unmarshal(m, b)
}
func (m *DataplaneUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneUpdate.Marshal(b, m, deterministic)
}
func (dst *DataplaneUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneUpdate.Merge(dst, src)
}
func (m *DataplaneUpdate) XXX_Size() int {
	return xxx_messageInfo_DataplaneUpdate.Size(m)
}
func (m *DataplaneUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneUpdate proto.InternalMessageInfo

func (m *DataplaneUpdate) GetRemoteMechanism() []*common.RemoteMechanism {
	if m != nil {
		return m.RemoteMechanism
	}
	return nil
}

func init() {
	proto.RegisterType((*DataplaneUpdate)(nil), "dataplaneinterface.DataplaneUpdate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataplaneOperationsClient is the client API for DataplaneOperations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneOperationsClient interface {
	UpdateDataplane(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (DataplaneOperations_UpdateDataplaneClient, error)
}

type dataplaneOperationsClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneOperationsClient(cc *grpc.ClientConn) DataplaneOperationsClient {
	return &dataplaneOperationsClient{cc}
}

func (c *dataplaneOperationsClient) UpdateDataplane(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (DataplaneOperations_UpdateDataplaneClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataplaneOperations_serviceDesc.Streams[0], "/dataplaneinterface.DataplaneOperations/UpdateDataplane", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataplaneOperationsUpdateDataplaneClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataplaneOperations_UpdateDataplaneClient interface {
	Recv() (*DataplaneUpdate, error)
	grpc.ClientStream
}

type dataplaneOperationsUpdateDataplaneClient struct {
	grpc.ClientStream
}

func (x *dataplaneOperationsUpdateDataplaneClient) Recv() (*DataplaneUpdate, error) {
	m := new(DataplaneUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataplaneOperationsServer is the server API for DataplaneOperations service.
type DataplaneOperationsServer interface {
	UpdateDataplane(*common.Empty, DataplaneOperations_UpdateDataplaneServer) error
}

func RegisterDataplaneOperationsServer(s *grpc.Server, srv DataplaneOperationsServer) {
	s.RegisterService(&_DataplaneOperations_serviceDesc, srv)
}

func _DataplaneOperations_UpdateDataplane_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(common.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataplaneOperationsServer).UpdateDataplane(m, &dataplaneOperationsUpdateDataplaneServer{stream})
}

type DataplaneOperations_UpdateDataplaneServer interface {
	Send(*DataplaneUpdate) error
	grpc.ServerStream
}

type dataplaneOperationsUpdateDataplaneServer struct {
	grpc.ServerStream
}

func (x *dataplaneOperationsUpdateDataplaneServer) Send(m *DataplaneUpdate) error {
	return x.ServerStream.SendMsg(m)
}

var _DataplaneOperations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneinterface.DataplaneOperations",
	HandlerType: (*DataplaneOperationsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateDataplane",
			Handler:       _DataplaneOperations_UpdateDataplane_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dataplaneinterface.proto",
}

func init() {
	proto.RegisterFile("dataplaneinterface.proto", fileDescriptor_dataplaneinterface_327f578c9150d455)
}

var fileDescriptor_dataplaneinterface_327f578c9150d455 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0x59, 0x04, 0x0f, 0x11, 0xa9, 0xc4, 0x83, 0xcb, 0x9e, 0x44, 0x2f, 0x9e, 0x12, 0x59,
	0xff, 0x81, 0xb8, 0x78, 0x12, 0xa1, 0xd0, 0xab, 0x32, 0x4d, 0xc7, 0x36, 0xb4, 0x93, 0x09, 0xc9,
	0xa8, 0xf8, 0xef, 0x85, 0x7e, 0x81, 0x1f, 0xa7, 0x81, 0x79, 0x98, 0x67, 0x5e, 0x5e, 0xb5, 0x6d,
	0x40, 0x20, 0x0e, 0x10, 0xd0, 0x07, 0xc1, 0xf4, 0x06, 0x0e, 0x4d, 0x4c, 0x2c, 0xac, 0xf5, 0x5f,
	0xb2, 0x3b, 0xb4, 0x5e, 0xba, 0xf7, 0xda, 0x38, 0x26, 0x3b, 0xf8, 0x16, 0x84, 0x6d, 0x40, 0xf9,
	0xe4, 0xd4, 0x67, 0x4c, 0x1f, 0xde, 0x21, 0x61, 0xee, 0x6c, 0xec, 0x5b, 0x1b, 0x32, 0x59, 0x88,
	0x3e, 0x5b, 0xc7, 0x44, 0x1c, 0xe6, 0x31, 0xa9, 0xaf, 0x2a, 0x55, 0x3c, 0x2c, 0xf2, 0x2a, 0x36,
	0x20, 0xa8, 0xef, 0xd5, 0x59, 0x42, 0x62, 0xc1, 0x57, 0x42, 0xd7, 0x41, 0xf0, 0x99, 0xb6, 0x9b,
	0xcb, 0xa3, 0x9b, 0x93, 0xfd, 0x85, 0x99, 0x6f, 0xcb, 0x91, 0x3f, 0x2d, 0xb8, 0x2c, 0xd2, 0xcf,
	0xc5, 0xfe, 0x45, 0x9d, 0xaf, 0xda, 0xe7, 0x88, 0x09, 0xc4, 0x73, 0xc8, 0xfa, 0x51, 0x15, 0xd3,
	0x93, 0x15, 0xea, 0xd3, 0xc5, 0x79, 0xa0, 0x28, 0x5f, 0xbb, 0x6b, 0xf3, 0x4f, 0x0b, 0xbf, 0x12,
	0xde, 0x6e, 0xea, 0xe3, 0x31, 0xfd, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xf9, 0x11,
	0x75, 0x34, 0x01, 0x00, 0x00,
}
