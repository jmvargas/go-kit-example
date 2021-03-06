// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type GetRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetReply struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{1}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListReply struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListReply) Reset()         { *m = ListReply{} }
func (m *ListReply) String() string { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()    {}
func (*ListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{3}
}

func (m *ListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReply.Unmarshal(m, b)
}
func (m *ListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReply.Marshal(b, m, deterministic)
}
func (m *ListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReply.Merge(m, src)
}
func (m *ListReply) XXX_Size() int {
	return xxx_messageInfo_ListReply.Size(m)
}
func (m *ListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListReply proto.InternalMessageInfo

func (m *ListReply) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type Product struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price                float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{4}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetReply)(nil), "pb.GetReply")
	proto.RegisterType((*ListRequest)(nil), "pb.ListRequest")
	proto.RegisterType((*ListReply)(nil), "pb.ListReply")
	proto.RegisterType((*Product)(nil), "pb.Product")
}

func init() { proto.RegisterFile("catalog.proto", fileDescriptor_0abbfcf058acdf89) }

var fileDescriptor_0abbfcf058acdf89 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x14, 0xc4, 0x49, 0xea, 0xda, 0xdd, 0x57, 0xbb, 0x42, 0xf0, 0x10, 0x16, 0x0f, 0x21, 0xa2, 0xf6,
	0x54, 0x70, 0xf5, 0x1b, 0x88, 0xec, 0xc5, 0x83, 0xe4, 0xe0, 0x7d, 0xdb, 0x06, 0x09, 0x04, 0x12,
	0xdb, 0xb7, 0x87, 0x7e, 0x7b, 0x49, 0xd3, 0x3f, 0xd2, 0xe3, 0xfc, 0xe6, 0x31, 0x99, 0x09, 0xe4,
	0xf5, 0x19, 0xcf, 0xd6, 0xfd, 0x94, 0xbe, 0x75, 0xe8, 0x18, 0xf5, 0x95, 0xbc, 0x07, 0x38, 0x69,
	0x54, 0xfa, 0xf7, 0xa2, 0x3b, 0x64, 0x7b, 0xa0, 0xa6, 0xe1, 0x44, 0x90, 0x62, 0xa3, 0xa8, 0x69,
	0xe4, 0x0b, 0x6c, 0x07, 0xd7, 0xdb, 0x9e, 0x3d, 0x42, 0xea, 0x5b, 0xd7, 0x5c, 0x6a, 0x1c, 0x0e,
	0xb2, 0x63, 0x56, 0xfa, 0xaa, 0xfc, 0x8a, 0x48, 0x4d, 0x9e, 0xcc, 0x21, 0xfb, 0x34, 0xdd, 0x94,
	0x28, 0xdf, 0x60, 0x17, 0x65, 0x88, 0x78, 0x86, 0xed, 0x78, 0xd6, 0x71, 0x22, 0x92, 0x75, 0xc6,
	0x6c, 0xca, 0x0f, 0x48, 0x47, 0xb8, 0xae, 0xc4, 0xee, 0x60, 0x83, 0x06, 0xad, 0xe6, 0x54, 0x90,
	0x62, 0xa7, 0xa2, 0x08, 0xd4, 0xb7, 0xa6, 0xd6, 0x3c, 0x11, 0xa4, 0xa0, 0x2a, 0x8a, 0xe3, 0x37,
	0xa4, 0xef, 0x71, 0x31, 0x7b, 0x80, 0xe4, 0xa4, 0x91, 0xed, 0xc3, 0x7b, 0xcb, 0xe0, 0xc3, 0xcd,
	0xac, 0x43, 0xbf, 0x27, 0xb8, 0x0a, 0x65, 0xd9, 0x6d, 0xa0, 0xff, 0x56, 0x1c, 0xf2, 0x05, 0x78,
	0xdb, 0x57, 0xd7, 0xc3, 0xff, 0xbd, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x2b, 0xe0, 0x18,
	0x50, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
}

type catalogClient struct {
	cc *grpc.ClientConn
}

func NewCatalogClient(cc *grpc.ClientConn) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/pb.Catalog/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/pb.Catalog/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
type CatalogServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
}

func RegisterCatalogServer(s *grpc.Server, srv CatalogServer) {
	s.RegisterService(&_Catalog_serviceDesc, srv)
}

func _Catalog_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catalog/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catalog/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Catalog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Catalog_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Catalog_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog.proto",
}
