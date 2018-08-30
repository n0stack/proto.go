// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provisioning/v0/image.proto

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

type Image struct {
	Metadata             *v0.Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec                 *ImageSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status               *ImageStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_ec488839fae0cb5a, []int{0}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Image) GetSpec() *ImageSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Image) GetStatus() *ImageStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// 途中
type ImageSpec struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageSpec) Reset()         { *m = ImageSpec{} }
func (m *ImageSpec) String() string { return proto.CompactTextString(m) }
func (*ImageSpec) ProtoMessage()    {}
func (*ImageSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_ec488839fae0cb5a, []int{1}
}
func (m *ImageSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageSpec.Unmarshal(m, b)
}
func (m *ImageSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageSpec.Marshal(b, m, deterministic)
}
func (dst *ImageSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageSpec.Merge(dst, src)
}
func (m *ImageSpec) XXX_Size() int {
	return xxx_messageInfo_ImageSpec.Size(m)
}
func (m *ImageSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ImageSpec proto.InternalMessageInfo

type ImageStatus struct {
	Revisions            map[string]string `protobuf:"bytes,1,rep,name=revisions" json:"revisions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ImageStatus) Reset()         { *m = ImageStatus{} }
func (m *ImageStatus) String() string { return proto.CompactTextString(m) }
func (*ImageStatus) ProtoMessage()    {}
func (*ImageStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_ec488839fae0cb5a, []int{2}
}
func (m *ImageStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageStatus.Unmarshal(m, b)
}
func (m *ImageStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageStatus.Marshal(b, m, deterministic)
}
func (dst *ImageStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageStatus.Merge(dst, src)
}
func (m *ImageStatus) XXX_Size() int {
	return xxx_messageInfo_ImageStatus.Size(m)
}
func (m *ImageStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ImageStatus proto.InternalMessageInfo

func (m *ImageStatus) GetRevisions() map[string]string {
	if m != nil {
		return m.Revisions
	}
	return nil
}

type ListImagesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListImagesRequest) Reset()         { *m = ListImagesRequest{} }
func (m *ListImagesRequest) String() string { return proto.CompactTextString(m) }
func (*ListImagesRequest) ProtoMessage()    {}
func (*ListImagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_ec488839fae0cb5a, []int{3}
}
func (m *ListImagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListImagesRequest.Unmarshal(m, b)
}
func (m *ListImagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListImagesRequest.Marshal(b, m, deterministic)
}
func (dst *ListImagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListImagesRequest.Merge(dst, src)
}
func (m *ListImagesRequest) XXX_Size() int {
	return xxx_messageInfo_ListImagesRequest.Size(m)
}
func (m *ListImagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListImagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListImagesRequest proto.InternalMessageInfo

type ListImagesResponse struct {
	Images               []*Image `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListImagesResponse) Reset()         { *m = ListImagesResponse{} }
func (m *ListImagesResponse) String() string { return proto.CompactTextString(m) }
func (*ListImagesResponse) ProtoMessage()    {}
func (*ListImagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_ec488839fae0cb5a, []int{4}
}
func (m *ListImagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListImagesResponse.Unmarshal(m, b)
}
func (m *ListImagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListImagesResponse.Marshal(b, m, deterministic)
}
func (dst *ListImagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListImagesResponse.Merge(dst, src)
}
func (m *ListImagesResponse) XXX_Size() int {
	return xxx_messageInfo_ListImagesResponse.Size(m)
}
func (m *ListImagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListImagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListImagesResponse proto.InternalMessageInfo

func (m *ListImagesResponse) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type GetImageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetImageRequest) Reset()         { *m = GetImageRequest{} }
func (m *GetImageRequest) String() string { return proto.CompactTextString(m) }
func (*GetImageRequest) ProtoMessage()    {}
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_ec488839fae0cb5a, []int{5}
}
func (m *GetImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetImageRequest.Unmarshal(m, b)
}
func (m *GetImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetImageRequest.Marshal(b, m, deterministic)
}
func (dst *GetImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetImageRequest.Merge(dst, src)
}
func (m *GetImageRequest) XXX_Size() int {
	return xxx_messageInfo_GetImageRequest.Size(m)
}
func (m *GetImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetImageRequest proto.InternalMessageInfo

func (m *GetImageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ApplyImageRequest struct {
	Metadata             *v0.Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec                 *ImageSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ApplyImageRequest) Reset()         { *m = ApplyImageRequest{} }
func (m *ApplyImageRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyImageRequest) ProtoMessage()    {}
func (*ApplyImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_ec488839fae0cb5a, []int{6}
}
func (m *ApplyImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyImageRequest.Unmarshal(m, b)
}
func (m *ApplyImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyImageRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyImageRequest.Merge(dst, src)
}
func (m *ApplyImageRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyImageRequest.Size(m)
}
func (m *ApplyImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyImageRequest proto.InternalMessageInfo

func (m *ApplyImageRequest) GetMetadata() *v0.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ApplyImageRequest) GetSpec() *ImageSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type DeleteImageRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteImageRequest) Reset()         { *m = DeleteImageRequest{} }
func (m *DeleteImageRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteImageRequest) ProtoMessage()    {}
func (*DeleteImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_ec488839fae0cb5a, []int{7}
}
func (m *DeleteImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteImageRequest.Unmarshal(m, b)
}
func (m *DeleteImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteImageRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteImageRequest.Merge(dst, src)
}
func (m *DeleteImageRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteImageRequest.Size(m)
}
func (m *DeleteImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteImageRequest proto.InternalMessageInfo

func (m *DeleteImageRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Image)(nil), "n0stack.provisioning.Image")
	proto.RegisterType((*ImageSpec)(nil), "n0stack.provisioning.ImageSpec")
	proto.RegisterType((*ImageStatus)(nil), "n0stack.provisioning.ImageStatus")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.provisioning.ImageStatus.RevisionsEntry")
	proto.RegisterType((*ListImagesRequest)(nil), "n0stack.provisioning.ListImagesRequest")
	proto.RegisterType((*ListImagesResponse)(nil), "n0stack.provisioning.ListImagesResponse")
	proto.RegisterType((*GetImageRequest)(nil), "n0stack.provisioning.GetImageRequest")
	proto.RegisterType((*ApplyImageRequest)(nil), "n0stack.provisioning.ApplyImageRequest")
	proto.RegisterType((*DeleteImageRequest)(nil), "n0stack.provisioning.DeleteImageRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ImageService service

type ImageServiceClient interface {
	ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error)
	ApplyImage(ctx context.Context, in *ApplyImageRequest, opts ...grpc.CallOption) (*Image, error)
	DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type imageServiceClient struct {
	cc *grpc.ClientConn
}

func NewImageServiceClient(cc *grpc.ClientConn) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	out := new(ListImagesResponse)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.ImageService/ListImages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.ImageService/GetImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ApplyImage(ctx context.Context, in *ApplyImageRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.ImageService/ApplyImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.provisioning.ImageService/DeleteImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImageService service

type ImageServiceServer interface {
	ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	GetImage(context.Context, *GetImageRequest) (*Image, error)
	ApplyImage(context.Context, *ApplyImageRequest) (*Image, error)
	DeleteImage(context.Context, *DeleteImageRequest) (*empty.Empty, error)
}

func RegisterImageServiceServer(s *grpc.Server, srv ImageServiceServer) {
	s.RegisterService(&_ImageService_serviceDesc, srv)
}

func _ImageService_ListImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ListImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.ImageService/ListImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ListImages(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.ImageService/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).GetImage(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ApplyImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ApplyImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.ImageService/ApplyImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ApplyImage(ctx, req.(*ApplyImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.provisioning.ImageService/DeleteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).DeleteImage(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.provisioning.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListImages",
			Handler:    _ImageService_ListImages_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _ImageService_GetImage_Handler,
		},
		{
			MethodName: "ApplyImage",
			Handler:    _ImageService_ApplyImage_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _ImageService_DeleteImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provisioning/v0/image.proto",
}

func init() { proto.RegisterFile("provisioning/v0/image.proto", fileDescriptor_image_ec488839fae0cb5a) }

var fileDescriptor_image_ec488839fae0cb5a = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0x36, 0x6a, 0xc6, 0x08, 0xc8, 0x50, 0xa1, 0xc8, 0x39, 0x50, 0x56, 0xaa, 0xc8,
	0x85, 0x75, 0x94, 0x5c, 0x28, 0xe5, 0x02, 0xa2, 0x42, 0x95, 0x00, 0x81, 0xe1, 0xc4, 0xcd, 0x31,
	0x83, 0xb1, 0x1a, 0x7b, 0x17, 0xef, 0xda, 0x28, 0xbf, 0x85, 0x3b, 0xbf, 0x8e, 0x1f, 0x81, 0xb2,
	0x6b, 0x37, 0x6e, 0xf3, 0x75, 0xeb, 0x6d, 0x3d, 0xf3, 0xde, 0x9b, 0xe7, 0x37, 0xd2, 0xc0, 0x50,
	0xe6, 0xa2, 0x4c, 0x54, 0x22, 0xb2, 0x24, 0x8b, 0xfd, 0x72, 0xec, 0x27, 0x69, 0x18, 0x13, 0x97,
	0xb9, 0xd0, 0x02, 0x8f, 0xb3, 0xb1, 0xd2, 0x61, 0x74, 0xc5, 0x9b, 0x20, 0x6f, 0x18, 0x0b, 0x11,
	0xcf, 0xc9, 0x37, 0x98, 0x59, 0xf1, 0xc3, 0xa7, 0x54, 0xea, 0x85, 0xa5, 0x78, 0xfd, 0x72, 0xec,
	0xa7, 0xa4, 0xc3, 0xef, 0xa1, 0x0e, 0x6d, 0x89, 0xfd, 0x75, 0xe0, 0xf0, 0x72, 0xa9, 0x8a, 0xcf,
	0xe1, 0xa8, 0xee, 0x0d, 0x9c, 0x13, 0x67, 0xe4, 0x4e, 0xfa, 0xbc, 0x1e, 0xf1, 0xa1, 0x6a, 0x04,
	0xd7, 0x10, 0x9c, 0xc2, 0x81, 0x92, 0x14, 0x0d, 0xda, 0x06, 0xfa, 0x84, 0x6f, 0x72, 0xc3, 0x8d,
	0xf2, 0x17, 0x49, 0x51, 0x60, 0xc0, 0x78, 0x06, 0x5d, 0xa5, 0x43, 0x5d, 0xa8, 0x41, 0xc7, 0xd0,
	0x9e, 0xee, 0xa2, 0x19, 0x60, 0x50, 0x11, 0x98, 0x0b, 0xbd, 0x6b, 0x35, 0xf6, 0xc7, 0x01, 0xb7,
	0x01, 0xc2, 0x8f, 0xd0, 0xcb, 0xc9, 0xf2, 0xd5, 0xc0, 0x39, 0xe9, 0x8c, 0xdc, 0xc9, 0x78, 0xaf,
	0x34, 0x0f, 0x6a, 0xca, 0x45, 0xa6, 0xf3, 0x45, 0xb0, 0x92, 0xf0, 0x5e, 0xc1, 0xfd, 0x9b, 0x4d,
	0x7c, 0x08, 0x9d, 0x2b, 0x5a, 0x98, 0x60, 0x7a, 0xc1, 0xf2, 0x89, 0xc7, 0x70, 0x58, 0x86, 0xf3,
	0x82, 0x4c, 0x02, 0xbd, 0xc0, 0x7e, 0xbc, 0x6c, 0xbf, 0x70, 0xd8, 0x23, 0xe8, 0xbf, 0x4f, 0x94,
	0x36, 0xa3, 0x54, 0x40, 0xbf, 0x0a, 0x52, 0x9a, 0x5d, 0x02, 0x36, 0x8b, 0x4a, 0x8a, 0x4c, 0x11,
	0x4e, 0xa1, 0x6b, 0x76, 0x5a, 0xbb, 0x1e, 0xee, 0x70, 0x1d, 0x54, 0x50, 0x76, 0x0a, 0x0f, 0xde,
	0x91, 0x55, 0xaa, 0xd4, 0x11, 0xe1, 0x20, 0x0b, 0x53, 0xaa, 0xfc, 0x99, 0x37, 0xfb, 0x0d, 0xfd,
	0xd7, 0x52, 0xce, 0x17, 0x37, 0x80, 0x77, 0xb0, 0x65, 0x36, 0x02, 0x7c, 0x4b, 0x73, 0xd2, 0xb4,
	0xcf, 0xe2, 0xe4, 0x5f, 0x1b, 0xee, 0x59, 0x36, 0xe5, 0x65, 0x12, 0x11, 0x86, 0x00, 0xab, 0x94,
	0xf0, 0xd9, 0xe6, 0x79, 0x6b, 0xe1, 0x7a, 0xa3, 0xfd, 0x40, 0x1b, 0x38, 0x6b, 0xe1, 0x27, 0x38,
	0xaa, 0xd3, 0xc3, 0xd3, 0xcd, 0xbc, 0x5b, 0xe9, 0x7a, 0xbb, 0xb6, 0xc2, 0x5a, 0xf8, 0x15, 0x60,
	0x15, 0xf4, 0x36, 0xd3, 0x6b, 0xab, 0xd8, 0xa7, 0xfa, 0x19, 0xdc, 0x46, 0x8a, 0xb8, 0xe5, 0x17,
	0xd7, 0x83, 0xf6, 0x1e, 0x73, 0x7b, 0x03, 0x78, 0x7d, 0x03, 0xf8, 0xc5, 0xf2, 0x06, 0xb0, 0xd6,
	0x9b, 0xf3, 0x6f, 0x67, 0x71, 0xa2, 0x7f, 0x16, 0x33, 0x1e, 0x89, 0xd4, 0xaf, 0xf4, 0xec, 0xa9,
	0xe0, 0xb1, 0xf0, 0x6f, 0x5d, 0x9b, 0x73, 0xd9, 0x2c, 0xcc, 0xba, 0x06, 0x37, 0xfd, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0x82, 0x1d, 0xd2, 0xb5, 0x95, 0x04, 0x00, 0x00,
}
