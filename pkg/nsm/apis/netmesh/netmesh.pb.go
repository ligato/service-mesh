// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netmesh.proto

package netmesh

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/ligato/networkservicemesh/pkg/nsm/apis/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NetworkServiceEndpoint struct {
	NetworkServiceName   string                   `protobuf:"bytes,1,opt,name=network_service_name,json=networkServiceName" json:"network_service_name,omitempty"`
	NetworkServiceHost   string                   `protobuf:"bytes,2,opt,name=network_service_host,json=networkServiceHost" json:"network_service_host,omitempty"`
	NseProviderName      string                   `protobuf:"bytes,3,opt,name=nse_provider_name,json=nseProviderName" json:"nse_provider_name,omitempty"`
	SocketLocation       string                   `protobuf:"bytes,4,opt,name=socket_location,json=socketLocation" json:"socket_location,omitempty"`
	LocalMechanisms      []*common.LocalMechanism `protobuf:"bytes,5,rep,name=local_mechanisms,json=localMechanisms" json:"local_mechanisms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NetworkServiceEndpoint) Reset()         { *m = NetworkServiceEndpoint{} }
func (m *NetworkServiceEndpoint) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceEndpoint) ProtoMessage()    {}
func (*NetworkServiceEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_215a6331178c758f, []int{0}
}
func (m *NetworkServiceEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceEndpoint.Unmarshal(m, b)
}
func (m *NetworkServiceEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceEndpoint.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceEndpoint.Merge(dst, src)
}
func (m *NetworkServiceEndpoint) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceEndpoint.Size(m)
}
func (m *NetworkServiceEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceEndpoint proto.InternalMessageInfo

func (m *NetworkServiceEndpoint) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetNetworkServiceHost() string {
	if m != nil {
		return m.NetworkServiceHost
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetNseProviderName() string {
	if m != nil {
		return m.NseProviderName
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetSocketLocation() string {
	if m != nil {
		return m.SocketLocation
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetLocalMechanisms() []*common.LocalMechanism {
	if m != nil {
		return m.LocalMechanisms
	}
	return nil
}

type NetworkService struct {
	NetworkServiceName   string   `protobuf:"bytes,1,opt,name=network_service_name,json=networkServiceName" json:"network_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkService) Reset()         { *m = NetworkService{} }
func (m *NetworkService) String() string { return proto.CompactTextString(m) }
func (*NetworkService) ProtoMessage()    {}
func (*NetworkService) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_215a6331178c758f, []int{1}
}
func (m *NetworkService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkService.Unmarshal(m, b)
}
func (m *NetworkService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkService.Marshal(b, m, deterministic)
}
func (dst *NetworkService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkService.Merge(dst, src)
}
func (m *NetworkService) XXX_Size() int {
	return xxx_messageInfo_NetworkService.Size(m)
}
func (m *NetworkService) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkService.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkService proto.InternalMessageInfo

func (m *NetworkService) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkServiceEndpoint)(nil), "netmesh.NetworkServiceEndpoint")
	proto.RegisterType((*NetworkService)(nil), "netmesh.NetworkService")
}

func init() { proto.RegisterFile("netmesh.proto", fileDescriptor_netmesh_215a6331178c758f) }

var fileDescriptor_netmesh_215a6331178c758f = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x69, 0xeb, 0x1f, 0x5c, 0xb1, 0xd1, 0x20, 0x25, 0x78, 0x2a, 0xbd, 0x58, 0x3c, 0x64,
	0x45, 0x9f, 0x40, 0xa1, 0xe0, 0x41, 0x8b, 0xd4, 0x07, 0x08, 0xdb, 0xed, 0x90, 0x2c, 0xc9, 0xce,
	0x84, 0xcc, 0x58, 0x5f, 0xc4, 0x07, 0x96, 0x24, 0xeb, 0x21, 0xd0, 0x4b, 0x4f, 0xcb, 0xcc, 0xfe,
	0x7e, 0xdf, 0xb2, 0x9f, 0xba, 0x42, 0x10, 0x0f, 0x5c, 0xa4, 0x75, 0x43, 0x42, 0xf1, 0x79, 0x18,
	0xef, 0x56, 0xb9, 0x93, 0xe2, 0x7b, 0x9b, 0x5a, 0xf2, 0xba, 0x72, 0xb9, 0x11, 0xd2, 0x08, 0xf2,
	0x43, 0x4d, 0xc9, 0xd0, 0xec, 0x9d, 0x85, 0x96, 0xd2, 0x75, 0x99, 0x6b, 0x64, 0xaf, 0x4d, 0xed,
	0x58, 0x5b, 0xf2, 0x9e, 0x30, 0x1c, 0x7d, 0xde, 0xe2, 0x77, 0xac, 0x66, 0xeb, 0xde, 0xfb, 0xea,
	0xbd, 0x15, 0xee, 0x6a, 0x72, 0x28, 0xf1, 0xa3, 0xba, 0x0d, 0x89, 0x59, 0x88, 0xcc, 0xd0, 0x78,
	0x48, 0x46, 0xf3, 0xd1, 0xf2, 0x62, 0x13, 0xe3, 0xc0, 0x5a, 0x1b, 0x0f, 0x87, 0x8c, 0x82, 0x58,
	0x92, 0xf1, 0x21, 0xe3, 0x8d, 0x58, 0xe2, 0x07, 0x75, 0x83, 0x0c, 0x59, 0xdd, 0xd0, 0xde, 0xed,
	0xa0, 0xe9, 0x1f, 0x98, 0x74, 0x78, 0x84, 0x0c, 0x9f, 0x61, 0xdf, 0xa5, 0xdf, 0xab, 0x88, 0xc9,
	0x96, 0x20, 0x59, 0x45, 0xd6, 0x88, 0x23, 0x4c, 0x4e, 0x3a, 0x72, 0xda, 0xaf, 0xdf, 0xc3, 0x36,
	0x7e, 0x51, 0xd7, 0x2d, 0x51, 0x65, 0x1e, 0x6c, 0x61, 0xd0, 0xb1, 0xe7, 0xe4, 0x74, 0x3e, 0x59,
	0x5e, 0x3e, 0xcd, 0xd2, 0xf0, 0xf9, 0x96, 0xad, 0x3e, 0xfe, 0xaf, 0x37, 0x51, 0x35, 0x98, 0x79,
	0xf1, 0xaa, 0xa6, 0xc3, 0x56, 0x8e, 0x6f, 0x63, 0x7b, 0xd6, 0x35, 0xfc, 0xfc, 0x17, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0xb0, 0x4e, 0xed, 0xc2, 0x01, 0x00, 0x00,
}
