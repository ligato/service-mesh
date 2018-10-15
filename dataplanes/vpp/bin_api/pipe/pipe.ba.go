// Code generated by GoVPP binapi-generator. DO NOT EDIT.
// source: api/pipe.api.json

/*
Package pipe is a generated VPP binary API of the 'pipe' VPP module.

It is generated from this file:
	pipe.api.json

It contains these VPP binary API objects:
	6 messages
	3 services
*/
package pipe

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Messages */

// PipeCreate represents the VPP binary API message 'pipe_create'.
// Generated from 'pipe.api.json', line 4:
//
//            "pipe_create",
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
//                "is_specified"
//            ],
//            [
//                "u32",
//                "user_instance"
//            ],
//            {
//                "crc": "0xdb413409"
//            }
//
type PipeCreate struct {
	IsSpecified  uint8
	UserInstance uint32
}

func (*PipeCreate) GetMessageName() string {
	return "pipe_create"
}
func (*PipeCreate) GetCrcString() string {
	return "db413409"
}
func (*PipeCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewPipeCreate() api.Message {
	return &PipeCreate{}
}

// PipeCreateReply represents the VPP binary API message 'pipe_create_reply'.
// Generated from 'pipe.api.json', line 30:
//
//            "pipe_create_reply",
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
//            [
//                "u32",
//                "pipe_sw_if_index",
//                2
//            ],
//            {
//                "crc": "0x9f0eef14"
//            }
//
type PipeCreateReply struct {
	Retval        int32
	SwIfIndex     uint32
	PipeSwIfIndex []uint32 `struc:"[2]uint32"`
}

func (*PipeCreateReply) GetMessageName() string {
	return "pipe_create_reply"
}
func (*PipeCreateReply) GetCrcString() string {
	return "9f0eef14"
}
func (*PipeCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewPipeCreateReply() api.Message {
	return &PipeCreateReply{}
}

// PipeDelete represents the VPP binary API message 'pipe_delete'.
// Generated from 'pipe.api.json', line 57:
//
//            "pipe_delete",
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
//                "sw_if_index"
//            ],
//            {
//                "crc": "0x529cb13f"
//            }
//
type PipeDelete struct {
	SwIfIndex uint32
}

func (*PipeDelete) GetMessageName() string {
	return "pipe_delete"
}
func (*PipeDelete) GetCrcString() string {
	return "529cb13f"
}
func (*PipeDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewPipeDelete() api.Message {
	return &PipeDelete{}
}

// PipeDeleteReply represents the VPP binary API message 'pipe_delete_reply'.
// Generated from 'pipe.api.json', line 79:
//
//            "pipe_delete_reply",
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
type PipeDeleteReply struct {
	Retval int32
}

func (*PipeDeleteReply) GetMessageName() string {
	return "pipe_delete_reply"
}
func (*PipeDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*PipeDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewPipeDeleteReply() api.Message {
	return &PipeDeleteReply{}
}

// PipeDump represents the VPP binary API message 'pipe_dump'.
// Generated from 'pipe.api.json', line 97:
//
//            "pipe_dump",
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
//            {
//                "crc": "0x51077d14"
//            }
//
type PipeDump struct{}

func (*PipeDump) GetMessageName() string {
	return "pipe_dump"
}
func (*PipeDump) GetCrcString() string {
	return "51077d14"
}
func (*PipeDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewPipeDump() api.Message {
	return &PipeDump{}
}

// PipeDetails represents the VPP binary API message 'pipe_details'.
// Generated from 'pipe.api.json', line 115:
//
//            "pipe_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u32",
//                "pipe_sw_if_index",
//                2
//            ],
//            [
//                "u32",
//                "instance"
//            ],
//            {
//                "crc": "0x91286b09"
//            }
//
type PipeDetails struct {
	SwIfIndex     uint32
	PipeSwIfIndex []uint32 `struc:"[2]uint32"`
	Instance      uint32
}

func (*PipeDetails) GetMessageName() string {
	return "pipe_details"
}
func (*PipeDetails) GetCrcString() string {
	return "91286b09"
}
func (*PipeDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewPipeDetails() api.Message {
	return &PipeDetails{}
}

/* Services */

type Services interface {
	DumpPipe(*PipeDump) (*PipeDetails, error)
	PipeCreate(*PipeCreate) (*PipeCreateReply, error)
	PipeDelete(*PipeDelete) (*PipeDeleteReply, error)
}

func init() {
	api.RegisterMessage((*PipeCreate)(nil), "pipe.PipeCreate")
	api.RegisterMessage((*PipeCreateReply)(nil), "pipe.PipeCreateReply")
	api.RegisterMessage((*PipeDelete)(nil), "pipe.PipeDelete")
	api.RegisterMessage((*PipeDeleteReply)(nil), "pipe.PipeDeleteReply")
	api.RegisterMessage((*PipeDump)(nil), "pipe.PipeDump")
	api.RegisterMessage((*PipeDetails)(nil), "pipe.PipeDetails")
}
