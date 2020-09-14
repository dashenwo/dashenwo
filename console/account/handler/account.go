package handler

import (
	"context"
	"encoding/json"
	conf "github.com/dashenwo/dashenwo/console/account/config"
	"github.com/dashenwo/dashenwo/console/account/global"
	"github.com/dashenwo/dashenwo/console/account/internal/service"
	"github.com/dashenwo/dashenwo/console/account/proto"
	"github.com/dashenwo/dashenwo/pkg/crypto"
	"github.com/dashenwo/dashenwo/pkg/utils/validate"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/util/log"
	"strconv"
	"time"
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
	//1.验证数据
	if err := validate.Validate(req, conf.AppId); err != nil {
		return err
	}
	user, err := a.accountService.Login(req.Username, req.Password)
	if err != nil {
		return err
	}
	// 生成token
	token := crypto.Md5("userToken_" + req.Source + "_" + strconv.Itoa(user.ID))
	log.Info(global.Redis)
	// 保存数据到redis
	data, _ := json.Marshal(user)
	if err := global.Redis.Set(token, string(data), time.Hour*2).Err(); err != nil {
		return errors.New(conf.AppId, "设置用户登录状态失败", 504)
	}
	now := time.Now()
	hh, _ := time.ParseDuration("1h")
	// 返回token
	res.Token = token
	res.Expires = now.Add(2 * hh).Format("2006-01-02 15:04:05")
	log.Info(user, err)
	return nil
}

// 注册handler
func (a *Account) Register(ctx context.Context, req *proto.RegisterRequest, res *proto.RegisterResponse) error {
	log.Info("进入了注册方法", req)
	//1.验证数据
	if err := validate.Validate(req, conf.AppId); err != nil {
		return err
	}
	// 1.调用短信服务，验证当前短信验证码是否正确

	// 2.进入注册
	account, err := a.accountService.Register(req.Nickname, req.Password, req.Phone)
	if err != nil {
		return err
	}
	res.Id = int32(account.ID)
	return nil
}
