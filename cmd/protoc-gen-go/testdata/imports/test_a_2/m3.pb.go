// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_2/m3.proto

package test_a_2

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

type M3 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M3) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_imports_test_a_2_m3_messageTypes[0].MessageOf(m)
}
func (m *M3) Reset()         { *m = M3{} }
func (m *M3) String() string { return proto.CompactTextString(m) }
func (*M3) ProtoMessage()    {}
func (*M3) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9d8f834875c9c5_gzipped, []int{0}
}

func (m *M3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M3.Unmarshal(m, b)
}
func (m *M3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M3.Marshal(b, m, deterministic)
}
func (m *M3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M3.Merge(m, src)
}
func (m *M3) XXX_Size() int {
	return xxx_messageInfo_M3.Size(m)
}
func (m *M3) XXX_DiscardUnknown() {
	xxx_messageInfo_M3.DiscardUnknown(m)
}

var xxx_messageInfo_M3 proto.InternalMessageInfo

func init() {
	proto.RegisterFile("imports/test_a_2/m3.proto", fileDescriptor_ff9d8f834875c9c5_gzipped)
	proto.RegisterType((*M3)(nil), "test.a.M3")
}

var fileDescriptor_ff9d8f834875c9c5 = []byte{
	// 119 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x5f, 0x32, 0x2f, 0x6d, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x22, 0x04, 0x0a, 0x02, 0x4d, 0x33, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var fileDescriptor_ff9d8f834875c9c5_gzipped = protoapi.CompressGZIP(fileDescriptor_ff9d8f834875c9c5)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_imports_test_a_2_m3 protoreflect.FileDescriptor

var xxx_ProtoFile_imports_test_a_2_m3_messageTypes [1]protoimpl.MessageType
var xxx_ProtoFile_imports_test_a_2_m3_goTypes = []interface{}{
	(*M3)(nil), // 0: test.a.M3
}
var xxx_ProtoFile_imports_test_a_2_m3_depIdxs = []int32{}

func init() {
	var messageTypes [1]protoreflect.MessageType
	ProtoFile_imports_test_a_2_m3 = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_ff9d8f834875c9c5,
		GoTypes:            xxx_ProtoFile_imports_test_a_2_m3_goTypes,
		DependencyIndexes:  xxx_ProtoFile_imports_test_a_2_m3_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_imports_test_a_2_m3_goTypes[0:][:1]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_imports_test_a_2_m3_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_imports_test_a_2_m3_messageTypes[i].PBType = mt
	}
	xxx_ProtoFile_imports_test_a_2_m3_goTypes = nil
	xxx_ProtoFile_imports_test_a_2_m3_depIdxs = nil
}
