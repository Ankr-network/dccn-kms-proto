// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token/token.proto

package kms_token

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GenReq struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenReq) Reset()         { *m = GenReq{} }
func (m *GenReq) String() string { return proto.CompactTextString(m) }
func (*GenReq) ProtoMessage()    {}
func (*GenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e2ef433bb3fdc80, []int{0}
}

func (m *GenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenReq.Unmarshal(m, b)
}
func (m *GenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenReq.Marshal(b, m, deterministic)
}
func (m *GenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenReq.Merge(m, src)
}
func (m *GenReq) XXX_Size() int {
	return xxx_messageInfo_GenReq.Size(m)
}
func (m *GenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GenReq.DiscardUnknown(m)
}

var xxx_messageInfo_GenReq proto.InternalMessageInfo

func (m *GenReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type GenRsp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenRsp) Reset()         { *m = GenRsp{} }
func (m *GenRsp) String() string { return proto.CompactTextString(m) }
func (*GenRsp) ProtoMessage()    {}
func (*GenRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e2ef433bb3fdc80, []int{1}
}

func (m *GenRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenRsp.Unmarshal(m, b)
}
func (m *GenRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenRsp.Marshal(b, m, deterministic)
}
func (m *GenRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenRsp.Merge(m, src)
}
func (m *GenRsp) XXX_Size() int {
	return xxx_messageInfo_GenRsp.Size(m)
}
func (m *GenRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GenRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GenRsp proto.InternalMessageInfo

func (m *GenRsp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GenRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *GenRsp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e2ef433bb3fdc80, []int{2}
}

func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (m *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(m, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetRsp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRsp) Reset()         { *m = GetRsp{} }
func (m *GetRsp) String() string { return proto.CompactTextString(m) }
func (*GetRsp) ProtoMessage()    {}
func (*GetRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e2ef433bb3fdc80, []int{3}
}

func (m *GetRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRsp.Unmarshal(m, b)
}
func (m *GetRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRsp.Marshal(b, m, deterministic)
}
func (m *GetRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRsp.Merge(m, src)
}
func (m *GetRsp) XXX_Size() int {
	return xxx_messageInfo_GetRsp.Size(m)
}
func (m *GetRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetRsp proto.InternalMessageInfo

func (m *GetRsp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GetRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *GetRsp) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type CheckReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckReq) Reset()         { *m = CheckReq{} }
func (m *CheckReq) String() string { return proto.CompactTextString(m) }
func (*CheckReq) ProtoMessage()    {}
func (*CheckReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e2ef433bb3fdc80, []int{4}
}

func (m *CheckReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckReq.Unmarshal(m, b)
}
func (m *CheckReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckReq.Marshal(b, m, deterministic)
}
func (m *CheckReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckReq.Merge(m, src)
}
func (m *CheckReq) XXX_Size() int {
	return xxx_messageInfo_CheckReq.Size(m)
}
func (m *CheckReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckReq proto.InternalMessageInfo

func (m *CheckReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CheckRsp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRsp) Reset()         { *m = CheckRsp{} }
func (m *CheckRsp) String() string { return proto.CompactTextString(m) }
func (*CheckRsp) ProtoMessage()    {}
func (*CheckRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e2ef433bb3fdc80, []int{5}
}

func (m *CheckRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRsp.Unmarshal(m, b)
}
func (m *CheckRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRsp.Marshal(b, m, deterministic)
}
func (m *CheckRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRsp.Merge(m, src)
}
func (m *CheckRsp) XXX_Size() int {
	return xxx_messageInfo_CheckRsp.Size(m)
}
func (m *CheckRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRsp proto.InternalMessageInfo

func (m *CheckRsp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CheckRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*GenReq)(nil), "kms.token.GenReq")
	proto.RegisterType((*GenRsp)(nil), "kms.token.GenRsp")
	proto.RegisterType((*GetReq)(nil), "kms.token.GetReq")
	proto.RegisterType((*GetRsp)(nil), "kms.token.GetRsp")
	proto.RegisterType((*CheckReq)(nil), "kms.token.CheckReq")
	proto.RegisterType((*CheckRsp)(nil), "kms.token.CheckRsp")
}

func init() { proto.RegisterFile("token/token.proto", fileDescriptor_6e2ef433bb3fdc80) }

var fileDescriptor_6e2ef433bb3fdc80 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xc9, 0xcf, 0x4e,
	0xcd, 0xd3, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0xd9, 0xb9, 0xc5, 0x7a,
	0x60, 0x01, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4,
	0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x42, 0x25, 0x19, 0x2e, 0x36,
	0xf7, 0xd4, 0xbc, 0xa0, 0xd4, 0x42, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x0d, 0x22, 0x5b, 0x5c, 0x00, 0x92, 0x4d, 0xce,
	0x4f, 0x49, 0x85, 0xc9, 0x82, 0xd8, 0x60, 0x1d, 0xa9, 0xc5, 0xc9, 0x12, 0x4c, 0x50, 0x1d, 0xa9,
	0xc5, 0xc9, 0x42, 0x22, 0x5c, 0xac, 0x60, 0x6b, 0x25, 0x98, 0xc1, 0x82, 0x10, 0x8e, 0x92, 0x1c,
	0xc8, 0x9c, 0x12, 0x90, 0x2d, 0x70, 0x79, 0x46, 0x64, 0x79, 0x17, 0x88, 0x3c, 0x09, 0xf6, 0xc0,
	0x5c, 0xcb, 0x8c, 0xe4, 0x5a, 0x05, 0x2e, 0x0e, 0xe7, 0x8c, 0xd4, 0xe4, 0x6c, 0xdc, 0xf6, 0x18,
	0xc1, 0x54, 0x10, 0x6f, 0x93, 0xd1, 0x4b, 0x46, 0x2e, 0x8e, 0x10, 0x90, 0xee, 0xe0, 0xa2, 0x32,
	0x21, 0x6f, 0x2e, 0x4e, 0xf7, 0xd4, 0xbc, 0xd4, 0xa2, 0xc4, 0x92, 0xfc, 0x22, 0x21, 0x41, 0x3d,
	0x78, 0x28, 0xeb, 0x41, 0x02, 0x51, 0x0a, 0x5d, 0xa8, 0xb8, 0x40, 0x49, 0xa2, 0xe9, 0xf2, 0x93,
	0xc9, 0x4c, 0x42, 0x4a, 0xbc, 0xfa, 0xd9, 0xb9, 0xc5, 0xfa, 0x65, 0x86, 0x90, 0x78, 0xb2, 0x62,
	0xd4, 0x12, 0x72, 0xe1, 0x62, 0x76, 0x4f, 0x2d, 0x41, 0x33, 0xa6, 0x04, 0xd3, 0x18, 0x50, 0xc0,
	0x28, 0x89, 0x83, 0x8d, 0x11, 0x14, 0xe2, 0x87, 0x19, 0x53, 0x0d, 0x96, 0xae, 0x15, 0xf2, 0xe2,
	0x62, 0x05, 0xfb, 0x49, 0x48, 0x18, 0x49, 0x13, 0x2c, 0x1c, 0xa4, 0x30, 0x05, 0xf1, 0x98, 0x95,
	0xc4, 0x06, 0x4e, 0x14, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x71, 0xb4, 0x95, 0x52,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokenSrvClient is the client API for TokenSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenSrvClient interface {
	Generator(ctx context.Context, in *GenReq, opts ...grpc.CallOption) (*GenRsp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error)
	Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckRsp, error)
}

type tokenSrvClient struct {
	cc *grpc.ClientConn
}

func NewTokenSrvClient(cc *grpc.ClientConn) TokenSrvClient {
	return &tokenSrvClient{cc}
}

func (c *tokenSrvClient) Generator(ctx context.Context, in *GenReq, opts ...grpc.CallOption) (*GenRsp, error) {
	out := new(GenRsp)
	err := c.cc.Invoke(ctx, "/kms.token.TokenSrv/Generator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenSrvClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error) {
	out := new(GetRsp)
	err := c.cc.Invoke(ctx, "/kms.token.TokenSrv/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenSrvClient) Check(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckRsp, error) {
	out := new(CheckRsp)
	err := c.cc.Invoke(ctx, "/kms.token.TokenSrv/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenSrvServer is the server API for TokenSrv service.
type TokenSrvServer interface {
	Generator(context.Context, *GenReq) (*GenRsp, error)
	Get(context.Context, *GetReq) (*GetRsp, error)
	Check(context.Context, *CheckReq) (*CheckRsp, error)
}

// UnimplementedTokenSrvServer can be embedded to have forward compatible implementations.
type UnimplementedTokenSrvServer struct {
}

func (*UnimplementedTokenSrvServer) Generator(ctx context.Context, req *GenReq) (*GenRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generator not implemented")
}
func (*UnimplementedTokenSrvServer) Get(ctx context.Context, req *GetReq) (*GetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTokenSrvServer) Check(ctx context.Context, req *CheckReq) (*CheckRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterTokenSrvServer(s *grpc.Server, srv TokenSrvServer) {
	s.RegisterService(&_TokenSrv_serviceDesc, srv)
}

func _TokenSrv_Generator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenSrvServer).Generator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kms.token.TokenSrv/Generator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenSrvServer).Generator(ctx, req.(*GenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenSrv_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenSrvServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kms.token.TokenSrv/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenSrvServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenSrv_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenSrvServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kms.token.TokenSrv/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenSrvServer).Check(ctx, req.(*CheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kms.token.TokenSrv",
	HandlerType: (*TokenSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generator",
			Handler:    _TokenSrv_Generator_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TokenSrv_Get_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _TokenSrv_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token/token.proto",
}