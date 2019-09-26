// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/article.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ArticleRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	UrlTitle             string   `protobuf:"bytes,3,opt,name=url_title,json=urlTitle,proto3" json:"url_title,omitempty"`
	Author               string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Year                 int32    `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	Month                int32    `protobuf:"varint,6,opt,name=month,proto3" json:"month,omitempty"`
	Day                  int32    `protobuf:"varint,7,opt,name=day,proto3" json:"day,omitempty"`
	Hour                 float64  `protobuf:"fixed64,8,opt,name=hour,proto3" json:"hour,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleRequest) Reset()         { *m = ArticleRequest{} }
func (m *ArticleRequest) String() string { return proto.CompactTextString(m) }
func (*ArticleRequest) ProtoMessage()    {}
func (*ArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d61e4162546950e, []int{0}
}

func (m *ArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleRequest.Unmarshal(m, b)
}
func (m *ArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleRequest.Marshal(b, m, deterministic)
}
func (m *ArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleRequest.Merge(m, src)
}
func (m *ArticleRequest) XXX_Size() int {
	return xxx_messageInfo_ArticleRequest.Size(m)
}
func (m *ArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleRequest proto.InternalMessageInfo

func (m *ArticleRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ArticleRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArticleRequest) GetUrlTitle() string {
	if m != nil {
		return m.UrlTitle
	}
	return ""
}

func (m *ArticleRequest) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *ArticleRequest) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *ArticleRequest) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *ArticleRequest) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *ArticleRequest) GetHour() float64 {
	if m != nil {
		return m.Hour
	}
	return 0
}

// Multiple Articles
type ArticleList struct {
	Articles             []*Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ArticleList) Reset()         { *m = ArticleList{} }
func (m *ArticleList) String() string { return proto.CompactTextString(m) }
func (*ArticleList) ProtoMessage()    {}
func (*ArticleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d61e4162546950e, []int{1}
}

func (m *ArticleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleList.Unmarshal(m, b)
}
func (m *ArticleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleList.Marshal(b, m, deterministic)
}
func (m *ArticleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleList.Merge(m, src)
}
func (m *ArticleList) XXX_Size() int {
	return xxx_messageInfo_ArticleList.Size(m)
}
func (m *ArticleList) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleList.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleList proto.InternalMessageInfo

func (m *ArticleList) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type Article struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	UrlTitle             string   `protobuf:"bytes,3,opt,name=url_title,json=urlTitle,proto3" json:"url_title,omitempty"`
	Author               string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Year                 int32    `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	Month                int32    `protobuf:"varint,6,opt,name=month,proto3" json:"month,omitempty"`
	Day                  int32    `protobuf:"varint,7,opt,name=day,proto3" json:"day,omitempty"`
	Hour                 int32    `protobuf:"varint,8,opt,name=hour,proto3" json:"hour,omitempty"`
	Minute               int32    `protobuf:"varint,9,opt,name=minute,proto3" json:"minute,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d61e4162546950e, []int{2}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetUrlTitle() string {
	if m != nil {
		return m.UrlTitle
	}
	return ""
}

func (m *Article) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Article) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Article) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Article) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *Article) GetHour() int32 {
	if m != nil {
		return m.Hour
	}
	return 0
}

func (m *Article) GetMinute() int32 {
	if m != nil {
		return m.Minute
	}
	return 0
}

type ArticleOffset struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleOffset) Reset()         { *m = ArticleOffset{} }
func (m *ArticleOffset) String() string { return proto.CompactTextString(m) }
func (*ArticleOffset) ProtoMessage()    {}
func (*ArticleOffset) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d61e4162546950e, []int{3}
}

func (m *ArticleOffset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleOffset.Unmarshal(m, b)
}
func (m *ArticleOffset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleOffset.Marshal(b, m, deterministic)
}
func (m *ArticleOffset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleOffset.Merge(m, src)
}
func (m *ArticleOffset) XXX_Size() int {
	return xxx_messageInfo_ArticleOffset.Size(m)
}
func (m *ArticleOffset) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleOffset.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleOffset proto.InternalMessageInfo

func (m *ArticleOffset) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ArticleOffset) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*ArticleRequest)(nil), "v1.ArticleRequest")
	proto.RegisterType((*ArticleList)(nil), "v1.ArticleList")
	proto.RegisterType((*Article)(nil), "v1.Article")
	proto.RegisterType((*ArticleOffset)(nil), "v1.ArticleOffset")
}

func init() { proto.RegisterFile("proto/v1/article.proto", fileDescriptor_1d61e4162546950e) }

var fileDescriptor_1d61e4162546950e = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x92, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0xc6, 0xa3, 0x24, 0x76, 0x92, 0x33, 0xfd, 0x27, 0x4a, 0x10, 0xe9, 0x62, 0xb4, 0xd4, 0x53,
	0x82, 0x13, 0x28, 0x74, 0xec, 0xd4, 0xa5, 0x50, 0x70, 0xbb, 0x17, 0x37, 0x55, 0xb0, 0xc0, 0x89,
	0x52, 0x59, 0x32, 0xe4, 0xe1, 0xfa, 0x0a, 0x7d, 0xa6, 0xa2, 0x93, 0x92, 0x7a, 0xe8, 0x0b, 0x74,
	0xbb, 0xef, 0xf7, 0x21, 0xfb, 0xbe, 0xbb, 0x83, 0xe9, 0x5e, 0x2b, 0xa3, 0x16, 0x6d, 0xbe, 0x28,
	0xb5, 0x91, 0xeb, 0x5a, 0xcc, 0x11, 0xd0, 0x7e, 0x9b, 0xf3, 0x2f, 0x02, 0xe7, 0x0f, 0x9e, 0x16,
	0xe2, 0xd3, 0x8a, 0xc6, 0xd0, 0x4b, 0x18, 0x94, 0x7b, 0xc9, 0x48, 0x4a, 0xb2, 0x49, 0xe1, 0x4a,
	0x7a, 0x0d, 0x91, 0x91, 0xa6, 0x16, 0xac, 0x8f, 0xcc, 0x0b, 0x7a, 0x03, 0x13, 0xab, 0xeb, 0x37,
	0xef, 0x0c, 0xd0, 0x19, 0x5b, 0x5d, 0xbf, 0xa2, 0x39, 0x85, 0xb8, 0xb4, 0xa6, 0x52, 0x9a, 0x0d,
	0xd1, 0x09, 0x8a, 0x52, 0x18, 0x1e, 0x44, 0xa9, 0x59, 0x94, 0x92, 0x2c, 0x2a, 0xb0, 0x76, 0x9f,
	0xdf, 0xaa, 0x9d, 0xa9, 0x58, 0x8c, 0xd0, 0x0b, 0xd7, 0xc6, 0x47, 0x79, 0x60, 0x23, 0x64, 0xae,
	0x74, 0x6f, 0x2b, 0x65, 0x35, 0x1b, 0xa7, 0x24, 0x23, 0x05, 0xd6, 0xfc, 0x0e, 0x92, 0xd0, 0xfe,
	0x93, 0x6c, 0x0c, 0xbd, 0x85, 0x71, 0xc8, 0xd8, 0x30, 0x92, 0x0e, 0xb2, 0x64, 0x99, 0xcc, 0xdb,
	0x7c, 0x7e, 0x4c, 0x78, 0x32, 0xf9, 0x37, 0x81, 0x51, 0xa0, 0xff, 0x23, 0x70, 0xe4, 0x03, 0xbb,
	0xff, 0x6c, 0xe5, 0xce, 0x1a, 0xc1, 0x26, 0x48, 0x83, 0xe2, 0xf7, 0x70, 0x16, 0xf2, 0x3c, 0x6f,
	0x36, 0x8d, 0xf8, 0x6b, 0x8d, 0x53, 0x88, 0x15, 0x7a, 0x18, 0x2b, 0x2a, 0x82, 0x5a, 0xb6, 0xa7,
	0x13, 0x78, 0x11, 0xba, 0x95, 0x6b, 0x41, 0x17, 0x00, 0x8f, 0xc2, 0x1c, 0xe7, 0x43, 0xbb, 0x23,
	0xf4, 0x47, 0x32, 0xeb, 0x8e, 0x95, 0xf7, 0xe8, 0x0a, 0x92, 0xdf, 0x07, 0x0d, 0xbd, 0xea, 0xb8,
	0xbe, 0x9d, 0xd9, 0x45, 0x07, 0xb9, 0x55, 0xf1, 0xde, 0x7b, 0x8c, 0x67, 0xb8, 0xfa, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x52, 0x7b, 0xaf, 0x9f, 0xa0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleServiceClient interface {
	GetArticle(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*Article, error)
	GetArticles(ctx context.Context, in *ArticleOffset, opts ...grpc.CallOption) (*ArticleList, error)
}

type articleServiceClient struct {
	cc *grpc.ClientConn
}

func NewArticleServiceClient(cc *grpc.ClientConn) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) GetArticle(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/v1.ArticleService/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticles(ctx context.Context, in *ArticleOffset, opts ...grpc.CallOption) (*ArticleList, error) {
	out := new(ArticleList)
	err := c.cc.Invoke(ctx, "/v1.ArticleService/GetArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
type ArticleServiceServer interface {
	GetArticle(context.Context, *ArticleRequest) (*Article, error)
	GetArticles(context.Context, *ArticleOffset) (*ArticleList, error)
}

// UnimplementedArticleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (*UnimplementedArticleServiceServer) GetArticle(ctx context.Context, req *ArticleRequest) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (*UnimplementedArticleServiceServer) GetArticles(ctx context.Context, req *ArticleOffset) (*ArticleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}

func RegisterArticleServiceServer(s *grpc.Server, srv ArticleServiceServer) {
	s.RegisterService(&_ArticleService_serviceDesc, srv)
}

func _ArticleService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ArticleService/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticle(ctx, req.(*ArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleOffset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ArticleService/GetArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticles(ctx, req.(*ArticleOffset))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArticleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticle",
			Handler:    _ArticleService_GetArticle_Handler,
		},
		{
			MethodName: "GetArticles",
			Handler:    _ArticleService_GetArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/article.proto",
}
