// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: messages.proto

package test

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/annotations"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MessageWithoutFieldSetter struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageWithoutFieldSetter) Reset()         { *m = MessageWithoutFieldSetter{} }
func (m *MessageWithoutFieldSetter) String() string { return proto.CompactTextString(m) }
func (*MessageWithoutFieldSetter) ProtoMessage()    {}
func (*MessageWithoutFieldSetter) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}
func (m *MessageWithoutFieldSetter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithoutFieldSetter.Unmarshal(m, b)
}
func (m *MessageWithoutFieldSetter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithoutFieldSetter.Marshal(b, m, deterministic)
}
func (m *MessageWithoutFieldSetter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithoutFieldSetter.Merge(m, src)
}
func (m *MessageWithoutFieldSetter) XXX_Size() int {
	return xxx_messageInfo_MessageWithoutFieldSetter.Size(m)
}
func (m *MessageWithoutFieldSetter) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithoutFieldSetter.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithoutFieldSetter proto.InternalMessageInfo

func (m *MessageWithoutFieldSetter) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SubMessage struct {
	Foo                  string   `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar                  string   `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubMessage) Reset()         { *m = SubMessage{} }
func (m *SubMessage) String() string { return proto.CompactTextString(m) }
func (*SubMessage) ProtoMessage()    {}
func (*SubMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}
func (m *SubMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubMessage.Unmarshal(m, b)
}
func (m *SubMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubMessage.Marshal(b, m, deterministic)
}
func (m *SubMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubMessage.Merge(m, src)
}
func (m *SubMessage) XXX_Size() int {
	return xxx_messageInfo_SubMessage.Size(m)
}
func (m *SubMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubMessage proto.InternalMessageInfo

func (m *SubMessage) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

func (m *SubMessage) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

type MessageWithSubMessages struct {
	A                    *SubMessage `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    *SubMessage `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MessageWithSubMessages) Reset()         { *m = MessageWithSubMessages{} }
func (m *MessageWithSubMessages) String() string { return proto.CompactTextString(m) }
func (*MessageWithSubMessages) ProtoMessage()    {}
func (*MessageWithSubMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}
func (m *MessageWithSubMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithSubMessages.Unmarshal(m, b)
}
func (m *MessageWithSubMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithSubMessages.Marshal(b, m, deterministic)
}
func (m *MessageWithSubMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithSubMessages.Merge(m, src)
}
func (m *MessageWithSubMessages) XXX_Size() int {
	return xxx_messageInfo_MessageWithSubMessages.Size(m)
}
func (m *MessageWithSubMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithSubMessages.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithSubMessages proto.InternalMessageInfo

func (m *MessageWithSubMessages) GetA() *SubMessage {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *MessageWithSubMessages) GetB() *SubMessage {
	if m != nil {
		return m.B
	}
	return nil
}

type MessageWithOneofSubMessages struct {
	// Types that are valid to be assigned to Sub:
	//	*MessageWithOneofSubMessages_A
	//	*MessageWithOneofSubMessages_B
	Sub                  isMessageWithOneofSubMessages_Sub `protobuf_oneof:"sub"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *MessageWithOneofSubMessages) Reset()         { *m = MessageWithOneofSubMessages{} }
func (m *MessageWithOneofSubMessages) String() string { return proto.CompactTextString(m) }
func (*MessageWithOneofSubMessages) ProtoMessage()    {}
func (*MessageWithOneofSubMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}
func (m *MessageWithOneofSubMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithOneofSubMessages.Unmarshal(m, b)
}
func (m *MessageWithOneofSubMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithOneofSubMessages.Marshal(b, m, deterministic)
}
func (m *MessageWithOneofSubMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithOneofSubMessages.Merge(m, src)
}
func (m *MessageWithOneofSubMessages) XXX_Size() int {
	return xxx_messageInfo_MessageWithOneofSubMessages.Size(m)
}
func (m *MessageWithOneofSubMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithOneofSubMessages.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithOneofSubMessages proto.InternalMessageInfo

type isMessageWithOneofSubMessages_Sub interface {
	isMessageWithOneofSubMessages_Sub()
}

type MessageWithOneofSubMessages_A struct {
	A *SubMessage `protobuf:"bytes,1,opt,name=a,proto3,oneof" json:"a,omitempty"`
}
type MessageWithOneofSubMessages_B struct {
	B *SubMessage `protobuf:"bytes,2,opt,name=b,proto3,oneof" json:"b,omitempty"`
}

func (*MessageWithOneofSubMessages_A) isMessageWithOneofSubMessages_Sub() {}
func (*MessageWithOneofSubMessages_B) isMessageWithOneofSubMessages_Sub() {}

func (m *MessageWithOneofSubMessages) GetSub() isMessageWithOneofSubMessages_Sub {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *MessageWithOneofSubMessages) GetA() *SubMessage {
	if x, ok := m.GetSub().(*MessageWithOneofSubMessages_A); ok {
		return x.A
	}
	return nil
}

func (m *MessageWithOneofSubMessages) GetB() *SubMessage {
	if x, ok := m.GetSub().(*MessageWithOneofSubMessages_B); ok {
		return x.B
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MessageWithOneofSubMessages) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MessageWithOneofSubMessages_A)(nil),
		(*MessageWithOneofSubMessages_B)(nil),
	}
}

func init() {
	proto.RegisterType((*MessageWithoutFieldSetter)(nil), "thethings.fieldmask.test.MessageWithoutFieldSetter")
	proto.RegisterType((*SubMessage)(nil), "thethings.fieldmask.test.SubMessage")
	proto.RegisterType((*MessageWithSubMessages)(nil), "thethings.fieldmask.test.MessageWithSubMessages")
	proto.RegisterType((*MessageWithOneofSubMessages)(nil), "thethings.fieldmask.test.MessageWithOneofSubMessages")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x97, 0x4d, 0xa7, 0x46, 0x10, 0xed, 0x41, 0x52, 0xbd, 0x48, 0xf1, 0xe0, 0xa5, 0xa9,
	0xcc, 0x9d, 0x14, 0x2f, 0x03, 0x45, 0x0f, 0x22, 0x6c, 0x03, 0xc1, 0x5b, 0xb2, 0xbd, 0xb6, 0x41,
	0x9b, 0x48, 0xdf, 0xeb, 0xdd, 0x8f, 0xe0, 0xe7, 0xd8, 0x47, 0xdc, 0x49, 0xd2, 0x55, 0xd7, 0x8b,
	0xe0, 0x6e, 0xc9, 0xcb, 0xff, 0xf7, 0xcb, 0x3f, 0x84, 0x1f, 0x14, 0x80, 0xa8, 0x32, 0x40, 0xf9,
	0x51, 0x3a, 0x72, 0x81, 0xa0, 0x1c, 0x28, 0x37, 0x36, 0x43, 0x99, 0x1a, 0x78, 0x9f, 0x17, 0x0a,
	0xdf, 0x24, 0x01, 0xd2, 0xc9, 0x91, 0xb2, 0xd6, 0x91, 0x22, 0xe3, 0x6c, 0x13, 0x8e, 0x6e, 0x79,
	0xf8, 0xb4, 0xc2, 0x5f, 0x0c, 0xe5, 0xae, 0xa2, 0x7b, 0x8f, 0x4c, 0x80, 0x08, 0xca, 0x40, 0xf0,
	0x9d, 0xc6, 0x2d, 0xd8, 0x19, 0xbb, 0xd8, 0x1b, 0xff, 0x6c, 0xaf, 0xfb, 0xcb, 0x45, 0xd8, 0x15,
	0x9d, 0xe8, 0x92, 0xf3, 0x49, 0xa5, 0x1b, 0x43, 0x70, 0xc8, 0x7b, 0xa9, 0x73, 0x4d, 0xd6, 0x2f,
	0xfd, 0x44, 0xab, 0x52, 0x74, 0x57, 0x13, 0xad, 0xca, 0xe8, 0x93, 0xf1, 0xe3, 0xd6, 0x8d, 0x6b,
	0x1a, 0x83, 0x01, 0x67, 0xaa, 0x86, 0xf7, 0x07, 0xe7, 0xf2, 0xaf, 0x47, 0xc8, 0x35, 0x31, 0x66,
	0xca, 0x33, 0xba, 0xd6, 0xff, 0x9b, 0xd1, 0xd1, 0x17, 0xe3, 0xa7, 0xad, 0x0a, 0xcf, 0x16, 0x5c,
	0xda, 0xee, 0x31, 0xdc, 0xb0, 0xc7, 0x43, 0xc7, 0x37, 0x19, 0x6e, 0xd8, 0xc4, 0x53, 0x7a, 0xb4,
	0xcd, 0x7b, 0x58, 0xe9, 0xd1, 0xdd, 0x72, 0x11, 0x6e, 0xed, 0x32, 0xc1, 0x5e, 0x6f, 0x32, 0x43,
	0x79, 0xa5, 0xe5, 0xcc, 0x15, 0xc9, 0x34, 0x87, 0x69, 0x6d, 0x79, 0xb4, 0xf3, 0x0a, 0xa9, 0x34,
	0x80, 0x49, 0xfd, 0x6d, 0xb3, 0x38, 0x03, 0x1b, 0x67, 0x2e, 0xfe, 0xb5, 0x27, 0xde, 0xae, 0xfb,
	0xf5, 0xe9, 0xd5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x28, 0x62, 0x95, 0x13, 0x02, 0x00,
	0x00,
}
