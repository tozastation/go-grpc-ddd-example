// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user // import "github.com/tozastation/gRPC-Training-Golang/interfaces/rpc/user"

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

type PostUser struct {
	// tozastation
	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	// Hakodate
	CityName string `protobuf:"bytes,2,opt,name=CityName,proto3" json:"CityName,omitempty"`
	// test
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	// RyoTozawa
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostUser) Reset()         { *m = PostUser{} }
func (m *PostUser) String() string { return proto.CompactTextString(m) }
func (*PostUser) ProtoMessage()    {}
func (*PostUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_60d31a5782020fab, []int{0}
}
func (m *PostUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostUser.Unmarshal(m, b)
}
func (m *PostUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostUser.Marshal(b, m, deterministic)
}
func (dst *PostUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostUser.Merge(dst, src)
}
func (m *PostUser) XXX_Size() int {
	return xxx_messageInfo_PostUser.Size(m)
}
func (m *PostUser) XXX_DiscardUnknown() {
	xxx_messageInfo_PostUser.DiscardUnknown(m)
}

var xxx_messageInfo_PostUser proto.InternalMessageInfo

func (m *PostUser) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *PostUser) GetCityName() string {
	if m != nil {
		return m.CityName
	}
	return ""
}

func (m *PostUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PostUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetUser struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	CityName             string   `protobuf:"bytes,2,opt,name=CityName,proto3" json:"CityName,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUser) Reset()         { *m = GetUser{} }
func (m *GetUser) String() string { return proto.CompactTextString(m) }
func (*GetUser) ProtoMessage()    {}
func (*GetUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_60d31a5782020fab, []int{1}
}
func (m *GetUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUser.Unmarshal(m, b)
}
func (m *GetUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUser.Marshal(b, m, deterministic)
}
func (dst *GetUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUser.Merge(dst, src)
}
func (m *GetUser) XXX_Size() int {
	return xxx_messageInfo_GetUser.Size(m)
}
func (m *GetUser) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUser.DiscardUnknown(m)
}

var xxx_messageInfo_GetUser proto.InternalMessageInfo

func (m *GetUser) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetUser) GetCityName() string {
	if m != nil {
		return m.CityName
	}
	return ""
}

func (m *GetUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_60d31a5782020fab, []int{2}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetResponse struct {
	User                 *GetUser `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_60d31a5782020fab, []int{3}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetUser() *GetUser {
	if m != nil {
		return m.User
	}
	return nil
}

type LoginRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_60d31a5782020fab, []int{4}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	CityName             string   `protobuf:"bytes,1,opt,name=CityName,proto3" json:"CityName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_60d31a5782020fab, []int{5}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetCityName() string {
	if m != nil {
		return m.CityName
	}
	return ""
}

type PostRequest struct {
	User                 *PostUser `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PostRequest) Reset()         { *m = PostRequest{} }
func (m *PostRequest) String() string { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()    {}
func (*PostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_60d31a5782020fab, []int{6}
}
func (m *PostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostRequest.Unmarshal(m, b)
}
func (m *PostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostRequest.Marshal(b, m, deterministic)
}
func (dst *PostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostRequest.Merge(dst, src)
}
func (m *PostRequest) XXX_Size() int {
	return xxx_messageInfo_PostRequest.Size(m)
}
func (m *PostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostRequest proto.InternalMessageInfo

func (m *PostRequest) GetUser() *PostUser {
	if m != nil {
		return m.User
	}
	return nil
}

type PostResponse struct {
	// string Token = 1;
	CityName             string   `protobuf:"bytes,1,opt,name=CityName,proto3" json:"CityName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResponse) Reset()         { *m = PostResponse{} }
func (m *PostResponse) String() string { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()    {}
func (*PostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_60d31a5782020fab, []int{7}
}
func (m *PostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResponse.Unmarshal(m, b)
}
func (m *PostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResponse.Marshal(b, m, deterministic)
}
func (dst *PostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResponse.Merge(dst, src)
}
func (m *PostResponse) XXX_Size() int {
	return xxx_messageInfo_PostResponse.Size(m)
}
func (m *PostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostResponse proto.InternalMessageInfo

func (m *PostResponse) GetCityName() string {
	if m != nil {
		return m.CityName
	}
	return ""
}

func init() {
	proto.RegisterType((*PostUser)(nil), "user.PostUser")
	proto.RegisterType((*GetUser)(nil), "user.GetUser")
	proto.RegisterType((*GetRequest)(nil), "user.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "user.GetResponse")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
	proto.RegisterType((*PostRequest)(nil), "user.PostRequest")
	proto.RegisterType((*PostResponse)(nil), "user.PostResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user.Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/user.Users/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.Users/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Post(context.Context, *PostRequest) (*PostResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Users_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _Users_Post_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Users_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_60d31a5782020fab) }

var fileDescriptor_user_60d31a5782020fab = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xdf, 0x4e, 0xf2, 0x30,
	0x14, 0xff, 0x06, 0x83, 0x0f, 0x0f, 0x60, 0xa4, 0x1a, 0x43, 0xb8, 0xd2, 0x5e, 0x19, 0x15, 0xa6,
	0xf8, 0x00, 0x46, 0x30, 0x59, 0x4c, 0x8c, 0xc1, 0x05, 0x6f, 0xbc, 0x2b, 0x58, 0x67, 0xa3, 0xb4,
	0xd8, 0x76, 0x31, 0xfa, 0x28, 0x3e, 0xad, 0x69, 0xd7, 0xce, 0xed, 0x46, 0x13, 0xaf, 0xd6, 0x73,
	0x7e, 0xfd, 0xfd, 0x39, 0xa7, 0x03, 0xc8, 0x14, 0x95, 0xa3, 0xb5, 0x14, 0x5a, 0xa0, 0xd0, 0x9c,
	0x31, 0x87, 0xd6, 0x4c, 0x28, 0x7d, 0xa7, 0xa8, 0x44, 0xbb, 0xd0, 0x34, 0xdf, 0xab, 0xcb, 0x7e,
	0xb0, 0x17, 0x1c, 0x6c, 0x24, 0xae, 0x42, 0x03, 0x68, 0x4d, 0x99, 0x7e, 0xbf, 0x21, 0x2b, 0xda,
	0xaf, 0x59, 0xa4, 0xa8, 0x0d, 0x36, 0x23, 0x4a, 0xbd, 0x09, 0xf9, 0xd0, 0xaf, 0xe7, 0x98, 0xaf,
	0x11, 0x82, 0xd0, 0x72, 0x42, 0xdb, 0xb7, 0x67, 0x7c, 0x0b, 0xff, 0x63, 0xfa, 0x77, 0x3b, 0x2f,
	0x59, 0x2f, 0x49, 0x62, 0x80, 0x98, 0xea, 0x84, 0xbe, 0x66, 0x54, 0x69, 0xb4, 0x03, 0x8d, 0xb9,
	0x78, 0xa6, 0xdc, 0x89, 0xe6, 0x05, 0x3e, 0x81, 0xb6, 0xbd, 0xa3, 0xd6, 0x82, 0x2b, 0x8a, 0xf6,
	0x21, 0x34, 0x66, 0xf6, 0x4e, 0x7b, 0xdc, 0x1d, 0xd9, 0xb5, 0xb8, 0x5c, 0x89, 0x85, 0xf0, 0x04,
	0x3a, 0xd7, 0x22, 0x65, 0xdc, 0xeb, 0xfe, 0x90, 0xb6, 0x58, 0x40, 0xad, 0xba, 0x00, 0x7c, 0x04,
	0x5d, 0xa7, 0xe1, 0x7c, 0xcb, 0xa3, 0x05, 0xd5, 0xd1, 0xf0, 0x29, 0xb4, 0xcd, 0x4b, 0x78, 0x3f,
	0x5c, 0x89, 0xb8, 0x99, 0x47, 0xf4, 0x4f, 0xe5, 0x32, 0x1e, 0x42, 0x27, 0xa7, 0xfc, 0x2e, 0x3f,
	0xfe, 0x0c, 0xa0, 0x61, 0x48, 0x0a, 0x1d, 0x43, 0x3d, 0xa6, 0x1a, 0x6d, 0x15, 0x53, 0x3b, 0xcb,
	0x41, 0xaf, 0xd4, 0xc9, 0x15, 0xf1, 0x3f, 0x14, 0x41, 0x68, 0x3c, 0x50, 0xef, 0x3b, 0x81, 0xbf,
	0x8f, 0xca, 0xad, 0x82, 0x30, 0x86, 0x86, 0x1d, 0x1a, 0x39, 0xb8, 0xbc, 0xc5, 0xc1, 0x76, 0xa5,
	0xe7, 0x39, 0x93, 0x8b, 0xfb, 0xf3, 0x94, 0xe9, 0xa7, 0x6c, 0x31, 0x5a, 0x8a, 0x55, 0xa4, 0xc5,
	0x07, 0x51, 0x9a, 0x68, 0x26, 0x78, 0x94, 0x26, 0xb3, 0xe9, 0x70, 0x2e, 0x09, 0xe3, 0x8c, 0xa7,
	0xc3, 0x58, 0xbc, 0x10, 0x9e, 0x46, 0x8c, 0x6b, 0x2a, 0x1f, 0xc9, 0x92, 0xaa, 0x48, 0xae, 0x97,
	0x91, 0x91, 0x5c, 0x34, 0xed, 0x5f, 0x7d, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x0a, 0x13,
	0xc9, 0xe3, 0x02, 0x00, 0x00,
}