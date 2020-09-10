package handler

import (
	"context"
	"github.com/dashenwo/dashenwo/console/account/internal/service"
	"github.com/dashenwo/dashenwo/console/account/proto"
	"github.com/micro/go-micro/v2/util/log"
)

type Account struct {
	accountService *service.AccountService
}

// 实例化方法
func NewAccountHandler(accountService *service.AccountService) *Account {
	return &Account{
		accountService: accountService,
	}
}

// 登录handler
func (a *Account) Login(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {
	log.Info("进来了登录方法", req)
	user, err := a.accountService.Login(req.Username, req.Password)

	return nil
}

// 注册handler
func (a *Account) Register(ctx context.Context, req *proto.RegisterRequest, res *proto.RegisterResponse) error {
	return nil
}
