// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reflect/protoregistry/testprotos/test.proto

package testprotos

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Enum1 int32

const (
	Enum1_ONE Enum1 = 1
)

func (e Enum1) Type() protoreflect.EnumType {
	return xxx_ProtoFile_reflect_protoregistry_testprotos_test_enumTypes[0]
}
func (e Enum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var Enum1_name = map[int32]string{
	1: "ONE",
}

var Enum1_value = map[string]int32{
	"ONE": 1,
}

func (x Enum1) Enum() *Enum1 {
	p := new(Enum1)
	*p = x
	return p
}

func (x Enum1) String() string {
	return proto.EnumName(Enum1_name, int32(x))
}

func (x *Enum1) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Enum1_value, data, "Enum1")
	if err != nil {
		return err
	}
	*x = Enum1(value)
	return nil
}

func (Enum1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3628d63611f7063d_gzipped, []int{0}
}

type Enum2 int32

const (
	Enum2_UNO Enum2 = 1
)

func (e Enum2) Type() protoreflect.EnumType {
	return xxx_ProtoFile_reflect_protoregistry_testprotos_test_enumTypes[1]
}
func (e Enum2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var Enum2_name = map[int32]string{
	1: "UNO",
}

var Enum2_value = map[string]int32{
	"UNO": 1,
}

func (x Enum2) Enum() *Enum2 {
	p := new(Enum2)
	*p = x
	return p
}

func (x Enum2) String() string {
	return proto.EnumName(Enum2_name, int32(x))
}

func (x *Enum2) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Enum2_value, data, "Enum2")
	if err != nil {
		return err
	}
	*x = Enum2(value)
	return nil
}

func (Enum2) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3628d63611f7063d_gzipped, []int{1}
}

type Enum3 int32

const (
	Enum3_YI Enum3 = 1
)

func (e Enum3) Type() protoreflect.EnumType {
	return xxx_ProtoFile_reflect_protoregistry_testprotos_test_enumTypes[2]
}
func (e Enum3) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var Enum3_name = map[int32]string{
	1: "YI",
}

var Enum3_value = map[string]int32{
	"YI": 1,
}

func (x Enum3) Enum() *Enum3 {
	p := new(Enum3)
	*p = x
	return p
}

func (x Enum3) String() string {
	return proto.EnumName(Enum3_name, int32(x))
}

func (x *Enum3) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Enum3_value, data, "Enum3")
	if err != nil {
		return err
	}
	*x = Enum3(value)
	return nil
}

func (Enum3) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3628d63611f7063d_gzipped, []int{2}
}

type Message1 struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *Message1) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_reflect_protoregistry_testprotos_test_messageTypes[0].MessageOf(m)
}
func (m *Message1) Reset()         { *m = Message1{} }
func (m *Message1) String() string { return proto.CompactTextString(m) }
func (*Message1) ProtoMessage()    {}
func (*Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_3628d63611f7063d_gzipped, []int{0}
}

var extRange_Message1 = []proto.ExtensionRange{
	{Start: 10, End: 536870911},
}

func (*Message1) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Message1
}

func (m *Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1.Unmarshal(m, b)
}
func (m *Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1.Marshal(b, m, deterministic)
}
func (m *Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1.Merge(m, src)
}
func (m *Message1) XXX_Size() int {
	return xxx_messageInfo_Message1.Size(m)
}
func (m *Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1.DiscardUnknown(m)
}

var xxx_messageInfo_Message1 proto.InternalMessageInfo

type Message2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message2) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_reflect_protoregistry_testprotos_test_messageTypes[1].MessageOf(m)
}
func (m *Message2) Reset()         { *m = Message2{} }
func (m *Message2) String() string { return proto.CompactTextString(m) }
func (*Message2) ProtoMessage()    {}
func (*Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_3628d63611f7063d_gzipped, []int{1}
}

func (m *Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2.Unmarshal(m, b)
}
func (m *Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2.Marshal(b, m, deterministic)
}
func (m *Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2.Merge(m, src)
}
func (m *Message2) XXX_Size() int {
	return xxx_messageInfo_Message2.Size(m)
}
func (m *Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2.DiscardUnknown(m)
}

var xxx_messageInfo_Message2 proto.InternalMessageInfo

type Message3 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message3) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_reflect_protoregistry_testprotos_test_messageTypes[2].MessageOf(m)
}
func (m *Message3) Reset()         { *m = Message3{} }
func (m *Message3) String() string { return proto.CompactTextString(m) }
func (*Message3) ProtoMessage()    {}
func (*Message3) Descriptor() ([]byte, []int) {
	return fileDescriptor_3628d63611f7063d_gzipped, []int{2}
}

func (m *Message3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message3.Unmarshal(m, b)
}
func (m *Message3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message3.Marshal(b, m, deterministic)
}
func (m *Message3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message3.Merge(m, src)
}
func (m *Message3) XXX_Size() int {
	return xxx_messageInfo_Message3.Size(m)
}
func (m *Message3) XXX_DiscardUnknown() {
	xxx_messageInfo_Message3.DiscardUnknown(m)
}

var xxx_messageInfo_Message3 proto.InternalMessageInfo

type Message4 struct {
	BoolField            *bool    `protobuf:"varint,30,opt,name=bool_field,json=boolField" json:"bool_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message4) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_reflect_protoregistry_testprotos_test_messageTypes[3].MessageOf(m)
}
func (m *Message4) Reset()         { *m = Message4{} }
func (m *Message4) String() string { return proto.CompactTextString(m) }
func (*Message4) ProtoMessage()    {}
func (*Message4) Descriptor() ([]byte, []int) {
	return fileDescriptor_3628d63611f7063d_gzipped, []int{3}
}

func (m *Message4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message4.Unmarshal(m, b)
}
func (m *Message4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message4.Marshal(b, m, deterministic)
}
func (m *Message4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message4.Merge(m, src)
}
func (m *Message4) XXX_Size() int {
	return xxx_messageInfo_Message4.Size(m)
}
func (m *Message4) XXX_DiscardUnknown() {
	xxx_messageInfo_Message4.DiscardUnknown(m)
}

var xxx_messageInfo_Message4 proto.InternalMessageInfo

func (m *Message4) GetBoolField() bool {
	if m != nil && m.BoolField != nil {
		return *m.BoolField
	}
	return false
}

var E_StringField = &proto.ExtensionDesc{
	ExtendedType:  (*Message1)(nil),
	ExtensionType: (*string)(nil),
	Field:         11,
	Name:          "testprotos.string_field",
	Tag:           "bytes,11,opt,name=string_field",
	Filename:      "reflect/protoregistry/testprotos/test.proto",
}

var E_EnumField = &proto.ExtensionDesc{
	ExtendedType:  (*Message1)(nil),
	ExtensionType: (*Enum1)(nil),
	Field:         12,
	Name:          "testprotos.enum_field",
	Tag:           "varint,12,opt,name=enum_field,enum=testprotos.Enum1",
	Filename:      "reflect/protoregistry/testprotos/test.proto",
}

var E_MessageField = &proto.ExtensionDesc{
	ExtendedType:  (*Message1)(nil),
	ExtensionType: (*Message2)(nil),
	Field:         13,
	Name:          "testprotos.message_field",
	Tag:           "bytes,13,opt,name=message_field",
	Filename:      "reflect/protoregistry/testprotos/test.proto",
}

var E_Message4_MessageField = &proto.ExtensionDesc{
	ExtendedType:  (*Message1)(nil),
	ExtensionType: (*Message2)(nil),
	Field:         21,
	Name:          "testprotos.Message4.message_field",
	Tag:           "bytes,21,opt,name=message_field",
	Filename:      "reflect/protoregistry/testprotos/test.proto",
}

var E_Message4_EnumField = &proto.ExtensionDesc{
	ExtendedType:  (*Message1)(nil),
	ExtensionType: (*Enum1)(nil),
	Field:         22,
	Name:          "testprotos.Message4.enum_field",
	Tag:           "varint,22,opt,name=enum_field,enum=testprotos.Enum1",
	Filename:      "reflect/protoregistry/testprotos/test.proto",
}

var E_Message4_StringField = &proto.ExtensionDesc{
	ExtendedType:  (*Message1)(nil),
	ExtensionType: (*string)(nil),
	Field:         23,
	Name:          "testprotos.Message4.string_field",
	Tag:           "bytes,23,opt,name=string_field",
	Filename:      "reflect/protoregistry/testprotos/test.proto",
}

func init() {
	proto.RegisterFile("reflect/protoregistry/testprotos/test.proto", fileDescriptor_3628d63611f7063d_gzipped)
	proto.RegisterEnum("testprotos.Enum1", Enum1_name, Enum1_value)
	proto.RegisterEnum("testprotos.Enum2", Enum2_name, Enum2_value)
	proto.RegisterEnum("testprotos.Enum3", Enum3_name, Enum3_value)
	proto.RegisterType((*Message1)(nil), "testprotos.Message1")
	proto.RegisterType((*Message2)(nil), "testprotos.Message2")
	proto.RegisterType((*Message3)(nil), "testprotos.Message3")
	proto.RegisterType((*Message4)(nil), "testprotos.Message4")
	proto.RegisterExtension(E_StringField)
	proto.RegisterExtension(E_EnumField)
	proto.RegisterExtension(E_MessageField)
	proto.RegisterExtension(E_Message4_MessageField)
	proto.RegisterExtension(E_Message4_EnumField)
	proto.RegisterExtension(E_Message4_StringField)
}

var fileDescriptor_3628d63611f7063d = []byte{
	// 686 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x2b, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x14, 0x0a, 0x08, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x31, 0x2a, 0x08, 0x08, 0x0a, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x22,
	0x0a, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x22, 0x0a, 0x0a, 0x08, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x22, 0xfb, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x34, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x4f, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x32, 0x46, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x31, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x0a, 0x0c,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x31, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x2a, 0x10, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x07,
	0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x2a, 0x10, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x32,
	0x12, 0x07, 0x0a, 0x03, 0x55, 0x4e, 0x4f, 0x10, 0x01, 0x2a, 0x0f, 0x0a, 0x05, 0x45, 0x6e, 0x75,
	0x6d, 0x33, 0x12, 0x06, 0x0a, 0x02, 0x59, 0x49, 0x10, 0x01, 0x3a, 0x37, 0x0a, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x3a, 0x46, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31,
	0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x4f, 0x0a, 0x0d, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x31, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x52, 0x0c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x40, 0x5a, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
}

var fileDescriptor_3628d63611f7063d_gzipped = protoapi.CompressGZIP(fileDescriptor_3628d63611f7063d)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_reflect_protoregistry_testprotos_test protoreflect.FileDescriptor

var xxx_ProtoFile_reflect_protoregistry_testprotos_test_enumTypes [3]protoreflect.EnumType
var xxx_ProtoFile_reflect_protoregistry_testprotos_test_messageTypes [4]protoimpl.MessageType
var xxx_ProtoFile_reflect_protoregistry_testprotos_test_goTypes = []interface{}{
	(Enum1)(0),       // 0: testprotos.Enum1
	(Enum2)(0),       // 1: testprotos.Enum2
	(Enum3)(0),       // 2: testprotos.Enum3
	(*Message1)(nil), // 3: testprotos.Message1
	(*Message2)(nil), // 4: testprotos.Message2
	(*Message3)(nil), // 5: testprotos.Message3
	(*Message4)(nil), // 6: testprotos.Message4
}
var xxx_ProtoFile_reflect_protoregistry_testprotos_test_depIdxs = []int32{
	3, // testprotos.string_field:extendee -> testprotos.Message1
	3, // testprotos.enum_field:extendee -> testprotos.Message1
	3, // testprotos.message_field:extendee -> testprotos.Message1
	3, // testprotos.Message4.message_field:extendee -> testprotos.Message1
	3, // testprotos.Message4.enum_field:extendee -> testprotos.Message1
	3, // testprotos.Message4.string_field:extendee -> testprotos.Message1
	0, // testprotos.enum_field:type_name -> testprotos.Enum1
	4, // testprotos.message_field:type_name -> testprotos.Message2
	4, // testprotos.Message4.message_field:type_name -> testprotos.Message2
	0, // testprotos.Message4.enum_field:type_name -> testprotos.Enum1
}

func init() {
	var messageTypes [4]protoreflect.MessageType
	var extensionTypes [6]protoreflect.ExtensionType
	ProtoFile_reflect_protoregistry_testprotos_test = protoimpl.FileBuilder{
		RawDescriptor:        fileDescriptor_3628d63611f7063d,
		GoTypes:              xxx_ProtoFile_reflect_protoregistry_testprotos_test_goTypes,
		DependencyIndexes:    xxx_ProtoFile_reflect_protoregistry_testprotos_test_depIdxs,
		EnumOutputTypes:      xxx_ProtoFile_reflect_protoregistry_testprotos_test_enumTypes[:],
		MessageOutputTypes:   messageTypes[:],
		ExtensionOutputTypes: extensionTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_reflect_protoregistry_testprotos_test_goTypes[3:][:4]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_reflect_protoregistry_testprotos_test_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_reflect_protoregistry_testprotos_test_messageTypes[i].PBType = mt
	}
	E_StringField.Type = extensionTypes[0]
	E_EnumField.Type = extensionTypes[1]
	E_MessageField.Type = extensionTypes[2]
	E_Message4_MessageField.Type = extensionTypes[3]
	E_Message4_EnumField.Type = extensionTypes[4]
	E_Message4_StringField.Type = extensionTypes[5]
	xxx_ProtoFile_reflect_protoregistry_testprotos_test_goTypes = nil
	xxx_ProtoFile_reflect_protoregistry_testprotos_test_depIdxs = nil
}
