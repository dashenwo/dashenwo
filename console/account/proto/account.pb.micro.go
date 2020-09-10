// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: account.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Api Endpoints for Account service

func NewAccountEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Account service

type AccountService interface {
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) error {
	type account interface {
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
	}
	type Account struct {
		account
	}
	h := &accountHandler{hdlr}
	return s.Handle(s.NewHandler(&Account{h}, opts...))
}

type accountHandler struct {
	AccountHandler
}

func (h *accountHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *accountHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.AccountHandler.Register(ctx, in, out)
}
