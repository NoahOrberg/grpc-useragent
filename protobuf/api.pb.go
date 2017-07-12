// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	UserAgentRes
	UserAgentReq
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type UserAgentRes struct {
	UserAgent string `protobuf:"bytes,1,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
}

func (m *UserAgentRes) Reset()                    { *m = UserAgentRes{} }
func (m *UserAgentRes) String() string            { return proto1.CompactTextString(m) }
func (*UserAgentRes) ProtoMessage()               {}
func (*UserAgentRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserAgentRes) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

type UserAgentReq struct {
}

func (m *UserAgentReq) Reset()                    { *m = UserAgentReq{} }
func (m *UserAgentReq) String() string            { return proto1.CompactTextString(m) }
func (*UserAgentReq) ProtoMessage()               {}
func (*UserAgentReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto1.RegisterType((*UserAgentRes)(nil), "proto.UserAgentRes")
	proto1.RegisterType((*UserAgentReq)(nil), "proto.UserAgentReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserAgent service

type UserAgentClient interface {
	GetUA(ctx context.Context, in *UserAgentReq, opts ...grpc.CallOption) (*UserAgentRes, error)
}

type userAgentClient struct {
	cc *grpc.ClientConn
}

func NewUserAgentClient(cc *grpc.ClientConn) UserAgentClient {
	return &userAgentClient{cc}
}

func (c *userAgentClient) GetUA(ctx context.Context, in *UserAgentReq, opts ...grpc.CallOption) (*UserAgentRes, error) {
	out := new(UserAgentRes)
	err := grpc.Invoke(ctx, "/proto.UserAgent/GetUA", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserAgent service

type UserAgentServer interface {
	GetUA(context.Context, *UserAgentReq) (*UserAgentRes, error)
}

func RegisterUserAgentServer(s *grpc.Server, srv UserAgentServer) {
	s.RegisterService(&_UserAgent_serviceDesc, srv)
}

func _UserAgent_GetUA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAgentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAgentServer).GetUA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserAgent/GetUA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAgentServer).GetUA(ctx, req.(*UserAgentReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserAgent",
	HandlerType: (*UserAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUA",
			Handler:    _UserAgent_GetUA_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto1.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xba, 0x5c, 0x3c, 0xa1, 0xc5, 0xa9,
	0x45, 0x8e, 0xe9, 0xa9, 0x79, 0x25, 0x41, 0xa9, 0xc5, 0x42, 0xb2, 0x5c, 0x5c, 0xa5, 0xc5, 0xa9,
	0x45, 0xf1, 0x89, 0x20, 0x01, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xce, 0x52, 0x98, 0x0a,
	0x25, 0x3e, 0x14, 0xe5, 0x85, 0x46, 0x76, 0x5c, 0x9c, 0x70, 0xbe, 0x90, 0x21, 0x17, 0xab, 0x7b,
	0x6a, 0x49, 0xa8, 0xa3, 0x90, 0x30, 0xc4, 0x0e, 0x3d, 0x64, 0xa5, 0x52, 0x58, 0x04, 0x8b, 0x93,
	0xd8, 0xc0, 0x62, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x0b, 0x90, 0xdd, 0x99, 0x00,
	0x00, 0x00,
}
