// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/category/category.proto

/*
Package go_micro_service_category is a generated protocol buffer package.

It is generated from these files:
	proto/category/category.proto

It has these top-level messages:
	CategoryRequest
	CreateCategoryResponse
	UpdateCategoryResponse
	DeleteCategoryRequest
	DeleteCategoryResponse
	FindByNameRequest
	CategoryResponse
	FindByIdRequest
	FindByLevelRequest
	FindByParentRequest
	FindAllRequest
	FindAllResponse
*/
package service_category

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

type CategoryRequest struct {
	CategoryName        string `protobuf:"bytes,1,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,2,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,3,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
	CategoryImage       string `protobuf:"bytes,4,opt,name=category_image,json=categoryImage" json:"category_image,omitempty"`
	CategoryDescription string `protobuf:"bytes,5,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *CategoryRequest) Reset()                    { *m = CategoryRequest{} }
func (m *CategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*CategoryRequest) ProtoMessage()               {}
func (*CategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CategoryRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *CategoryRequest) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *CategoryRequest) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *CategoryRequest) GetCategoryImage() string {
	if m != nil {
		return m.CategoryImage
	}
	return ""
}

func (m *CategoryRequest) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

type CreateCategoryResponse struct {
	Message    string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	CategoryId int64  `protobuf:"varint,2,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
}

func (m *CreateCategoryResponse) Reset()                    { *m = CreateCategoryResponse{} }
func (m *CreateCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateCategoryResponse) ProtoMessage()               {}
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateCategoryResponse) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type UpdateCategoryResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UpdateCategoryResponse) Reset()                    { *m = UpdateCategoryResponse{} }
func (m *UpdateCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateCategoryResponse) ProtoMessage()               {}
func (*UpdateCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteCategoryRequest struct {
	CategoryId int64 `protobuf:"varint,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
}

func (m *DeleteCategoryRequest) Reset()                    { *m = DeleteCategoryRequest{} }
func (m *DeleteCategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoryRequest) ProtoMessage()               {}
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteCategoryRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type DeleteCategoryResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteCategoryResponse) Reset()                    { *m = DeleteCategoryResponse{} }
func (m *DeleteCategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteCategoryResponse) ProtoMessage()               {}
func (*DeleteCategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteCategoryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type FindByNameRequest struct {
	CategoryName string `protobuf:"bytes,1,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
}

func (m *FindByNameRequest) Reset()                    { *m = FindByNameRequest{} }
func (m *FindByNameRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByNameRequest) ProtoMessage()               {}
func (*FindByNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FindByNameRequest) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

type CategoryResponse struct {
	Id                  int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CategoryName        string `protobuf:"bytes,2,opt,name=category_name,json=categoryName" json:"category_name,omitempty"`
	CategoryLevel       uint32 `protobuf:"varint,3,opt,name=category_level,json=categoryLevel" json:"category_level,omitempty"`
	CategoryParent      int64  `protobuf:"varint,4,opt,name=category_parent,json=categoryParent" json:"category_parent,omitempty"`
	CategoryImages      string `protobuf:"bytes,5,opt,name=category_images,json=categoryImages" json:"category_images,omitempty"`
	CategoryDescription string `protobuf:"bytes,6,opt,name=category_description,json=categoryDescription" json:"category_description,omitempty"`
}

func (m *CategoryResponse) Reset()                    { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()               {}
func (*CategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CategoryResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategoryResponse) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *CategoryResponse) GetCategoryLevel() uint32 {
	if m != nil {
		return m.CategoryLevel
	}
	return 0
}

func (m *CategoryResponse) GetCategoryParent() int64 {
	if m != nil {
		return m.CategoryParent
	}
	return 0
}

func (m *CategoryResponse) GetCategoryImages() string {
	if m != nil {
		return m.CategoryImages
	}
	return ""
}

func (m *CategoryResponse) GetCategoryDescription() string {
	if m != nil {
		return m.CategoryDescription
	}
	return ""
}

type FindByIdRequest struct {
	CategoryId int64 `protobuf:"varint,1,opt,name=category_id,json=categoryId" json:"category_id,omitempty"`
}

func (m *FindByIdRequest) Reset()                    { *m = FindByIdRequest{} }
func (m *FindByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByIdRequest) ProtoMessage()               {}
func (*FindByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FindByIdRequest) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

type FindByLevelRequest struct {
	Level uint32 `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
}

func (m *FindByLevelRequest) Reset()                    { *m = FindByLevelRequest{} }
func (m *FindByLevelRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByLevelRequest) ProtoMessage()               {}
func (*FindByLevelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FindByLevelRequest) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type FindByParentRequest struct {
	ParentId int64 `protobuf:"varint,1,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
}

func (m *FindByParentRequest) Reset()                    { *m = FindByParentRequest{} }
func (m *FindByParentRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByParentRequest) ProtoMessage()               {}
func (*FindByParentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FindByParentRequest) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type FindAllRequest struct {
}

func (m *FindAllRequest) Reset()                    { *m = FindAllRequest{} }
func (m *FindAllRequest) String() string            { return proto.CompactTextString(m) }
func (*FindAllRequest) ProtoMessage()               {}
func (*FindAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type FindAllResponse struct {
	Category []*CategoryResponse `protobuf:"bytes,1,rep,name=category" json:"category,omitempty"`
}

func (m *FindAllResponse) Reset()                    { *m = FindAllResponse{} }
func (m *FindAllResponse) String() string            { return proto.CompactTextString(m) }
func (*FindAllResponse) ProtoMessage()               {}
func (*FindAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FindAllResponse) GetCategory() []*CategoryResponse {
	if m != nil {
		return m.Category
	}
	return nil
}

func init() {
	proto.RegisterType((*CategoryRequest)(nil), "go.micro.service.category.CategoryRequest")
	proto.RegisterType((*CreateCategoryResponse)(nil), "go.micro.service.category.CreateCategoryResponse")
	proto.RegisterType((*UpdateCategoryResponse)(nil), "go.micro.service.category.UpdateCategoryResponse")
	proto.RegisterType((*DeleteCategoryRequest)(nil), "go.micro.service.category.DeleteCategoryRequest")
	proto.RegisterType((*DeleteCategoryResponse)(nil), "go.micro.service.category.DeleteCategoryResponse")
	proto.RegisterType((*FindByNameRequest)(nil), "go.micro.service.category.FindByNameRequest")
	proto.RegisterType((*CategoryResponse)(nil), "go.micro.service.category.CategoryResponse")
	proto.RegisterType((*FindByIdRequest)(nil), "go.micro.service.category.FindByIdRequest")
	proto.RegisterType((*FindByLevelRequest)(nil), "go.micro.service.category.FindByLevelRequest")
	proto.RegisterType((*FindByParentRequest)(nil), "go.micro.service.category.FindByParentRequest")
	proto.RegisterType((*FindAllRequest)(nil), "go.micro.service.category.FindAllRequest")
	proto.RegisterType((*FindAllResponse)(nil), "go.micro.service.category.FindAllResponse")
}

func init() { proto.RegisterFile("proto/category/category.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0x26, 0xbd, 0xa4, 0x53, 0xea, 0x84, 0x6d, 0xa8, 0x4c, 0x10, 0x22, 0x5a, 0x84, 0x08,
	0x05, 0x0c, 0x35, 0x2f, 0xbc, 0x42, 0x23, 0x50, 0x24, 0x84, 0x90, 0x11, 0x2f, 0xbc, 0x20, 0x13,
	0x8f, 0x22, 0xa3, 0xf8, 0x52, 0xaf, 0x29, 0xe2, 0x73, 0xf8, 0x33, 0xfe, 0x81, 0x1f, 0x40, 0xde,
	0xcd, 0xae, 0xe3, 0xc4, 0x71, 0x6c, 0xc4, 0x5b, 0x76, 0x76, 0x66, 0xce, 0x9c, 0x9d, 0x73, 0x62,
	0xb8, 0x1b, 0x27, 0x51, 0x1a, 0x3d, 0x9b, 0xb9, 0x29, 0xce, 0xa3, 0xe4, 0xa7, 0xfe, 0x61, 0x89,
	0x38, 0xbd, 0x3d, 0x8f, 0xac, 0xc0, 0x9f, 0x25, 0x91, 0xc5, 0x31, 0xb9, 0xf6, 0x67, 0x68, 0xa9,
	0x04, 0xf6, 0x9b, 0x40, 0xef, 0x72, 0x79, 0x70, 0xf0, 0xea, 0x3b, 0xf2, 0x94, 0xde, 0x87, 0x13,
	0x75, 0xff, 0x25, 0x74, 0x03, 0x34, 0xc9, 0x88, 0x8c, 0x8f, 0x9c, 0x1b, 0x2a, 0xf8, 0xde, 0x0d,
	0x90, 0x3e, 0x00, 0x43, 0x27, 0x2d, 0xf0, 0x1a, 0x17, 0x66, 0x7b, 0x44, 0xc6, 0x27, 0x8e, 0x2e,
	0x7d, 0x97, 0x05, 0xe9, 0x43, 0xe8, 0xe9, 0xb4, 0xd8, 0x4d, 0x30, 0x4c, 0xcd, 0xce, 0x88, 0x8c,
	0x3b, 0x8e, 0xae, 0xfe, 0x20, 0xa2, 0x85, 0x7e, 0x7e, 0xe0, 0xce, 0xd1, 0xdc, 0x13, 0xa8, 0xba,
	0xdf, 0x34, 0x0b, 0xd2, 0x0b, 0x18, 0xe8, 0x34, 0x0f, 0xf9, 0x2c, 0xf1, 0xe3, 0xd4, 0x8f, 0x42,
	0x73, 0x5f, 0x24, 0x9f, 0xaa, 0xbb, 0x49, 0x7e, 0xc5, 0x3e, 0xc2, 0xd9, 0x65, 0x82, 0x6e, 0x8a,
	0x39, 0x4f, 0x1e, 0x47, 0x21, 0x47, 0x6a, 0xc2, 0x61, 0x80, 0x9c, 0x67, 0x60, 0x92, 0xa2, 0x3a,
	0xd2, 0x7b, 0x70, 0x9c, 0x4f, 0xe3, 0x09, 0x6a, 0x1d, 0x07, 0xf4, 0x28, 0x1e, 0xb3, 0xe1, 0xec,
	0x53, 0xec, 0x35, 0x6a, 0xca, 0x5e, 0xc2, 0xad, 0x09, 0x2e, 0x70, 0xb5, 0x46, 0x3e, 0xf8, 0x1a,
	0x1a, 0x29, 0x43, 0x5b, 0xaf, 0xac, 0x81, 0x76, 0xf3, 0x8d, 0x1f, 0x7a, 0xaf, 0xc5, 0xba, 0x9a,
	0xac, 0x96, 0xfd, 0x21, 0xd0, 0xdf, 0x00, 0x32, 0xa0, 0xad, 0x47, 0x6b, 0xfb, 0xde, 0x66, 0xa7,
	0x76, 0x2d, 0x91, 0x74, 0x6a, 0x8a, 0x64, 0xaf, 0x54, 0x24, 0xab, 0x89, 0x42, 0x24, 0x7c, 0xb9,
	0x78, 0xa3, 0xa0, 0x12, 0xbe, 0x55, 0x26, 0x07, 0xdb, 0x65, 0x62, 0x43, 0x4f, 0xbe, 0xd7, 0xd4,
	0xab, 0xbd, 0x97, 0x73, 0xa0, 0xb2, 0x46, 0xf0, 0x50, 0x65, 0x03, 0xd8, 0x97, 0x64, 0x89, 0x20,
	0x2b, 0x0f, 0xcc, 0x86, 0x53, 0x99, 0x2b, 0xb9, 0xa8, 0xe4, 0x3b, 0x70, 0x24, 0x29, 0xe7, 0x08,
	0x5d, 0x19, 0x98, 0x7a, 0xac, 0x0f, 0x46, 0x56, 0xf3, 0x6a, 0xa1, 0x7a, 0xb3, 0xcf, 0x72, 0x4a,
	0x11, 0x59, 0x6e, 0xe6, 0x2d, 0x74, 0xd5, 0x48, 0x26, 0x19, 0x75, 0xc6, 0xc7, 0xf6, 0x63, 0x6b,
	0xab, 0xe1, 0xad, 0xf5, 0xc5, 0x3a, 0xba, 0xd8, 0xfe, 0x75, 0x08, 0x5d, 0x75, 0x4d, 0xaf, 0xc0,
	0x28, 0xba, 0x86, 0x9e, 0xd7, 0xea, 0x2a, 0xc6, 0x1c, 0x5e, 0x54, 0xe5, 0x96, 0x9a, 0x91, 0xb5,
	0x32, 0xc8, 0xa2, 0xa7, 0xfe, 0x1b, 0x64, 0xb9, 0x55, 0x59, 0x8b, 0xfe, 0x00, 0xa3, 0x68, 0x2c,
	0xfa, 0xbc, 0xa2, 0x4d, 0xa9, 0x7b, 0x2b, 0x81, 0xcb, 0x5d, 0x2b, 0xb8, 0x0a, 0xe5, 0xa8, 0x1b,
	0xe9, 0x52, 0xfa, 0xa4, 0xa2, 0xd5, 0x86, 0x99, 0x87, 0x4d, 0xd6, 0xcc, 0x5a, 0x34, 0x80, 0x7e,
	0x11, 0x72, 0x3a, 0xa9, 0x7c, 0xe0, 0x35, 0x37, 0x34, 0x85, 0x4b, 0xa4, 0xde, 0x73, 0x38, 0xe9,
	0xf5, 0xa7, 0x3b, 0x11, 0x57, 0xbd, 0x34, 0xdc, 0x35, 0xe0, 0x8a, 0x11, 0x58, 0x8b, 0xa6, 0x30,
	0x28, 0x62, 0x2e, 0xff, 0x37, 0xac, 0x9d, 0xa0, 0x05, 0x53, 0x36, 0x44, 0xfd, 0xa6, 0x3d, 0xa9,
	0x55, 0xf4, 0xa8, 0x4e, 0x83, 0x7f, 0xc0, 0xfa, 0x7a, 0x20, 0xbe, 0xe8, 0x2f, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x58, 0xb8, 0x11, 0x43, 0xf2, 0x07, 0x00, 0x00,
}
