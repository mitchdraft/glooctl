// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transformation_filter.proto

/*
Package transformation is a generated protocol buffer package.

It is generated from these files:
	transformation_filter.proto

It has these top-level messages:
	Transformations
	Transformation
	Extraction
	RequestTemplate
	InjaTemplate
*/
package transformation

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Transformations struct {
	Transformations map[string]*Transformation `protobuf:"bytes,1,rep,name=transformations" json:"transformations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Transformations) Reset()         { *m = Transformations{} }
func (m *Transformations) String() string { return proto.CompactTextString(m) }
func (*Transformations) ProtoMessage()    {}
func (*Transformations) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{0}
}

func (m *Transformations) GetTransformations() map[string]*Transformation {
	if m != nil {
		return m.Transformations
	}
	return nil
}

// [#proto-status: experimental]
type Transformation struct {
	// Extractors are in the origin request language domain
	Extractors map[string]*Extraction `protobuf:"bytes,1,rep,name=extractors" json:"extractors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// Template is in the transformed request language domain
	// currently both are JSON
	RequestTemplate *RequestTemplate `protobuf:"bytes,2,opt,name=request_template,json=requestTemplate" json:"request_template,omitempty"`
}

func (m *Transformation) Reset()         { *m = Transformation{} }
func (m *Transformation) String() string { return proto.CompactTextString(m) }
func (*Transformation) ProtoMessage()    {}
func (*Transformation) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{1}
}

func (m *Transformation) GetExtractors() map[string]*Extraction {
	if m != nil {
		return m.Extractors
	}
	return nil
}

func (m *Transformation) GetRequestTemplate() *RequestTemplate {
	if m != nil {
		return m.RequestTemplate
	}
	return nil
}

type Extraction struct {
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// what information to extract. if extraction fails the result is
	// an empty value.
	Regex    string `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	Subgroup uint32 `protobuf:"varint,3,opt,name=subgroup,proto3" json:"subgroup,omitempty"`
}

func (m *Extraction) Reset()                    { *m = Extraction{} }
func (m *Extraction) String() string            { return proto.CompactTextString(m) }
func (*Extraction) ProtoMessage()               {}
func (*Extraction) Descriptor() ([]byte, []int) { return fileDescriptorTransformationFilter, []int{2} }

func (m *Extraction) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Extraction) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

func (m *Extraction) GetSubgroup() uint32 {
	if m != nil {
		return m.Subgroup
	}
	return 0
}

type RequestTemplate struct {
	Headers map[string]*InjaTemplate `protobuf:"bytes,1,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Body    *InjaTemplate            `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *RequestTemplate) Reset()         { *m = RequestTemplate{} }
func (m *RequestTemplate) String() string { return proto.CompactTextString(m) }
func (*RequestTemplate) ProtoMessage()    {}
func (*RequestTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{3}
}

func (m *RequestTemplate) GetHeaders() map[string]*InjaTemplate {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *RequestTemplate) GetBody() *InjaTemplate {
	if m != nil {
		return m.Body
	}
	return nil
}

//
// custom functions:
// header_value(name) -> from the original headers
// extracted_value(name, index) -> from the extracted values
type InjaTemplate struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *InjaTemplate) Reset()                    { *m = InjaTemplate{} }
func (m *InjaTemplate) String() string            { return proto.CompactTextString(m) }
func (*InjaTemplate) ProtoMessage()               {}
func (*InjaTemplate) Descriptor() ([]byte, []int) { return fileDescriptorTransformationFilter, []int{4} }

func (m *InjaTemplate) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Transformations)(nil), "transformation.Transformations")
	proto.RegisterType((*Transformation)(nil), "transformation.Transformation")
	proto.RegisterType((*Extraction)(nil), "transformation.Extraction")
	proto.RegisterType((*RequestTemplate)(nil), "transformation.RequestTemplate")
	proto.RegisterType((*InjaTemplate)(nil), "transformation.InjaTemplate")
}

func init() { proto.RegisterFile("transformation_filter.proto", fileDescriptorTransformationFilter) }

var fileDescriptorTransformationFilter = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x4e, 0xc2, 0x30,
	0x14, 0x87, 0x53, 0xfe, 0x29, 0x07, 0x64, 0xa4, 0x21, 0x66, 0x99, 0x46, 0x97, 0x5d, 0xed, 0xc2,
	0x2c, 0x04, 0xb9, 0x30, 0xde, 0x63, 0xd4, 0x0b, 0x2f, 0x1a, 0x62, 0xf4, 0x46, 0xd2, 0x49, 0x41,
	0x14, 0x56, 0xec, 0x3a, 0x03, 0x4f, 0xe0, 0x13, 0xfa, 0x08, 0xbe, 0x87, 0x71, 0xdd, 0x70, 0x6d,
	0x08, 0xde, 0xf5, 0xf4, 0xfc, 0xfa, 0x9d, 0x9d, 0x2f, 0x19, 0x1c, 0x49, 0x41, 0xa3, 0x78, 0xc2,
	0xc5, 0x82, 0xca, 0x19, 0x8f, 0x46, 0x93, 0xd9, 0x5c, 0x32, 0x11, 0x2c, 0x05, 0x97, 0x1c, 0xb7,
	0xf4, 0xa6, 0xf7, 0x85, 0xc0, 0x1a, 0x6a, 0x57, 0x31, 0x7e, 0x02, 0x4b, 0x4f, 0xc5, 0x36, 0x72,
	0xcb, 0x7e, 0xa3, 0xd7, 0x0f, 0xf4, 0xfb, 0xc0, 0x78, 0x69, 0xd6, 0x83, 0x48, 0x8a, 0x35, 0x31,
	0x61, 0x4e, 0x08, 0x9d, 0x6d, 0x41, 0xdc, 0x86, 0xf2, 0x1b, 0x5b, 0xdb, 0xc8, 0x45, 0x7e, 0x9d,
	0xfc, 0x1e, 0x71, 0x1f, 0xaa, 0x1f, 0x74, 0x9e, 0x30, 0xbb, 0xe4, 0x22, 0xbf, 0xd1, 0x3b, 0xd9,
	0x3d, 0x9f, 0xa8, 0xf0, 0x65, 0xe9, 0x02, 0x79, 0x9f, 0x25, 0x68, 0xe9, 0x5d, 0x7c, 0x07, 0xc0,
	0x56, 0x52, 0xd0, 0x67, 0xc9, 0x45, 0xbe, 0x51, 0xb0, 0x9b, 0x18, 0x0c, 0x36, 0x0f, 0xd4, 0x2e,
	0x05, 0x02, 0xbe, 0x85, 0xb6, 0x60, 0xef, 0x09, 0x8b, 0xe5, 0x48, 0xb2, 0xc5, 0x72, 0x4e, 0x65,
	0xfe, 0x9d, 0xa7, 0x26, 0x95, 0xa8, 0xdc, 0x30, 0x8b, 0x11, 0x4b, 0xe8, 0x17, 0xce, 0x23, 0x58,
	0xc6, 0xa8, 0x2d, 0x36, 0xba, 0xba, 0x0d, 0xc7, 0x9c, 0x92, 0x11, 0x0c, 0x13, 0xf7, 0x00, 0x7f,
	0x0d, 0x7c, 0x08, 0xb5, 0x17, 0x46, 0xc7, 0x4c, 0x64, 0xe0, 0xac, 0xc2, 0x1d, 0xa8, 0x0a, 0x36,
	0x65, 0xab, 0x94, 0x5d, 0x27, 0xaa, 0xc0, 0x0e, 0xec, 0xc7, 0x49, 0x38, 0x15, 0x3c, 0x59, 0xda,
	0x65, 0x17, 0xf9, 0x07, 0x64, 0x53, 0x7b, 0xdf, 0x08, 0x2c, 0x63, 0x2f, 0x7c, 0x05, 0x7b, 0x8a,
	0x97, 0xfb, 0x3d, 0xfb, 0xc7, 0x44, 0x70, 0xad, 0xe2, 0xca, 0x6e, 0xfe, 0x18, 0x77, 0xa1, 0x12,
	0xf2, 0xf1, 0x3a, 0x5b, 0xf4, 0xd8, 0x84, 0xdc, 0x44, 0xaf, 0x74, 0xe3, 0x32, 0x4d, 0x3a, 0x0f,
	0xd0, 0x2c, 0xa2, 0xb6, 0xd8, 0xeb, 0xe9, 0xf6, 0x76, 0x43, 0x0b, 0xfe, 0x3c, 0x68, 0x16, 0x5b,
	0x18, 0x43, 0x45, 0xb2, 0x95, 0xcc, 0xd0, 0xe9, 0x39, 0xac, 0xa5, 0x3f, 0xd7, 0xf9, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0xb5, 0x94, 0x87, 0x7b, 0x03, 0x00, 0x00,
}
