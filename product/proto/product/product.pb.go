// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/product/product.proto

/*
Package go_micro_service_product is a generated protocol buffer package.

It is generated from these files:
	proto/product/product.proto

It has these top-level messages:
	ProductInfo
	ProductImage
	ProductSize
	ProductSeo
	ResponseProduct
	RequestID
	Response
	RequestAll
	AllProduct
*/
package service_product

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

type ProductInfo struct {
	Id                 int64           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ProductName        string          `protobuf:"bytes,2,opt,name=product_name,json=productName" json:"product_name,omitempty"`
	ProductSku         string          `protobuf:"bytes,3,opt,name=product_sku,json=productSku" json:"product_sku,omitempty"`
	ProductPrice       float64         `protobuf:"fixed64,4,opt,name=product_price,json=productPrice" json:"product_price,omitempty"`
	ProductDescription string          `protobuf:"bytes,5,opt,name=product_description,json=productDescription" json:"product_description,omitempty"`
	ProductCategoryId  int64           `protobuf:"varint,6,opt,name=product_category_id,json=productCategoryId" json:"product_category_id,omitempty"`
	ProductImage       []*ProductImage `protobuf:"bytes,7,rep,name=product_image,json=productImage" json:"product_image,omitempty"`
	ProductSize        []*ProductSize  `protobuf:"bytes,8,rep,name=product_size,json=productSize" json:"product_size,omitempty"`
	ProductSeo         *ProductSeo     `protobuf:"bytes,9,opt,name=product_seo,json=productSeo" json:"product_seo,omitempty"`
}

func (m *ProductInfo) Reset()                    { *m = ProductInfo{} }
func (m *ProductInfo) String() string            { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()               {}
func (*ProductInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProductInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductInfo) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *ProductInfo) GetProductSku() string {
	if m != nil {
		return m.ProductSku
	}
	return ""
}

func (m *ProductInfo) GetProductPrice() float64 {
	if m != nil {
		return m.ProductPrice
	}
	return 0
}

func (m *ProductInfo) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *ProductInfo) GetProductCategoryId() int64 {
	if m != nil {
		return m.ProductCategoryId
	}
	return 0
}

func (m *ProductInfo) GetProductImage() []*ProductImage {
	if m != nil {
		return m.ProductImage
	}
	return nil
}

func (m *ProductInfo) GetProductSize() []*ProductSize {
	if m != nil {
		return m.ProductSize
	}
	return nil
}

func (m *ProductInfo) GetProductSeo() *ProductSeo {
	if m != nil {
		return m.ProductSeo
	}
	return nil
}

type ProductImage struct {
	Id        int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ImageName string `protobuf:"bytes,2,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	ImageCode string `protobuf:"bytes,3,opt,name=image_code,json=imageCode" json:"image_code,omitempty"`
	ImageUrl  string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
}

func (m *ProductImage) Reset()                    { *m = ProductImage{} }
func (m *ProductImage) String() string            { return proto.CompactTextString(m) }
func (*ProductImage) ProtoMessage()               {}
func (*ProductImage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProductImage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductImage) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *ProductImage) GetImageCode() string {
	if m != nil {
		return m.ImageCode
	}
	return ""
}

func (m *ProductImage) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type ProductSize struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SizeName string `protobuf:"bytes,2,opt,name=size_name,json=sizeName" json:"size_name,omitempty"`
	SizeCode string `protobuf:"bytes,3,opt,name=size_code,json=sizeCode" json:"size_code,omitempty"`
}

func (m *ProductSize) Reset()                    { *m = ProductSize{} }
func (m *ProductSize) String() string            { return proto.CompactTextString(m) }
func (*ProductSize) ProtoMessage()               {}
func (*ProductSize) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProductSize) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSize) GetSizeName() string {
	if m != nil {
		return m.SizeName
	}
	return ""
}

func (m *ProductSize) GetSizeCode() string {
	if m != nil {
		return m.SizeCode
	}
	return ""
}

type ProductSeo struct {
	Id             int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SeoTitle       string `protobuf:"bytes,2,opt,name=seo_title,json=seoTitle" json:"seo_title,omitempty"`
	SeoKeywords    string `protobuf:"bytes,3,opt,name=seo_keywords,json=seoKeywords" json:"seo_keywords,omitempty"`
	SeoDescription string `protobuf:"bytes,4,opt,name=seo_description,json=seoDescription" json:"seo_description,omitempty"`
	SeoCode        string `protobuf:"bytes,6,opt,name=seo_code,json=seoCode" json:"seo_code,omitempty"`
}

func (m *ProductSeo) Reset()                    { *m = ProductSeo{} }
func (m *ProductSeo) String() string            { return proto.CompactTextString(m) }
func (*ProductSeo) ProtoMessage()               {}
func (*ProductSeo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProductSeo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSeo) GetSeoTitle() string {
	if m != nil {
		return m.SeoTitle
	}
	return ""
}

func (m *ProductSeo) GetSeoKeywords() string {
	if m != nil {
		return m.SeoKeywords
	}
	return ""
}

func (m *ProductSeo) GetSeoDescription() string {
	if m != nil {
		return m.SeoDescription
	}
	return ""
}

func (m *ProductSeo) GetSeoCode() string {
	if m != nil {
		return m.SeoCode
	}
	return ""
}

type ResponseProduct struct {
	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *ResponseProduct) Reset()                    { *m = ResponseProduct{} }
func (m *ResponseProduct) String() string            { return proto.CompactTextString(m) }
func (*ResponseProduct) ProtoMessage()               {}
func (*ResponseProduct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResponseProduct) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type RequestID struct {
	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *RequestID) Reset()                    { *m = RequestID{} }
func (m *RequestID) String() string            { return proto.CompactTextString(m) }
func (*RequestID) ProtoMessage()               {}
func (*RequestID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RequestID) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RequestAll struct {
}

func (m *RequestAll) Reset()                    { *m = RequestAll{} }
func (m *RequestAll) String() string            { return proto.CompactTextString(m) }
func (*RequestAll) ProtoMessage()               {}
func (*RequestAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AllProduct struct {
	ProductInfo []*ProductInfo `protobuf:"bytes,1,rep,name=product_info,json=productInfo" json:"product_info,omitempty"`
}

func (m *AllProduct) Reset()                    { *m = AllProduct{} }
func (m *AllProduct) String() string            { return proto.CompactTextString(m) }
func (*AllProduct) ProtoMessage()               {}
func (*AllProduct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AllProduct) GetProductInfo() []*ProductInfo {
	if m != nil {
		return m.ProductInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductInfo)(nil), "go.micro.service.product.ProductInfo")
	proto.RegisterType((*ProductImage)(nil), "go.micro.service.product.ProductImage")
	proto.RegisterType((*ProductSize)(nil), "go.micro.service.product.ProductSize")
	proto.RegisterType((*ProductSeo)(nil), "go.micro.service.product.ProductSeo")
	proto.RegisterType((*ResponseProduct)(nil), "go.micro.service.product.ResponseProduct")
	proto.RegisterType((*RequestID)(nil), "go.micro.service.product.RequestID")
	proto.RegisterType((*Response)(nil), "go.micro.service.product.Response")
	proto.RegisterType((*RequestAll)(nil), "go.micro.service.product.RequestAll")
	proto.RegisterType((*AllProduct)(nil), "go.micro.service.product.AllProduct")
}

func init() { proto.RegisterFile("proto/product/product.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x6b, 0x68, 0xe3, 0x49, 0x7f, 0xe8, 0xf2, 0x62, 0x12, 0x10, 0x61, 0x5b, 0x20, 0xf0,
	0xe0, 0xa2, 0x70, 0x82, 0xd0, 0x80, 0x88, 0x2a, 0xa1, 0xca, 0xa5, 0xf0, 0x82, 0x30, 0xc1, 0x3b,
	0x8d, 0x56, 0x71, 0xbc, 0xc6, 0xeb, 0x80, 0xd2, 0xd3, 0x70, 0x19, 0x2e, 0xc3, 0x29, 0xd0, 0x6e,
	0x76, 0x6d, 0xab, 0xd0, 0x26, 0x3c, 0xc5, 0x33, 0xf3, 0xf9, 0x9b, 0x6f, 0xbe, 0x19, 0xc5, 0xd0,
	0xc9, 0x72, 0x51, 0x88, 0xa3, 0x2c, 0x17, 0x6c, 0x1e, 0x17, 0xf6, 0x37, 0xd0, 0x59, 0xe2, 0x4f,
	0x44, 0x30, 0xe3, 0x71, 0x2e, 0x02, 0x89, 0xf9, 0x77, 0x1e, 0x63, 0x60, 0xea, 0xf4, 0x97, 0x0b,
	0xad, 0xd3, 0xe5, 0xf3, 0x28, 0xbd, 0x10, 0x64, 0x17, 0x36, 0x38, 0xf3, 0x9d, 0xae, 0xd3, 0x73,
	0xc3, 0x0d, 0xce, 0xc8, 0x23, 0xd8, 0x36, 0xd0, 0x28, 0x1d, 0xcf, 0xd0, 0xdf, 0xe8, 0x3a, 0x3d,
	0x2f, 0x6c, 0x99, 0xdc, 0xbb, 0xf1, 0x0c, 0xc9, 0x43, 0xb0, 0x61, 0x24, 0xa7, 0x73, 0xdf, 0xd5,
	0x08, 0x30, 0xa9, 0xb3, 0xe9, 0x9c, 0x1c, 0xc0, 0x8e, 0x05, 0x64, 0x39, 0x8f, 0xd1, 0xbf, 0xd5,
	0x75, 0x7a, 0x4e, 0x68, 0x89, 0x4f, 0x55, 0x8e, 0x1c, 0xc1, 0x5d, 0x0b, 0x62, 0x28, 0xe3, 0x9c,
	0x67, 0x05, 0x17, 0xa9, 0x7f, 0x5b, 0xb3, 0x11, 0x53, 0x1a, 0x56, 0x15, 0x12, 0x54, 0x2f, 0xc4,
	0xe3, 0x02, 0x27, 0x22, 0x5f, 0x44, 0x9c, 0xf9, 0x9b, 0x5a, 0xfa, 0xbe, 0x29, 0x1d, 0x9b, 0xca,
	0x88, 0x91, 0x93, 0x4a, 0x05, 0x9f, 0x8d, 0x27, 0xe8, 0x6f, 0x75, 0xdd, 0x5e, 0xab, 0xff, 0x24,
	0xb8, 0xce, 0x9b, 0xc0, 0xfa, 0xa2, 0xd0, 0xa5, 0x5a, 0x1d, 0x91, 0xb7, 0x95, 0x2d, 0x92, 0x5f,
	0xa2, 0xdf, 0xd4, 0x5c, 0x8f, 0x57, 0x72, 0x9d, 0xf1, 0x4b, 0x2c, 0xdd, 0x53, 0x01, 0x79, 0x5d,
	0x73, 0x0f, 0x85, 0xef, 0x75, 0x9d, 0x5e, 0xab, 0x7f, 0xb8, 0x9a, 0x08, 0x45, 0xe5, 0x31, 0x0a,
	0xba, 0x80, 0xed, 0xba, 0xdc, 0xbf, 0xf6, 0xf8, 0x00, 0x40, 0x4f, 0x5d, 0xdf, 0xa2, 0xa7, 0x33,
	0x7a, 0x87, 0x65, 0x39, 0x16, 0x0c, 0xcd, 0x0a, 0x97, 0xe5, 0x63, 0xc1, 0x90, 0x74, 0x60, 0x19,
	0x44, 0xf3, 0x3c, 0xd1, 0xdb, 0xf3, 0xc2, 0xa6, 0x4e, 0x9c, 0xe7, 0x09, 0xfd, 0x58, 0x5e, 0x90,
	0x1e, 0xe8, 0x6a, 0xe7, 0x0e, 0x78, 0xca, 0xa2, 0x7a, 0xe3, 0xa6, 0x4a, 0xe8, 0xbe, 0xb6, 0x58,
	0x6b, 0xab, 0x8b, 0xaa, 0x2b, 0xfd, 0xe9, 0x00, 0x54, 0xe3, 0xfe, 0x93, 0x18, 0x45, 0x54, 0xf0,
	0x22, 0xa9, 0x88, 0x51, 0xbc, 0x57, 0xb1, 0xba, 0x5b, 0x55, 0x9c, 0xe2, 0xe2, 0x87, 0xc8, 0x99,
	0x34, 0xdc, 0x2d, 0x89, 0xe2, 0xc4, 0xa4, 0xc8, 0x53, 0xd8, 0x53, 0x90, 0xfa, 0xb5, 0x2d, 0x47,
	0xdb, 0x95, 0x28, 0xea, 0x97, 0x76, 0x0f, 0x14, 0xef, 0x52, 0xe3, 0xa6, 0x46, 0x6c, 0x49, 0x14,
	0x5a, 0xe2, 0x0b, 0xd8, 0x0b, 0x51, 0x66, 0x22, 0x95, 0x68, 0x94, 0x2a, 0x2b, 0xcb, 0x3b, 0xb3,
	0x72, 0x3d, 0x7b, 0x3c, 0x8c, 0x3e, 0x07, 0x2f, 0xc4, 0x6f, 0x73, 0x94, 0xc5, 0x68, 0xb8, 0x0a,
	0x7b, 0x1f, 0x9a, 0x96, 0x9d, 0xdc, 0x01, 0x77, 0x26, 0x27, 0x1a, 0xe3, 0x85, 0xea, 0x91, 0x6e,
	0x03, 0x18, 0xa6, 0x41, 0x92, 0xd0, 0x0f, 0x00, 0x83, 0x24, 0xb1, 0x22, 0x6a, 0xf7, 0xc9, 0xd3,
	0x0b, 0xe1, 0x3b, 0x6b, 0xde, 0xa7, 0xfa, 0x0f, 0x28, 0xef, 0x53, 0x05, 0xfd, 0xdf, 0x2e, 0x6c,
	0x59, 0xd6, 0x2f, 0x00, 0x03, 0xc6, 0x6c, 0xb4, 0x1e, 0x5b, 0xfb, 0xd9, 0xf5, 0xb0, 0x2b, 0xd6,
	0xd1, 0x06, 0x89, 0x60, 0xef, 0x0d, 0x4f, 0x6d, 0x8b, 0x57, 0x8b, 0xd1, 0x90, 0x1c, 0xdc, 0xf4,
	0xbe, 0x31, 0xb2, 0xbd, 0x9e, 0x16, 0xda, 0x20, 0x9f, 0x60, 0xe7, 0x3c, 0x63, 0xe3, 0x02, 0xff,
	0x73, 0x0a, 0xba, 0x7a, 0x0a, 0xda, 0x20, 0x9f, 0x61, 0x7f, 0x88, 0x09, 0x96, 0xec, 0xeb, 0x0f,
	0xb0, 0x2e, 0xff, 0xae, 0xb2, 0xa7, 0xb6, 0xe8, 0xc3, 0x95, 0xe4, 0x83, 0x24, 0x69, 0xdf, 0x80,
	0xaa, 0xb8, 0x68, 0xe3, 0xeb, 0xa6, 0xfe, 0x5c, 0xbc, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xac,
	0x03, 0x99, 0x33, 0x4d, 0x06, 0x00, 0x00,
}
