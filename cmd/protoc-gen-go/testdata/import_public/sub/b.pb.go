// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub/b.proto

package sub

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

type M2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_import_public_sub_b_messageTypes[0].MessageOf(m)
}
func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc66afda3d7c2232_gzipped, []int{0}
}

func (m *M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2.Unmarshal(m, b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2.Marshal(b, m, deterministic)
}
func (m *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(m, src)
}
func (m *M2) XXX_Size() int {
	return xxx_messageInfo_M2.Size(m)
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func init() {
	proto.RegisterFile("import_public/sub/b.proto", fileDescriptor_fc66afda3d7c2232_gzipped)
	proto.RegisterType((*M2)(nil), "goproto.protoc.import_public.sub.M2")
}

var fileDescriptor_fc66afda3d7c2232 = []byte{
	// 145 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x73, 0x75, 0x62, 0x2f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x22, 0x04, 0x0a,
	0x02, 0x4d, 0x32, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x75,
	0x62,
}

var fileDescriptor_fc66afda3d7c2232_gzipped = protoapi.CompressGZIP(fileDescriptor_fc66afda3d7c2232)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_import_public_sub_b protoreflect.FileDescriptor

var xxx_ProtoFile_import_public_sub_b_messageTypes [1]protoimpl.MessageType
var xxx_ProtoFile_import_public_sub_b_goTypes = []interface{}{
	(*M2)(nil), // 0: goproto.protoc.import_public.sub.M2
}
var xxx_ProtoFile_import_public_sub_b_depIdxs = []int32{}

func init() {
	var messageTypes [1]protoreflect.MessageType
	ProtoFile_import_public_sub_b = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_fc66afda3d7c2232,
		GoTypes:            xxx_ProtoFile_import_public_sub_b_goTypes,
		DependencyIndexes:  xxx_ProtoFile_import_public_sub_b_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_import_public_sub_b_goTypes[0:][:1]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_import_public_sub_b_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_import_public_sub_b_messageTypes[i].PBType = mt
	}
	xxx_ProtoFile_import_public_sub_b_goTypes = nil
	xxx_ProtoFile_import_public_sub_b_depIdxs = nil
}
