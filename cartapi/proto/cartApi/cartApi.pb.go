// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/cartApi/cartApi.proto

/*
Package go_micro_api_cartApi is a generated protocol buffer package.

It is generated from these files:
	proto/cartApi/cartApi.proto

It has these top-level messages:
	Pair
	Request
	Response
*/
package api_cartApi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Pair struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Request struct {
	Method string           `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   string           `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
	Url    string           `protobuf:"bytes,7,opt,name=url" json:"url,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Response struct {
	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       string           `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.micro.api.cartApi.Pair")
	proto.RegisterType((*Request)(nil), "go.micro.api.cartApi.Request")
	proto.RegisterType((*Response)(nil), "go.micro.api.cartApi.Response")
}

func init() { proto.RegisterFile("proto/cartApi/cartApi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xc7, 0xed, 0x07, 0x2d, 0x0c, 0x17, 0xb3, 0x31, 0x66, 0x53, 0x23, 0x21, 0x3d, 0x28, 0x7a,
	0xa8, 0x04, 0x2f, 0x44, 0x4f, 0x48, 0xfc, 0x88, 0x27, 0xb2, 0x86, 0x07, 0x28, 0x74, 0x03, 0x8d,
	0x85, 0xad, 0xbb, 0x5b, 0x13, 0x1e, 0xc4, 0x87, 0xf3, 0x6d, 0xcc, 0x6e, 0x17, 0xe4, 0x50, 0x7b,
	0xc2, 0x53, 0x77, 0x67, 0xe7, 0x37, 0x99, 0xf9, 0xff, 0xa7, 0x70, 0x96, 0x73, 0x26, 0xd9, 0xcd,
	0x3c, 0xe6, 0x72, 0x94, 0xa7, 0xdb, 0x6f, 0xa4, 0xa3, 0xe8, 0x64, 0xc1, 0xa2, 0x55, 0x3a, 0xe7,
	0x2c, 0x8a, 0xf3, 0x34, 0x32, 0x6f, 0x61, 0x1f, 0xdc, 0x49, 0x9c, 0x72, 0x74, 0x0c, 0xce, 0x3b,
	0xdd, 0x60, 0xab, 0x6b, 0xf5, 0x5a, 0x44, 0x1d, 0xd1, 0x29, 0x78, 0x9f, 0x71, 0x56, 0x50, 0x81,
	0xed, 0xae, 0xd3, 0x6b, 0x11, 0x73, 0x0b, 0xbf, 0x5c, 0xf0, 0x09, 0xfd, 0x28, 0xa8, 0x90, 0x2a,
	0x67, 0x45, 0xe5, 0x92, 0x25, 0x06, 0x34, 0x37, 0x84, 0xc0, 0xcd, 0x63, 0xb9, 0xc4, 0xb6, 0x8e,
	0xea, 0x33, 0x1a, 0x81, 0xb7, 0xa4, 0x71, 0x42, 0x39, 0x76, 0xba, 0x4e, 0xaf, 0x3d, 0xb8, 0x8a,
	0xaa, 0x1a, 0x8a, 0x4c, 0xe9, 0xe8, 0x45, 0xe7, 0x3e, 0xae, 0x25, 0xdf, 0x10, 0x03, 0xa2, 0x21,
	0x38, 0x0b, 0x2a, 0xb1, 0xab, 0xf9, 0x8b, 0x7a, 0xfe, 0x99, 0xca, 0x12, 0x56, 0x08, 0xba, 0x07,
	0x37, 0x67, 0x42, 0xe2, 0x86, 0x46, 0x2f, 0xeb, 0xd1, 0x09, 0x13, 0x86, 0xd5, 0x90, 0x9a, 0x66,
	0xc6, 0x92, 0x0d, 0xf6, 0xca, 0x69, 0xd4, 0x59, 0xe9, 0x55, 0xf0, 0x0c, 0xfb, 0xa5, 0x5e, 0x05,
	0xcf, 0x82, 0x29, 0xb4, 0xf7, 0x7a, 0xae, 0x10, 0xb4, 0x0f, 0x0d, 0x2d, 0xa1, 0x56, 0xa5, 0x3d,
	0x08, 0xaa, 0x9b, 0x50, 0x6e, 0x90, 0x32, 0xf1, 0xce, 0x1e, 0x5a, 0x01, 0x81, 0xe6, 0x76, 0x94,
	0x83, 0xd5, 0x7c, 0x83, 0xd6, 0x6e, 0xc6, 0x43, 0x15, 0x0d, 0xbf, 0x2d, 0x68, 0x12, 0x2a, 0x72,
	0xb6, 0x16, 0x14, 0x75, 0x00, 0x84, 0x8c, 0x65, 0x21, 0xc6, 0x2c, 0xa1, 0xba, 0x76, 0x83, 0xec,
	0x45, 0xd0, 0xc3, 0x6e, 0x19, 0x6c, 0xed, 0xc8, 0xf5, 0x5f, 0x8e, 0x94, 0xf5, 0x2a, 0xb7, 0x61,
	0x6b, 0x8b, 0xf3, 0x6b, 0xcb, 0x3f, 0x99, 0x30, 0x98, 0x82, 0x3f, 0x2e, 0x9f, 0xd0, 0x2b, 0xf8,
	0x4f, 0xe9, 0x3a, 0x19, 0x65, 0x19, 0x3a, 0xaf, 0x5d, 0xa3, 0xa0, 0x53, 0x3f, 0x53, 0x78, 0x34,
	0xf3, 0xf4, 0x9f, 0x79, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x33, 0xf6, 0x02, 0xb8, 0x03,
	0x00, 0x00,
}
