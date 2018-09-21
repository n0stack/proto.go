// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pool/v0/node.proto

package ppool // import "github.com/n0stack/proto.go/pool/v0"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import v01 "github.com/n0stack/proto.go/resource/v0"
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

type NodeStatus_NodeState int32

const (
	NodeStatus_NotReady NodeStatus_NodeState = 0
	NodeStatus_Ready    NodeStatus_NodeState = 1
)

var NodeStatus_NodeState_name = map[int32]string{
	0: "NotReady",
	1: "Ready",
}
var NodeStatus_NodeState_value = map[string]int32{
	"NotReady": 0,
	"Ready":    1,
}

func (x NodeStatus_NodeState) String() string {
	return proto.EnumName(NodeStatus_NodeState_name, int32(x))
}
func (NodeStatus_NodeState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{2, 0}
}

type Node struct {
	Metadata             *v0.Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec                 *NodeSpec    `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status               *NodeStatus  `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetSpec() *NodeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Node) GetStatus() *NodeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type NodeSpec struct {
	Serial               string   `protobuf:"bytes,1,opt,name=serial" json:"serial,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	IpmiAddress          string   `protobuf:"bytes,3,opt,name=ipmi_address,json=ipmiAddress" json:"ipmi_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeSpec) Reset()         { *m = NodeSpec{} }
func (m *NodeSpec) String() string { return proto.CompactTextString(m) }
func (*NodeSpec) ProtoMessage()    {}
func (*NodeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{1}
}
func (m *NodeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeSpec.Unmarshal(m, b)
}
func (m *NodeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeSpec.Marshal(b, m, deterministic)
}
func (dst *NodeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSpec.Merge(dst, src)
}
func (m *NodeSpec) XXX_Size() int {
	return xxx_messageInfo_NodeSpec.Size(m)
}
func (m *NodeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSpec proto.InternalMessageInfo

func (m *NodeSpec) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *NodeSpec) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NodeSpec) GetIpmiAddress() string {
	if m != nil {
		return m.IpmiAddress
	}
	return ""
}

type NodeStatus struct {
	State                NodeStatus_NodeState    `protobuf:"varint,1,opt,name=state,enum=n0stack.pool.NodeStatus_NodeState" json:"state,omitempty"`
	Compute              *v01.Compute            `protobuf:"bytes,2,opt,name=compute" json:"compute,omitempty"`
	ReservedComputes     map[string]*v01.Compute `protobuf:"bytes,3,rep,name=reserved_computes,json=reservedComputes" json:"reserved_computes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Storage              *v01.Storage            `protobuf:"bytes,4,opt,name=storage" json:"storage,omitempty"`
	ReservedStorages     map[string]*v01.Storage `protobuf:"bytes,5,rep,name=reserved_storages,json=reservedStorages" json:"reserved_storages,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NodeStatus) Reset()         { *m = NodeStatus{} }
func (m *NodeStatus) String() string { return proto.CompactTextString(m) }
func (*NodeStatus) ProtoMessage()    {}
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{2}
}
func (m *NodeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStatus.Unmarshal(m, b)
}
func (m *NodeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStatus.Marshal(b, m, deterministic)
}
func (dst *NodeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStatus.Merge(dst, src)
}
func (m *NodeStatus) XXX_Size() int {
	return xxx_messageInfo_NodeStatus.Size(m)
}
func (m *NodeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStatus proto.InternalMessageInfo

func (m *NodeStatus) GetState() NodeStatus_NodeState {
	if m != nil {
		return m.State
	}
	return NodeStatus_NotReady
}

func (m *NodeStatus) GetCompute() *v01.Compute {
	if m != nil {
		return m.Compute
	}
	return nil
}

func (m *NodeStatus) GetReservedComputes() map[string]*v01.Compute {
	if m != nil {
		return m.ReservedComputes
	}
	return nil
}

func (m *NodeStatus) GetStorage() *v01.Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *NodeStatus) GetReservedStorages() map[string]*v01.Storage {
	if m != nil {
		return m.ReservedStorages
	}
	return nil
}

type ListNodesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNodesRequest) Reset()         { *m = ListNodesRequest{} }
func (m *ListNodesRequest) String() string { return proto.CompactTextString(m) }
func (*ListNodesRequest) ProtoMessage()    {}
func (*ListNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{3}
}
func (m *ListNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNodesRequest.Unmarshal(m, b)
}
func (m *ListNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNodesRequest.Marshal(b, m, deterministic)
}
func (dst *ListNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodesRequest.Merge(dst, src)
}
func (m *ListNodesRequest) XXX_Size() int {
	return xxx_messageInfo_ListNodesRequest.Size(m)
}
func (m *ListNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodesRequest proto.InternalMessageInfo

type ListNodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNodesResponse) Reset()         { *m = ListNodesResponse{} }
func (m *ListNodesResponse) String() string { return proto.CompactTextString(m) }
func (*ListNodesResponse) ProtoMessage()    {}
func (*ListNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{4}
}
func (m *ListNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNodesResponse.Unmarshal(m, b)
}
func (m *ListNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNodesResponse.Marshal(b, m, deterministic)
}
func (dst *ListNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNodesResponse.Merge(dst, src)
}
func (m *ListNodesResponse) XXX_Size() int {
	return xxx_messageInfo_ListNodesResponse.Size(m)
}
func (m *ListNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNodesResponse proto.InternalMessageInfo

func (m *ListNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type GetNodeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeRequest) Reset()         { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()    {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{5}
}
func (m *GetNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeRequest.Unmarshal(m, b)
}
func (m *GetNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeRequest.Marshal(b, m, deterministic)
}
func (dst *GetNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeRequest.Merge(dst, src)
}
func (m *GetNodeRequest) XXX_Size() int {
	return xxx_messageInfo_GetNodeRequest.Size(m)
}
func (m *GetNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeRequest proto.InternalMessageInfo

func (m *GetNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplyNodeRequest struct {
	Metadata             *v0.Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec                 *NodeSpec    `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ApplyNodeRequest) Reset()         { *m = ApplyNodeRequest{} }
func (m *ApplyNodeRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyNodeRequest) ProtoMessage()    {}
func (*ApplyNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{6}
}
func (m *ApplyNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyNodeRequest.Unmarshal(m, b)
}
func (m *ApplyNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyNodeRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyNodeRequest.Merge(dst, src)
}
func (m *ApplyNodeRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyNodeRequest.Size(m)
}
func (m *ApplyNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyNodeRequest proto.InternalMessageInfo

func (m *ApplyNodeRequest) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ApplyNodeRequest) GetSpec() *NodeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type DeleteNodeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNodeRequest) Reset()         { *m = DeleteNodeRequest{} }
func (m *DeleteNodeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNodeRequest) ProtoMessage()    {}
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{7}
}
func (m *DeleteNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNodeRequest.Unmarshal(m, b)
}
func (m *DeleteNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNodeRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNodeRequest.Merge(dst, src)
}
func (m *DeleteNodeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNodeRequest.Size(m)
}
func (m *DeleteNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNodeRequest proto.InternalMessageInfo

func (m *DeleteNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReserveComputeRequest struct {
	// optional
	Name        string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ComputeName string            `protobuf:"bytes,2,opt,name=compute_name,json=computeName" json:"compute_name,omitempty"`
	Annotations map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// CPU
	// request_
	// limit_
	Vcpus uint32 `protobuf:"varint,4,opt,name=vcpus" json:"vcpus,omitempty"`
	// Memory
	// request_memory_bytes
	// limit_memory_bytes
	MemoryBytes          uint64   `protobuf:"varint,5,opt,name=memory_bytes,json=memoryBytes" json:"memory_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReserveComputeRequest) Reset()         { *m = ReserveComputeRequest{} }
func (m *ReserveComputeRequest) String() string { return proto.CompactTextString(m) }
func (*ReserveComputeRequest) ProtoMessage()    {}
func (*ReserveComputeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{8}
}
func (m *ReserveComputeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveComputeRequest.Unmarshal(m, b)
}
func (m *ReserveComputeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveComputeRequest.Marshal(b, m, deterministic)
}
func (dst *ReserveComputeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveComputeRequest.Merge(dst, src)
}
func (m *ReserveComputeRequest) XXX_Size() int {
	return xxx_messageInfo_ReserveComputeRequest.Size(m)
}
func (m *ReserveComputeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveComputeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveComputeRequest proto.InternalMessageInfo

func (m *ReserveComputeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReserveComputeRequest) GetComputeName() string {
	if m != nil {
		return m.ComputeName
	}
	return ""
}

func (m *ReserveComputeRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ReserveComputeRequest) GetVcpus() uint32 {
	if m != nil {
		return m.Vcpus
	}
	return 0
}

func (m *ReserveComputeRequest) GetMemoryBytes() uint64 {
	if m != nil {
		return m.MemoryBytes
	}
	return 0
}

type ReleaseComputeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ComputeName          string   `protobuf:"bytes,2,opt,name=compute_name,json=computeName" json:"compute_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseComputeRequest) Reset()         { *m = ReleaseComputeRequest{} }
func (m *ReleaseComputeRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseComputeRequest) ProtoMessage()    {}
func (*ReleaseComputeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{9}
}
func (m *ReleaseComputeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseComputeRequest.Unmarshal(m, b)
}
func (m *ReleaseComputeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseComputeRequest.Marshal(b, m, deterministic)
}
func (dst *ReleaseComputeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseComputeRequest.Merge(dst, src)
}
func (m *ReleaseComputeRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseComputeRequest.Size(m)
}
func (m *ReleaseComputeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseComputeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseComputeRequest proto.InternalMessageInfo

func (m *ReleaseComputeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReleaseComputeRequest) GetComputeName() string {
	if m != nil {
		return m.ComputeName
	}
	return ""
}

type ReserveStorageRequest struct {
	// optional
	Name                 string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StorageName          string            `protobuf:"bytes,2,opt,name=storage_name,json=storageName" json:"storage_name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Bytes                uint64            `protobuf:"varint,4,opt,name=bytes" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReserveStorageRequest) Reset()         { *m = ReserveStorageRequest{} }
func (m *ReserveStorageRequest) String() string { return proto.CompactTextString(m) }
func (*ReserveStorageRequest) ProtoMessage()    {}
func (*ReserveStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{10}
}
func (m *ReserveStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveStorageRequest.Unmarshal(m, b)
}
func (m *ReserveStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveStorageRequest.Marshal(b, m, deterministic)
}
func (dst *ReserveStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveStorageRequest.Merge(dst, src)
}
func (m *ReserveStorageRequest) XXX_Size() int {
	return xxx_messageInfo_ReserveStorageRequest.Size(m)
}
func (m *ReserveStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveStorageRequest proto.InternalMessageInfo

func (m *ReserveStorageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReserveStorageRequest) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

func (m *ReserveStorageRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ReserveStorageRequest) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

type ReleaseStorageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StorageName          string   `protobuf:"bytes,2,opt,name=storage_name,json=storageName" json:"storage_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseStorageRequest) Reset()         { *m = ReleaseStorageRequest{} }
func (m *ReleaseStorageRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseStorageRequest) ProtoMessage()    {}
func (*ReleaseStorageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_23bf8e8d1486b69e, []int{11}
}
func (m *ReleaseStorageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseStorageRequest.Unmarshal(m, b)
}
func (m *ReleaseStorageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseStorageRequest.Marshal(b, m, deterministic)
}
func (dst *ReleaseStorageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseStorageRequest.Merge(dst, src)
}
func (m *ReleaseStorageRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseStorageRequest.Size(m)
}
func (m *ReleaseStorageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseStorageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseStorageRequest proto.InternalMessageInfo

func (m *ReleaseStorageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReleaseStorageRequest) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

func init() {
	proto.RegisterType((*Node)(nil), "n0stack.pool.Node")
	proto.RegisterType((*NodeSpec)(nil), "n0stack.pool.NodeSpec")
	proto.RegisterType((*NodeStatus)(nil), "n0stack.pool.NodeStatus")
	proto.RegisterMapType((map[string]*v01.Compute)(nil), "n0stack.pool.NodeStatus.ReservedComputesEntry")
	proto.RegisterMapType((map[string]*v01.Storage)(nil), "n0stack.pool.NodeStatus.ReservedStoragesEntry")
	proto.RegisterType((*ListNodesRequest)(nil), "n0stack.pool.ListNodesRequest")
	proto.RegisterType((*ListNodesResponse)(nil), "n0stack.pool.ListNodesResponse")
	proto.RegisterType((*GetNodeRequest)(nil), "n0stack.pool.GetNodeRequest")
	proto.RegisterType((*ApplyNodeRequest)(nil), "n0stack.pool.ApplyNodeRequest")
	proto.RegisterType((*DeleteNodeRequest)(nil), "n0stack.pool.DeleteNodeRequest")
	proto.RegisterType((*ReserveComputeRequest)(nil), "n0stack.pool.ReserveComputeRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.ReserveComputeRequest.AnnotationsEntry")
	proto.RegisterType((*ReleaseComputeRequest)(nil), "n0stack.pool.ReleaseComputeRequest")
	proto.RegisterType((*ReserveStorageRequest)(nil), "n0stack.pool.ReserveStorageRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.pool.ReserveStorageRequest.AnnotationsEntry")
	proto.RegisterType((*ReleaseStorageRequest)(nil), "n0stack.pool.ReleaseStorageRequest")
	proto.RegisterEnum("n0stack.pool.NodeStatus_NodeState", NodeStatus_NodeState_name, NodeStatus_NodeState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeService service

type NodeServiceClient interface {
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error)
	ApplyNode(ctx context.Context, in *ApplyNodeRequest, opts ...grpc.CallOption) (*Node, error)
	DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ReserveCompute(ctx context.Context, in *ReserveComputeRequest, opts ...grpc.CallOption) (*v01.Compute, error)
	ReleaseCompute(ctx context.Context, in *ReleaseComputeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ReserveStorage(ctx context.Context, in *ReserveStorageRequest, opts ...grpc.CallOption) (*v01.Storage, error)
	ReleaseStorage(ctx context.Context, in *ReleaseStorageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type nodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeServiceClient(cc *grpc.ClientConn) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := grpc.Invoke(ctx, "/n0stack.pool.NodeService/ListNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/n0stack.pool.NodeService/GetNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ApplyNode(ctx context.Context, in *ApplyNodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/n0stack.pool.NodeService/ApplyNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.pool.NodeService/DeleteNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ReserveCompute(ctx context.Context, in *ReserveComputeRequest, opts ...grpc.CallOption) (*v01.Compute, error) {
	out := new(v01.Compute)
	err := grpc.Invoke(ctx, "/n0stack.pool.NodeService/ReserveCompute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ReleaseCompute(ctx context.Context, in *ReleaseComputeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.pool.NodeService/ReleaseCompute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ReserveStorage(ctx context.Context, in *ReserveStorageRequest, opts ...grpc.CallOption) (*v01.Storage, error) {
	out := new(v01.Storage)
	err := grpc.Invoke(ctx, "/n0stack.pool.NodeService/ReserveStorage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) ReleaseStorage(ctx context.Context, in *ReleaseStorageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.pool.NodeService/ReleaseStorage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeService service

type NodeServiceServer interface {
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
	GetNode(context.Context, *GetNodeRequest) (*Node, error)
	ApplyNode(context.Context, *ApplyNodeRequest) (*Node, error)
	DeleteNode(context.Context, *DeleteNodeRequest) (*empty.Empty, error)
	ReserveCompute(context.Context, *ReserveComputeRequest) (*v01.Compute, error)
	ReleaseCompute(context.Context, *ReleaseComputeRequest) (*empty.Empty, error)
	ReserveStorage(context.Context, *ReserveStorageRequest) (*v01.Storage, error)
	ReleaseStorage(context.Context, *ReleaseStorageRequest) (*empty.Empty, error)
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NodeService/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NodeService/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ApplyNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ApplyNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NodeService/ApplyNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ApplyNode(ctx, req.(*ApplyNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NodeService/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).DeleteNode(ctx, req.(*DeleteNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ReserveCompute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ReserveCompute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NodeService/ReserveCompute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ReserveCompute(ctx, req.(*ReserveComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ReleaseCompute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ReleaseCompute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NodeService/ReleaseCompute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ReleaseCompute(ctx, req.(*ReleaseComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ReserveStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ReserveStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NodeService/ReserveStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ReserveStorage(ctx, req.(*ReserveStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_ReleaseStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ReleaseStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.pool.NodeService/ReleaseStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ReleaseStorage(ctx, req.(*ReleaseStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.pool.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNodes",
			Handler:    _NodeService_ListNodes_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _NodeService_GetNode_Handler,
		},
		{
			MethodName: "ApplyNode",
			Handler:    _NodeService_ApplyNode_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _NodeService_DeleteNode_Handler,
		},
		{
			MethodName: "ReserveCompute",
			Handler:    _NodeService_ReserveCompute_Handler,
		},
		{
			MethodName: "ReleaseCompute",
			Handler:    _NodeService_ReleaseCompute_Handler,
		},
		{
			MethodName: "ReserveStorage",
			Handler:    _NodeService_ReserveStorage_Handler,
		},
		{
			MethodName: "ReleaseStorage",
			Handler:    _NodeService_ReleaseStorage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pool/v0/node.proto",
}

func init() { proto.RegisterFile("pool/v0/node.proto", fileDescriptor_node_23bf8e8d1486b69e) }

var fileDescriptor_node_23bf8e8d1486b69e = []byte{
	// 816 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x18, 0x8d, 0x1b, 0xa7, 0x6d, 0xbe, 0x94, 0x2a, 0x19, 0x95, 0xca, 0x35, 0x88, 0x1f, 0x53, 0x89,
	0x00, 0xc2, 0x8e, 0x52, 0x2e, 0x2a, 0x50, 0x91, 0x52, 0xa8, 0x7a, 0x03, 0x41, 0x9a, 0x4a, 0x5c,
	0x80, 0x44, 0xe4, 0x38, 0x43, 0x36, 0x6a, 0xec, 0xf1, 0x7a, 0xc6, 0x91, 0xf2, 0x1c, 0xfb, 0x00,
	0xfb, 0x4a, 0xfb, 0x2c, 0x7b, 0xbd, 0x17, 0xab, 0xf9, 0xb1, 0x63, 0xa7, 0xf6, 0x36, 0xbb, 0xea,
	0xde, 0xd9, 0xdf, 0x77, 0xe6, 0xcc, 0xc9, 0x99, 0xef, 0x38, 0x03, 0x28, 0xa6, 0x74, 0xe9, 0xad,
	0x06, 0x5e, 0x44, 0x67, 0xc4, 0x8d, 0x13, 0xca, 0x29, 0x3a, 0x8a, 0x06, 0x8c, 0xfb, 0xc1, 0xbd,
	0x2b, 0x7a, 0xf6, 0x67, 0x73, 0x4a, 0xe7, 0x4b, 0xe2, 0xc9, 0xde, 0x34, 0xfd, 0xdf, 0x23, 0x61,
	0xcc, 0xd7, 0x0a, 0x6a, 0xf7, 0x56, 0x03, 0x2f, 0x24, 0xdc, 0x9f, 0xf9, 0xdc, 0xd7, 0xa5, 0xb3,
	0x84, 0x30, 0x9a, 0x26, 0x01, 0x11, 0xac, 0x01, 0x0d, 0xe3, 0x94, 0x93, 0xaa, 0x16, 0xe3, 0x34,
	0xf1, 0xe7, 0xba, 0xe5, 0xbc, 0x30, 0xc0, 0x1c, 0xd3, 0x19, 0x41, 0x3f, 0xc2, 0x61, 0x46, 0x68,
	0x19, 0x5f, 0x19, 0xfd, 0xce, 0xb0, 0xe7, 0x66, 0x7a, 0xfe, 0xd4, 0x0d, 0x9c, 0x43, 0xd0, 0xf7,
	0x60, 0xb2, 0x98, 0x04, 0xd6, 0x9e, 0x84, 0x9e, 0xba, 0x45, 0xe9, 0xae, 0x20, 0xbc, 0x8b, 0x49,
	0x80, 0x25, 0x06, 0x0d, 0x60, 0x9f, 0x71, 0x9f, 0xa7, 0xcc, 0x6a, 0x4a, 0xb4, 0x55, 0x81, 0x96,
	0x7d, 0xac, 0x71, 0xce, 0x04, 0x0e, 0x33, 0x0e, 0x74, 0x0a, 0xfb, 0x8c, 0x24, 0x0b, 0x7f, 0x29,
	0x65, 0xb5, 0xb1, 0x7e, 0x43, 0x16, 0x1c, 0xf8, 0xb3, 0x59, 0x42, 0x18, 0x93, 0x22, 0xda, 0x38,
	0x7b, 0x45, 0x5f, 0xc3, 0xd1, 0x22, 0x0e, 0x17, 0x93, 0xac, 0xdd, 0x94, 0xed, 0x8e, 0xa8, 0x8d,
	0x54, 0xc9, 0x79, 0x65, 0x02, 0x6c, 0xf6, 0x45, 0x97, 0xd0, 0x12, 0x3b, 0x13, 0xb9, 0xc5, 0xf1,
	0xd0, 0xa9, 0x13, 0x98, 0x3f, 0x12, 0xac, 0x16, 0xa0, 0x0b, 0x38, 0xd0, 0x5e, 0x6b, 0x2b, 0xce,
	0xf2, 0xb5, 0x99, 0xe9, 0xee, 0x6f, 0x0a, 0x80, 0x33, 0x24, 0xfa, 0x17, 0x7a, 0x09, 0x61, 0x24,
	0x59, 0x91, 0xd9, 0x44, 0xd7, 0x84, 0xca, 0x66, 0xbf, 0x33, 0x74, 0x6b, 0xb7, 0xc6, 0x7a, 0x85,
	0x66, 0x63, 0x37, 0x11, 0x4f, 0xd6, 0xb8, 0x9b, 0x6c, 0x95, 0x85, 0x22, 0x7d, 0xc4, 0x96, 0x59,
	0xa7, 0xe8, 0x4e, 0x01, 0x70, 0x86, 0x2c, 0x29, 0xd2, 0x35, 0x66, 0xb5, 0x76, 0x54, 0xa4, 0xd9,
	0xb6, 0x15, 0x65, 0x65, 0xfb, 0x3f, 0xf8, 0xb4, 0x52, 0x3c, 0xea, 0x42, 0xf3, 0x9e, 0xac, 0xf5,
	0xb9, 0x8a, 0x47, 0xe4, 0x41, 0x6b, 0xe5, 0x2f, 0xd3, 0x1d, 0xcc, 0x54, 0xb8, 0x9f, 0xf7, 0x2e,
	0x8d, 0x22, 0x7f, 0x49, 0xca, 0x07, 0xf1, 0x67, 0xd6, 0x6c, 0xf8, 0x9d, 0x73, 0x68, 0xe7, 0xe7,
	0x8e, 0x8e, 0xc4, 0x68, 0x72, 0x4c, 0xfc, 0xd9, 0xba, 0xdb, 0x40, 0x6d, 0x68, 0xa9, 0x47, 0xc3,
	0x41, 0xd0, 0xfd, 0x63, 0xc1, 0xb8, 0x40, 0x32, 0x4c, 0x9e, 0xa7, 0x84, 0x71, 0xe7, 0x0a, 0x7a,
	0x85, 0x1a, 0x8b, 0x69, 0xc4, 0x08, 0xea, 0x43, 0x4b, 0x84, 0x9e, 0x59, 0x86, 0xf4, 0x17, 0x3d,
	0xf4, 0x17, 0x2b, 0x80, 0x73, 0x0e, 0xc7, 0xb7, 0x44, 0xae, 0xd6, 0x84, 0x08, 0x81, 0x19, 0xf9,
	0x21, 0xd1, 0x3f, 0x49, 0x3e, 0x3b, 0x21, 0x74, 0x47, 0x71, 0xbc, 0x5c, 0x17, 0x71, 0x1f, 0x2f,
	0xcd, 0xce, 0xb7, 0xd0, 0xfb, 0x9d, 0x2c, 0x09, 0x27, 0x8f, 0xe9, 0x7a, 0xb9, 0x97, 0x9f, 0x4b,
	0x76, 0x68, 0xf5, 0x68, 0x11, 0x5a, 0x1d, 0x85, 0x89, 0xec, 0xa9, 0x4c, 0x77, 0x74, 0x6d, 0x2c,
	0x20, 0x7f, 0x43, 0xc7, 0x8f, 0x22, 0xca, 0x7d, 0xbe, 0xa0, 0x51, 0x16, 0x98, 0x9f, 0xca, 0x62,
	0x2b, 0x37, 0x74, 0x47, 0x9b, 0x65, 0x6a, 0x48, 0x8b, 0x44, 0xe8, 0x04, 0x5a, 0xab, 0x20, 0x4e,
	0x99, 0xcc, 0xcb, 0x27, 0x58, 0xbd, 0x08, 0x41, 0x21, 0x09, 0x69, 0xb2, 0x9e, 0x4c, 0xd7, 0x5c,
	0xa6, 0xc1, 0xe8, 0x9b, 0xb8, 0xa3, 0x6a, 0xd7, 0xa2, 0x64, 0xff, 0x0a, 0xdd, 0x6d, 0xe6, 0x8a,
	0x99, 0x3b, 0x29, 0xce, 0x5c, 0xbb, 0x38, 0x58, 0x63, 0x61, 0xd0, 0x92, 0xf8, 0xec, 0x69, 0x0c,
	0x72, 0xde, 0x18, 0xb9, 0xe3, 0xd9, 0x18, 0xbf, 0x9b, 0x50, 0x47, 0xbd, 0x44, 0xa8, 0x6b, 0xef,
	0xeb, 0x78, 0x79, 0xc3, 0xc7, 0x1d, 0x57, 0xa6, 0x9a, 0xd2, 0x54, 0xf5, 0xf2, 0x84, 0x76, 0x3e,
	0xc9, 0xaf, 0x1f, 0xbe, 0x36, 0xa1, 0x23, 0x87, 0x9f, 0x24, 0xab, 0x45, 0x40, 0xd0, 0x18, 0xda,
	0x79, 0x9a, 0xd1, 0x17, 0x65, 0x17, 0xb6, 0xa3, 0x6f, 0x7f, 0x59, 0xdb, 0x57, 0x9f, 0x01, 0xa7,
	0x81, 0xae, 0xe0, 0x40, 0xc7, 0x1b, 0x7d, 0x5e, 0x46, 0x97, 0x53, 0x6f, 0x57, 0x7c, 0x22, 0x9c,
	0x06, 0x1a, 0x41, 0x3b, 0xcf, 0xfd, 0xb6, 0x9c, 0xed, 0x0f, 0x42, 0x0d, 0xc5, 0x2d, 0xc0, 0x26,
	0xcb, 0x68, 0x4b, 0xf2, 0x83, 0x94, 0xdb, 0xa7, 0xae, 0xba, 0x93, 0xb8, 0xd9, 0x9d, 0xc4, 0xbd,
	0x11, 0x77, 0x12, 0xa7, 0x81, 0x30, 0x1c, 0x97, 0x93, 0x87, 0xbe, 0xd9, 0x21, 0x97, 0x76, 0xfd,
	0xf7, 0xdd, 0x69, 0xa0, 0xbf, 0x04, 0x67, 0x31, 0x1d, 0x0f, 0x39, 0x2b, 0xb2, 0xb3, 0x93, 0x48,
	0x3d, 0x1f, 0x35, 0x22, 0xcb, 0xd3, 0x63, 0xd7, 0xff, 0x49, 0x94, 0x44, 0xd6, 0x72, 0x56, 0x4c,
	0x64, 0xbd, 0xc8, 0xeb, 0x1f, 0xfe, 0xf9, 0x6e, 0xbe, 0xe0, 0xcf, 0xd2, 0xa9, 0x1b, 0xd0, 0xd0,
	0xd3, 0x54, 0xea, 0x12, 0xe8, 0xce, 0xa9, 0xa7, 0xaf, 0x8d, 0xbf, 0xc4, 0xe2, 0x61, 0xba, 0x2f,
	0xeb, 0x17, 0x6f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xe0, 0x6f, 0xf2, 0x4e, 0x0a, 0x00, 0x00,
}
