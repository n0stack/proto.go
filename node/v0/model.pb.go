// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node/v0/model.proto

/*
Package pnode is a generated protocol buffer package.

It is generated from these files:
	node/v0/model.proto

It has these top-level messages:
	Node
	Status
	ListRequest
	ListResponse
	GetRequest
*/
package pnode

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Node struct {
	// MotherBoardのUUIDを用いる
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Status *Status `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type Status struct {
	// Architectureは全部x64だろうという気持ち
	Cpus uint32 `protobuf:"varint,1,opt,name=cpus" json:"cpus,omitempty"`
	// Memory
	MemoryBytes uint64 `protobuf:"varint,2,opt,name=memoryBytes" json:"memoryBytes,omitempty"`
	// Storage
	StorageBytes uint64 `protobuf:"varint,3,opt,name=storageBytes" json:"storageBytes,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Status) GetCpus() uint32 {
	if m != nil {
		return m.Cpus
	}
	return 0
}

func (m *Status) GetMemoryBytes() uint64 {
	if m != nil {
		return m.MemoryBytes
	}
	return 0
}

func (m *Status) GetStorageBytes() uint64 {
	if m != nil {
		return m.StorageBytes
	}
	return 0
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListResponse struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Node)(nil), "n0stack.node.Node")
	proto.RegisterType((*Status)(nil), "n0stack.node.Status")
	proto.RegisterType((*ListRequest)(nil), "n0stack.node.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "n0stack.node.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "n0stack.node.GetRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeService service

type NodeServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Node, error)
}

type nodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeServiceClient(cc *grpc.ClientConn) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/n0stack.node.NodeService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/n0stack.node.NodeService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeService service

type NodeServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Get(context.Context, *GetRequest) (*Node, error)
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.NodeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.NodeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.node.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _NodeService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NodeService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/v0/model.proto",
}

func init() { proto.RegisterFile("node/v0/model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0x36, 0x16, 0x9c, 0xa4, 0x1e, 0x46, 0x0f, 0xb1, 0xa7, 0xb0, 0xa7, 0x88, 0x92,
	0x94, 0x88, 0x20, 0x78, 0x10, 0x8a, 0xd0, 0x8b, 0x78, 0xd8, 0xde, 0xbc, 0xe5, 0xcf, 0x10, 0x83,
	0x26, 0x1b, 0xb3, 0x9b, 0x42, 0x3f, 0x80, 0xdf, 0x5b, 0x36, 0x1b, 0x31, 0x81, 0xde, 0x26, 0xef,
	0xfd, 0xf2, 0xe6, 0x2d, 0x03, 0x97, 0xb5, 0xc8, 0x29, 0x3a, 0x6c, 0xa2, 0x4a, 0xe4, 0xf4, 0x15,
	0x36, 0xad, 0x50, 0x02, 0xdd, 0x7a, 0x23, 0x55, 0x92, 0x7d, 0x86, 0xda, 0x64, 0x2f, 0x60, 0xbf,
	0x89, 0x9c, 0xf0, 0x02, 0xe6, 0x65, 0xee, 0x59, 0xbe, 0x15, 0x9c, 0xf3, 0x79, 0x99, 0xe3, 0x1d,
	0x2c, 0xa5, 0x4a, 0x54, 0x27, 0x3d, 0xdb, 0xb7, 0x02, 0x27, 0xbe, 0x0a, 0xc7, 0xbf, 0x85, 0xfb,
	0xde, 0xe3, 0x03, 0xc3, 0x52, 0x58, 0x1a, 0x05, 0x11, 0xec, 0xac, 0xe9, 0x64, 0x9f, 0xb4, 0xe2,
	0xfd, 0x8c, 0x3e, 0x38, 0x15, 0x55, 0xa2, 0x3d, 0x6e, 0x8f, 0x8a, 0xa4, 0x37, 0xf7, 0xad, 0xc0,
	0xe6, 0x63, 0x09, 0x19, 0xb8, 0x52, 0x89, 0x36, 0x29, 0xc8, 0x20, 0x8b, 0x1e, 0x99, 0x68, 0x6c,
	0x05, 0xce, 0x6b, 0x29, 0x15, 0xa7, 0xef, 0x8e, 0xa4, 0x62, 0x8f, 0xe0, 0x9a, 0x4f, 0xd9, 0x88,
	0x5a, 0x12, 0x06, 0x70, 0xa6, 0x9b, 0xe9, 0xcd, 0x8b, 0xc0, 0x89, 0x71, 0xda, 0x57, 0xbf, 0x91,
	0x1b, 0x80, 0xb9, 0x00, 0x3b, 0xfa, 0xcb, 0x89, 0x7f, 0x2c, 0x70, 0xb4, 0xbb, 0xa7, 0xf6, 0x50,
	0x66, 0x84, 0xcf, 0x60, 0xeb, 0x5c, 0xbc, 0x9e, 0x06, 0x8c, 0x56, 0xaf, 0xd7, 0xa7, 0x2c, 0x53,
	0x83, 0xcd, 0xf0, 0x01, 0x16, 0x3b, 0x52, 0xe8, 0x4d, 0xa1, 0xff, 0x8d, 0xeb, 0x13, 0xd5, 0xd8,
	0x6c, 0x7b, 0xfb, 0x7e, 0x53, 0x94, 0xea, 0xa3, 0x4b, 0xc3, 0x4c, 0x54, 0xd1, 0x40, 0x44, 0x85,
	0x30, 0x57, 0x8b, 0x86, 0x4b, 0x3e, 0x35, 0x7a, 0x48, 0x97, 0xbd, 0x78, 0xff, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x6f, 0x33, 0xa1, 0x38, 0xe1, 0x01, 0x00, 0x00,
}
