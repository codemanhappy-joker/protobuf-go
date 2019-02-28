// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extensions/extra/extra.proto

package extra

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

type ExtraMessage struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraMessage) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_extensions_extra_extra_messageTypes[0].MessageOf(m)
}
func (m *ExtraMessage) Reset()         { *m = ExtraMessage{} }
func (m *ExtraMessage) String() string { return proto.CompactTextString(m) }
func (*ExtraMessage) ProtoMessage()    {}
func (*ExtraMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_496c2a5e9c1e8739_gzipped, []int{0}
}

func (m *ExtraMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraMessage.Unmarshal(m, b)
}
func (m *ExtraMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraMessage.Marshal(b, m, deterministic)
}
func (m *ExtraMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraMessage.Merge(m, src)
}
func (m *ExtraMessage) XXX_Size() int {
	return xxx_messageInfo_ExtraMessage.Size(m)
}
func (m *ExtraMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraMessage proto.InternalMessageInfo

func (m *ExtraMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterFile("extensions/extra/extra.proto", fileDescriptor_496c2a5e9c1e8739_gzipped)
	proto.RegisterType((*ExtraMessage)(nil), "goproto.protoc.extension.extra.ExtraMessage")
}

var fileDescriptor_496c2a5e9c1e8739 = []byte{
	// 175 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x1c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x22,
	0x0a, 0x0c, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x72, 0x61,
}

var fileDescriptor_496c2a5e9c1e8739_gzipped = protoapi.CompressGZIP(fileDescriptor_496c2a5e9c1e8739)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_extensions_extra_extra protoreflect.FileDescriptor

var xxx_ProtoFile_extensions_extra_extra_messageTypes [1]protoimpl.MessageType
var xxx_ProtoFile_extensions_extra_extra_goTypes = []interface{}{
	(*ExtraMessage)(nil), // 0: goproto.protoc.extension.extra.ExtraMessage
}
var xxx_ProtoFile_extensions_extra_extra_depIdxs = []int32{}

func init() {
	var messageTypes [1]protoreflect.MessageType
	ProtoFile_extensions_extra_extra = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_496c2a5e9c1e8739,
		GoTypes:            xxx_ProtoFile_extensions_extra_extra_goTypes,
		DependencyIndexes:  xxx_ProtoFile_extensions_extra_extra_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_extensions_extra_extra_goTypes[0:][:1]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_extensions_extra_extra_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_extensions_extra_extra_messageTypes[i].PBType = mt
	}
	xxx_ProtoFile_extensions_extra_extra_goTypes = nil
	xxx_ProtoFile_extensions_extra_extra_depIdxs = nil
}
