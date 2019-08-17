// Code generated by protoc-gen-go. DO NOT EDIT.
// source: identity/identity.proto

package kms_identity

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

type DataNode struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val                  string   `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataNode) Reset()         { *m = DataNode{} }
func (m *DataNode) String() string { return proto.CompactTextString(m) }
func (*DataNode) ProtoMessage()    {}
func (*DataNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_2433c1f46177a3e0, []int{0}
}

func (m *DataNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNode.Unmarshal(m, b)
}
func (m *DataNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNode.Marshal(b, m, deterministic)
}
func (m *DataNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNode.Merge(m, src)
}
func (m *DataNode) XXX_Size() int {
	return xxx_messageInfo_DataNode.Size(m)
}
func (m *DataNode) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNode.DiscardUnknown(m)
}

var xxx_messageInfo_DataNode proto.InternalMessageInfo

func (m *DataNode) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataNode) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type PutReq struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Meta                 []*DataNode `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty"`
	Path                 []string    `protobuf:"bytes,3,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PutReq) Reset()         { *m = PutReq{} }
func (m *PutReq) String() string { return proto.CompactTextString(m) }
func (*PutReq) ProtoMessage()    {}
func (*PutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2433c1f46177a3e0, []int{1}
}

func (m *PutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutReq.Unmarshal(m, b)
}
func (m *PutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutReq.Marshal(b, m, deterministic)
}
func (m *PutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutReq.Merge(m, src)
}
func (m *PutReq) XXX_Size() int {
	return xxx_messageInfo_PutReq.Size(m)
}
func (m *PutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PutReq.DiscardUnknown(m)
}

var xxx_messageInfo_PutReq proto.InternalMessageInfo

func (m *PutReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PutReq) GetMeta() []*DataNode {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PutReq) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

type PutRsp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRsp) Reset()         { *m = PutRsp{} }
func (m *PutRsp) String() string { return proto.CompactTextString(m) }
func (*PutRsp) ProtoMessage()    {}
func (*PutRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2433c1f46177a3e0, []int{2}
}

func (m *PutRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRsp.Unmarshal(m, b)
}
func (m *PutRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRsp.Marshal(b, m, deterministic)
}
func (m *PutRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRsp.Merge(m, src)
}
func (m *PutRsp) XXX_Size() int {
	return xxx_messageInfo_PutRsp.Size(m)
}
func (m *PutRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PutRsp proto.InternalMessageInfo

func (m *PutRsp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PutRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type GetReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2433c1f46177a3e0, []int{3}
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

func (m *GetReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRsp struct {
	Code                 string      `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Desc                 string      `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Id                   string      `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Meta                 []*DataNode `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty"`
	Path                 []string    `protobuf:"bytes,5,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetRsp) Reset()         { *m = GetRsp{} }
func (m *GetRsp) String() string { return proto.CompactTextString(m) }
func (*GetRsp) ProtoMessage()    {}
func (*GetRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2433c1f46177a3e0, []int{4}
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

func (m *GetRsp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetRsp) GetMeta() []*DataNode {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetRsp) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

type PatchReq struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Dns                  []*DataNode `protobuf:"bytes,3,rep,name=dns,proto3" json:"dns,omitempty"`
	Meta                 []*DataNode `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty"`
	Path                 []string    `protobuf:"bytes,5,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PatchReq) Reset()         { *m = PatchReq{} }
func (m *PatchReq) String() string { return proto.CompactTextString(m) }
func (*PatchReq) ProtoMessage()    {}
func (*PatchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2433c1f46177a3e0, []int{5}
}

func (m *PatchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchReq.Unmarshal(m, b)
}
func (m *PatchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchReq.Marshal(b, m, deterministic)
}
func (m *PatchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchReq.Merge(m, src)
}
func (m *PatchReq) XXX_Size() int {
	return xxx_messageInfo_PatchReq.Size(m)
}
func (m *PatchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchReq.DiscardUnknown(m)
}

var xxx_messageInfo_PatchReq proto.InternalMessageInfo

func (m *PatchReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PatchReq) GetDns() []*DataNode {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *PatchReq) GetMeta() []*DataNode {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PatchReq) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

type PatchRsp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchRsp) Reset()         { *m = PatchRsp{} }
func (m *PatchRsp) String() string { return proto.CompactTextString(m) }
func (*PatchRsp) ProtoMessage()    {}
func (*PatchRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2433c1f46177a3e0, []int{6}
}

func (m *PatchRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchRsp.Unmarshal(m, b)
}
func (m *PatchRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchRsp.Marshal(b, m, deterministic)
}
func (m *PatchRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchRsp.Merge(m, src)
}
func (m *PatchRsp) XXX_Size() int {
	return xxx_messageInfo_PatchRsp.Size(m)
}
func (m *PatchRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PatchRsp proto.InternalMessageInfo

func (m *PatchRsp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PatchRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type DeleteReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReq) Reset()         { *m = DeleteReq{} }
func (m *DeleteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()    {}
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2433c1f46177a3e0, []int{7}
}

func (m *DeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReq.Unmarshal(m, b)
}
func (m *DeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReq.Marshal(b, m, deterministic)
}
func (m *DeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReq.Merge(m, src)
}
func (m *DeleteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteReq.Size(m)
}
func (m *DeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReq proto.InternalMessageInfo

func (m *DeleteReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteRsp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRsp) Reset()         { *m = DeleteRsp{} }
func (m *DeleteRsp) String() string { return proto.CompactTextString(m) }
func (*DeleteRsp) ProtoMessage()    {}
func (*DeleteRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2433c1f46177a3e0, []int{8}
}

func (m *DeleteRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRsp.Unmarshal(m, b)
}
func (m *DeleteRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRsp.Marshal(b, m, deterministic)
}
func (m *DeleteRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRsp.Merge(m, src)
}
func (m *DeleteRsp) XXX_Size() int {
	return xxx_messageInfo_DeleteRsp.Size(m)
}
func (m *DeleteRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRsp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRsp proto.InternalMessageInfo

func (m *DeleteRsp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *DeleteRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*DataNode)(nil), "kms.identity.DataNode")
	proto.RegisterType((*PutReq)(nil), "kms.identity.PutReq")
	proto.RegisterType((*PutRsp)(nil), "kms.identity.PutRsp")
	proto.RegisterType((*GetReq)(nil), "kms.identity.GetReq")
	proto.RegisterType((*GetRsp)(nil), "kms.identity.GetRsp")
	proto.RegisterType((*PatchReq)(nil), "kms.identity.PatchReq")
	proto.RegisterType((*PatchRsp)(nil), "kms.identity.PatchRsp")
	proto.RegisterType((*DeleteReq)(nil), "kms.identity.DeleteReq")
	proto.RegisterType((*DeleteRsp)(nil), "kms.identity.DeleteRsp")
}

func init() { proto.RegisterFile("identity/identity.proto", fileDescriptor_2433c1f46177a3e0) }

var fileDescriptor_2433c1f46177a3e0 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0xaa, 0xda, 0x40,
	0x14, 0x25, 0x19, 0x0d, 0x7a, 0x2d, 0xa5, 0x0c, 0xa5, 0xa6, 0x41, 0x50, 0x66, 0x25, 0x2e, 0x92,
	0x36, 0xee, 0xba, 0x16, 0xa4, 0x9b, 0x22, 0x29, 0x5d, 0xd9, 0xcd, 0xd4, 0x0c, 0x1a, 0x4c, 0x32,
	0xa9, 0x33, 0x06, 0xa4, 0x14, 0x4a, 0x57, 0xdd, 0xf7, 0xd3, 0xfa, 0x0b, 0xfd, 0x82, 0xf7, 0x05,
	0x8f, 0x99, 0x98, 0xe8, 0xf3, 0x25, 0x0f, 0xe5, 0xed, 0x0e, 0xe7, 0xde, 0x39, 0x73, 0xee, 0xb9,
	0x33, 0xd0, 0x8f, 0x42, 0x96, 0xca, 0x48, 0x1e, 0xbc, 0x12, 0xb8, 0xd9, 0x8e, 0x4b, 0x8e, 0x5f,
	0x6c, 0x13, 0xe1, 0x96, 0x9c, 0x33, 0x58, 0x73, 0xbe, 0x8e, 0x99, 0x47, 0xb3, 0xc8, 0xa3, 0x69,
	0xca, 0x25, 0x95, 0x11, 0x4f, 0x45, 0xd1, 0x4b, 0x5c, 0xe8, 0xcc, 0xa8, 0xa4, 0x9f, 0x78, 0xc8,
	0xf0, 0x2b, 0x40, 0x5b, 0x76, 0xb0, 0x8d, 0x91, 0x31, 0xee, 0x06, 0x0a, 0x2a, 0x26, 0xa7, 0xb1,
	0x6d, 0x16, 0x4c, 0x4e, 0x63, 0xf2, 0x15, 0xac, 0xc5, 0x5e, 0x06, 0xec, 0x3b, 0xc6, 0xd0, 0x4a,
	0x69, 0xc2, 0x8e, 0xed, 0x1a, 0xe3, 0x09, 0xb4, 0x12, 0x26, 0xa9, 0x6d, 0x8e, 0xd0, 0xb8, 0xe7,
	0xbf, 0x71, 0xcf, 0x8d, 0xb8, 0xe5, 0x3d, 0x81, 0xee, 0x51, 0xe7, 0x33, 0x2a, 0x37, 0x36, 0x1a,
	0x21, 0x75, 0x5e, 0x61, 0xf2, 0xae, 0x50, 0x17, 0x99, 0xaa, 0xae, 0x78, 0x58, 0xa9, 0x2b, 0xac,
	0xb8, 0x90, 0x89, 0xd5, 0xd1, 0x8e, 0xc6, 0x64, 0x00, 0xd6, 0x9c, 0x35, 0xf9, 0x21, 0xbf, 0x8c,
	0xa2, 0x7c, 0xbd, 0x20, 0x7e, 0x09, 0x66, 0x14, 0xda, 0x48, 0x33, 0x66, 0x14, 0x56, 0x23, 0xb5,
	0x6e, 0x18, 0xa9, 0x7d, 0x36, 0xd2, 0x1f, 0x03, 0x3a, 0x0b, 0x2a, 0x57, 0x9b, 0xa6, 0xcc, 0xc6,
	0x80, 0xc2, 0x54, 0xe8, 0x18, 0x9a, 0xf5, 0x55, 0xcb, 0xb3, 0xad, 0xf8, 0xa5, 0x93, 0x1b, 0xf2,
	0x1d, 0x42, 0x77, 0xc6, 0x62, 0x26, 0x59, 0x53, 0xc4, 0xd3, 0xaa, 0xe1, 0x7a, 0x55, 0xff, 0xce,
	0x84, 0xde, 0xc7, 0xa3, 0xf3, 0xcf, 0xbb, 0x1c, 0x7f, 0x01, 0xb4, 0xd8, 0x4b, 0xfc, 0xfa, 0xe1,
	0x48, 0xc5, 0x43, 0x73, 0x6a, 0x58, 0x91, 0x11, 0xf2, 0xfb, 0xdf, 0xff, 0xbf, 0xe6, 0x80, 0xf4,
	0xbd, 0x6d, 0x22, 0xbc, 0xfc, 0x7d, 0xf5, 0x09, 0xbc, 0x1f, 0xca, 0xd7, 0xcf, 0x0f, 0xc6, 0x04,
	0x07, 0x80, 0xe6, 0xec, 0x91, 0x6c, 0xf1, 0x5e, 0x9c, 0x1a, 0x56, 0x64, 0x64, 0xa8, 0x65, 0xdf,
	0xe2, 0x26, 0x59, 0xbc, 0x84, 0xb6, 0x0e, 0x11, 0x5f, 0xe4, 0x5f, 0xee, 0xd8, 0xa9, 0xe5, 0x4f,
	0x86, 0xfd, 0xa7, 0x0c, 0x2f, 0xc1, 0x2a, 0xc2, 0xc4, 0xfd, 0x8b, 0xed, 0x96, 0x3b, 0x70, 0xea,
	0x0b, 0x27, 0xe7, 0x93, 0x26, 0xfd, 0x6f, 0x96, 0xfe, 0xf1, 0xd3, 0xfb, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8b, 0x1f, 0x3e, 0x86, 0x38, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentitySrvClient is the client API for IdentitySrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentitySrvClient interface {
	Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutRsp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error)
	Patch(ctx context.Context, in *PatchReq, opts ...grpc.CallOption) (*PatchRsp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRsp, error)
}

type identitySrvClient struct {
	cc *grpc.ClientConn
}

func NewIdentitySrvClient(cc *grpc.ClientConn) IdentitySrvClient {
	return &identitySrvClient{cc}
}

func (c *identitySrvClient) Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutRsp, error) {
	out := new(PutRsp)
	err := c.cc.Invoke(ctx, "/kms.identity.IdentitySrv/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identitySrvClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error) {
	out := new(GetRsp)
	err := c.cc.Invoke(ctx, "/kms.identity.IdentitySrv/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identitySrvClient) Patch(ctx context.Context, in *PatchReq, opts ...grpc.CallOption) (*PatchRsp, error) {
	out := new(PatchRsp)
	err := c.cc.Invoke(ctx, "/kms.identity.IdentitySrv/Patch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identitySrvClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRsp, error) {
	out := new(DeleteRsp)
	err := c.cc.Invoke(ctx, "/kms.identity.IdentitySrv/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentitySrvServer is the server API for IdentitySrv service.
type IdentitySrvServer interface {
	Put(context.Context, *PutReq) (*PutRsp, error)
	Get(context.Context, *GetReq) (*GetRsp, error)
	Patch(context.Context, *PatchReq) (*PatchRsp, error)
	Delete(context.Context, *DeleteReq) (*DeleteRsp, error)
}

// UnimplementedIdentitySrvServer can be embedded to have forward compatible implementations.
type UnimplementedIdentitySrvServer struct {
}

func (*UnimplementedIdentitySrvServer) Put(ctx context.Context, req *PutReq) (*PutRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedIdentitySrvServer) Get(ctx context.Context, req *GetReq) (*GetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedIdentitySrvServer) Patch(ctx context.Context, req *PatchReq) (*PatchRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (*UnimplementedIdentitySrvServer) Delete(ctx context.Context, req *DeleteReq) (*DeleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterIdentitySrvServer(s *grpc.Server, srv IdentitySrvServer) {
	s.RegisterService(&_IdentitySrv_serviceDesc, srv)
}

func _IdentitySrv_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentitySrvServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kms.identity.IdentitySrv/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentitySrvServer).Put(ctx, req.(*PutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentitySrv_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentitySrvServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kms.identity.IdentitySrv/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentitySrvServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentitySrv_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentitySrvServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kms.identity.IdentitySrv/Patch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentitySrvServer).Patch(ctx, req.(*PatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentitySrv_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentitySrvServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kms.identity.IdentitySrv/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentitySrvServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentitySrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kms.identity.IdentitySrv",
	HandlerType: (*IdentitySrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _IdentitySrv_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _IdentitySrv_Get_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _IdentitySrv_Patch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _IdentitySrv_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/identity.proto",
}
