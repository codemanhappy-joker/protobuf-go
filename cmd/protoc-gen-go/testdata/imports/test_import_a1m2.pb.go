// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_import_a1m2.proto

package imports

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	test_a_1 "github.com/golang/protobuf/protoc-gen-go/testdata/imports/test_a_1"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type A1M2 struct {
	F                    *test_a_1.M2 `protobuf:"bytes,1,opt,name=f,proto3" json:"f,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *A1M2) ProtoReflect() protoreflect.Message {
	return xxx_ProtoFile_imports_test_import_a1m2_messageTypes[0].MessageOf(m)
}
func (m *A1M2) Reset()         { *m = A1M2{} }
func (m *A1M2) String() string { return proto.CompactTextString(m) }
func (*A1M2) ProtoMessage()    {}
func (*A1M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb27b114687957d_gzipped, []int{0}
}

func (m *A1M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A1M2.Unmarshal(m, b)
}
func (m *A1M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A1M2.Marshal(b, m, deterministic)
}
func (m *A1M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A1M2.Merge(m, src)
}
func (m *A1M2) XXX_Size() int {
	return xxx_messageInfo_A1M2.Size(m)
}
func (m *A1M2) XXX_DiscardUnknown() {
	xxx_messageInfo_A1M2.DiscardUnknown(m)
}

var xxx_messageInfo_A1M2 proto.InternalMessageInfo

func (m *A1M2) GetF() *test_a_1.M2 {
	if m != nil {
		return m.F
	}
	return nil
}

func init() {
	proto.RegisterFile("imports/test_import_a1m2.proto", fileDescriptor_bdb27b114687957d_gzipped)
	proto.RegisterType((*A1M2)(nil), "test.A1M2")
}

var fileDescriptor_bdb27b114687957d = []byte{
	// 168 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x1e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x31, 0x6d, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x31, 0x2f, 0x6d, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x20, 0x0a, 0x04, 0x41, 0x31, 0x4d, 0x32, 0x12, 0x18, 0x0a, 0x01, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x2e, 0x4d, 0x32,
	0x52, 0x01, 0x66, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var fileDescriptor_bdb27b114687957d_gzipped = protoapi.CompressGZIP(fileDescriptor_bdb27b114687957d)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var ProtoFile_imports_test_import_a1m2 protoreflect.FileDescriptor

var xxx_ProtoFile_imports_test_import_a1m2_messageTypes [1]protoimpl.MessageType
var xxx_ProtoFile_imports_test_import_a1m2_goTypes = []interface{}{
	(*A1M2)(nil),        // 0: test.A1M2
	(*test_a_1.M2)(nil), // 1: test.a.M2
}
var xxx_ProtoFile_imports_test_import_a1m2_depIdxs = []int32{
	1, // test.A1M2.f:type_name -> test.a.M2
}

func init() {
	var messageTypes [1]protoreflect.MessageType
	ProtoFile_imports_test_import_a1m2 = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_bdb27b114687957d,
		GoTypes:            xxx_ProtoFile_imports_test_import_a1m2_goTypes,
		DependencyIndexes:  xxx_ProtoFile_imports_test_import_a1m2_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_ProtoFile_imports_test_import_a1m2_goTypes[0:][:1]
	for i, mt := range messageTypes[:] {
		xxx_ProtoFile_imports_test_import_a1m2_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_ProtoFile_imports_test_import_a1m2_messageTypes[i].PBType = mt
	}
	xxx_ProtoFile_imports_test_import_a1m2_goTypes = nil
	xxx_ProtoFile_imports_test_import_a1m2_depIdxs = nil
}
