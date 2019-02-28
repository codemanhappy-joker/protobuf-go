// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test_public.proto

package test

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

type PublicImportMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicImportMessage) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_test_public_messageTypes[0].MessageOf(m)
}
func (m *PublicImportMessage) Reset()         { *m = PublicImportMessage{} }
func (m *PublicImportMessage) String() string { return proto.CompactTextString(m) }
func (*PublicImportMessage) ProtoMessage()    {}
func (*PublicImportMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_36dd44afd5b47374_gzipped, []int{0}
}

func (m *PublicImportMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicImportMessage.Unmarshal(m, b)
}
func (m *PublicImportMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicImportMessage.Marshal(b, m, deterministic)
}
func (m *PublicImportMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicImportMessage.Merge(m, src)
}
func (m *PublicImportMessage) XXX_Size() int {
	return xxx_messageInfo_PublicImportMessage.Size(m)
}
func (m *PublicImportMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicImportMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PublicImportMessage proto.InternalMessageInfo

func init() {
	proto.RegisterFile("test_public.proto", fileDescriptor_36dd44afd5b47374_gzipped)
	proto.RegisterType((*PublicImportMessage)(nil), "goproto.proto.test.PublicImportMessage")
}

var fileDescriptor_36dd44afd5b47374 = []byte{
	// 120 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x38,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74,
}

var fileDescriptor_36dd44afd5b47374_gzipped = protoapi.CompressGZIP(fileDescriptor_36dd44afd5b47374)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_test_public protoreflect.FileDescriptor

var xxx_ProtoFile_test_public_messageTypes [1]protoimpl.MessageType
var xxx_ProtoFile_test_public_goTypes = []interface{}{
	(*PublicImportMessage)(nil), // 0: goproto.proto.test.PublicImportMessage
}
var xxx_ProtoFile_test_public_depIdxs = []int32{}

func init() {
	var messageTypes [1]protoreflect.MessageType
	ProtoFile_test_public = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_36dd44afd5b47374,
		GoTypes:            xxx_ProtoFile_test_public_goTypes,
		DependencyIndexes:  xxx_ProtoFile_test_public_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_test_public_goTypes[0:][:1]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_test_public_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_test_public_messageTypes[i].PBType = mt
	}
	xxx_ProtoFile_test_public_goTypes = nil
	xxx_ProtoFile_test_public_depIdxs = nil
}
