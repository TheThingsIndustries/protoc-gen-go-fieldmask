// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: messages.proto

package test

import (
	_ "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageWithoutFieldSetter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageWithoutFieldSetter) Reset() {
	*x = MessageWithoutFieldSetter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithoutFieldSetter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithoutFieldSetter) ProtoMessage() {}

func (x *MessageWithoutFieldSetter) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithoutFieldSetter.ProtoReflect.Descriptor instead.
func (*MessageWithoutFieldSetter) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *MessageWithoutFieldSetter) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

type SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar string `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *SubMessage) Reset() {
	*x = SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubMessage) ProtoMessage() {}

func (x *SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubMessage.ProtoReflect.Descriptor instead.
func (*SubMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *SubMessage) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

func (x *SubMessage) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

type MessageWithSubMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *SubMessage `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B *SubMessage `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *MessageWithSubMessages) Reset() {
	*x = MessageWithSubMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithSubMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithSubMessages) ProtoMessage() {}

func (x *MessageWithSubMessages) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithSubMessages.ProtoReflect.Descriptor instead.
func (*MessageWithSubMessages) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *MessageWithSubMessages) GetA() *SubMessage {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *MessageWithSubMessages) GetB() *SubMessage {
	if x != nil {
		return x.B
	}
	return nil
}

type MessageWithOneofSubMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sub:
	//
	//	*MessageWithOneofSubMessages_A
	//	*MessageWithOneofSubMessages_B
	Sub isMessageWithOneofSubMessages_Sub `protobuf_oneof:"sub"`
}

func (x *MessageWithOneofSubMessages) Reset() {
	*x = MessageWithOneofSubMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithOneofSubMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithOneofSubMessages) ProtoMessage() {}

func (x *MessageWithOneofSubMessages) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithOneofSubMessages.ProtoReflect.Descriptor instead.
func (*MessageWithOneofSubMessages) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{4}
}

func (m *MessageWithOneofSubMessages) GetSub() isMessageWithOneofSubMessages_Sub {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (x *MessageWithOneofSubMessages) GetA() *SubMessage {
	if x, ok := x.GetSub().(*MessageWithOneofSubMessages_A); ok {
		return x.A
	}
	return nil
}

func (x *MessageWithOneofSubMessages) GetB() *SubMessage {
	if x, ok := x.GetSub().(*MessageWithOneofSubMessages_B); ok {
		return x.B
	}
	return nil
}

type isMessageWithOneofSubMessages_Sub interface {
	isMessageWithOneofSubMessages_Sub()
}

type MessageWithOneofSubMessages_A struct {
	A *SubMessage `protobuf:"bytes,1,opt,name=a,proto3,oneof"`
}

type MessageWithOneofSubMessages_B struct {
	B *SubMessage `protobuf:"bytes,2,opt,name=b,proto3,oneof"`
}

func (*MessageWithOneofSubMessages_A) isMessageWithOneofSubMessages_Sub() {}

func (*MessageWithOneofSubMessages_B) isMessageWithOneofSubMessages_Sub() {}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a,
	0x19, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x3a, 0x06, 0xfa, 0xaa, 0x19, 0x02, 0x18, 0x00, 0x22, 0x0e, 0x0a, 0x0c,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x0a,
	0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x61, 0x72, 0x22, 0x80,
	0x01, 0x0a, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x53, 0x75,
	0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x01, 0x61, 0x12, 0x32, 0x0a,
	0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x01,
	0x62, 0x22, 0x90, 0x01, 0x0a, 0x1b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x34, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74,
	0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x01, 0x61, 0x12, 0x34, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75,
	0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x01, 0x62, 0x42, 0x05, 0x0a,
	0x03, 0x73, 0x75, 0x62, 0x42, 0x47, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75,
	0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0xfa, 0xaa, 0x19, 0x06, 0x08, 0x01, 0x10, 0x01, 0x18, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_messages_proto_goTypes = []interface{}{
	(*MessageWithoutFieldSetter)(nil),   // 0: thethings.fieldmask.test.MessageWithoutFieldSetter
	(*EmptyMessage)(nil),                // 1: thethings.fieldmask.test.EmptyMessage
	(*SubMessage)(nil),                  // 2: thethings.fieldmask.test.SubMessage
	(*MessageWithSubMessages)(nil),      // 3: thethings.fieldmask.test.MessageWithSubMessages
	(*MessageWithOneofSubMessages)(nil), // 4: thethings.fieldmask.test.MessageWithOneofSubMessages
}
var file_messages_proto_depIdxs = []int32{
	2, // 0: thethings.fieldmask.test.MessageWithSubMessages.a:type_name -> thethings.fieldmask.test.SubMessage
	2, // 1: thethings.fieldmask.test.MessageWithSubMessages.b:type_name -> thethings.fieldmask.test.SubMessage
	2, // 2: thethings.fieldmask.test.MessageWithOneofSubMessages.a:type_name -> thethings.fieldmask.test.SubMessage
	2, // 3: thethings.fieldmask.test.MessageWithOneofSubMessages.b:type_name -> thethings.fieldmask.test.SubMessage
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithoutFieldSetter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithSubMessages); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithOneofSubMessages); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_messages_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*MessageWithOneofSubMessages_A)(nil),
		(*MessageWithOneofSubMessages_B)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
