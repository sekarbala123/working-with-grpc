// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product_info.proto

package ecommerce

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

type Product struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_info_de3a8245360e059f, []int{0}
}
func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (dst *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(dst, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ProductID struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductID) Reset()         { *m = ProductID{} }
func (m *ProductID) String() string { return proto.CompactTextString(m) }
func (*ProductID) ProtoMessage()    {}
func (*ProductID) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_info_de3a8245360e059f, []int{1}
}
func (m *ProductID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductID.Unmarshal(m, b)
}
func (m *ProductID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductID.Marshal(b, m, deterministic)
}
func (dst *ProductID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductID.Merge(dst, src)
}
func (m *ProductID) XXX_Size() int {
	return xxx_messageInfo_ProductID.Size(m)
}
func (m *ProductID) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductID.DiscardUnknown(m)
}

var xxx_messageInfo_ProductID proto.InternalMessageInfo

func (m *ProductID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Product)(nil), "ecommerce.Product")
	proto.RegisterType((*ProductID)(nil), "ecommerce.ProductID")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductInfoClient is the client API for ProductInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductInfoClient interface {
	AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductID, error)
	GetProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Product, error)
}

type productInfoClient struct {
	cc *grpc.ClientConn
}

func NewProductInfoClient(cc *grpc.ClientConn) ProductInfoClient {
	return &productInfoClient{cc}
}

func (c *productInfoClient) AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductID, error) {
	out := new(ProductID)
	err := c.cc.Invoke(ctx, "/ecommerce.ProductInfo/addProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productInfoClient) GetProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/ecommerce.ProductInfo/getProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductInfoServer is the server API for ProductInfo service.
type ProductInfoServer interface {
	AddProduct(context.Context, *Product) (*ProductID, error)
	GetProduct(context.Context, *ProductID) (*Product, error)
}

func RegisterProductInfoServer(s *grpc.Server, srv ProductInfoServer) {
	s.RegisterService(&_ProductInfo_serviceDesc, srv)
}

func _ProductInfo_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.ProductInfo/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).AddProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductInfo_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.ProductInfo/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).GetProduct(ctx, req.(*ProductID))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.ProductInfo",
	HandlerType: (*ProductInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addProduct",
			Handler:    _ProductInfo_AddProduct_Handler,
		},
		{
			MethodName: "getProduct",
			Handler:    _ProductInfo_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_info.proto",
}

func init() { proto.RegisterFile("product_info.proto", fileDescriptor_product_info_de3a8245360e059f) }

var fileDescriptor_product_info_de3a8245360e059f = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0x89, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x4c, 0x4d, 0xce, 0xcf, 0xcd, 0x4d, 0x2d, 0x4a, 0x4e, 0x55, 0xf2, 0xe7, 0x62, 0x0f, 0x80, 0x28,
	0x10, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c,
	0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0x0a,
	0x5c, 0xdc, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x12, 0xcc, 0x60,
	0x29, 0x64, 0x21, 0x25, 0x45, 0x2e, 0x4e, 0xa8, 0x81, 0x9e, 0x2e, 0x42, 0x22, 0x5c, 0xac, 0x65,
	0x89, 0x39, 0xa5, 0xa9, 0x50, 0x53, 0x21, 0x1c, 0xa3, 0x5a, 0x2e, 0x6e, 0x98, 0x92, 0xbc, 0xb4,
	0x7c, 0x21, 0x33, 0x2e, 0xae, 0xc4, 0x94, 0x14, 0x98, 0x2b, 0x84, 0xf4, 0xe0, 0x8e, 0xd3, 0x83,
	0x8a, 0x49, 0x89, 0x60, 0x8a, 0x79, 0xba, 0x80, 0xf4, 0xa5, 0xa7, 0x96, 0xc0, 0xf4, 0x61, 0x55,
	0x23, 0x85, 0xc5, 0xb4, 0x24, 0x36, 0x70, 0x20, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3b,
	0xb6, 0x06, 0x95, 0x1a, 0x01, 0x00, 0x00,
}
