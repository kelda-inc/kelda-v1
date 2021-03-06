// Code generated by protoc-gen-go. DO NOT EDIT.
// source: _proto/kelda/messages/v0/messages.proto

package messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MessageMtype int32

const (
	Message_WARNING MessageMtype = 0
)

var MessageMtype_name = map[int32]string{
	0: "WARNING",
}

var MessageMtype_value = map[string]int32{
	"WARNING": 0,
}

func (x MessageMtype) String() string {
	return proto.EnumName(MessageMtype_name, int32(x))
}

func (MessageMtype) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c9834b77ae54d0d4, []int{0, 0}
}

// Message is a generic container for text messages to be sent
// from server to client.
type Message struct {
	Text                 string       `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Type                 MessageMtype `protobuf:"varint,2,opt,name=type,proto3,enum=kelda.messages.v0.MessageMtype" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9834b77ae54d0d4, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetType() MessageMtype {
	if m != nil {
		return m.Type
	}
	return Message_WARNING
}

func init() {
	proto.RegisterEnum("kelda.messages.v0.MessageMtype", MessageMtype_name, MessageMtype_value)
	proto.RegisterType((*Message)(nil), "kelda.messages.v0.Message")
}

func init() {
	proto.RegisterFile("_proto/kelda/messages/v0/messages.proto", fileDescriptor_c9834b77ae54d0d4)
}

var fileDescriptor_c9834b77ae54d0d4 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x8f, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0xd7, 0xcf, 0x4e, 0xcd, 0x49, 0x49, 0xd4, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x2d, 0xd6, 0x2f, 0x33, 0x80, 0xb3, 0xf5, 0xc0, 0x0a, 0x84, 0x04, 0xc1, 0x2a, 0xf4, 0xe0, 0xa2,
	0x65, 0x06, 0x4a, 0x99, 0x5c, 0xec, 0xbe, 0x10, 0xae, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x6a, 0x45,
	0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x64, 0xc2, 0xc5, 0x52, 0x52, 0x59,
	0x90, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x67, 0xa4, 0xa0, 0x87, 0x61, 0x80, 0x1e, 0x54, 0xb7,
	0x5e, 0x2e, 0x48, 0x5d, 0x10, 0x58, 0xb5, 0x92, 0x08, 0x17, 0x2b, 0x98, 0x2b, 0xc4, 0xcd, 0xc5,
	0x1e, 0xee, 0x18, 0xe4, 0xe7, 0xe9, 0xe7, 0x2e, 0xc0, 0xe0, 0xa4, 0x1f, 0xa5, 0x9b, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0x0b, 0x71, 0xac, 0x6e, 0x66, 0x5e, 0x32, 0xd4, 0xd9,
	0x05, 0xd9, 0xe9, 0xfa, 0x10, 0x6f, 0xc0, 0x4c, 0x4f, 0x62, 0x03, 0xf3, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xa9, 0xf4, 0x70, 0x50, 0xe0, 0x00, 0x00, 0x00,
}
