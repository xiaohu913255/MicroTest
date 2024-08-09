// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/payment/payment.proto

package service_payment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Payment service

func NewPaymentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Payment service

type PaymentService interface {
	AddPayment(ctx context.Context, in *PaymentInfo, opts ...client.CallOption) (*PaymentID, error)
	UpdatePayment(ctx context.Context, in *PaymentInfo, opts ...client.CallOption) (*Response, error)
	DeletePaymentByID(ctx context.Context, in *PaymentID, opts ...client.CallOption) (*Response, error)
	FindPaymentByID(ctx context.Context, in *PaymentID, opts ...client.CallOption) (*PaymentInfo, error)
	FindAllPayment(ctx context.Context, in *All, opts ...client.CallOption) (*PaymentAll, error)
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	return &paymentService{
		c:    c,
		name: name,
	}
}

func (c *paymentService) AddPayment(ctx context.Context, in *PaymentInfo, opts ...client.CallOption) (*PaymentID, error) {
	req := c.c.NewRequest(c.name, "Payment.AddPayment", in)
	out := new(PaymentID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) UpdatePayment(ctx context.Context, in *PaymentInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Payment.UpdatePayment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) DeletePaymentByID(ctx context.Context, in *PaymentID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Payment.DeletePaymentByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) FindPaymentByID(ctx context.Context, in *PaymentID, opts ...client.CallOption) (*PaymentInfo, error) {
	req := c.c.NewRequest(c.name, "Payment.FindPaymentByID", in)
	out := new(PaymentInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) FindAllPayment(ctx context.Context, in *All, opts ...client.CallOption) (*PaymentAll, error) {
	req := c.c.NewRequest(c.name, "Payment.FindAllPayment", in)
	out := new(PaymentAll)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Payment service

type PaymentHandler interface {
	AddPayment(context.Context, *PaymentInfo, *PaymentID) error
	UpdatePayment(context.Context, *PaymentInfo, *Response) error
	DeletePaymentByID(context.Context, *PaymentID, *Response) error
	FindPaymentByID(context.Context, *PaymentID, *PaymentInfo) error
	FindAllPayment(context.Context, *All, *PaymentAll) error
}

func RegisterPaymentHandler(s server.Server, hdlr PaymentHandler, opts ...server.HandlerOption) error {
	type payment interface {
		AddPayment(ctx context.Context, in *PaymentInfo, out *PaymentID) error
		UpdatePayment(ctx context.Context, in *PaymentInfo, out *Response) error
		DeletePaymentByID(ctx context.Context, in *PaymentID, out *Response) error
		FindPaymentByID(ctx context.Context, in *PaymentID, out *PaymentInfo) error
		FindAllPayment(ctx context.Context, in *All, out *PaymentAll) error
	}
	type Payment struct {
		payment
	}
	h := &paymentHandler{hdlr}
	return s.Handle(s.NewHandler(&Payment{h}, opts...))
}

type paymentHandler struct {
	PaymentHandler
}

func (h *paymentHandler) AddPayment(ctx context.Context, in *PaymentInfo, out *PaymentID) error {
	return h.PaymentHandler.AddPayment(ctx, in, out)
}

func (h *paymentHandler) UpdatePayment(ctx context.Context, in *PaymentInfo, out *Response) error {
	return h.PaymentHandler.UpdatePayment(ctx, in, out)
}

func (h *paymentHandler) DeletePaymentByID(ctx context.Context, in *PaymentID, out *Response) error {
	return h.PaymentHandler.DeletePaymentByID(ctx, in, out)
}

func (h *paymentHandler) FindPaymentByID(ctx context.Context, in *PaymentID, out *PaymentInfo) error {
	return h.PaymentHandler.FindPaymentByID(ctx, in, out)
}

func (h *paymentHandler) FindAllPayment(ctx context.Context, in *All, out *PaymentAll) error {
	return h.PaymentHandler.FindAllPayment(ctx, in, out)
}
