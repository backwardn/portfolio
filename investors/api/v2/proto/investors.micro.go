// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/investors.proto

package investors

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Investors service

type InvestorsService interface {
	Discover(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Connections(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	Follow(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	Unfollow(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
}

type investorsService struct {
	c    client.Client
	name string
}

func NewInvestorsService(name string, c client.Client) InvestorsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "investors"
	}
	return &investorsService{
		c:    c,
		name: name,
	}
}

func (c *investorsService) Discover(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Investors.Discover", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorsService) Connections(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Investors.Connections", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorsService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Investors.Search", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorsService) Get(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Investors.Get", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorsService) Follow(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Investors.Follow", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorsService) Unfollow(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "Investors.Unfollow", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Investors service

type InvestorsHandler interface {
	Discover(context.Context, *ListRequest, *ListResponse) error
	Connections(context.Context, *ListRequest, *ListResponse) error
	Search(context.Context, *SearchRequest, *ListResponse) error
	Get(context.Context, *User, *User) error
	Follow(context.Context, *User, *User) error
	Unfollow(context.Context, *User, *User) error
}

func RegisterInvestorsHandler(s server.Server, hdlr InvestorsHandler, opts ...server.HandlerOption) error {
	type investors interface {
		Discover(ctx context.Context, in *ListRequest, out *ListResponse) error
		Connections(ctx context.Context, in *ListRequest, out *ListResponse) error
		Search(ctx context.Context, in *SearchRequest, out *ListResponse) error
		Get(ctx context.Context, in *User, out *User) error
		Follow(ctx context.Context, in *User, out *User) error
		Unfollow(ctx context.Context, in *User, out *User) error
	}
	type Investors struct {
		investors
	}
	h := &investorsHandler{hdlr}
	return s.Handle(s.NewHandler(&Investors{h}, opts...))
}

type investorsHandler struct {
	InvestorsHandler
}

func (h *investorsHandler) Discover(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.InvestorsHandler.Discover(ctx, in, out)
}

func (h *investorsHandler) Connections(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.InvestorsHandler.Connections(ctx, in, out)
}

func (h *investorsHandler) Search(ctx context.Context, in *SearchRequest, out *ListResponse) error {
	return h.InvestorsHandler.Search(ctx, in, out)
}

func (h *investorsHandler) Get(ctx context.Context, in *User, out *User) error {
	return h.InvestorsHandler.Get(ctx, in, out)
}

func (h *investorsHandler) Follow(ctx context.Context, in *User, out *User) error {
	return h.InvestorsHandler.Follow(ctx, in, out)
}

func (h *investorsHandler) Unfollow(ctx context.Context, in *User, out *User) error {
	return h.InvestorsHandler.Unfollow(ctx, in, out)
}