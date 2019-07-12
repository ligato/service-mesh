// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crossconnect.proto

package crossconnect

import (
	context "context"
	connection "controlplane/pkg/apis/local/connection"
	connection1 "controlplane/pkg/apis/remote/connection"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type CrossConnectEventType int32

const (
	CrossConnectEventType_INITIAL_STATE_TRANSFER CrossConnectEventType = 0
	CrossConnectEventType_UPDATE                 CrossConnectEventType = 1
	CrossConnectEventType_DELETE                 CrossConnectEventType = 2
)

var CrossConnectEventType_name = map[int32]string{
	0: "INITIAL_STATE_TRANSFER",
	1: "UPDATE",
	2: "DELETE",
}

var CrossConnectEventType_value = map[string]int32{
	"INITIAL_STATE_TRANSFER": 0,
	"UPDATE":                 1,
	"DELETE":                 2,
}

func (x CrossConnectEventType) String() string {
	return proto.EnumName(CrossConnectEventType_name, int32(x))
}

func (CrossConnectEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_97acf85fcaabb3f6, []int{0}
}

type Metrics struct {
	Metrics              map[string]string `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_97acf85fcaabb3f6, []int{0}
}

func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (m *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(m, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetMetrics() map[string]string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type CrossConnectEvent struct {
	Type                 CrossConnectEventType    `protobuf:"varint,1,opt,name=type,proto3,enum=crossconnect.CrossConnectEventType" json:"type,omitempty"`
	CrossConnects        map[string]*CrossConnect `protobuf:"bytes,2,rep,name=cross_connects,json=crossConnects,proto3" json:"cross_connects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metrics              map[string]*Metrics      `protobuf:"bytes,3,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CrossConnectEvent) Reset()         { *m = CrossConnectEvent{} }
func (m *CrossConnectEvent) String() string { return proto.CompactTextString(m) }
func (*CrossConnectEvent) ProtoMessage()    {}
func (*CrossConnectEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_97acf85fcaabb3f6, []int{1}
}

func (m *CrossConnectEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossConnectEvent.Unmarshal(m, b)
}
func (m *CrossConnectEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossConnectEvent.Marshal(b, m, deterministic)
}
func (m *CrossConnectEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossConnectEvent.Merge(m, src)
}
func (m *CrossConnectEvent) XXX_Size() int {
	return xxx_messageInfo_CrossConnectEvent.Size(m)
}
func (m *CrossConnectEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossConnectEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CrossConnectEvent proto.InternalMessageInfo

func (m *CrossConnectEvent) GetType() CrossConnectEventType {
	if m != nil {
		return m.Type
	}
	return CrossConnectEventType_INITIAL_STATE_TRANSFER
}

func (m *CrossConnectEvent) GetCrossConnects() map[string]*CrossConnect {
	if m != nil {
		return m.CrossConnects
	}
	return nil
}

func (m *CrossConnectEvent) GetMetrics() map[string]*Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type CrossConnect struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Types that are valid to be assigned to Source:
	//	*CrossConnect_LocalSource
	//	*CrossConnect_RemoteSource
	Source isCrossConnect_Source `protobuf_oneof:"source"`
	// Types that are valid to be assigned to Destination:
	//	*CrossConnect_LocalDestination
	//	*CrossConnect_RemoteDestination
	Destination          isCrossConnect_Destination `protobuf_oneof:"destination"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CrossConnect) Reset()         { *m = CrossConnect{} }
func (m *CrossConnect) String() string { return proto.CompactTextString(m) }
func (*CrossConnect) ProtoMessage()    {}
func (*CrossConnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_97acf85fcaabb3f6, []int{2}
}

func (m *CrossConnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossConnect.Unmarshal(m, b)
}
func (m *CrossConnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossConnect.Marshal(b, m, deterministic)
}
func (m *CrossConnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossConnect.Merge(m, src)
}
func (m *CrossConnect) XXX_Size() int {
	return xxx_messageInfo_CrossConnect.Size(m)
}
func (m *CrossConnect) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossConnect.DiscardUnknown(m)
}

var xxx_messageInfo_CrossConnect proto.InternalMessageInfo

func (m *CrossConnect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CrossConnect) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type isCrossConnect_Source interface {
	isCrossConnect_Source()
}

type CrossConnect_LocalSource struct {
	LocalSource *connection.Connection `protobuf:"bytes,4,opt,name=local_source,json=localSource,proto3,oneof"`
}

type CrossConnect_RemoteSource struct {
	RemoteSource *connection1.Connection `protobuf:"bytes,5,opt,name=remote_source,json=remoteSource,proto3,oneof"`
}

func (*CrossConnect_LocalSource) isCrossConnect_Source() {}

func (*CrossConnect_RemoteSource) isCrossConnect_Source() {}

func (m *CrossConnect) GetSource() isCrossConnect_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *CrossConnect) GetLocalSource() *connection.Connection {
	if x, ok := m.GetSource().(*CrossConnect_LocalSource); ok {
		return x.LocalSource
	}
	return nil
}

func (m *CrossConnect) GetRemoteSource() *connection1.Connection {
	if x, ok := m.GetSource().(*CrossConnect_RemoteSource); ok {
		return x.RemoteSource
	}
	return nil
}

type isCrossConnect_Destination interface {
	isCrossConnect_Destination()
}

type CrossConnect_LocalDestination struct {
	LocalDestination *connection.Connection `protobuf:"bytes,6,opt,name=local_destination,json=localDestination,proto3,oneof"`
}

type CrossConnect_RemoteDestination struct {
	RemoteDestination *connection1.Connection `protobuf:"bytes,7,opt,name=remote_destination,json=remoteDestination,proto3,oneof"`
}

func (*CrossConnect_LocalDestination) isCrossConnect_Destination() {}

func (*CrossConnect_RemoteDestination) isCrossConnect_Destination() {}

func (m *CrossConnect) GetDestination() isCrossConnect_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *CrossConnect) GetLocalDestination() *connection.Connection {
	if x, ok := m.GetDestination().(*CrossConnect_LocalDestination); ok {
		return x.LocalDestination
	}
	return nil
}

func (m *CrossConnect) GetRemoteDestination() *connection1.Connection {
	if x, ok := m.GetDestination().(*CrossConnect_RemoteDestination); ok {
		return x.RemoteDestination
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CrossConnect) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CrossConnect_LocalSource)(nil),
		(*CrossConnect_RemoteSource)(nil),
		(*CrossConnect_LocalDestination)(nil),
		(*CrossConnect_RemoteDestination)(nil),
	}
}

func init() {
	proto.RegisterEnum("crossconnect.CrossConnectEventType", CrossConnectEventType_name, CrossConnectEventType_value)
	proto.RegisterType((*Metrics)(nil), "crossconnect.Metrics")
	proto.RegisterMapType((map[string]string)(nil), "crossconnect.Metrics.MetricsEntry")
	proto.RegisterType((*CrossConnectEvent)(nil), "crossconnect.CrossConnectEvent")
	proto.RegisterMapType((map[string]*CrossConnect)(nil), "crossconnect.CrossConnectEvent.CrossConnectsEntry")
	proto.RegisterMapType((map[string]*Metrics)(nil), "crossconnect.CrossConnectEvent.MetricsEntry")
	proto.RegisterType((*CrossConnect)(nil), "crossconnect.CrossConnect")
}

func init() { proto.RegisterFile("crossconnect.proto", fileDescriptor_97acf85fcaabb3f6) }

var fileDescriptor_97acf85fcaabb3f6 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xda, 0x4c,
	0x10, 0x8e, 0x4d, 0x02, 0xef, 0x3b, 0x7c, 0x08, 0xb6, 0x49, 0x64, 0x59, 0xad, 0x1a, 0xd1, 0x4b,
	0xd4, 0x56, 0xeb, 0xd4, 0x39, 0x24, 0x42, 0x3d, 0x94, 0x80, 0xa3, 0xa2, 0x24, 0xa8, 0x31, 0xee,
	0xa1, 0x52, 0x25, 0x64, 0xcc, 0xd6, 0xb1, 0x62, 0xbc, 0x96, 0xbd, 0x20, 0xf9, 0xdc, 0x3f, 0xd2,
	0x7f, 0xd9, 0x6b, 0xe5, 0x5d, 0xd3, 0x2c, 0x8a, 0x53, 0x72, 0x81, 0xf1, 0x33, 0xf3, 0x3c, 0xf3,
	0x69, 0x03, 0xf2, 0x12, 0x9a, 0xa6, 0x1e, 0x8d, 0x22, 0xe2, 0x31, 0x1c, 0x27, 0x94, 0x51, 0xd4,
	0x90, 0x31, 0xfd, 0xcc, 0xa3, 0x11, 0x4b, 0x68, 0x18, 0x87, 0x6e, 0x44, 0x8c, 0xf8, 0xde, 0x37,
	0xdc, 0x38, 0x48, 0x8d, 0x90, 0x7a, 0x6e, 0x68, 0x14, 0x41, 0x01, 0x8d, 0x24, 0x53, 0xc8, 0xe8,
	0xe7, 0xe5, 0xc4, 0x84, 0x2c, 0x28, 0x23, 0xff, 0x64, 0xf6, 0xfc, 0x80, 0xdd, 0x2d, 0x67, 0xd8,
	0xa3, 0x0b, 0xc3, 0xa7, 0xa1, 0x1b, 0xf9, 0x06, 0x77, 0xcc, 0x96, 0x3f, 0x3e, 0xad, 0x3e, 0xe0,
	0x53, 0x6c, 0x1a, 0x31, 0xcb, 0x62, 0x92, 0x1a, 0x64, 0x11, 0xb3, 0x4c, 0xfc, 0x0a, 0x6e, 0xf7,
	0xa7, 0x02, 0xb5, 0x1b, 0xc2, 0x92, 0xc0, 0x4b, 0xd1, 0x47, 0xa8, 0x2d, 0x84, 0xa9, 0x29, 0x47,
	0x95, 0xe3, 0xba, 0xd9, 0xc5, 0x1b, 0xed, 0x16, 0x71, 0xeb, 0x7f, 0x2b, 0x62, 0x49, 0x66, 0xaf,
	0x29, 0x7a, 0x0f, 0x1a, 0xb2, 0x03, 0xb5, 0xa1, 0x72, 0x4f, 0x32, 0x4d, 0x39, 0x52, 0x8e, 0xff,
	0xb7, 0x73, 0x13, 0xed, 0xc3, 0xde, 0xca, 0x0d, 0x97, 0x44, 0x53, 0x39, 0x26, 0x1e, 0x7a, 0xea,
	0xb9, 0xd2, 0xfd, 0x55, 0x81, 0xce, 0x20, 0x4f, 0x35, 0x10, 0xa9, 0xac, 0x15, 0x89, 0x18, 0x3a,
	0x83, 0xdd, 0xbc, 0x6c, 0x2e, 0xd1, 0x32, 0xdf, 0x6c, 0x16, 0xf3, 0x28, 0xdc, 0xc9, 0x62, 0x62,
	0x73, 0x02, 0xfa, 0x06, 0x2d, 0x1e, 0x3b, 0x2d, 0x82, 0x53, 0x4d, 0xe5, 0xfd, 0x98, 0x5b, 0x24,
	0x36, 0x90, 0xa2, 0xbf, 0xa6, 0x27, 0x63, 0xe8, 0xf2, 0x61, 0x46, 0x15, 0xae, 0xf9, 0x7e, 0x9b,
	0x66, 0xf9, 0xb4, 0xbe, 0x03, 0x7a, 0x9c, 0xac, 0x64, 0x66, 0x27, 0xf2, 0xcc, 0xea, 0xa6, 0xfe,
	0x74, 0x36, 0x69, 0x9e, 0xfa, 0xed, 0xd6, 0x5d, 0xbc, 0xdb, 0xd4, 0x3d, 0x28, 0xdd, 0xb4, 0xbc,
	0xa2, 0xdf, 0x2a, 0x34, 0xe4, 0x74, 0xa8, 0x05, 0x6a, 0x30, 0x2f, 0x24, 0xd5, 0x60, 0x8e, 0x34,
	0xa8, 0xc5, 0x6e, 0x16, 0x52, 0x77, 0x5e, 0xec, 0x77, 0xfd, 0x88, 0xfa, 0xd0, 0xe0, 0xe7, 0x3f,
	0x4d, 0xe9, 0x32, 0xf1, 0x88, 0xb6, 0xcb, 0x53, 0xbe, 0xc4, 0x1c, 0xc4, 0xd2, 0x39, 0x0f, 0xfe,
	0x9a, 0x9f, 0x77, 0xec, 0x3a, 0x77, 0x4f, 0x38, 0x05, 0x0d, 0xa1, 0x29, 0x5e, 0x84, 0xb5, 0xc6,
	0x1e, 0xd7, 0x78, 0x85, 0x05, 0xfa, 0xa4, 0x48, 0x43, 0xf8, 0x0b, 0x95, 0x2b, 0xe8, 0x88, 0x42,
	0xe6, 0x24, 0x65, 0x41, 0xe4, 0xe6, 0x41, 0x5a, 0xf5, 0x19, 0xd5, 0x28, 0x76, 0x9b, 0xbb, 0x87,
	0x0f, 0x3c, 0x34, 0x06, 0x54, 0x94, 0x24, 0xab, 0xd5, 0x9e, 0x53, 0x97, 0x62, 0x77, 0x84, 0x5f,
	0xd2, 0xbb, 0xf8, 0x0f, 0xaa, 0xa2, 0xb7, 0x8b, 0x26, 0xd4, 0x25, 0xc9, 0xb7, 0x57, 0x70, 0x50,
	0x7a, 0xec, 0x48, 0x87, 0xc3, 0xd1, 0x78, 0xe4, 0x8c, 0xfa, 0xd7, 0xd3, 0x89, 0xd3, 0x77, 0xac,
	0xa9, 0x63, 0xf7, 0xc7, 0x93, 0x4b, 0xcb, 0x6e, 0xef, 0x20, 0x80, 0xea, 0xd7, 0x2f, 0xc3, 0xbe,
	0x63, 0xb5, 0x95, 0xdc, 0x1e, 0x5a, 0xd7, 0x96, 0x63, 0xb5, 0x55, 0xf3, 0x0e, 0x5e, 0xdc, 0xd0,
	0x28, 0x60, 0x34, 0xd9, 0x58, 0xe6, 0x2d, 0xec, 0x97, 0xc0, 0x29, 0x3a, 0xc4, 0x3e, 0xa5, 0x7e,
	0x48, 0xf0, 0xfa, 0x83, 0x82, 0xad, 0xfc, 0xe3, 0xa1, 0xbf, 0xde, 0x72, 0xf5, 0x27, 0xca, 0xac,
	0xca, 0x29, 0xa7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x80, 0x46, 0xd5, 0x8f, 0x33, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitorCrossConnectClient is the client API for MonitorCrossConnect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorCrossConnectClient interface {
	MonitorCrossConnects(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MonitorCrossConnect_MonitorCrossConnectsClient, error)
}

type monitorCrossConnectClient struct {
	cc *grpc.ClientConn
}

func NewMonitorCrossConnectClient(cc *grpc.ClientConn) MonitorCrossConnectClient {
	return &monitorCrossConnectClient{cc}
}

func (c *monitorCrossConnectClient) MonitorCrossConnects(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MonitorCrossConnect_MonitorCrossConnectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MonitorCrossConnect_serviceDesc.Streams[0], "/crossconnect.MonitorCrossConnect/MonitorCrossConnects", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorCrossConnectMonitorCrossConnectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MonitorCrossConnect_MonitorCrossConnectsClient interface {
	Recv() (*CrossConnectEvent, error)
	grpc.ClientStream
}

type monitorCrossConnectMonitorCrossConnectsClient struct {
	grpc.ClientStream
}

func (x *monitorCrossConnectMonitorCrossConnectsClient) Recv() (*CrossConnectEvent, error) {
	m := new(CrossConnectEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitorCrossConnectServer is the server API for MonitorCrossConnect service.
type MonitorCrossConnectServer interface {
	MonitorCrossConnects(*empty.Empty, MonitorCrossConnect_MonitorCrossConnectsServer) error
}

// UnimplementedMonitorCrossConnectServer can be embedded to have forward compatible implementations.
type UnimplementedMonitorCrossConnectServer struct {
}

func (*UnimplementedMonitorCrossConnectServer) MonitorCrossConnects(req *empty.Empty, srv MonitorCrossConnect_MonitorCrossConnectsServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorCrossConnects not implemented")
}

func RegisterMonitorCrossConnectServer(s *grpc.Server, srv MonitorCrossConnectServer) {
	s.RegisterService(&_MonitorCrossConnect_serviceDesc, srv)
}

func _MonitorCrossConnect_MonitorCrossConnects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitorCrossConnectServer).MonitorCrossConnects(m, &monitorCrossConnectMonitorCrossConnectsServer{stream})
}

type MonitorCrossConnect_MonitorCrossConnectsServer interface {
	Send(*CrossConnectEvent) error
	grpc.ServerStream
}

type monitorCrossConnectMonitorCrossConnectsServer struct {
	grpc.ServerStream
}

func (x *monitorCrossConnectMonitorCrossConnectsServer) Send(m *CrossConnectEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _MonitorCrossConnect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crossconnect.MonitorCrossConnect",
	HandlerType: (*MonitorCrossConnectServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorCrossConnects",
			Handler:       _MonitorCrossConnect_MonitorCrossConnects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crossconnect.proto",
}
