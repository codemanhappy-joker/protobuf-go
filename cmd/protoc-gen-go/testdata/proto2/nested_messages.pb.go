// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/nested_messages.proto

package proto2

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

type Layer1 struct {
	L2                   *Layer1_Layer2        `protobuf:"bytes,1,opt,name=l2" json:"l2,omitempty"`
	L3                   *Layer1_Layer2_Layer3 `protobuf:"bytes,2,opt,name=l3" json:"l3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Layer1) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_proto2_nested_messages_messageTypes[0].MessageOf(m)
}
func (m *Layer1) Reset()         { *m = Layer1{} }
func (m *Layer1) String() string { return proto.CompactTextString(m) }
func (*Layer1) ProtoMessage()    {}
func (*Layer1) Descriptor() ([]byte, []int) {
	return fileDescriptor_7417ee157699d191_gzipped, []int{0}
}

func (m *Layer1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer1.Unmarshal(m, b)
}
func (m *Layer1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer1.Marshal(b, m, deterministic)
}
func (m *Layer1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer1.Merge(m, src)
}
func (m *Layer1) XXX_Size() int {
	return xxx_messageInfo_Layer1.Size(m)
}
func (m *Layer1) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer1.DiscardUnknown(m)
}

var xxx_messageInfo_Layer1 proto.InternalMessageInfo

func (m *Layer1) GetL2() *Layer1_Layer2 {
	if m != nil {
		return m.L2
	}
	return nil
}

func (m *Layer1) GetL3() *Layer1_Layer2_Layer3 {
	if m != nil {
		return m.L3
	}
	return nil
}

type Layer1_Layer2 struct {
	L3                   *Layer1_Layer2_Layer3 `protobuf:"bytes,1,opt,name=l3" json:"l3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Layer1_Layer2) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_proto2_nested_messages_messageTypes[1].MessageOf(m)
}
func (m *Layer1_Layer2) Reset()         { *m = Layer1_Layer2{} }
func (m *Layer1_Layer2) String() string { return proto.CompactTextString(m) }
func (*Layer1_Layer2) ProtoMessage()    {}
func (*Layer1_Layer2) Descriptor() ([]byte, []int) {
	return fileDescriptor_7417ee157699d191_gzipped, []int{0, 0}
}

func (m *Layer1_Layer2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer1_Layer2.Unmarshal(m, b)
}
func (m *Layer1_Layer2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer1_Layer2.Marshal(b, m, deterministic)
}
func (m *Layer1_Layer2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer1_Layer2.Merge(m, src)
}
func (m *Layer1_Layer2) XXX_Size() int {
	return xxx_messageInfo_Layer1_Layer2.Size(m)
}
func (m *Layer1_Layer2) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer1_Layer2.DiscardUnknown(m)
}

var xxx_messageInfo_Layer1_Layer2 proto.InternalMessageInfo

func (m *Layer1_Layer2) GetL3() *Layer1_Layer2_Layer3 {
	if m != nil {
		return m.L3
	}
	return nil
}

type Layer1_Layer2_Layer3 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Layer1_Layer2_Layer3) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_proto2_nested_messages_messageTypes[2].MessageOf(m)
}
func (m *Layer1_Layer2_Layer3) Reset()         { *m = Layer1_Layer2_Layer3{} }
func (m *Layer1_Layer2_Layer3) String() string { return proto.CompactTextString(m) }
func (*Layer1_Layer2_Layer3) ProtoMessage()    {}
func (*Layer1_Layer2_Layer3) Descriptor() ([]byte, []int) {
	return fileDescriptor_7417ee157699d191_gzipped, []int{0, 0, 0}
}

func (m *Layer1_Layer2_Layer3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer1_Layer2_Layer3.Unmarshal(m, b)
}
func (m *Layer1_Layer2_Layer3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer1_Layer2_Layer3.Marshal(b, m, deterministic)
}
func (m *Layer1_Layer2_Layer3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer1_Layer2_Layer3.Merge(m, src)
}
func (m *Layer1_Layer2_Layer3) XXX_Size() int {
	return xxx_messageInfo_Layer1_Layer2_Layer3.Size(m)
}
func (m *Layer1_Layer2_Layer3) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer1_Layer2_Layer3.DiscardUnknown(m)
}

var xxx_messageInfo_Layer1_Layer2_Layer3 proto.InternalMessageInfo

func init() {
	proto.RegisterFile("proto2/nested_messages.proto", fileDescriptor_7417ee157699d191_gzipped)
	proto.RegisterType((*Layer1)(nil), "goproto.protoc.proto2.Layer1")
	proto.RegisterType((*Layer1_Layer2)(nil), "goproto.protoc.proto2.Layer1.Layer2")
	proto.RegisterType((*Layer1_Layer2_Layer3)(nil), "goproto.protoc.proto2.Layer1.Layer2.Layer3")
}

var fileDescriptor_7417ee157699d191 = []byte{
	// 327 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x22, 0xcc, 0x01, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x31,
	0x12, 0x34, 0x0a, 0x02, 0x6c, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x32, 0x52, 0x02, 0x6c, 0x32, 0x12, 0x3b, 0x0a, 0x02, 0x6c, 0x33, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x33, 0x52,
	0x02, 0x6c, 0x33, 0x1a, 0x4f, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x12, 0x3b, 0x0a,
	0x02, 0x6c, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x33, 0x52, 0x02, 0x6c, 0x33, 0x1a, 0x08, 0x0a, 0x06, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x33, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
}

var fileDescriptor_7417ee157699d191_gzipped = protoapi.CompressGZIP(fileDescriptor_7417ee157699d191)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_proto2_nested_messages protoreflect.FileDescriptor

var xxx_ProtoFile_proto2_nested_messages_messageTypes [3]protoimpl.MessageType
var xxx_ProtoFile_proto2_nested_messages_goTypes = []interface{}{
	(*Layer1)(nil),               // 0: goproto.protoc.proto2.Layer1
	(*Layer1_Layer2)(nil),        // 1: goproto.protoc.proto2.Layer1.Layer2
	(*Layer1_Layer2_Layer3)(nil), // 2: goproto.protoc.proto2.Layer1.Layer2.Layer3
}
var xxx_ProtoFile_proto2_nested_messages_depIdxs = []int32{
	1, // goproto.protoc.proto2.Layer1.l2:type_name -> goproto.protoc.proto2.Layer1.Layer2
	2, // goproto.protoc.proto2.Layer1.l3:type_name -> goproto.protoc.proto2.Layer1.Layer2.Layer3
	2, // goproto.protoc.proto2.Layer1.Layer2.l3:type_name -> goproto.protoc.proto2.Layer1.Layer2.Layer3
}

func init() {
	var messageTypes [3]protoreflect.MessageType
	ProtoFile_proto2_nested_messages = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_7417ee157699d191,
		GoTypes:            xxx_ProtoFile_proto2_nested_messages_goTypes,
		DependencyIndexes:  xxx_ProtoFile_proto2_nested_messages_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_proto2_nested_messages_goTypes[0:][:3]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_proto2_nested_messages_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_proto2_nested_messages_messageTypes[i].PBType = mt
	}
	xxx_ProtoFile_proto2_nested_messages_goTypes = nil
	xxx_ProtoFile_proto2_nested_messages_depIdxs = nil
}
