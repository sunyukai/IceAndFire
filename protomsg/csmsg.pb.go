// Code generated by protoc-gen-go.
// source: csmsg.proto
// DO NOT EDIT!

/*
Package protomsg is a generated protocol buffer package.

It is generated from these files:
	csmsg.proto

It has these top-level messages:
	Head
	Chat
	CSMessage
*/
package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MSG_ID int32

const (
	MSG_ID_MSG_ID_NONE MSG_ID = 0
	MSG_ID_MSG_ID_TEST MSG_ID = 1
)

var MSG_ID_name = map[int32]string{
	0: "MSG_ID_NONE",
	1: "MSG_ID_TEST",
}
var MSG_ID_value = map[string]int32{
	"MSG_ID_NONE": 0,
	"MSG_ID_TEST": 1,
}

func (x MSG_ID) String() string {
	return proto.EnumName(MSG_ID_name, int32(x))
}
func (MSG_ID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Head struct {
	Id MSG_ID `protobuf:"varint,1,opt,name=Id,enum=protomsg.MSG_ID" json:"Id,omitempty"`
}

func (m *Head) Reset()                    { *m = Head{} }
func (m *Head) String() string            { return proto.CompactTextString(m) }
func (*Head) ProtoMessage()               {}
func (*Head) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Head) GetId() MSG_ID {
	if m != nil {
		return m.Id
	}
	return MSG_ID_MSG_ID_NONE
}

type Chat struct {
	Send string `protobuf:"bytes,1,opt,name=Send" json:"Send,omitempty"`
}

func (m *Chat) Reset()                    { *m = Chat{} }
func (m *Chat) String() string            { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()               {}
func (*Chat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Chat) GetSend() string {
	if m != nil {
		return m.Send
	}
	return ""
}

type CSMessage struct {
	Head *Head `protobuf:"bytes,1,opt,name=Head" json:"Head,omitempty"`
	// Types that are valid to be assigned to Body:
	//	*CSMessage_Chat
	Body isCSMessage_Body `protobuf_oneof:"Body"`
}

func (m *CSMessage) Reset()                    { *m = CSMessage{} }
func (m *CSMessage) String() string            { return proto.CompactTextString(m) }
func (*CSMessage) ProtoMessage()               {}
func (*CSMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isCSMessage_Body interface {
	isCSMessage_Body()
}

type CSMessage_Chat struct {
	Chat *Chat `protobuf:"bytes,2,opt,name=Chat,oneof"`
}

func (*CSMessage_Chat) isCSMessage_Body() {}

func (m *CSMessage) GetBody() isCSMessage_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *CSMessage) GetHead() *Head {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *CSMessage) GetChat() *Chat {
	if x, ok := m.GetBody().(*CSMessage_Chat); ok {
		return x.Chat
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CSMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CSMessage_OneofMarshaler, _CSMessage_OneofUnmarshaler, _CSMessage_OneofSizer, []interface{}{
		(*CSMessage_Chat)(nil),
	}
}

func _CSMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CSMessage)
	// Body
	switch x := m.Body.(type) {
	case *CSMessage_Chat:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Chat); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CSMessage.Body has unexpected type %T", x)
	}
	return nil
}

func _CSMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CSMessage)
	switch tag {
	case 2: // Body.Chat
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Chat)
		err := b.DecodeMessage(msg)
		m.Body = &CSMessage_Chat{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CSMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CSMessage)
	// Body
	switch x := m.Body.(type) {
	case *CSMessage_Chat:
		s := proto.Size(x.Chat)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Head)(nil), "protomsg.Head")
	proto.RegisterType((*Chat)(nil), "protomsg.Chat")
	proto.RegisterType((*CSMessage)(nil), "protomsg.CSMessage")
	proto.RegisterEnum("protomsg.MSG_ID", MSG_ID_name, MSG_ID_value)
}

func init() { proto.RegisterFile("csmsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2e, 0xce, 0x2d,
	0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xb9, 0xc5, 0xe9, 0x4a, 0x1a,
	0x5c, 0x2c, 0x1e, 0xa9, 0x89, 0x29, 0x42, 0x0a, 0x5c, 0x4c, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x7c, 0x46, 0x02, 0x7a, 0x30, 0x69, 0x3d, 0xdf, 0x60, 0xf7, 0x78, 0x4f, 0x97, 0x20, 0x26,
	0xcf, 0x14, 0x25, 0x29, 0x2e, 0x16, 0xe7, 0x8c, 0xc4, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xe0, 0xd4,
	0x3c, 0x88, 0x5a, 0xce, 0x20, 0x30, 0x5b, 0x29, 0x96, 0x8b, 0xd3, 0x39, 0xd8, 0x37, 0xb5, 0xb8,
	0x38, 0x31, 0x3d, 0x55, 0x48, 0x09, 0x62, 0x24, 0x58, 0x01, 0xb7, 0x11, 0x1f, 0xc2, 0x30, 0x90,
	0x68, 0x10, 0xc4, 0x3a, 0x15, 0x88, 0x61, 0x12, 0x4c, 0xe8, 0x6a, 0x40, 0xa2, 0x1e, 0x0c, 0x41,
	0x60, 0x59, 0x27, 0x36, 0x2e, 0x16, 0xa7, 0xfc, 0x94, 0x4a, 0x2d, 0x2d, 0x2e, 0x36, 0x88, 0x43,
	0x84, 0xf8, 0xb9, 0xb8, 0x21, 0xac, 0x78, 0x3f, 0x7f, 0x3f, 0x57, 0x01, 0x06, 0x24, 0x81, 0x10,
	0xd7, 0xe0, 0x10, 0x01, 0xc6, 0x24, 0x36, 0xb0, 0x51, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x91, 0x76, 0xca, 0xe6, 0xf0, 0x00, 0x00, 0x00,
}