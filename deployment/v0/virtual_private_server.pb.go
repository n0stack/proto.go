// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deployment/v0/virtual_private_server.proto

package pdeployment // import "github.com/n0stack/proto.go/deployment/v0"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import v0 "github.com/n0stack/proto.go/provisioning/v0"

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

type VirtualPrivateServer struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Version              uint64            `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	LimitCpuMilliCore    uint32            `protobuf:"varint,10,opt,name=limit_cpu_milli_core,json=limitCpuMilliCore" json:"limit_cpu_milli_core,omitempty"`
	LimitMemoryBytes     uint64            `protobuf:"varint,11,opt,name=limit_memory_bytes,json=limitMemoryBytes" json:"limit_memory_bytes,omitempty"`
	NetworkName          string            `protobuf:"bytes,12,opt,name=network_name,json=networkName" json:"network_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *VirtualPrivateServer) Reset()         { *m = VirtualPrivateServer{} }
func (m *VirtualPrivateServer) String() string { return proto.CompactTextString(m) }
func (*VirtualPrivateServer) ProtoMessage()    {}
func (*VirtualPrivateServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_private_server_d376265387d5e975, []int{0}
}
func (m *VirtualPrivateServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualPrivateServer.Unmarshal(m, b)
}
func (m *VirtualPrivateServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualPrivateServer.Marshal(b, m, deterministic)
}
func (dst *VirtualPrivateServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualPrivateServer.Merge(dst, src)
}
func (m *VirtualPrivateServer) XXX_Size() int {
	return xxx_messageInfo_VirtualPrivateServer.Size(m)
}
func (m *VirtualPrivateServer) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualPrivateServer.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualPrivateServer proto.InternalMessageInfo

func (m *VirtualPrivateServer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualPrivateServer) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *VirtualPrivateServer) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VirtualPrivateServer) GetLimitCpuMilliCore() uint32 {
	if m != nil {
		return m.LimitCpuMilliCore
	}
	return 0
}

func (m *VirtualPrivateServer) GetLimitMemoryBytes() uint64 {
	if m != nil {
		return m.LimitMemoryBytes
	}
	return 0
}

func (m *VirtualPrivateServer) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

type ListVirtualPrivateServersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVirtualPrivateServersRequest) Reset()         { *m = ListVirtualPrivateServersRequest{} }
func (m *ListVirtualPrivateServersRequest) String() string { return proto.CompactTextString(m) }
func (*ListVirtualPrivateServersRequest) ProtoMessage()    {}
func (*ListVirtualPrivateServersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_private_server_d376265387d5e975, []int{1}
}
func (m *ListVirtualPrivateServersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVirtualPrivateServersRequest.Unmarshal(m, b)
}
func (m *ListVirtualPrivateServersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVirtualPrivateServersRequest.Marshal(b, m, deterministic)
}
func (dst *ListVirtualPrivateServersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVirtualPrivateServersRequest.Merge(dst, src)
}
func (m *ListVirtualPrivateServersRequest) XXX_Size() int {
	return xxx_messageInfo_ListVirtualPrivateServersRequest.Size(m)
}
func (m *ListVirtualPrivateServersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVirtualPrivateServersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListVirtualPrivateServersRequest proto.InternalMessageInfo

type ListVirtualPrivateServersResponse struct {
	VirtualPrivateServers []*VirtualPrivateServer `protobuf:"bytes,1,rep,name=virtual_private_servers,json=virtualPrivateServers" json:"virtual_private_servers,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                `json:"-"`
	XXX_unrecognized      []byte                  `json:"-"`
	XXX_sizecache         int32                   `json:"-"`
}

func (m *ListVirtualPrivateServersResponse) Reset()         { *m = ListVirtualPrivateServersResponse{} }
func (m *ListVirtualPrivateServersResponse) String() string { return proto.CompactTextString(m) }
func (*ListVirtualPrivateServersResponse) ProtoMessage()    {}
func (*ListVirtualPrivateServersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_private_server_d376265387d5e975, []int{2}
}
func (m *ListVirtualPrivateServersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVirtualPrivateServersResponse.Unmarshal(m, b)
}
func (m *ListVirtualPrivateServersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVirtualPrivateServersResponse.Marshal(b, m, deterministic)
}
func (dst *ListVirtualPrivateServersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVirtualPrivateServersResponse.Merge(dst, src)
}
func (m *ListVirtualPrivateServersResponse) XXX_Size() int {
	return xxx_messageInfo_ListVirtualPrivateServersResponse.Size(m)
}
func (m *ListVirtualPrivateServersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVirtualPrivateServersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListVirtualPrivateServersResponse proto.InternalMessageInfo

func (m *ListVirtualPrivateServersResponse) GetVirtualPrivateServers() []*VirtualPrivateServer {
	if m != nil {
		return m.VirtualPrivateServers
	}
	return nil
}

type GetVirtualPrivateServerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVirtualPrivateServerRequest) Reset()         { *m = GetVirtualPrivateServerRequest{} }
func (m *GetVirtualPrivateServerRequest) String() string { return proto.CompactTextString(m) }
func (*GetVirtualPrivateServerRequest) ProtoMessage()    {}
func (*GetVirtualPrivateServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_private_server_d376265387d5e975, []int{3}
}
func (m *GetVirtualPrivateServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVirtualPrivateServerRequest.Unmarshal(m, b)
}
func (m *GetVirtualPrivateServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVirtualPrivateServerRequest.Marshal(b, m, deterministic)
}
func (dst *GetVirtualPrivateServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVirtualPrivateServerRequest.Merge(dst, src)
}
func (m *GetVirtualPrivateServerRequest) XXX_Size() int {
	return xxx_messageInfo_GetVirtualPrivateServerRequest.Size(m)
}
func (m *GetVirtualPrivateServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVirtualPrivateServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVirtualPrivateServerRequest proto.InternalMessageInfo

func (m *GetVirtualPrivateServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplyVirtualPrivateServerRequest struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Version              uint64            `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ApplyVirtualPrivateServerRequest) Reset()         { *m = ApplyVirtualPrivateServerRequest{} }
func (m *ApplyVirtualPrivateServerRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyVirtualPrivateServerRequest) ProtoMessage()    {}
func (*ApplyVirtualPrivateServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_private_server_d376265387d5e975, []int{4}
}
func (m *ApplyVirtualPrivateServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyVirtualPrivateServerRequest.Unmarshal(m, b)
}
func (m *ApplyVirtualPrivateServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyVirtualPrivateServerRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyVirtualPrivateServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyVirtualPrivateServerRequest.Merge(dst, src)
}
func (m *ApplyVirtualPrivateServerRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyVirtualPrivateServerRequest.Size(m)
}
func (m *ApplyVirtualPrivateServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyVirtualPrivateServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyVirtualPrivateServerRequest proto.InternalMessageInfo

func (m *ApplyVirtualPrivateServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApplyVirtualPrivateServerRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ApplyVirtualPrivateServerRequest) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type DeleteVirtualPrivateServerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteVirtualPrivateServerRequest) Reset()         { *m = DeleteVirtualPrivateServerRequest{} }
func (m *DeleteVirtualPrivateServerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteVirtualPrivateServerRequest) ProtoMessage()    {}
func (*DeleteVirtualPrivateServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_private_server_d376265387d5e975, []int{5}
}
func (m *DeleteVirtualPrivateServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteVirtualPrivateServerRequest.Unmarshal(m, b)
}
func (m *DeleteVirtualPrivateServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteVirtualPrivateServerRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteVirtualPrivateServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteVirtualPrivateServerRequest.Merge(dst, src)
}
func (m *DeleteVirtualPrivateServerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteVirtualPrivateServerRequest.Size(m)
}
func (m *DeleteVirtualPrivateServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteVirtualPrivateServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteVirtualPrivateServerRequest proto.InternalMessageInfo

func (m *DeleteVirtualPrivateServerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GenerateVirtualMachineRequest struct {
	VirtualPrivateServerName string            `protobuf:"bytes,1,opt,name=virtual_private_server_name,json=virtualPrivateServerName" json:"virtual_private_server_name,omitempty"`
	VirtualMachineName       string            `protobuf:"bytes,2,opt,name=virtual_machine_name,json=virtualMachineName" json:"virtual_machine_name,omitempty"`
	Annotations              map[string]string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	RequestCpuMilliCore      uint32            `protobuf:"varint,4,opt,name=request_cpu_milli_core,json=requestCpuMilliCore" json:"request_cpu_milli_core,omitempty"`
	RequestMemoryBytes       uint64            `protobuf:"varint,5,opt,name=request_memory_bytes,json=requestMemoryBytes" json:"request_memory_bytes,omitempty"`
	ImageName                string            `protobuf:"bytes,6,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *GenerateVirtualMachineRequest) Reset()         { *m = GenerateVirtualMachineRequest{} }
func (m *GenerateVirtualMachineRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateVirtualMachineRequest) ProtoMessage()    {}
func (*GenerateVirtualMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_private_server_d376265387d5e975, []int{6}
}
func (m *GenerateVirtualMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateVirtualMachineRequest.Unmarshal(m, b)
}
func (m *GenerateVirtualMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateVirtualMachineRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateVirtualMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateVirtualMachineRequest.Merge(dst, src)
}
func (m *GenerateVirtualMachineRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateVirtualMachineRequest.Size(m)
}
func (m *GenerateVirtualMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateVirtualMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateVirtualMachineRequest proto.InternalMessageInfo

func (m *GenerateVirtualMachineRequest) GetVirtualPrivateServerName() string {
	if m != nil {
		return m.VirtualPrivateServerName
	}
	return ""
}

func (m *GenerateVirtualMachineRequest) GetVirtualMachineName() string {
	if m != nil {
		return m.VirtualMachineName
	}
	return ""
}

func (m *GenerateVirtualMachineRequest) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *GenerateVirtualMachineRequest) GetRequestCpuMilliCore() uint32 {
	if m != nil {
		return m.RequestCpuMilliCore
	}
	return 0
}

func (m *GenerateVirtualMachineRequest) GetRequestMemoryBytes() uint64 {
	if m != nil {
		return m.RequestMemoryBytes
	}
	return 0
}

func (m *GenerateVirtualMachineRequest) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func init() {
	proto.RegisterType((*VirtualPrivateServer)(nil), "n0stack.deployment.VirtualPrivateServer")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.deployment.VirtualPrivateServer.AnnotationsEntry")
	proto.RegisterType((*ListVirtualPrivateServersRequest)(nil), "n0stack.deployment.ListVirtualPrivateServersRequest")
	proto.RegisterType((*ListVirtualPrivateServersResponse)(nil), "n0stack.deployment.ListVirtualPrivateServersResponse")
	proto.RegisterType((*GetVirtualPrivateServerRequest)(nil), "n0stack.deployment.GetVirtualPrivateServerRequest")
	proto.RegisterType((*ApplyVirtualPrivateServerRequest)(nil), "n0stack.deployment.ApplyVirtualPrivateServerRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.deployment.ApplyVirtualPrivateServerRequest.AnnotationsEntry")
	proto.RegisterType((*DeleteVirtualPrivateServerRequest)(nil), "n0stack.deployment.DeleteVirtualPrivateServerRequest")
	proto.RegisterType((*GenerateVirtualMachineRequest)(nil), "n0stack.deployment.GenerateVirtualMachineRequest")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.deployment.GenerateVirtualMachineRequest.AnnotationsEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VirtualPrivateServerService service

type VirtualPrivateServerServiceClient interface {
	ListVirtualPrivateServers(ctx context.Context, in *ListVirtualPrivateServersRequest, opts ...grpc.CallOption) (*ListVirtualPrivateServersResponse, error)
	GetVirtualPrivateServer(ctx context.Context, in *GetVirtualPrivateServerRequest, opts ...grpc.CallOption) (*VirtualPrivateServer, error)
	ApplyVirtualPrivateServer(ctx context.Context, in *ApplyVirtualPrivateServerRequest, opts ...grpc.CallOption) (*VirtualPrivateServer, error)
	DeleteVirtualPrivateServer(ctx context.Context, in *DeleteVirtualPrivateServerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GenerateVirtualMachine(ctx context.Context, in *GenerateVirtualMachineRequest, opts ...grpc.CallOption) (*v0.VirtualMachine, error)
}

type virtualPrivateServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewVirtualPrivateServerServiceClient(cc *grpc.ClientConn) VirtualPrivateServerServiceClient {
	return &virtualPrivateServerServiceClient{cc}
}

func (c *virtualPrivateServerServiceClient) ListVirtualPrivateServers(ctx context.Context, in *ListVirtualPrivateServersRequest, opts ...grpc.CallOption) (*ListVirtualPrivateServersResponse, error) {
	out := new(ListVirtualPrivateServersResponse)
	err := grpc.Invoke(ctx, "/n0stack.deployment.VirtualPrivateServerService/ListVirtualPrivateServers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualPrivateServerServiceClient) GetVirtualPrivateServer(ctx context.Context, in *GetVirtualPrivateServerRequest, opts ...grpc.CallOption) (*VirtualPrivateServer, error) {
	out := new(VirtualPrivateServer)
	err := grpc.Invoke(ctx, "/n0stack.deployment.VirtualPrivateServerService/GetVirtualPrivateServer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualPrivateServerServiceClient) ApplyVirtualPrivateServer(ctx context.Context, in *ApplyVirtualPrivateServerRequest, opts ...grpc.CallOption) (*VirtualPrivateServer, error) {
	out := new(VirtualPrivateServer)
	err := grpc.Invoke(ctx, "/n0stack.deployment.VirtualPrivateServerService/ApplyVirtualPrivateServer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualPrivateServerServiceClient) DeleteVirtualPrivateServer(ctx context.Context, in *DeleteVirtualPrivateServerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.deployment.VirtualPrivateServerService/DeleteVirtualPrivateServer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualPrivateServerServiceClient) GenerateVirtualMachine(ctx context.Context, in *GenerateVirtualMachineRequest, opts ...grpc.CallOption) (*v0.VirtualMachine, error) {
	out := new(v0.VirtualMachine)
	err := grpc.Invoke(ctx, "/n0stack.deployment.VirtualPrivateServerService/GenerateVirtualMachine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VirtualPrivateServerService service

type VirtualPrivateServerServiceServer interface {
	ListVirtualPrivateServers(context.Context, *ListVirtualPrivateServersRequest) (*ListVirtualPrivateServersResponse, error)
	GetVirtualPrivateServer(context.Context, *GetVirtualPrivateServerRequest) (*VirtualPrivateServer, error)
	ApplyVirtualPrivateServer(context.Context, *ApplyVirtualPrivateServerRequest) (*VirtualPrivateServer, error)
	DeleteVirtualPrivateServer(context.Context, *DeleteVirtualPrivateServerRequest) (*empty.Empty, error)
	GenerateVirtualMachine(context.Context, *GenerateVirtualMachineRequest) (*v0.VirtualMachine, error)
}

func RegisterVirtualPrivateServerServiceServer(s *grpc.Server, srv VirtualPrivateServerServiceServer) {
	s.RegisterService(&_VirtualPrivateServerService_serviceDesc, srv)
}

func _VirtualPrivateServerService_ListVirtualPrivateServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVirtualPrivateServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrivateServerServiceServer).ListVirtualPrivateServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.VirtualPrivateServerService/ListVirtualPrivateServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrivateServerServiceServer).ListVirtualPrivateServers(ctx, req.(*ListVirtualPrivateServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualPrivateServerService_GetVirtualPrivateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVirtualPrivateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrivateServerServiceServer).GetVirtualPrivateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.VirtualPrivateServerService/GetVirtualPrivateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrivateServerServiceServer).GetVirtualPrivateServer(ctx, req.(*GetVirtualPrivateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualPrivateServerService_ApplyVirtualPrivateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyVirtualPrivateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrivateServerServiceServer).ApplyVirtualPrivateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.VirtualPrivateServerService/ApplyVirtualPrivateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrivateServerServiceServer).ApplyVirtualPrivateServer(ctx, req.(*ApplyVirtualPrivateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualPrivateServerService_DeleteVirtualPrivateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVirtualPrivateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrivateServerServiceServer).DeleteVirtualPrivateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.VirtualPrivateServerService/DeleteVirtualPrivateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrivateServerServiceServer).DeleteVirtualPrivateServer(ctx, req.(*DeleteVirtualPrivateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualPrivateServerService_GenerateVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualPrivateServerServiceServer).GenerateVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.deployment.VirtualPrivateServerService/GenerateVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualPrivateServerServiceServer).GenerateVirtualMachine(ctx, req.(*GenerateVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualPrivateServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.deployment.VirtualPrivateServerService",
	HandlerType: (*VirtualPrivateServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVirtualPrivateServers",
			Handler:    _VirtualPrivateServerService_ListVirtualPrivateServers_Handler,
		},
		{
			MethodName: "GetVirtualPrivateServer",
			Handler:    _VirtualPrivateServerService_GetVirtualPrivateServer_Handler,
		},
		{
			MethodName: "ApplyVirtualPrivateServer",
			Handler:    _VirtualPrivateServerService_ApplyVirtualPrivateServer_Handler,
		},
		{
			MethodName: "DeleteVirtualPrivateServer",
			Handler:    _VirtualPrivateServerService_DeleteVirtualPrivateServer_Handler,
		},
		{
			MethodName: "GenerateVirtualMachine",
			Handler:    _VirtualPrivateServerService_GenerateVirtualMachine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deployment/v0/virtual_private_server.proto",
}

func init() {
	proto.RegisterFile("deployment/v0/virtual_private_server.proto", fileDescriptor_virtual_private_server_d376265387d5e975)
}

var fileDescriptor_virtual_private_server_d376265387d5e975 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x8e, 0x9b, 0xb6, 0xaf, 0x3a, 0xe9, 0x2b, 0x95, 0x25, 0xb4, 0xae, 0xab, 0xa2, 0xd4, 0x02,
	0x29, 0x42, 0xc8, 0x0e, 0xfd, 0x50, 0xf9, 0x10, 0x48, 0x6d, 0xa9, 0x7a, 0xa1, 0x08, 0x05, 0x89,
	0x03, 0x1c, 0x8c, 0x93, 0x0e, 0xee, 0xaa, 0xb6, 0xd7, 0xac, 0xd7, 0x46, 0x3e, 0x70, 0x44, 0x42,
	0xfc, 0x26, 0xfe, 0x0d, 0xbf, 0x82, 0x1b, 0xf2, 0xae, 0xd3, 0x26, 0x65, 0xd3, 0x34, 0xc0, 0x25,
	0xca, 0xee, 0xcc, 0x33, 0x3b, 0x33, 0xcf, 0x3c, 0xbb, 0x86, 0x7b, 0x27, 0x98, 0x84, 0xac, 0x88,
	0x30, 0x16, 0x6e, 0xde, 0x71, 0x73, 0xca, 0x45, 0xe6, 0x87, 0x5e, 0xc2, 0x69, 0xee, 0x0b, 0xf4,
	0x52, 0xe4, 0x39, 0x72, 0x27, 0xe1, 0x4c, 0x30, 0x42, 0xe2, 0x4e, 0x2a, 0xfc, 0xfe, 0x99, 0x73,
	0x81, 0xb1, 0xd6, 0x02, 0xc6, 0x82, 0x10, 0x5d, 0xe9, 0xd1, 0xcb, 0x3e, 0xb8, 0x18, 0x25, 0xa2,
	0x50, 0x00, 0xeb, 0x6e, 0xc2, 0x59, 0x4e, 0x53, 0xca, 0x62, 0x1a, 0x07, 0xc3, 0xe1, 0x23, 0xbf,
	0x7f, 0x4a, 0x63, 0x54, 0x6e, 0xf6, 0x8f, 0x19, 0x68, 0xbe, 0x51, 0x96, 0x57, 0xea, 0xdc, 0xd7,
	0xf2, 0x58, 0x42, 0x60, 0x36, 0xf6, 0x23, 0x34, 0x8d, 0x96, 0xd1, 0x5e, 0xe8, 0xca, 0xff, 0xe4,
	0x1d, 0x34, 0xfc, 0x38, 0x66, 0xc2, 0x17, 0x94, 0xc5, 0xa9, 0x59, 0x6f, 0xd5, 0xdb, 0x8d, 0xcd,
	0x47, 0xce, 0xef, 0xa9, 0x39, 0xba, 0x90, 0xce, 0xde, 0x05, 0xf6, 0x30, 0x16, 0xbc, 0xe8, 0x0e,
	0x47, 0x23, 0x26, 0xfc, 0x97, 0x23, 0x2f, 0x13, 0x36, 0xe7, 0x5a, 0x46, 0x7b, 0xb6, 0x3b, 0x58,
	0x12, 0x17, 0x9a, 0x21, 0x8d, 0xa8, 0xf0, 0xfa, 0x49, 0xe6, 0x45, 0x34, 0x0c, 0xa9, 0xd7, 0x67,
	0x1c, 0x4d, 0x68, 0x19, 0xed, 0xff, 0xbb, 0x37, 0xa4, 0xed, 0x20, 0xc9, 0x8e, 0x4b, 0xcb, 0x01,
	0xe3, 0x48, 0xee, 0x03, 0x51, 0x80, 0x08, 0x23, 0xc6, 0x0b, 0xaf, 0x57, 0x08, 0x4c, 0xcd, 0x86,
	0x8c, 0xba, 0x24, 0x2d, 0xc7, 0xd2, 0xb0, 0x5f, 0xee, 0x93, 0x0d, 0x58, 0x8c, 0x51, 0x7c, 0x62,
	0xfc, 0xcc, 0x93, 0x15, 0x2f, 0xca, 0x8a, 0x1b, 0xd5, 0xde, 0x4b, 0x3f, 0x42, 0xeb, 0x19, 0x2c,
	0x5d, 0x4e, 0x9e, 0x2c, 0x41, 0xfd, 0x0c, 0x8b, 0xaa, 0x3f, 0xe5, 0x5f, 0xd2, 0x84, 0xb9, 0xdc,
	0x0f, 0x33, 0x34, 0x67, 0xe4, 0x9e, 0x5a, 0x3c, 0x9e, 0x79, 0x68, 0xd8, 0x36, 0xb4, 0x5e, 0xd0,
	0x54, 0xe8, 0xba, 0x92, 0x76, 0xf1, 0x63, 0x86, 0xa9, 0xb0, 0xbf, 0x18, 0xb0, 0x71, 0x85, 0x53,
	0x9a, 0xb0, 0x38, 0x45, 0xf2, 0x1e, 0x56, 0xf4, 0x73, 0x92, 0x9a, 0x86, 0xa4, 0xa3, 0x7d, 0x5d,
	0x3a, 0xba, 0xb7, 0x72, 0xdd, 0x49, 0xf6, 0x36, 0xdc, 0x3e, 0x42, 0x6d, 0x16, 0x55, 0xa6, 0xba,
	0xd1, 0xb0, 0x7f, 0x1a, 0xd0, 0xda, 0x4b, 0x92, 0xb0, 0x98, 0x12, 0x48, 0x02, 0xdd, 0x4c, 0x1d,
	0xea, 0x8a, 0x98, 0x14, 0xfe, 0x4f, 0xe7, 0xeb, 0xaf, 0xd9, 0xdd, 0x85, 0x8d, 0xe7, 0x18, 0xa2,
	0xc0, 0x69, 0x9b, 0xf6, 0xbd, 0x0e, 0xeb, 0x47, 0x18, 0x23, 0xf7, 0xcf, 0xb1, 0xc7, 0x4a, 0x9d,
	0x03, 0xd4, 0x53, 0x58, 0xd3, 0xd3, 0xed, 0x0d, 0x05, 0x33, 0x75, 0x44, 0x96, 0x73, 0x4b, 0x3a,
	0xd0, 0xbc, 0x24, 0x7b, 0x85, 0x53, 0x25, 0x90, 0x7c, 0xe4, 0x4c, 0x89, 0x38, 0xd1, 0xd1, 0xb1,
	0xaf, 0xa3, 0xe3, 0xca, 0xc4, 0x27, 0x70, 0xb1, 0x05, 0xcb, 0x5c, 0x39, 0x5e, 0xd6, 0xf4, 0xac,
	0xd4, 0xf4, 0xcd, 0xca, 0x3a, 0xa2, 0xea, 0x0e, 0x34, 0x07, 0xa0, 0x11, 0x5d, 0x2b, 0x36, 0x49,
	0x65, 0x1b, 0x56, 0xf6, 0x3a, 0x00, 0x8d, 0xfc, 0xa0, 0x2a, 0x7a, 0x5e, 0x16, 0xbd, 0x20, 0x77,
	0xfe, 0x85, 0xaa, 0x37, 0xbf, 0xce, 0xc1, 0x9a, 0x8e, 0xf2, 0xf2, 0x97, 0xf6, 0x91, 0x7c, 0x33,
	0x60, 0x75, 0xac, 0xa2, 0xc9, 0xb6, 0xae, 0xa9, 0x93, 0x6e, 0x09, 0x6b, 0x67, 0x4a, 0x94, 0xba,
	0x36, 0xec, 0x1a, 0x29, 0x60, 0x65, 0x8c, 0xac, 0xc9, 0xa6, 0x9e, 0xde, 0xab, 0xee, 0x00, 0xeb,
	0xda, 0xd7, 0x8c, 0x5d, 0x23, 0x9f, 0x61, 0x75, 0xac, 0x76, 0xf5, 0x6d, 0x98, 0x24, 0xf5, 0xa9,
	0x8e, 0x8f, 0xc0, 0x1a, 0x2f, 0x4f, 0xa2, 0x6d, 0xe8, 0x44, 0x39, 0x5b, 0xcb, 0x8e, 0x7a, 0x7c,
	0x9d, 0xc1, 0xe3, 0xeb, 0x1c, 0x96, 0x8f, 0xaf, 0x5d, 0x23, 0x29, 0x2c, 0xeb, 0xa5, 0x41, 0x1e,
	0x4c, 0x2d, 0x23, 0xeb, 0xce, 0x39, 0x64, 0xf8, 0x39, 0x77, 0x46, 0x9d, 0xed, 0xda, 0xfe, 0xee,
	0xdb, 0x9d, 0x80, 0x8a, 0xd3, 0xac, 0xe7, 0xf4, 0x59, 0xe4, 0x56, 0x18, 0xf5, 0x61, 0xe0, 0x04,
	0xcc, 0x1d, 0xf9, 0xd0, 0x78, 0x92, 0x5c, 0x2c, 0x7b, 0xf3, 0xd2, 0x67, 0xeb, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xae, 0x58, 0x63, 0x04, 0x8c, 0x08, 0x00, 0x00,
}
