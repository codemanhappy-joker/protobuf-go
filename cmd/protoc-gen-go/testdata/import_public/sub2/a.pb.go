// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub2/a.proto

package sub2

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

type Sub2Message struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sub2Message) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_import_public_sub2_a_messageTypes[0].MessageOf(m)
}
func (m *Sub2Message) Reset()         { *m = Sub2Message{} }
func (m *Sub2Message) String() string { return proto.CompactTextString(m) }
func (*Sub2Message) ProtoMessage()    {}
func (*Sub2Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac53d99328778ac_gzipped, []int{0}
}

func (m *Sub2Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sub2Message.Unmarshal(m, b)
}
func (m *Sub2Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sub2Message.Marshal(b, m, deterministic)
}
func (m *Sub2Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sub2Message.Merge(m, src)
}
func (m *Sub2Message) XXX_Size() int {
	return xxx_messageInfo_Sub2Message.Size(m)
}
func (m *Sub2Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Sub2Message.DiscardUnknown(m)
}

var xxx_messageInfo_Sub2Message proto.InternalMessageInfo

func init() {
	proto.RegisterFile("import_public/sub2/a.proto", fileDescriptor_7ac53d99328778ac_gzipped)
	proto.RegisterType((*Sub2Message)(nil), "goproto.protoc.import_public.sub2.Sub2Message")
}

var fileDescriptor_7ac53d99328778ac = []byte{
	// 157 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x1a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x73, 0x75, 0x62, 0x32, 0x2f, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x32, 0x22,
	0x0d, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x4d,
	0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x75, 0x62, 0x32,
}

var fileDescriptor_7ac53d99328778ac_gzipped = protoapi.CompressGZIP(fileDescriptor_7ac53d99328778ac)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_import_public_sub2_a protoreflect.FileDescriptor

var xxx_ProtoFile_import_public_sub2_a_messageTypes [1]protoimpl.MessageType
var xxx_ProtoFile_import_public_sub2_a_goTypes = []interface{}{
	(*Sub2Message)(nil), // 0: goproto.protoc.import_public.sub2.Sub2Message
}
var xxx_ProtoFile_import_public_sub2_a_depIdxs = []int32{}

func init() {
	var messageTypes [1]protoreflect.MessageType
	ProtoFile_import_public_sub2_a = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_7ac53d99328778ac,
		GoTypes:            xxx_ProtoFile_import_public_sub2_a_goTypes,
		DependencyIndexes:  xxx_ProtoFile_import_public_sub2_a_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_import_public_sub2_a_goTypes[0:][:1]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_import_public_sub2_a_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_import_public_sub2_a_messageTypes[i].PBType = mt
	}
	xxx_ProtoFile_import_public_sub2_a_goTypes = nil
	xxx_ProtoFile_import_public_sub2_a_depIdxs = nil
}
