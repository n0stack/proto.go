// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provisioning/v0/network.proto

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

type NetworkStatus_NetworkState int32

const (
	// not scheduled
	NetworkStatus_PENDING   NetworkStatus_NetworkState = 0
	NetworkStatus_AVAILABLE NetworkStatus_NetworkState = 1
	// unknown state because failed to connect for scheduled node after RUNNING.
	NetworkStatus_FAILED NetworkStatus_NetworkState = 2
	// unknown state because failed to connect for scheduled node after RUNNING.
	NetworkStatus_UNKNOWN NetworkStatus_NetworkState = 3
)

var NetworkStatus_NetworkState_name = map[int32]string{
	0: "PENDING",
	1: "AVAILABLE",
	2: "FAILED",
	3: "UNKNOWN",
}
var NetworkStatus_NetworkState_value = map[string]int32{
	"PENDING":   0,
	"AVAILABLE": 1,
	"FAILED":    2,
	"UNKNOWN":   3,
}

func (x NetworkStatus_NetworkState) String() string {
	return proto.EnumName(NetworkStatus_NetworkState_name, int32(x))
}
func (NetworkStatus_NetworkState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{2, 0}
}

type WatchNetworksResponse_NetworkEvents int32

const (
	WatchNetworksResponse_APPLY  WatchNetworksResponse_NetworkEvents = 0
	WatchNetworksResponse_DELETE WatchNetworksResponse_NetworkEvents = 1
)

var WatchNetworksResponse_NetworkEvents_name = map[int32]string{
	0: "APPLY",
	1: "DELETE",
}
var WatchNetworksResponse_NetworkEvents_value = map[string]int32{
	"APPLY":  0,
	"DELETE": 1,
}

func (x WatchNetworksResponse_NetworkEvents) String() string {
	return proto.EnumName(WatchNetworksResponse_NetworkEvents_name, int32(x))
}
func (WatchNetworksResponse_NetworkEvents) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{9, 0}
}

type Network struct {
	Metadata             *v0.Metadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec                 *NetworkSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status               *NetworkStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}
func (*Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{0}
}
func (m *Network) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Network.Unmarshal(m, b)
}
func (m *Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Network.Marshal(b, m, deterministic)
}
func (dst *Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network.Merge(dst, src)
}
func (m *Network) XXX_Size() int {
	return xxx_messageInfo_Network.Size(m)
}
func (m *Network) XXX_DiscardUnknown() {
	xxx_messageInfo_Network.DiscardUnknown(m)
}

var xxx_messageInfo_Network proto.InternalMessageInfo

func (m *Network) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Network) GetSpec() *NetworkSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Network) GetStatus() *NetworkStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type NetworkSpec struct {
	Ipv4Cidr             string   `protobuf:"bytes,1,opt,name=ipv4_cidr,json=ipv4Cidr,proto3" json:"ipv4_cidr,omitempty"`
	Ipv4Gateway          string   `protobuf:"bytes,2,opt,name=ipv4_gateway,json=ipv4Gateway,proto3" json:"ipv4_gateway,omitempty"`
	NameServer           string   `protobuf:"bytes,5,opt,name=name_server,json=nameServer,proto3" json:"name_server,omitempty"`
	Domain               string   `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkSpec) Reset()         { *m = NetworkSpec{} }
func (m *NetworkSpec) String() string { return proto.CompactTextString(m) }
func (*NetworkSpec) ProtoMessage()    {}
func (*NetworkSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{1}
}
func (m *NetworkSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkSpec.Unmarshal(m, b)
}
func (m *NetworkSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkSpec.Marshal(b, m, deterministic)
}
func (dst *NetworkSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkSpec.Merge(dst, src)
}
func (m *NetworkSpec) XXX_Size() int {
	return xxx_messageInfo_NetworkSpec.Size(m)
}
func (m *NetworkSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkSpec.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkSpec proto.InternalMessageInfo

func (m *NetworkSpec) GetIpv4Cidr() string {
	if m != nil {
		return m.Ipv4Cidr
	}
	return ""
}

func (m *NetworkSpec) GetIpv4Gateway() string {
	if m != nil {
		return m.Ipv4Gateway
	}
	return ""
}

func (m *NetworkSpec) GetNameServer() string {
	if m != nil {
		return m.NameServer
	}
	return ""
}

func (m *NetworkSpec) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type NetworkStatus struct {
	State                NetworkStatus_NetworkState `protobuf:"varint,1,opt,name=state,proto3,enum=n0stack.provisioning.NetworkStatus_NetworkState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *NetworkStatus) Reset()         { *m = NetworkStatus{} }
func (m *NetworkStatus) String() string { return proto.CompactTextString(m) }
func (*NetworkStatus) ProtoMessage()    {}
func (*NetworkStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{2}
}
func (m *NetworkStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkStatus.Unmarshal(m, b)
}
func (m *NetworkStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkStatus.Marshal(b, m, deterministic)
}
func (dst *NetworkStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkStatus.Merge(dst, src)
}
func (m *NetworkStatus) XXX_Size() int {
	return xxx_messageInfo_NetworkStatus.Size(m)
}
func (m *NetworkStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkStatus.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkStatus proto.InternalMessageInfo

func (m *NetworkStatus) GetState() NetworkStatus_NetworkState {
	if m != nil {
		return m.State
	}
	return NetworkStatus_PENDING
}

type ListNetworksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNetworksRequest) Reset()         { *m = ListNetworksRequest{} }
func (m *ListNetworksRequest) String() string { return proto.CompactTextString(m) }
func (*ListNetworksRequest) ProtoMessage()    {}
func (*ListNetworksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{3}
}
func (m *ListNetworksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworksRequest.Unmarshal(m, b)
}
func (m *ListNetworksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworksRequest.Marshal(b, m, deterministic)
}
func (dst *ListNetworksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworksRequest.Merge(dst, src)
}
func (m *ListNetworksRequest) XXX_Size() int {
	return xxx_messageInfo_ListNetworksRequest.Size(m)
}
func (m *ListNetworksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworksRequest proto.InternalMessageInfo

type ListNetworksResponse struct {
	Networks             []*Network `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListNetworksResponse) Reset()         { *m = ListNetworksResponse{} }
func (m *ListNetworksResponse) String() string { return proto.CompactTextString(m) }
func (*ListNetworksResponse) ProtoMessage()    {}
func (*ListNetworksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{4}
}
func (m *ListNetworksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworksResponse.Unmarshal(m, b)
}
func (m *ListNetworksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworksResponse.Marshal(b, m, deterministic)
}
func (dst *ListNetworksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworksResponse.Merge(dst, src)
}
func (m *ListNetworksResponse) XXX_Size() int {
	return xxx_messageInfo_ListNetworksResponse.Size(m)
}
func (m *ListNetworksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworksResponse proto.InternalMessageInfo

func (m *ListNetworksResponse) GetNetworks() []*Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

type GetNetworkRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNetworkRequest) Reset()         { *m = GetNetworkRequest{} }
func (m *GetNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*GetNetworkRequest) ProtoMessage()    {}
func (*GetNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{5}
}
func (m *GetNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNetworkRequest.Unmarshal(m, b)
}
func (m *GetNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNetworkRequest.Marshal(b, m, deterministic)
}
func (dst *GetNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNetworkRequest.Merge(dst, src)
}
func (m *GetNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_GetNetworkRequest.Size(m)
}
func (m *GetNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNetworkRequest proto.InternalMessageInfo

func (m *GetNetworkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplyNetworkRequest struct {
	Metadata             *v0.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec                 *NetworkSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ApplyNetworkRequest) Reset()         { *m = ApplyNetworkRequest{} }
func (m *ApplyNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyNetworkRequest) ProtoMessage()    {}
func (*ApplyNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{6}
}
func (m *ApplyNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyNetworkRequest.Unmarshal(m, b)
}
func (m *ApplyNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyNetworkRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyNetworkRequest.Merge(dst, src)
}
func (m *ApplyNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyNetworkRequest.Size(m)
}
func (m *ApplyNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyNetworkRequest proto.InternalMessageInfo

func (m *ApplyNetworkRequest) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ApplyNetworkRequest) GetSpec() *NetworkSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type DeleteNetworkRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNetworkRequest) Reset()         { *m = DeleteNetworkRequest{} }
func (m *DeleteNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNetworkRequest) ProtoMessage()    {}
func (*DeleteNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{7}
}
func (m *DeleteNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNetworkRequest.Unmarshal(m, b)
}
func (m *DeleteNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNetworkRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNetworkRequest.Merge(dst, src)
}
func (m *DeleteNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNetworkRequest.Size(m)
}
func (m *DeleteNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNetworkRequest proto.InternalMessageInfo

func (m *DeleteNetworkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type WatchNetworksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchNetworksRequest) Reset()         { *m = WatchNetworksRequest{} }
func (m *WatchNetworksRequest) String() string { return proto.CompactTextString(m) }
func (*WatchNetworksRequest) ProtoMessage()    {}
func (*WatchNetworksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{8}
}
func (m *WatchNetworksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchNetworksRequest.Unmarshal(m, b)
}
func (m *WatchNetworksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchNetworksRequest.Marshal(b, m, deterministic)
}
func (dst *WatchNetworksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchNetworksRequest.Merge(dst, src)
}
func (m *WatchNetworksRequest) XXX_Size() int {
	return xxx_messageInfo_WatchNetworksRequest.Size(m)
}
func (m *WatchNetworksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchNetworksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchNetworksRequest proto.InternalMessageInfo

type WatchNetworksResponse struct {
	Event                WatchNetworksResponse_NetworkEvents `protobuf:"varint,1,opt,name=event,proto3,enum=n0stack.provisioning.WatchNetworksResponse_NetworkEvents" json:"event,omitempty"`
	Network              *Network                            `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *WatchNetworksResponse) Reset()         { *m = WatchNetworksResponse{} }
func (m *WatchNetworksResponse) String() string { return proto.CompactTextString(m) }
func (*WatchNetworksResponse) ProtoMessage()    {}
func (*WatchNetworksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_263284cba1af3705, []int{9}
}
func (m *WatchNetworksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchNetworksResponse.Unmarshal(m, b)
}
func (m *WatchNetworksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchNetworksResponse.Marshal(b, m, deterministic)
}
func (dst *WatchNetworksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchNetworksResponse.Merge(dst, src)
}
func (m *WatchNetworksResponse) XXX_Size() int {
	return xxx_messageInfo_WatchNetworksResponse.Size(m)
}
func (m *WatchNetworksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchNetworksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchNetworksResponse proto.InternalMessageInfo

func (m *WatchNetworksResponse) GetEvent() WatchNetworksResponse_NetworkEvents {
	if m != nil {
		return m.Event
	}
	return WatchNetworksResponse_APPLY
}

func (m *WatchNetworksResponse) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func init() {
	proto.RegisterType((*Network)(nil), "n0stack.provisioning.Network")
	proto.RegisterType((*NetworkSpec)(nil), "n0stack.provisioning.NetworkSpec")
	proto.RegisterType((*NetworkStatus)(nil), "n0stack.provisioning.NetworkStatus")
	proto.RegisterType((*ListNetworksRequest)(nil), "n0stack.provisioning.ListNetworksRequest")
	proto.RegisterType((*ListNetworksResponse)(nil), "n0stack.provisioning.ListNetworksResponse")
	proto.RegisterType((*GetNetworkRequest)(nil), "n0stack.provisioning.GetNetworkRequest")
	proto.RegisterType((*ApplyNetworkRequest)(nil), "n0stack.provisioning.ApplyNetworkRequest")
	proto.RegisterType((*DeleteNetworkRequest)(nil), "n0stack.provisioning.DeleteNetworkRequest")
	proto.RegisterType((*WatchNetworksRequest)(nil), "n0stack.provisioning.WatchNetworksRequest")
	proto.RegisterType((*WatchNetworksResponse)(nil), "n0stack.provisioning.WatchNetworksResponse")
	proto.RegisterEnum("n0stack.provisioning.NetworkStatus_NetworkState", NetworkStatus_NetworkState_name, NetworkStatus_NetworkState_value)
	proto.RegisterEnum("n0stack.provisioning.WatchNetworksResponse_NetworkEvents", WatchNetworksResponse_NetworkEvents_name, WatchNetworksResponse_NetworkEvents_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error)
	GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error)
	ApplyNetwork(ctx context.Context, in *ApplyNetworkRequest, opts ...grpc.CallOption) (*Network, error)
	DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	WatchNetworks(ctx context.Context, in *WatchNetworksRequest, opts ...grpc.CallOption) (NetworkService_WatchNetworksClient, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error) {
	out := new(ListNetworksResponse)
	err := c.cc.Invoke(ctx, "/n0stack.provisioning.NetworkService/ListNetworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := c.cc.Invoke(ctx, "/n0stack.provisioning.NetworkService/GetNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ApplyNetwork(ctx context.Context, in *ApplyNetworkRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := c.cc.Invoke(ctx, "/n0stack.provisioning.NetworkService/ApplyNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.provisioning.NetworkService/DeleteNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) WatchNetworks(ctx context.Context, in *WatchNetworksRequest, opts ...grpc.CallOption) (NetworkService_WatchNetworksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkService_serviceDesc.Streams[0], "/n0stack.provisioning.NetworkService/WatchNetworks", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkServiceWatchNetworksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetworkService_WatchNetworksClient interface {
	Recv() (*WatchNetworksResponse, error)
	grpc.ClientStream
}

type networkServiceWatchNetworksClient struct {
	grpc.ClientStream
}

func (x *networkServiceWatchNetworksClient) Recv() (*WatchNetworksResponse, error) {
	m := new(WatchNetworksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error)
	GetNetwork(context.Context, *GetNetworkRequest) (*Network, error)
	ApplyNetwork(context.Context, *ApplyNetworkRequest) (*Network, error)
	DeleteNetwork(context.Context, *DeleteNetworkRequest) (*empty.Empty, error)
	WatchNetworks(*WatchNetworksRequest, NetworkService_WatchNetworksServer) error
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_ListNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ListNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.NetworkService/ListNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ListNetworks(ctx, req.(*ListNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.NetworkService/GetNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetNetwork(ctx, req.(*GetNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ApplyNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ApplyNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.NetworkService/ApplyNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ApplyNetwork(ctx, req.(*ApplyNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_DeleteNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).DeleteNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.NetworkService/DeleteNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).DeleteNetwork(ctx, req.(*DeleteNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_WatchNetworks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchNetworksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetworkServiceServer).WatchNetworks(m, &networkServiceWatchNetworksServer{stream})
}

type NetworkService_WatchNetworksServer interface {
	Send(*WatchNetworksResponse) error
	grpc.ServerStream
}

type networkServiceWatchNetworksServer struct {
	grpc.ServerStream
}

func (x *networkServiceWatchNetworksServer) Send(m *WatchNetworksResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.provisioning.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNetworks",
			Handler:    _NetworkService_ListNetworks_Handler,
		},
		{
			MethodName: "GetNetwork",
			Handler:    _NetworkService_GetNetwork_Handler,
		},
		{
			MethodName: "ApplyNetwork",
			Handler:    _NetworkService_ApplyNetwork_Handler,
		},
		{
			MethodName: "DeleteNetwork",
			Handler:    _NetworkService_DeleteNetwork_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchNetworks",
			Handler:       _NetworkService_WatchNetworks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "provisioning/v0/network.proto",
}

func init() {
	proto.RegisterFile("provisioning/v0/network.proto", fileDescriptor_network_263284cba1af3705)
}

var fileDescriptor_network_263284cba1af3705 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0xd1, 0x4e, 0xdb, 0x3c,
	0x18, 0x6d, 0x28, 0x2d, 0xf4, 0x6b, 0x8b, 0x8a, 0x29, 0xa8, 0x2a, 0x42, 0x3f, 0xf8, 0x97, 0x06,
	0x63, 0x5a, 0x52, 0x75, 0x9b, 0x26, 0xc4, 0x55, 0xa1, 0x01, 0xa1, 0x65, 0x81, 0x85, 0x01, 0x1a,
	0x37, 0x28, 0xa4, 0x5e, 0x88, 0xd6, 0x26, 0x59, 0xec, 0x06, 0xa1, 0xdd, 0xef, 0x29, 0xf6, 0x08,
	0x7b, 0x91, 0xbd, 0xce, 0x9e, 0x60, 0x8a, 0xe3, 0x74, 0x69, 0x17, 0xa0, 0xbb, 0xd9, 0x5d, 0xfc,
	0xf9, 0x9c, 0xe3, 0xf3, 0x7d, 0x3e, 0x72, 0x60, 0xcd, 0x0f, 0xbc, 0xd0, 0xa1, 0x8e, 0xe7, 0x3a,
	0xae, 0xad, 0x84, 0x2d, 0xc5, 0x25, 0xec, 0xd6, 0x0b, 0x3e, 0xc9, 0x7e, 0xe0, 0x31, 0x0f, 0xd5,
	0xdd, 0x16, 0x65, 0xa6, 0xc5, 0x97, 0x23, 0x58, 0x73, 0xd5, 0xf6, 0x3c, 0xbb, 0x4f, 0x14, 0x8e,
	0xb9, 0x1e, 0x7e, 0x54, 0xc8, 0xc0, 0x67, 0x77, 0x31, 0xa5, 0xb9, 0x18, 0xb6, 0x94, 0x01, 0x61,
	0x66, 0xcf, 0x64, 0x66, 0x5c, 0xc2, 0xdf, 0x25, 0x98, 0xd3, 0x63, 0x5d, 0xf4, 0x1c, 0xe6, 0x93,
	0xdd, 0x86, 0xb4, 0x2e, 0x6d, 0x95, 0xdb, 0x8b, 0x72, 0x72, 0xc8, 0x5b, 0xb1, 0x61, 0x8c, 0x20,
	0xe8, 0x15, 0xcc, 0x52, 0x9f, 0x58, 0x8d, 0x19, 0x0e, 0xdd, 0x90, 0xb3, 0xfc, 0xc8, 0x42, 0xfb,
	0xd4, 0x27, 0x96, 0xc1, 0xe1, 0x68, 0x17, 0x8a, 0x94, 0x99, 0x6c, 0x48, 0x1b, 0x79, 0x4e, 0xfc,
	0xff, 0x61, 0x22, 0x87, 0x1a, 0x82, 0x82, 0xbf, 0x4a, 0x50, 0x4e, 0x49, 0xa2, 0x55, 0x28, 0x39,
	0x7e, 0xf8, 0xf2, 0xca, 0x72, 0x7a, 0x01, 0xf7, 0x5c, 0x32, 0xe6, 0xa3, 0xc2, 0xbe, 0xd3, 0x0b,
	0xd0, 0x06, 0x54, 0xf8, 0xa6, 0x6d, 0x32, 0x72, 0x6b, 0xde, 0x71, 0xa3, 0x25, 0xa3, 0x1c, 0xd5,
	0x0e, 0xe3, 0x12, 0xfa, 0x0f, 0xca, 0xae, 0x39, 0x20, 0x57, 0x94, 0x04, 0x21, 0x09, 0x1a, 0x05,
	0x8e, 0x80, 0xa8, 0x74, 0xca, 0x2b, 0x68, 0x05, 0x8a, 0x3d, 0x6f, 0x60, 0x3a, 0x6e, 0xa3, 0xc8,
	0xf7, 0xc4, 0x0a, 0x7f, 0x93, 0xa0, 0x3a, 0x66, 0x11, 0x1d, 0x40, 0x21, 0x32, 0x49, 0xb8, 0x8d,
	0x85, 0x76, 0x6b, 0x8a, 0xb6, 0xd2, 0x2b, 0x62, 0xc4, 0x74, 0xbc, 0x0f, 0x95, 0x74, 0x19, 0x95,
	0x61, 0xee, 0x44, 0xd5, 0xbb, 0x47, 0xfa, 0x61, 0x2d, 0x87, 0xaa, 0x50, 0xea, 0x9c, 0x77, 0x8e,
	0xb4, 0xce, 0x9e, 0xa6, 0xd6, 0x24, 0x04, 0x50, 0x3c, 0xe8, 0x1c, 0x69, 0x6a, 0xb7, 0x36, 0x13,
	0xe1, 0xce, 0xf4, 0x37, 0xfa, 0xf1, 0x85, 0x5e, 0xcb, 0xe3, 0x65, 0x58, 0xd2, 0x1c, 0xca, 0x84,
	0x10, 0x35, 0xc8, 0xe7, 0x21, 0xa1, 0x0c, 0xbf, 0x83, 0xfa, 0x78, 0x99, 0xfa, 0x9e, 0x4b, 0x09,
	0xda, 0x81, 0x79, 0x11, 0x2e, 0xda, 0x90, 0xd6, 0xf3, 0x5b, 0xe5, 0xf6, 0xda, 0x83, 0xf6, 0x8d,
	0x11, 0x1c, 0x6f, 0xc2, 0xe2, 0x21, 0x49, 0x14, 0xc5, 0x39, 0x08, 0xc1, 0x6c, 0x34, 0x43, 0x71,
	0x23, 0xfc, 0x1b, 0x7f, 0x81, 0xa5, 0x8e, 0xef, 0xf7, 0xef, 0x26, 0xa0, 0xff, 0x24, 0x74, 0x78,
	0x1b, 0xea, 0x5d, 0xd2, 0x27, 0x8c, 0x4c, 0x61, 0x74, 0x05, 0xea, 0x17, 0x26, 0xb3, 0x6e, 0x26,
	0x87, 0xf7, 0x43, 0x82, 0xe5, 0x89, 0x0d, 0x31, 0xbe, 0x63, 0x28, 0x90, 0x90, 0xb8, 0x4c, 0x5c,
	0xfd, 0x4e, 0xb6, 0xab, 0x4c, 0x6e, 0xe2, 0x55, 0x8d, 0xf8, 0xd4, 0x88, 0x75, 0xd0, 0x6b, 0x98,
	0x13, 0x03, 0x16, 0x8d, 0x3e, 0x72, 0x1d, 0x09, 0x1a, 0x3f, 0x19, 0xa5, 0x32, 0x16, 0x44, 0x25,
	0x28, 0x74, 0x4e, 0x4e, 0xb4, 0x0f, 0xb5, 0x5c, 0x14, 0x96, 0xae, 0xaa, 0xa9, 0xef, 0xd5, 0x9a,
	0xd4, 0xfe, 0x99, 0x87, 0x85, 0x64, 0x4a, 0x24, 0x08, 0x1d, 0x8b, 0x20, 0x1b, 0x2a, 0xe9, 0x6c,
	0xa0, 0xa7, 0xd9, 0x47, 0x66, 0xc4, 0xaa, 0xb9, 0x3d, 0x0d, 0x34, 0xee, 0x17, 0xe7, 0xd0, 0x39,
	0xc0, 0xef, 0xc4, 0xa0, 0xcd, 0x6c, 0xee, 0x1f, 0x99, 0x6a, 0x3e, 0x3c, 0x02, 0x9c, 0x43, 0x97,
	0x50, 0x49, 0x07, 0xec, 0xbe, 0x06, 0x32, 0x42, 0xf8, 0xb8, 0xf6, 0x19, 0x54, 0xc7, 0xf2, 0x83,
	0xee, 0x69, 0x39, 0x2b, 0x64, 0xcd, 0x15, 0x39, 0x7e, 0x94, 0xe5, 0xe4, 0x51, 0x96, 0xd5, 0xe8,
	0x51, 0xc6, 0x39, 0xd4, 0x87, 0xea, 0x58, 0x2a, 0xee, 0x93, 0xcd, 0xca, 0x63, 0xf3, 0xd9, 0x5f,
	0xc4, 0x0c, 0xe7, 0x5a, 0xd2, 0xde, 0xee, 0xe5, 0x8e, 0xed, 0xb0, 0x9b, 0xe1, 0xb5, 0x6c, 0x79,
	0x03, 0x45, 0x90, 0xe3, 0x3f, 0x85, 0x6c, 0x7b, 0xca, 0xc4, 0xef, 0x66, 0xd7, 0x4f, 0x17, 0xae,
	0x8b, 0x1c, 0xf7, 0xe2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x3c, 0x7c, 0x7b, 0x96, 0x06,
	0x00, 0x00,
}
