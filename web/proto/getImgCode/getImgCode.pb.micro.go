// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: getImgCode.proto

package getImgCode

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for GetImgCode service

func NewGetImgCodeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GetImgCode service

type GetImgCodeService interface {
	Call(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*Response, error)
}

type getImgCodeService struct {
	c    client.Client
	name string
}

func NewGetImgCodeService(name string, c client.Client) GetImgCodeService {
	return &getImgCodeService{
		c:    c,
		name: name,
	}
}

func (c *getImgCodeService) Call(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetImgCode.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetImgCode service

type GetImgCodeHandler interface {
	Call(context.Context, *emptypb.Empty, *Response) error
}

func RegisterGetImgCodeHandler(s server.Server, hdlr GetImgCodeHandler, opts ...server.HandlerOption) error {
	type getImgCode interface {
		Call(ctx context.Context, in *emptypb.Empty, out *Response) error
	}
	type GetImgCode struct {
		getImgCode
	}
	h := &getImgCodeHandler{hdlr}
	return s.Handle(s.NewHandler(&GetImgCode{h}, opts...))
}

type getImgCodeHandler struct {
	GetImgCodeHandler
}

func (h *getImgCodeHandler) Call(ctx context.Context, in *emptypb.Empty, out *Response) error {
	return h.GetImgCodeHandler.Call(ctx, in, out)
}
