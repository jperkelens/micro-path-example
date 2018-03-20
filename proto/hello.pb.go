// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	proto/hello.proto

It has these top-level messages:
	GreetRequest
	GreetResponse
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type GreetRequest struct {
}

func (m *GreetRequest) Reset()                    { *m = GreetRequest{} }
func (m *GreetRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()               {}
func (*GreetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GreetResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *GreetResponse) Reset()                    { *m = GreetResponse{} }
func (m *GreetResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()               {}
func (*GreetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "GreetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Hello service

type HelloClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...client.CallOption) (*GreetResponse, error)
}

type helloClient struct {
	c           client.Client
	serviceName string
}

func NewHelloClient(serviceName string, c client.Client) HelloClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "hello"
	}
	return &helloClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *helloClient) Greet(ctx context.Context, in *GreetRequest, opts ...client.CallOption) (*GreetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Hello.Greet", in)
	out := new(GreetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloHandler interface {
	Greet(context.Context, *GreetRequest, *GreetResponse) error
}

func RegisterHelloHandler(s server.Server, hdlr HelloHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Hello{hdlr}, opts...))
}

type Hello struct {
	HelloHandler
}

func (h *Hello) Greet(ctx context.Context, in *GreetRequest, out *GreetResponse) error {
	return h.HelloHandler.Greet(ctx, in, out)
}

func init() { proto.RegisterFile("proto/hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0xb3, 0x95, 0xf8, 0xb8, 0x78, 0xdc, 0x8b,
	0x52, 0x53, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x34, 0xb9, 0x78, 0xa1, 0xfc,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0xd7, 0x48, 0x9f, 0x8b, 0xd5, 0x03, 0x64,
	0x92, 0x90, 0x1a, 0x17, 0x2b, 0x58, 0x8f, 0x10, 0xaf, 0x1e, 0xb2, 0x59, 0x52, 0x7c, 0x7a, 0x28,
	0x46, 0x25, 0xb1, 0x81, 0xad, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x45, 0x21, 0x8e,
	0x87, 0x00, 0x00, 0x00,
}
