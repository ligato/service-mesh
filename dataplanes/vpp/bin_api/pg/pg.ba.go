// Code generated by GoVPP binapi-generator. DO NOT EDIT.
// source: api/pg.api.json

/*
Package pg is a generated VPP binary API of the 'pg' VPP module.

It is generated from this file:
	pg.api.json

It contains these VPP binary API objects:
	6 messages
	3 services
*/
package pg

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Messages */

// PgCreateInterface represents the VPP binary API message 'pg_create_interface'.
// Generated from 'pg.api.json', line 4:
//
//            "pg_create_interface",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "interface_id"
//            ],
//            {
//                "crc": "0xef1dc980"
//            }
//
type PgCreateInterface struct {
	InterfaceID uint32
}

func (*PgCreateInterface) GetMessageName() string {
	return "pg_create_interface"
}
func (*PgCreateInterface) GetCrcString() string {
	return "ef1dc980"
}
func (*PgCreateInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewPgCreateInterface() api.Message {
	return &PgCreateInterface{}
}

// PgCreateInterfaceReply represents the VPP binary API message 'pg_create_interface_reply'.
// Generated from 'pg.api.json', line 26:
//
//            "pg_create_interface_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0xfda5941f"
//            }
//
type PgCreateInterfaceReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*PgCreateInterfaceReply) GetMessageName() string {
	return "pg_create_interface_reply"
}
func (*PgCreateInterfaceReply) GetCrcString() string {
	return "fda5941f"
}
func (*PgCreateInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewPgCreateInterfaceReply() api.Message {
	return &PgCreateInterfaceReply{}
}

// PgCapture represents the VPP binary API message 'pg_capture'.
// Generated from 'pg.api.json', line 48:
//
//            "pg_capture",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "interface_id"
//            ],
//            [
//                "u8",
//                "is_enabled"
//            ],
//            [
//                "u32",
//                "count"
//            ],
//            [
//                "u32",
//                "pcap_name_length"
//            ],
//            [
//                "u8",
//                "pcap_file_name",
//                0,
//                "pcap_name_length"
//            ],
//            {
//                "crc": "0x453da78d"
//            }
//
type PgCapture struct {
	InterfaceID    uint32
	IsEnabled      uint8
	Count          uint32
	PcapNameLength uint32 `struc:"sizeof=PcapFileName"`
	PcapFileName   []byte
}

func (*PgCapture) GetMessageName() string {
	return "pg_capture"
}
func (*PgCapture) GetCrcString() string {
	return "453da78d"
}
func (*PgCapture) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewPgCapture() api.Message {
	return &PgCapture{}
}

// PgCaptureReply represents the VPP binary API message 'pg_capture_reply'.
// Generated from 'pg.api.json', line 88:
//
//            "pg_capture_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type PgCaptureReply struct {
	Retval int32
}

func (*PgCaptureReply) GetMessageName() string {
	return "pg_capture_reply"
}
func (*PgCaptureReply) GetCrcString() string {
	return "e8d4e804"
}
func (*PgCaptureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewPgCaptureReply() api.Message {
	return &PgCaptureReply{}
}

// PgEnableDisable represents the VPP binary API message 'pg_enable_disable'.
// Generated from 'pg.api.json', line 106:
//
//            "pg_enable_disable",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_enabled"
//            ],
//            [
//                "u32",
//                "stream_name_length"
//            ],
//            [
//                "u8",
//                "stream_name",
//                0,
//                "stream_name_length"
//            ],
//            {
//                "crc": "0x0cb71d10"
//            }
//
type PgEnableDisable struct {
	IsEnabled        uint8
	StreamNameLength uint32 `struc:"sizeof=StreamName"`
	StreamName       []byte
}

func (*PgEnableDisable) GetMessageName() string {
	return "pg_enable_disable"
}
func (*PgEnableDisable) GetCrcString() string {
	return "0cb71d10"
}
func (*PgEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewPgEnableDisable() api.Message {
	return &PgEnableDisable{}
}

// PgEnableDisableReply represents the VPP binary API message 'pg_enable_disable_reply'.
// Generated from 'pg.api.json', line 138:
//
//            "pg_enable_disable_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type PgEnableDisableReply struct {
	Retval int32
}

func (*PgEnableDisableReply) GetMessageName() string {
	return "pg_enable_disable_reply"
}
func (*PgEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*PgEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewPgEnableDisableReply() api.Message {
	return &PgEnableDisableReply{}
}

/* Services */

type Services interface {
	PgCapture(*PgCapture) (*PgCaptureReply, error)
	PgCreateInterface(*PgCreateInterface) (*PgCreateInterfaceReply, error)
	PgEnableDisable(*PgEnableDisable) (*PgEnableDisableReply, error)
}

func init() {
	api.RegisterMessage((*PgCreateInterface)(nil), "pg.PgCreateInterface")
	api.RegisterMessage((*PgCreateInterfaceReply)(nil), "pg.PgCreateInterfaceReply")
	api.RegisterMessage((*PgCapture)(nil), "pg.PgCapture")
	api.RegisterMessage((*PgCaptureReply)(nil), "pg.PgCaptureReply")
	api.RegisterMessage((*PgEnableDisable)(nil), "pg.PgEnableDisable")
	api.RegisterMessage((*PgEnableDisableReply)(nil), "pg.PgEnableDisableReply")
}
