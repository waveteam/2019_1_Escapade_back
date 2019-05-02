// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

package session

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

type SessionID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionID) Reset()         { *m = SessionID{} }
func (m *SessionID) String() string { return proto.CompactTextString(m) }
func (*SessionID) ProtoMessage()    {}
func (*SessionID) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_f0d54a363c561023, []int{0}
}
func (m *SessionID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionID.Unmarshal(m, b)
}
func (m *SessionID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionID.Marshal(b, m, deterministic)
}
func (dst *SessionID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionID.Merge(dst, src)
}
func (m *SessionID) XXX_Size() int {
	return xxx_messageInfo_SessionID.Size(m)
}
func (m *SessionID) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionID.DiscardUnknown(m)
}

var xxx_messageInfo_SessionID proto.InternalMessageInfo

func (m *SessionID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Session struct {
	UserID               int32    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Login                string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_f0d54a363c561023, []int{1}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *Session) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

type Nothing struct {
	Dummy                bool     `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_f0d54a363c561023, []int{2}
}
func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (dst *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(dst, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func (m *Nothing) GetDummy() bool {
	if m != nil {
		return m.Dummy
	}
	return false
}

func init() {
	proto.RegisterType((*SessionID)(nil), "session.SessionID")
	proto.RegisterType((*Session)(nil), "session.Session")
	proto.RegisterType((*Nothing)(nil), "session.Nothing")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthCheckerClient is the client API for AuthChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthCheckerClient interface {
	Create(ctx context.Context, in *Session, opts ...grpc.CallOption) (*SessionID, error)
	Check(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Session, error)
	Delete(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Nothing, error)
}

type authCheckerClient struct {
	cc *grpc.ClientConn
}

func NewAuthCheckerClient(cc *grpc.ClientConn) AuthCheckerClient {
	return &authCheckerClient{cc}
}

func (c *authCheckerClient) Create(ctx context.Context, in *Session, opts ...grpc.CallOption) (*SessionID, error) {
	out := new(SessionID)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) Check(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) Delete(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthCheckerServer is the server API for AuthChecker service.
type AuthCheckerServer interface {
	Create(context.Context, *Session) (*SessionID, error)
	Check(context.Context, *SessionID) (*Session, error)
	Delete(context.Context, *SessionID) (*Nothing, error)
}

func RegisterAuthCheckerServer(s *grpc.Server, srv AuthCheckerServer) {
	s.RegisterService(&_AuthChecker_serviceDesc, srv)
}

func _AuthChecker_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Create(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Check(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Delete(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session.AuthChecker",
	HandlerType: (*AuthCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AuthChecker_Create_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _AuthChecker_Check_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AuthChecker_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}

func init() { proto.RegisterFile("session.proto", fileDescriptor_session_f0d54a363c561023) }

var fileDescriptor_session_f0d54a363c561023 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xa4, 0xb9,
	0x38, 0x83, 0x21, 0x4c, 0x4f, 0x17, 0x21, 0x3e, 0x2e, 0x26, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0xa6, 0x4c, 0x17, 0x25, 0x73, 0x2e, 0x76, 0xa8, 0xa4, 0x90, 0x18, 0x17, 0x5b,
	0x69, 0x71, 0x6a, 0x11, 0x54, 0x9a, 0x35, 0x08, 0xca, 0x13, 0x12, 0xe1, 0x62, 0xcd, 0xc9, 0x4f,
	0xcf, 0xcc, 0x93, 0x60, 0x02, 0xeb, 0x82, 0x70, 0x94, 0xe4, 0xb9, 0xd8, 0xfd, 0xf2, 0x4b, 0x32,
	0x32, 0xf3, 0xd2, 0x41, 0x0a, 0x52, 0x4a, 0x73, 0x73, 0x2b, 0xc1, 0xfa, 0x38, 0x82, 0x20, 0x1c,
	0xa3, 0x45, 0x8c, 0x5c, 0xdc, 0x8e, 0xa5, 0x25, 0x19, 0xce, 0x19, 0xa9, 0xc9, 0xd9, 0xa9, 0x45,
	0x42, 0x06, 0x5c, 0x6c, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x42, 0x02, 0x7a, 0x30, 0x97, 0x42,
	0xad, 0x96, 0x12, 0x42, 0x17, 0xf1, 0x74, 0x51, 0x62, 0x10, 0xd2, 0xe7, 0x62, 0x05, 0x6b, 0x16,
	0xc2, 0x22, 0x2d, 0x85, 0x61, 0x08, 0x50, 0x03, 0xd0, 0x0a, 0x97, 0xd4, 0x9c, 0x54, 0xa0, 0x15,
	0xf8, 0x75, 0x40, 0x1d, 0xae, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x2b, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x06, 0x25, 0x70, 0x62, 0x3c, 0x01, 0x00, 0x00,
}
