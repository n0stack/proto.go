// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v0/metadata.proto

package pn0stack // import "github.com/n0stack/proto.go/v0"

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Metadata struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,2,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_4e341ddd36aef163, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Metadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "n0stack.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.Metadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.Metadata.LabelsEntry")
}

func init() { proto.RegisterFile("v0/metadata.proto", fileDescriptor_metadata_4e341ddd36aef163) }

var fileDescriptor_metadata_4e341ddd36aef163 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x33, 0xd0, 0xcf,
	0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf,
	0x33, 0x28, 0x2e, 0x49, 0x4c, 0xce, 0x56, 0x9a, 0xc6, 0xc4, 0xc5, 0xe1, 0x0b, 0x95, 0x13, 0x12,
	0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85,
	0x5c, 0xb8, 0xb8, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x25, 0x98,
	0x14, 0x98, 0x35, 0xb8, 0x8d, 0x94, 0xf4, 0xa0, 0xfa, 0xf5, 0x60, 0x7a, 0xf5, 0x1c, 0x11, 0x8a,
	0x5c, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0x90, 0xb5, 0x09, 0x99, 0x72, 0xb1, 0xe5, 0x24, 0x26, 0xa5,
	0xe6, 0x14, 0x4b, 0x30, 0x83, 0x0d, 0x90, 0xc5, 0x34, 0xc0, 0x07, 0x2c, 0x0f, 0xd1, 0x0b, 0x55,
	0x2c, 0x65, 0xc7, 0x25, 0x80, 0x6e, 0xae, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0x25, 0xd4, 0x8d,
	0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0x13, 0x58, 0x0c, 0xc2,
	0xb1, 0x62, 0xb2, 0x60, 0x94, 0xb2, 0xe4, 0xe2, 0x46, 0x32, 0x96, 0x14, 0xad, 0x4e, 0x9a, 0x51,
	0xea, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x50, 0xd7, 0xea, 0x83,
	0x43, 0x4f, 0x2f, 0x3d, 0x5f, 0xbf, 0xcc, 0xc0, 0xba, 0x00, 0x2a, 0x98, 0xc4, 0x06, 0x16, 0x35,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x7c, 0xdf, 0x5a, 0x68, 0x01, 0x00, 0x00,
}
