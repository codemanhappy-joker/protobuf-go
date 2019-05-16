// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conformance/conformance.proto

package conformance

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type WireFormat int32

const (
	WireFormat_UNSPECIFIED WireFormat = 0
	WireFormat_PROTOBUF    WireFormat = 1
	WireFormat_JSON        WireFormat = 2
	WireFormat_JSPB        WireFormat = 3
	WireFormat_TEXT_FORMAT WireFormat = 4
)

// Deprecated: Use WireFormat.Type.Values instead.
var WireFormat_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "PROTOBUF",
	2: "JSON",
	3: "JSPB",
	4: "TEXT_FORMAT",
}

// Deprecated: Use WireFormat.Type.Values instead.
var WireFormat_value = map[string]int32{
	"UNSPECIFIED": 0,
	"PROTOBUF":    1,
	"JSON":        2,
	"JSPB":        3,
	"TEXT_FORMAT": 4,
}

func (x WireFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WireFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_conformance_conformance_proto_enumTypes[0].Descriptor()
}

// Deprecated: Use Descriptor instead.
func (WireFormat) Type() protoreflect.EnumType {
	return file_conformance_conformance_proto_enumTypes[0]
}

func (x WireFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WireFormat.Type instead.
func (WireFormat) EnumDescriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{0}
}

type TestCategory int32

const (
	TestCategory_UNSPECIFIED_TEST TestCategory = 0
	TestCategory_BINARY_TEST      TestCategory = 1
	TestCategory_JSON_TEST        TestCategory = 2
	// Similar to JSON_TEST. However, during parsing json, testee should ignore
	// unknown fields. This feature is optional. Each implementation can descide
	// whether to support it.  See
	// https://developers.google.com/protocol-buffers/docs/proto3#json_options
	// for more detail.
	TestCategory_JSON_IGNORE_UNKNOWN_PARSING_TEST TestCategory = 3
	// Test jspb wire format. Google internal only. Opensource testees just skip it.
	TestCategory_JSPB_TEST TestCategory = 4
	// Test text format. For cpp, java and python, testees can already deal with
	// this type. Testees of other languages can simply skip it.
	TestCategory_TEXT_FORMAT_TEST TestCategory = 5
)

// Deprecated: Use TestCategory.Type.Values instead.
var TestCategory_name = map[int32]string{
	0: "UNSPECIFIED_TEST",
	1: "BINARY_TEST",
	2: "JSON_TEST",
	3: "JSON_IGNORE_UNKNOWN_PARSING_TEST",
	4: "JSPB_TEST",
	5: "TEXT_FORMAT_TEST",
}

// Deprecated: Use TestCategory.Type.Values instead.
var TestCategory_value = map[string]int32{
	"UNSPECIFIED_TEST":                 0,
	"BINARY_TEST":                      1,
	"JSON_TEST":                        2,
	"JSON_IGNORE_UNKNOWN_PARSING_TEST": 3,
	"JSPB_TEST":                        4,
	"TEXT_FORMAT_TEST":                 5,
}

func (x TestCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_conformance_conformance_proto_enumTypes[1].Descriptor()
}

// Deprecated: Use Descriptor instead.
func (TestCategory) Type() protoreflect.EnumType {
	return file_conformance_conformance_proto_enumTypes[1]
}

func (x TestCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestCategory.Type instead.
func (TestCategory) EnumDescriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{1}
}

// The conformance runner will request a list of failures as the first request.
// This will be known by message_type == "conformance.FailureSet", a conformance
// test should return a serialized FailureSet in protobuf_payload.
type FailureSet struct {
	Failure              []string                `protobuf:"bytes,1,rep,name=failure,proto3" json:"failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *FailureSet) Reset() {
	*x = FailureSet{}
}

func (x *FailureSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailureSet) ProtoMessage() {}

func (x *FailureSet) ProtoReflect() protoreflect.Message {
	return file_conformance_conformance_proto_msgTypes[0].MessageOf(x)
}

func (m *FailureSet) XXX_Methods() *protoiface.Methods {
	return file_conformance_conformance_proto_msgTypes[0].Methods()
}

// Deprecated: Use FailureSet.ProtoReflect.Type instead.
func (*FailureSet) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{0}
}

func (x *FailureSet) GetFailure() []string {
	if x != nil {
		return x.Failure
	}
	return nil
}

// Represents a single test case's input.  The testee should:
//
//   1. parse this proto (which should always succeed)
//   2. parse the protobuf or JSON payload in "payload" (which may fail)
//   3. if the parse succeeded, serialize the message in the requested format.
type ConformanceRequest struct {
	// The payload (whether protobuf of JSON) is always for a
	// protobuf_test_messages.proto3.TestAllTypes proto (as defined in
	// src/google/protobuf/proto3_test_messages.proto).
	//
	// TODO(haberman): if/when we expand the conformance tests to support proto2,
	// we will want to include a field that lets the payload/response be a
	// protobuf_test_messages.proto2.TestAllTypes message instead.
	//
	// Types that are valid to be assigned to Payload:
	//	*ConformanceRequest_ProtobufPayload
	//	*ConformanceRequest_JsonPayload
	// Google internal only.  Opensource testees just skip it.
	//	*ConformanceRequest_JspbPayload
	//	*ConformanceRequest_TextPayload
	Payload isConformanceRequest_Payload `protobuf_oneof:"payload"`
	// Which format should the testee serialize its message to?
	RequestedOutputFormat WireFormat `protobuf:"varint,3,opt,name=requested_output_format,json=requestedOutputFormat,proto3,enum=conformance.WireFormat" json:"requested_output_format,omitempty"`
	// The full name for the test message to use; for the moment, either:
	// protobuf_test_messages.proto3.TestAllTypesProto3 or
	// protobuf_test_messages.proto2.TestAllTypesProto2.
	MessageType string `protobuf:"bytes,4,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	// Each test is given a specific test category. Some category may need
	// spedific support in testee programs. Refer to the defintion of TestCategory
	// for more information.
	TestCategory TestCategory `protobuf:"varint,5,opt,name=test_category,json=testCategory,proto3,enum=conformance.TestCategory" json:"test_category,omitempty"`
	// Specify details for how to encode jspb.
	JspbEncodingOptions  *JspbEncodingConfig     `protobuf:"bytes,6,opt,name=jspb_encoding_options,json=jspbEncodingOptions,proto3" json:"jspb_encoding_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *ConformanceRequest) Reset() {
	*x = ConformanceRequest{}
}

func (x *ConformanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConformanceRequest) ProtoMessage() {}

func (x *ConformanceRequest) ProtoReflect() protoreflect.Message {
	return file_conformance_conformance_proto_msgTypes[1].MessageOf(x)
}

func (m *ConformanceRequest) XXX_Methods() *protoiface.Methods {
	return file_conformance_conformance_proto_msgTypes[1].Methods()
}

// Deprecated: Use ConformanceRequest.ProtoReflect.Type instead.
func (*ConformanceRequest) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{1}
}

func (m *ConformanceRequest) GetPayload() isConformanceRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *ConformanceRequest) GetProtobufPayload() []byte {
	if x, ok := x.GetPayload().(*ConformanceRequest_ProtobufPayload); ok {
		return x.ProtobufPayload
	}
	return nil
}

func (x *ConformanceRequest) GetJsonPayload() string {
	if x, ok := x.GetPayload().(*ConformanceRequest_JsonPayload); ok {
		return x.JsonPayload
	}
	return ""
}

func (x *ConformanceRequest) GetJspbPayload() string {
	if x, ok := x.GetPayload().(*ConformanceRequest_JspbPayload); ok {
		return x.JspbPayload
	}
	return ""
}

func (x *ConformanceRequest) GetTextPayload() string {
	if x, ok := x.GetPayload().(*ConformanceRequest_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (x *ConformanceRequest) GetRequestedOutputFormat() WireFormat {
	if x != nil {
		return x.RequestedOutputFormat
	}
	return WireFormat_UNSPECIFIED
}

func (x *ConformanceRequest) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *ConformanceRequest) GetTestCategory() TestCategory {
	if x != nil {
		return x.TestCategory
	}
	return TestCategory_UNSPECIFIED_TEST
}

func (x *ConformanceRequest) GetJspbEncodingOptions() *JspbEncodingConfig {
	if x != nil {
		return x.JspbEncodingOptions
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConformanceRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConformanceRequest_ProtobufPayload)(nil),
		(*ConformanceRequest_JsonPayload)(nil),
		(*ConformanceRequest_JspbPayload)(nil),
		(*ConformanceRequest_TextPayload)(nil),
	}
}

type isConformanceRequest_Payload interface {
	isConformanceRequest_Payload()
}

type ConformanceRequest_ProtobufPayload struct {
	ProtobufPayload []byte `protobuf:"bytes,1,opt,name=protobuf_payload,json=protobufPayload,proto3,oneof"`
}

type ConformanceRequest_JsonPayload struct {
	JsonPayload string `protobuf:"bytes,2,opt,name=json_payload,json=jsonPayload,proto3,oneof"`
}

type ConformanceRequest_JspbPayload struct {
	JspbPayload string `protobuf:"bytes,7,opt,name=jspb_payload,json=jspbPayload,proto3,oneof"`
}

type ConformanceRequest_TextPayload struct {
	TextPayload string `protobuf:"bytes,8,opt,name=text_payload,json=textPayload,proto3,oneof"`
}

func (*ConformanceRequest_ProtobufPayload) isConformanceRequest_Payload() {}

func (*ConformanceRequest_JsonPayload) isConformanceRequest_Payload() {}

func (*ConformanceRequest_JspbPayload) isConformanceRequest_Payload() {}

func (*ConformanceRequest_TextPayload) isConformanceRequest_Payload() {}

// Represents a single test case's output.
type ConformanceResponse struct {
	// Types that are valid to be assigned to Result:
	// This string should be set to indicate parsing failed.  The string can
	// provide more information about the parse error if it is available.
	//
	// Setting this string does not necessarily mean the testee failed the
	// test.  Some of the test cases are intentionally invalid input.
	//	*ConformanceResponse_ParseError
	// If the input was successfully parsed but errors occurred when
	// serializing it to the requested output format, set the error message in
	// this field.
	//	*ConformanceResponse_SerializeError
	// This should be set if some other error occurred.  This will always
	// indicate that the test failed.  The string can provide more information
	// about the failure.
	//	*ConformanceResponse_RuntimeError
	// If the input was successfully parsed and the requested output was
	// protobuf, serialize it to protobuf and set it in this field.
	//	*ConformanceResponse_ProtobufPayload
	// If the input was successfully parsed and the requested output was JSON,
	// serialize to JSON and set it in this field.
	//	*ConformanceResponse_JsonPayload
	// For when the testee skipped the test, likely because a certain feature
	// wasn't supported, like JSON input/output.
	//	*ConformanceResponse_Skipped
	// If the input was successfully parsed and the requested output was JSPB,
	// serialize to JSPB and set it in this field. JSPB is google internal only
	// format. Opensource testees can just skip it.
	//	*ConformanceResponse_JspbPayload
	// If the input was successfully parsed and the requested output was
	// TEXT_FORMAT, serialize to TEXT_FORMAT and set it in this field.
	//	*ConformanceResponse_TextPayload
	Result               isConformanceResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields      `json:"-"`
	XXX_sizecache        protoimpl.SizeCache          `json:"-"`
}

func (x *ConformanceResponse) Reset() {
	*x = ConformanceResponse{}
}

func (x *ConformanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConformanceResponse) ProtoMessage() {}

func (x *ConformanceResponse) ProtoReflect() protoreflect.Message {
	return file_conformance_conformance_proto_msgTypes[2].MessageOf(x)
}

func (m *ConformanceResponse) XXX_Methods() *protoiface.Methods {
	return file_conformance_conformance_proto_msgTypes[2].Methods()
}

// Deprecated: Use ConformanceResponse.ProtoReflect.Type instead.
func (*ConformanceResponse) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{2}
}

func (m *ConformanceResponse) GetResult() isConformanceResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ConformanceResponse) GetParseError() string {
	if x, ok := x.GetResult().(*ConformanceResponse_ParseError); ok {
		return x.ParseError
	}
	return ""
}

func (x *ConformanceResponse) GetSerializeError() string {
	if x, ok := x.GetResult().(*ConformanceResponse_SerializeError); ok {
		return x.SerializeError
	}
	return ""
}

func (x *ConformanceResponse) GetRuntimeError() string {
	if x, ok := x.GetResult().(*ConformanceResponse_RuntimeError); ok {
		return x.RuntimeError
	}
	return ""
}

func (x *ConformanceResponse) GetProtobufPayload() []byte {
	if x, ok := x.GetResult().(*ConformanceResponse_ProtobufPayload); ok {
		return x.ProtobufPayload
	}
	return nil
}

func (x *ConformanceResponse) GetJsonPayload() string {
	if x, ok := x.GetResult().(*ConformanceResponse_JsonPayload); ok {
		return x.JsonPayload
	}
	return ""
}

func (x *ConformanceResponse) GetSkipped() string {
	if x, ok := x.GetResult().(*ConformanceResponse_Skipped); ok {
		return x.Skipped
	}
	return ""
}

func (x *ConformanceResponse) GetJspbPayload() string {
	if x, ok := x.GetResult().(*ConformanceResponse_JspbPayload); ok {
		return x.JspbPayload
	}
	return ""
}

func (x *ConformanceResponse) GetTextPayload() string {
	if x, ok := x.GetResult().(*ConformanceResponse_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConformanceResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConformanceResponse_ParseError)(nil),
		(*ConformanceResponse_SerializeError)(nil),
		(*ConformanceResponse_RuntimeError)(nil),
		(*ConformanceResponse_ProtobufPayload)(nil),
		(*ConformanceResponse_JsonPayload)(nil),
		(*ConformanceResponse_Skipped)(nil),
		(*ConformanceResponse_JspbPayload)(nil),
		(*ConformanceResponse_TextPayload)(nil),
	}
}

type isConformanceResponse_Result interface {
	isConformanceResponse_Result()
}

type ConformanceResponse_ParseError struct {
	ParseError string `protobuf:"bytes,1,opt,name=parse_error,json=parseError,proto3,oneof"`
}

type ConformanceResponse_SerializeError struct {
	SerializeError string `protobuf:"bytes,6,opt,name=serialize_error,json=serializeError,proto3,oneof"`
}

type ConformanceResponse_RuntimeError struct {
	RuntimeError string `protobuf:"bytes,2,opt,name=runtime_error,json=runtimeError,proto3,oneof"`
}

type ConformanceResponse_ProtobufPayload struct {
	ProtobufPayload []byte `protobuf:"bytes,3,opt,name=protobuf_payload,json=protobufPayload,proto3,oneof"`
}

type ConformanceResponse_JsonPayload struct {
	JsonPayload string `protobuf:"bytes,4,opt,name=json_payload,json=jsonPayload,proto3,oneof"`
}

type ConformanceResponse_Skipped struct {
	Skipped string `protobuf:"bytes,5,opt,name=skipped,proto3,oneof"`
}

type ConformanceResponse_JspbPayload struct {
	JspbPayload string `protobuf:"bytes,7,opt,name=jspb_payload,json=jspbPayload,proto3,oneof"`
}

type ConformanceResponse_TextPayload struct {
	TextPayload string `protobuf:"bytes,8,opt,name=text_payload,json=textPayload,proto3,oneof"`
}

func (*ConformanceResponse_ParseError) isConformanceResponse_Result() {}

func (*ConformanceResponse_SerializeError) isConformanceResponse_Result() {}

func (*ConformanceResponse_RuntimeError) isConformanceResponse_Result() {}

func (*ConformanceResponse_ProtobufPayload) isConformanceResponse_Result() {}

func (*ConformanceResponse_JsonPayload) isConformanceResponse_Result() {}

func (*ConformanceResponse_Skipped) isConformanceResponse_Result() {}

func (*ConformanceResponse_JspbPayload) isConformanceResponse_Result() {}

func (*ConformanceResponse_TextPayload) isConformanceResponse_Result() {}

// Encoding options for jspb format.
type JspbEncodingConfig struct {
	// Encode the value field of Any as jspb array if ture, otherwise binary.
	UseJspbArrayAnyFormat bool                    `protobuf:"varint,1,opt,name=use_jspb_array_any_format,json=useJspbArrayAnyFormat,proto3" json:"use_jspb_array_any_format,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                `json:"-"`
	XXX_unrecognized      protoimpl.UnknownFields `json:"-"`
	XXX_sizecache         protoimpl.SizeCache     `json:"-"`
}

func (x *JspbEncodingConfig) Reset() {
	*x = JspbEncodingConfig{}
}

func (x *JspbEncodingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JspbEncodingConfig) ProtoMessage() {}

func (x *JspbEncodingConfig) ProtoReflect() protoreflect.Message {
	return file_conformance_conformance_proto_msgTypes[3].MessageOf(x)
}

func (m *JspbEncodingConfig) XXX_Methods() *protoiface.Methods {
	return file_conformance_conformance_proto_msgTypes[3].Methods()
}

// Deprecated: Use JspbEncodingConfig.ProtoReflect.Type instead.
func (*JspbEncodingConfig) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{3}
}

func (x *JspbEncodingConfig) GetUseJspbArrayAnyFormat() bool {
	if x != nil {
		return x.UseJspbArrayAnyFormat
	}
	return false
}

var File_conformance_conformance_proto protoreflect.FileDescriptor

var file_conformance_conformance_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x0a,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x22, 0xc4, 0x03, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x10, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a,
	0x0c, 0x6a, 0x73, 0x70, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6a, 0x73, 0x70, 0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4f, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0c, 0x74,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x53, 0x0a, 0x15, 0x6a,
	0x73, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4a, 0x73, 0x70, 0x62, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x13, 0x6a, 0x73, 0x70,
	0x62, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xcc, 0x02, 0x0a, 0x13,
	0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x25, 0x0a, 0x0d, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6a,
	0x73, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x07, 0x73, 0x6b,
	0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x73,
	0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x6a, 0x73, 0x70, 0x62, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x6a, 0x73, 0x70, 0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x4e, 0x0a, 0x12, 0x4a, 0x73,
	0x70, 0x62, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x38, 0x0a, 0x19, 0x75, 0x73, 0x65, 0x5f, 0x6a, 0x73, 0x70, 0x62, 0x5f, 0x61, 0x72, 0x72,
	0x61, 0x79, 0x5f, 0x61, 0x6e, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x75, 0x73, 0x65, 0x4a, 0x73, 0x70, 0x62, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x41, 0x6e, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2a, 0x50, 0x0a, 0x0a, 0x57, 0x69,
	0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f,
	0x54, 0x4f, 0x42, 0x55, 0x46, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x50, 0x42, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x54,
	0x45, 0x58, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x04, 0x2a, 0x8f, 0x01, 0x0a,
	0x0c, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x10, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x54, 0x45, 0x53,
	0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x5f, 0x54, 0x45,
	0x53, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x45, 0x53,
	0x54, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x47, 0x4e, 0x4f,
	0x52, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x41, 0x52, 0x53, 0x49,
	0x4e, 0x47, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x53, 0x50,
	0x42, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x45, 0x58, 0x54,
	0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x05, 0x42, 0x5d,
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x5a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conformance_conformance_proto_rawDescOnce sync.Once
	file_conformance_conformance_proto_rawDescData = file_conformance_conformance_proto_rawDesc
)

func file_conformance_conformance_proto_rawDescGZIP() []byte {
	file_conformance_conformance_proto_rawDescOnce.Do(func() {
		file_conformance_conformance_proto_rawDescData = protoimpl.X.CompressGZIP(file_conformance_conformance_proto_rawDescData)
	})
	return file_conformance_conformance_proto_rawDescData
}

var file_conformance_conformance_proto_enumTypes = make([]protoreflect.EnumType, 2)
var file_conformance_conformance_proto_msgTypes = make([]protoimpl.MessageType, 4)
var file_conformance_conformance_proto_goTypes = []interface{}{
	(WireFormat)(0),             // 0: conformance.WireFormat
	(TestCategory)(0),           // 1: conformance.TestCategory
	(*FailureSet)(nil),          // 2: conformance.FailureSet
	(*ConformanceRequest)(nil),  // 3: conformance.ConformanceRequest
	(*ConformanceResponse)(nil), // 4: conformance.ConformanceResponse
	(*JspbEncodingConfig)(nil),  // 5: conformance.JspbEncodingConfig
}
var file_conformance_conformance_proto_depIdxs = []int32{
	0, // conformance.ConformanceRequest.requested_output_format:type_name -> conformance.WireFormat
	1, // conformance.ConformanceRequest.test_category:type_name -> conformance.TestCategory
	5, // conformance.ConformanceRequest.jspb_encoding_options:type_name -> conformance.JspbEncodingConfig
}

func init() { file_conformance_conformance_proto_init() }
func file_conformance_conformance_proto_init() {
	if File_conformance_conformance_proto != nil {
		return
	}
	File_conformance_conformance_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_conformance_conformance_proto_rawDesc,
		GoTypes:            file_conformance_conformance_proto_goTypes,
		DependencyIndexes:  file_conformance_conformance_proto_depIdxs,
		EnumOutputTypes:    file_conformance_conformance_proto_enumTypes,
		MessageOutputTypes: file_conformance_conformance_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_conformance_conformance_proto_rawDesc = nil
	file_conformance_conformance_proto_goTypes = nil
	file_conformance_conformance_proto_depIdxs = nil
}
