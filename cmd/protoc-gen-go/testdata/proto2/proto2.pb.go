// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/proto2.proto

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

type Message struct {
	I32                  *int32   `protobuf:"varint,1,opt,name=i32" json:"i32,omitempty"`
	M                    *Message `protobuf:"bytes,2,opt,name=m" json:"m,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_proto2_proto2_messageTypes[0].MessageOf(m)
}
func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d756bbe8817c03c1_gzipped, []int{0}
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

func (m *Message) GetI32() int32 {
	if m != nil && m.I32 != nil {
		return *m.I32
	}
	return 0
}

func (m *Message) GetM() *Message {
	if m != nil {
		return m.M
	}
	return nil
}

func init() {
	proto.RegisterFile("proto2/proto2.proto", fileDescriptor_d756bbe8817c03c1_gzipped)
	proto.RegisterType((*Message)(nil), "goproto.protoc.proto2.Message")
}

var fileDescriptor_d756bbe8817c03c1 = []byte{
	// 186 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x22, 0x49, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x33, 0x32, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x33, 0x32, 0x12, 0x2c, 0x0a, 0x01, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x01, 0x6d, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
}

var fileDescriptor_d756bbe8817c03c1_gzipped = protoapi.CompressGZIP(fileDescriptor_d756bbe8817c03c1)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_proto2_proto2 protoreflect.FileDescriptor

var xxx_ProtoFile_proto2_proto2_messageTypes [1]protoimpl.MessageType
var xxx_ProtoFile_proto2_proto2_goTypes = []interface{}{
	(*Message)(nil), // 0: goproto.protoc.proto2.Message
}
var xxx_ProtoFile_proto2_proto2_depIdxs = []int32{
	0, // goproto.protoc.proto2.Message.m:type_name -> goproto.protoc.proto2.Message
}

func init() {
	var messageTypes [1]protoreflect.MessageType
	ProtoFile_proto2_proto2 = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_d756bbe8817c03c1,
		GoTypes:            xxx_ProtoFile_proto2_proto2_goTypes,
		DependencyIndexes:  xxx_ProtoFile_proto2_proto2_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_proto2_proto2_goTypes[0:][:1]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_proto2_proto2_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_proto2_proto2_messageTypes[i].PBType = mt
	}
	xxx_ProtoFile_proto2_proto2_goTypes = nil
	xxx_ProtoFile_proto2_proto2_depIdxs = nil
}
