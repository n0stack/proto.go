// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provisioning/v0/volume.proto

package pprovisioning // import "github.com/n0stack/proto.go/provisioning/v0"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import v0 "github.com/n0stack/proto.go/v0"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VolumeStatus_VolumeState int32

const (
	// falied state because failed some process on API.
	VolumeStatus_FAILED VolumeStatus_VolumeState = 0
	// unknown state because failed to connect for scheduled node after RUNNING.
	VolumeStatus_UNKNOWN   VolumeStatus_VolumeState = 1
	VolumeStatus_AVAILABLE VolumeStatus_VolumeState = 2
	VolumeStatus_IN_USE    VolumeStatus_VolumeState = 3
)

var VolumeStatus_VolumeState_name = map[int32]string{
	0: "FAILED",
	1: "UNKNOWN",
	2: "AVAILABLE",
	3: "IN_USE",
}
var VolumeStatus_VolumeState_value = map[string]int32{
	"FAILED":    0,
	"UNKNOWN":   1,
	"AVAILABLE": 2,
	"IN_USE":    3,
}

func (x VolumeStatus_VolumeState) String() string {
	return proto.EnumName(VolumeStatus_VolumeState_name, int32(x))
}
func (VolumeStatus_VolumeState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{2, 0}
}

type Volume struct {
	Metadata             *v0.Metadata  `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec                 *VolumeSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status               *VolumeStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Volume) Reset()         { *m = Volume{} }
func (m *Volume) String() string { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()    {}
func (*Volume) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{0}
}
func (m *Volume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Volume.Unmarshal(m, b)
}
func (m *Volume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Volume.Marshal(b, m, deterministic)
}
func (dst *Volume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Volume.Merge(dst, src)
}
func (m *Volume) XXX_Size() int {
	return xxx_messageInfo_Volume.Size(m)
}
func (m *Volume) XXX_DiscardUnknown() {
	xxx_messageInfo_Volume.DiscardUnknown(m)
}

var xxx_messageInfo_Volume proto.InternalMessageInfo

func (m *Volume) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Volume) GetSpec() *VolumeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Volume) GetStatus() *VolumeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type VolumeSpec struct {
	RequestBytes         uint64   `protobuf:"varint,1,opt,name=request_bytes,json=requestBytes" json:"request_bytes,omitempty"`
	LimitBytes           uint64   `protobuf:"varint,2,opt,name=limit_bytes,json=limitBytes" json:"limit_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeSpec) Reset()         { *m = VolumeSpec{} }
func (m *VolumeSpec) String() string { return proto.CompactTextString(m) }
func (*VolumeSpec) ProtoMessage()    {}
func (*VolumeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{1}
}
func (m *VolumeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeSpec.Unmarshal(m, b)
}
func (m *VolumeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeSpec.Marshal(b, m, deterministic)
}
func (dst *VolumeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeSpec.Merge(dst, src)
}
func (m *VolumeSpec) XXX_Size() int {
	return xxx_messageInfo_VolumeSpec.Size(m)
}
func (m *VolumeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeSpec proto.InternalMessageInfo

func (m *VolumeSpec) GetRequestBytes() uint64 {
	if m != nil {
		return m.RequestBytes
	}
	return 0
}

func (m *VolumeSpec) GetLimitBytes() uint64 {
	if m != nil {
		return m.LimitBytes
	}
	return 0
}

type VolumeStatus struct {
	State                VolumeStatus_VolumeState `protobuf:"varint,1,opt,name=state,enum=n0stack.provisioning.VolumeStatus_VolumeState" json:"state,omitempty"`
	NodeName             string                   `protobuf:"bytes,2,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	StorageName          string                   `protobuf:"bytes,3,opt,name=storage_name,json=storageName" json:"storage_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *VolumeStatus) Reset()         { *m = VolumeStatus{} }
func (m *VolumeStatus) String() string { return proto.CompactTextString(m) }
func (*VolumeStatus) ProtoMessage()    {}
func (*VolumeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{2}
}
func (m *VolumeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeStatus.Unmarshal(m, b)
}
func (m *VolumeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeStatus.Marshal(b, m, deterministic)
}
func (dst *VolumeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeStatus.Merge(dst, src)
}
func (m *VolumeStatus) XXX_Size() int {
	return xxx_messageInfo_VolumeStatus.Size(m)
}
func (m *VolumeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeStatus proto.InternalMessageInfo

func (m *VolumeStatus) GetState() VolumeStatus_VolumeState {
	if m != nil {
		return m.State
	}
	return VolumeStatus_FAILED
}

func (m *VolumeStatus) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *VolumeStatus) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

type CreateEmptyVolumeRequest struct {
	Metadata             *v0.Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec                 *VolumeSpec  `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateEmptyVolumeRequest) Reset()         { *m = CreateEmptyVolumeRequest{} }
func (m *CreateEmptyVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEmptyVolumeRequest) ProtoMessage()    {}
func (*CreateEmptyVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{3}
}
func (m *CreateEmptyVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEmptyVolumeRequest.Unmarshal(m, b)
}
func (m *CreateEmptyVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEmptyVolumeRequest.Marshal(b, m, deterministic)
}
func (dst *CreateEmptyVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEmptyVolumeRequest.Merge(dst, src)
}
func (m *CreateEmptyVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEmptyVolumeRequest.Size(m)
}
func (m *CreateEmptyVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEmptyVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEmptyVolumeRequest proto.InternalMessageInfo

func (m *CreateEmptyVolumeRequest) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateEmptyVolumeRequest) GetSpec() *VolumeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type CreateVolumeWithDownloadingRequest struct {
	Metadata             *v0.Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec                 *VolumeSpec  `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	SourceUrl            string       `protobuf:"bytes,3,opt,name=source_url,json=sourceUrl" json:"source_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateVolumeWithDownloadingRequest) Reset()         { *m = CreateVolumeWithDownloadingRequest{} }
func (m *CreateVolumeWithDownloadingRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVolumeWithDownloadingRequest) ProtoMessage()    {}
func (*CreateVolumeWithDownloadingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{4}
}
func (m *CreateVolumeWithDownloadingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVolumeWithDownloadingRequest.Unmarshal(m, b)
}
func (m *CreateVolumeWithDownloadingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVolumeWithDownloadingRequest.Marshal(b, m, deterministic)
}
func (dst *CreateVolumeWithDownloadingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVolumeWithDownloadingRequest.Merge(dst, src)
}
func (m *CreateVolumeWithDownloadingRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVolumeWithDownloadingRequest.Size(m)
}
func (m *CreateVolumeWithDownloadingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVolumeWithDownloadingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVolumeWithDownloadingRequest proto.InternalMessageInfo

func (m *CreateVolumeWithDownloadingRequest) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateVolumeWithDownloadingRequest) GetSpec() *VolumeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *CreateVolumeWithDownloadingRequest) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

type ListVolumesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVolumesRequest) Reset()         { *m = ListVolumesRequest{} }
func (m *ListVolumesRequest) String() string { return proto.CompactTextString(m) }
func (*ListVolumesRequest) ProtoMessage()    {}
func (*ListVolumesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{5}
}
func (m *ListVolumesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVolumesRequest.Unmarshal(m, b)
}
func (m *ListVolumesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVolumesRequest.Marshal(b, m, deterministic)
}
func (dst *ListVolumesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVolumesRequest.Merge(dst, src)
}
func (m *ListVolumesRequest) XXX_Size() int {
	return xxx_messageInfo_ListVolumesRequest.Size(m)
}
func (m *ListVolumesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVolumesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListVolumesRequest proto.InternalMessageInfo

type ListVolumesResponse struct {
	Volumes              []*Volume `protobuf:"bytes,1,rep,name=volumes" json:"volumes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListVolumesResponse) Reset()         { *m = ListVolumesResponse{} }
func (m *ListVolumesResponse) String() string { return proto.CompactTextString(m) }
func (*ListVolumesResponse) ProtoMessage()    {}
func (*ListVolumesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{6}
}
func (m *ListVolumesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVolumesResponse.Unmarshal(m, b)
}
func (m *ListVolumesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVolumesResponse.Marshal(b, m, deterministic)
}
func (dst *ListVolumesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVolumesResponse.Merge(dst, src)
}
func (m *ListVolumesResponse) XXX_Size() int {
	return xxx_messageInfo_ListVolumesResponse.Size(m)
}
func (m *ListVolumesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVolumesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListVolumesResponse proto.InternalMessageInfo

func (m *ListVolumesResponse) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type GetVolumeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVolumeRequest) Reset()         { *m = GetVolumeRequest{} }
func (m *GetVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*GetVolumeRequest) ProtoMessage()    {}
func (*GetVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{7}
}
func (m *GetVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVolumeRequest.Unmarshal(m, b)
}
func (m *GetVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVolumeRequest.Marshal(b, m, deterministic)
}
func (dst *GetVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVolumeRequest.Merge(dst, src)
}
func (m *GetVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_GetVolumeRequest.Size(m)
}
func (m *GetVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVolumeRequest proto.InternalMessageInfo

func (m *GetVolumeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateVolumeRequest struct {
	Metadata             *v0.Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec                 *VolumeSpec  `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateVolumeRequest) Reset()         { *m = UpdateVolumeRequest{} }
func (m *UpdateVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateVolumeRequest) ProtoMessage()    {}
func (*UpdateVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{8}
}
func (m *UpdateVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVolumeRequest.Unmarshal(m, b)
}
func (m *UpdateVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVolumeRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVolumeRequest.Merge(dst, src)
}
func (m *UpdateVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateVolumeRequest.Size(m)
}
func (m *UpdateVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVolumeRequest proto.InternalMessageInfo

func (m *UpdateVolumeRequest) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *UpdateVolumeRequest) GetSpec() *VolumeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type DeleteVolumeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteVolumeRequest) Reset()         { *m = DeleteVolumeRequest{} }
func (m *DeleteVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteVolumeRequest) ProtoMessage()    {}
func (*DeleteVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{9}
}
func (m *DeleteVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteVolumeRequest.Unmarshal(m, b)
}
func (m *DeleteVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteVolumeRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteVolumeRequest.Merge(dst, src)
}
func (m *DeleteVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteVolumeRequest.Size(m)
}
func (m *DeleteVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteVolumeRequest proto.InternalMessageInfo

func (m *DeleteVolumeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SetInuseVolumeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetInuseVolumeRequest) Reset()         { *m = SetInuseVolumeRequest{} }
func (m *SetInuseVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*SetInuseVolumeRequest) ProtoMessage()    {}
func (*SetInuseVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{10}
}
func (m *SetInuseVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetInuseVolumeRequest.Unmarshal(m, b)
}
func (m *SetInuseVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetInuseVolumeRequest.Marshal(b, m, deterministic)
}
func (dst *SetInuseVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetInuseVolumeRequest.Merge(dst, src)
}
func (m *SetInuseVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_SetInuseVolumeRequest.Size(m)
}
func (m *SetInuseVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetInuseVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetInuseVolumeRequest proto.InternalMessageInfo

func (m *SetInuseVolumeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SetAvailableVolumeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetAvailableVolumeRequest) Reset()         { *m = SetAvailableVolumeRequest{} }
func (m *SetAvailableVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*SetAvailableVolumeRequest) ProtoMessage()    {}
func (*SetAvailableVolumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_volume_b864912e45f02d3a, []int{11}
}
func (m *SetAvailableVolumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAvailableVolumeRequest.Unmarshal(m, b)
}
func (m *SetAvailableVolumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAvailableVolumeRequest.Marshal(b, m, deterministic)
}
func (dst *SetAvailableVolumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAvailableVolumeRequest.Merge(dst, src)
}
func (m *SetAvailableVolumeRequest) XXX_Size() int {
	return xxx_messageInfo_SetAvailableVolumeRequest.Size(m)
}
func (m *SetAvailableVolumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAvailableVolumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetAvailableVolumeRequest proto.InternalMessageInfo

func (m *SetAvailableVolumeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Volume)(nil), "n0stack.provisioning.Volume")
	proto.RegisterType((*VolumeSpec)(nil), "n0stack.provisioning.VolumeSpec")
	proto.RegisterType((*VolumeStatus)(nil), "n0stack.provisioning.VolumeStatus")
	proto.RegisterType((*CreateEmptyVolumeRequest)(nil), "n0stack.provisioning.CreateEmptyVolumeRequest")
	proto.RegisterType((*CreateVolumeWithDownloadingRequest)(nil), "n0stack.provisioning.CreateVolumeWithDownloadingRequest")
	proto.RegisterType((*ListVolumesRequest)(nil), "n0stack.provisioning.ListVolumesRequest")
	proto.RegisterType((*ListVolumesResponse)(nil), "n0stack.provisioning.ListVolumesResponse")
	proto.RegisterType((*GetVolumeRequest)(nil), "n0stack.provisioning.GetVolumeRequest")
	proto.RegisterType((*UpdateVolumeRequest)(nil), "n0stack.provisioning.UpdateVolumeRequest")
	proto.RegisterType((*DeleteVolumeRequest)(nil), "n0stack.provisioning.DeleteVolumeRequest")
	proto.RegisterType((*SetInuseVolumeRequest)(nil), "n0stack.provisioning.SetInuseVolumeRequest")
	proto.RegisterType((*SetAvailableVolumeRequest)(nil), "n0stack.provisioning.SetAvailableVolumeRequest")
	proto.RegisterEnum("n0stack.provisioning.VolumeStatus_VolumeState", VolumeStatus_VolumeState_name, VolumeStatus_VolumeState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VolumeService service

type VolumeServiceClient interface {
	CreateEmptyVolume(ctx context.Context, in *CreateEmptyVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	CreateVolumeWithDownloading(ctx context.Context, in *CreateVolumeWithDownloadingRequest, opts ...grpc.CallOption) (*Volume, error)
	ListVolumes(ctx context.Context, in *ListVolumesRequest, opts ...grpc.CallOption) (*ListVolumesResponse, error)
	GetVolume(ctx context.Context, in *GetVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	UpdateVolume(ctx context.Context, in *UpdateVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SetInuseVolume(ctx context.Context, in *SetInuseVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	SetAvailableVolume(ctx context.Context, in *SetAvailableVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
}

type volumeServiceClient struct {
	cc *grpc.ClientConn
}

func NewVolumeServiceClient(cc *grpc.ClientConn) VolumeServiceClient {
	return &volumeServiceClient{cc}
}

func (c *volumeServiceClient) CreateEmptyVolume(ctx context.Context, in *CreateEmptyVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.VolumeService/CreateEmptyVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) CreateVolumeWithDownloading(ctx context.Context, in *CreateVolumeWithDownloadingRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.VolumeService/CreateVolumeWithDownloading", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) ListVolumes(ctx context.Context, in *ListVolumesRequest, opts ...grpc.CallOption) (*ListVolumesResponse, error) {
	out := new(ListVolumesResponse)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.VolumeService/ListVolumes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) GetVolume(ctx context.Context, in *GetVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.VolumeService/GetVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) UpdateVolume(ctx context.Context, in *UpdateVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.VolumeService/UpdateVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.VolumeService/DeleteVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) SetInuseVolume(ctx context.Context, in *SetInuseVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.VolumeService/SetInuseVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) SetAvailableVolume(ctx context.Context, in *SetAvailableVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.VolumeService/SetAvailableVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VolumeService service

type VolumeServiceServer interface {
	CreateEmptyVolume(context.Context, *CreateEmptyVolumeRequest) (*Volume, error)
	CreateVolumeWithDownloading(context.Context, *CreateVolumeWithDownloadingRequest) (*Volume, error)
	ListVolumes(context.Context, *ListVolumesRequest) (*ListVolumesResponse, error)
	GetVolume(context.Context, *GetVolumeRequest) (*Volume, error)
	UpdateVolume(context.Context, *UpdateVolumeRequest) (*Volume, error)
	DeleteVolume(context.Context, *DeleteVolumeRequest) (*empty.Empty, error)
	SetInuseVolume(context.Context, *SetInuseVolumeRequest) (*Volume, error)
	SetAvailableVolume(context.Context, *SetAvailableVolumeRequest) (*Volume, error)
}

func RegisterVolumeServiceServer(s *grpc.Server, srv VolumeServiceServer) {
	s.RegisterService(&_VolumeService_serviceDesc, srv)
}

func _VolumeService_CreateEmptyVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmptyVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).CreateEmptyVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.VolumeService/CreateEmptyVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).CreateEmptyVolume(ctx, req.(*CreateEmptyVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_CreateVolumeWithDownloading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeWithDownloadingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).CreateVolumeWithDownloading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.VolumeService/CreateVolumeWithDownloading",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).CreateVolumeWithDownloading(ctx, req.(*CreateVolumeWithDownloadingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_ListVolumes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVolumesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).ListVolumes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.VolumeService/ListVolumes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).ListVolumes(ctx, req.(*ListVolumesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_GetVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).GetVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.VolumeService/GetVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).GetVolume(ctx, req.(*GetVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_UpdateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).UpdateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.VolumeService/UpdateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).UpdateVolume(ctx, req.(*UpdateVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_DeleteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).DeleteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.VolumeService/DeleteVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).DeleteVolume(ctx, req.(*DeleteVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_SetInuseVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInuseVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).SetInuseVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.VolumeService/SetInuseVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).SetInuseVolume(ctx, req.(*SetInuseVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_SetAvailableVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAvailableVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).SetAvailableVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.VolumeService/SetAvailableVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).SetAvailableVolume(ctx, req.(*SetAvailableVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VolumeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.provisioning.VolumeService",
	HandlerType: (*VolumeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmptyVolume",
			Handler:    _VolumeService_CreateEmptyVolume_Handler,
		},
		{
			MethodName: "CreateVolumeWithDownloading",
			Handler:    _VolumeService_CreateVolumeWithDownloading_Handler,
		},
		{
			MethodName: "ListVolumes",
			Handler:    _VolumeService_ListVolumes_Handler,
		},
		{
			MethodName: "GetVolume",
			Handler:    _VolumeService_GetVolume_Handler,
		},
		{
			MethodName: "UpdateVolume",
			Handler:    _VolumeService_UpdateVolume_Handler,
		},
		{
			MethodName: "DeleteVolume",
			Handler:    _VolumeService_DeleteVolume_Handler,
		},
		{
			MethodName: "SetInuseVolume",
			Handler:    _VolumeService_SetInuseVolume_Handler,
		},
		{
			MethodName: "SetAvailableVolume",
			Handler:    _VolumeService_SetAvailableVolume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provisioning/v0/volume.proto",
}

func init() {
	proto.RegisterFile("provisioning/v0/volume.proto", fileDescriptor_volume_b864912e45f02d3a)
}

var fileDescriptor_volume_b864912e45f02d3a = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0x61, 0x4f, 0xd3, 0x40,
	0x18, 0xc7, 0x37, 0x86, 0x83, 0x3d, 0x1b, 0x64, 0x3c, 0xa0, 0x99, 0x03, 0x23, 0x9e, 0x09, 0x81,
	0x10, 0x5b, 0x32, 0x8d, 0x51, 0x79, 0x35, 0xdc, 0x34, 0x8b, 0x63, 0x26, 0x5d, 0x06, 0xd1, 0xc4,
	0x2c, 0xdd, 0x76, 0x8e, 0xc6, 0xb6, 0x57, 0x7b, 0xd7, 0x11, 0x7c, 0xe3, 0x37, 0xf1, 0xb5, 0x7e,
	0x22, 0xbf, 0x8e, 0xe9, 0x5d, 0x87, 0x05, 0xba, 0x95, 0x57, 0xbc, 0xdb, 0x9e, 0xfb, 0x3f, 0xcf,
	0xff, 0x9f, 0xbb, 0xdf, 0xe5, 0x0a, 0x5b, 0x9e, 0xcf, 0x26, 0x16, 0xb7, 0x98, 0x6b, 0xb9, 0x63,
	0x7d, 0x72, 0xa0, 0x4f, 0x98, 0x1d, 0x38, 0x54, 0xf3, 0x7c, 0x26, 0x18, 0x6e, 0xb8, 0x07, 0x5c,
	0x98, 0xc3, 0x6f, 0x5a, 0x5c, 0x55, 0xdd, 0x1c, 0x33, 0x36, 0xb6, 0xa9, 0x2e, 0x35, 0x83, 0xe0,
	0xab, 0x4e, 0x1d, 0x4f, 0x5c, 0xa8, 0x96, 0xea, 0xda, 0xe4, 0x40, 0x77, 0xa8, 0x30, 0x47, 0xa6,
	0x30, 0x55, 0x89, 0xfc, 0xc9, 0x42, 0xfe, 0x44, 0x8e, 0xc5, 0x67, 0xb0, 0x3c, 0x5d, 0xac, 0x64,
	0xb7, 0xb3, 0xbb, 0xc5, 0xda, 0x9a, 0x36, 0xf5, 0x38, 0x8e, 0x16, 0x8c, 0x4b, 0x09, 0xbe, 0x80,
	0x45, 0xee, 0xd1, 0x61, 0x65, 0x41, 0x4a, 0xb7, 0xb5, 0xa4, 0x38, 0x9a, 0x1a, 0xdd, 0xf5, 0xe8,
	0xd0, 0x90, 0x6a, 0x7c, 0x03, 0x79, 0x2e, 0x4c, 0x11, 0xf0, 0x4a, 0x4e, 0xf6, 0x91, 0xb9, 0x7d,
	0x52, 0x69, 0x44, 0x1d, 0xc4, 0x00, 0xf8, 0x3f, 0x0f, 0x9f, 0xc2, 0x8a, 0x4f, 0xbf, 0x07, 0x94,
	0x8b, 0xfe, 0xe0, 0x42, 0x50, 0x2e, 0x33, 0x2f, 0x1a, 0xa5, 0xa8, 0x78, 0x14, 0xd6, 0xf0, 0x31,
	0x14, 0x6d, 0xcb, 0xb1, 0xa6, 0x92, 0x05, 0x29, 0x01, 0x59, 0x92, 0x02, 0xf2, 0x37, 0x0b, 0xa5,
	0xb8, 0x19, 0x36, 0xe0, 0x5e, 0x68, 0x47, 0xe5, 0xb8, 0xd5, 0x9a, 0x96, 0x9e, 0x2f, 0xf6, 0x87,
	0x1a, 0xaa, 0x19, 0x37, 0xa1, 0xe0, 0xb2, 0x11, 0xed, 0xbb, 0xa6, 0x43, 0xa5, 0x6b, 0xc1, 0x58,
	0x0e, 0x0b, 0x1d, 0xd3, 0xa1, 0xf8, 0x04, 0x4a, 0x5c, 0x30, 0xdf, 0x1c, 0x47, 0xeb, 0x39, 0xb9,
	0x5e, 0x8c, 0x6a, 0xa1, 0x84, 0xd4, 0xa1, 0x18, 0x9b, 0x8a, 0x00, 0xf9, 0x77, 0xf5, 0x56, 0xbb,
	0xd9, 0x28, 0x67, 0xb0, 0x08, 0x4b, 0xbd, 0xce, 0x87, 0xce, 0xc7, 0xd3, 0x4e, 0x39, 0x8b, 0x2b,
	0x50, 0xa8, 0x9f, 0xd4, 0x5b, 0xed, 0xfa, 0x51, 0xbb, 0x59, 0x5e, 0x08, 0x75, 0xad, 0x4e, 0xbf,
	0xd7, 0x6d, 0x96, 0x73, 0xe4, 0x27, 0x54, 0xde, 0xfa, 0xd4, 0x14, 0xb4, 0x19, 0x12, 0xa0, 0xa6,
	0x19, 0x6a, 0x6f, 0xee, 0xe4, 0xa8, 0xc9, 0xef, 0x2c, 0x10, 0x95, 0x40, 0x2d, 0x9d, 0x5a, 0xe2,
	0xac, 0xc1, 0xce, 0x5d, 0x9b, 0x99, 0x23, 0xcb, 0x1d, 0xdf, 0x65, 0x16, 0x7c, 0x04, 0xc0, 0x59,
	0xe0, 0x0f, 0x69, 0x3f, 0xf0, 0xed, 0x68, 0xc3, 0x0b, 0xaa, 0xd2, 0xf3, 0x6d, 0xb2, 0x01, 0xd8,
	0xb6, 0xb8, 0x50, 0x6d, 0x3c, 0x4a, 0x46, 0x8e, 0x61, 0xfd, 0x4a, 0x95, 0x7b, 0xcc, 0xe5, 0x14,
	0x5f, 0xc2, 0x92, 0xba, 0x88, 0x21, 0x72, 0xb9, 0xdd, 0x62, 0x6d, 0x6b, 0x5e, 0x08, 0x63, 0x2a,
	0x26, 0x3b, 0x50, 0x7e, 0x4f, 0xc5, 0xd5, 0x83, 0x40, 0x58, 0x94, 0x08, 0x64, 0x65, 0x22, 0xf9,
	0x9b, 0xfc, 0x80, 0xf5, 0x9e, 0x37, 0xba, 0xdc, 0xb6, 0x3b, 0x3d, 0xb3, 0x3d, 0x58, 0x6f, 0x50,
	0x9b, 0x5e, 0xf7, 0x4e, 0x8a, 0xb9, 0x0f, 0xf7, 0xbb, 0x54, 0xb4, 0xdc, 0x80, 0xdf, 0x42, 0xac,
	0xc3, 0xc3, 0x2e, 0x15, 0xf5, 0x89, 0x69, 0xd9, 0xe6, 0xc0, 0x4e, 0x6f, 0xa8, 0xfd, 0xca, 0xc3,
	0x4a, 0x94, 0x8e, 0xfa, 0x13, 0x6b, 0x48, 0x71, 0x08, 0x6b, 0x37, 0x78, 0xc6, 0x19, 0xd7, 0x73,
	0x16, 0xf8, 0xd5, 0xb9, 0x47, 0x45, 0x32, 0x78, 0x0e, 0x9b, 0x73, 0x90, 0xc5, 0x57, 0xf3, 0xec,
	0xe6, 0x51, 0x9e, 0x6a, 0x3c, 0x82, 0x62, 0x8c, 0x35, 0xdc, 0x4d, 0x96, 0xdf, 0x84, 0xb4, 0xba,
	0x77, 0x0b, 0xa5, 0x02, 0x97, 0x64, 0xb0, 0x0b, 0x85, 0x4b, 0x04, 0x71, 0x27, 0xb9, 0xf3, 0x3a,
	0xa3, 0xa9, 0xd1, 0x3f, 0x41, 0x29, 0xce, 0x2b, 0xce, 0x48, 0x94, 0xc0, 0x74, 0xea, 0xe8, 0x2e,
	0x94, 0xe2, 0x38, 0xce, 0x1a, 0x9d, 0x80, 0x6c, 0xf5, 0x81, 0xa6, 0x5e, 0x42, 0x6d, 0xfa, 0x12,
	0x6a, 0x12, 0x07, 0x92, 0xc1, 0x2f, 0xb0, 0x7a, 0x15, 0x5c, 0xdc, 0x4f, 0x1e, 0x9b, 0x88, 0x77,
	0x6a, 0x66, 0x0a, 0x78, 0x13, 0x75, 0xd4, 0x67, 0x5a, 0x24, 0x5f, 0x8a, 0x34, 0x9b, 0xa3, 0xc3,
	0xcf, 0xaf, 0xc7, 0x96, 0x38, 0x0b, 0x06, 0xda, 0x90, 0x39, 0x7a, 0xa4, 0x55, 0xcf, 0xbe, 0x36,
	0x66, 0xfa, 0xb5, 0x4f, 0x87, 0x43, 0x2f, 0x5e, 0x18, 0xe4, 0xa5, 0xee, 0xf9, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6b, 0xa9, 0x88, 0x72, 0x62, 0x08, 0x00, 0x00,
}
