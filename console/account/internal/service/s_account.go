package service

import (
	conf "github.com/dashenwo/dashenwo/console/account/config"
	"github.com/dashenwo/dashenwo/console/account/internal/model"
	"github.com/dashenwo/dashenwo/console/account/internal/repository"
	"github.com/dashenwo/dashenwo/console/account/schema"
	"github.com/dashenwo/dashenwo/pkg/crypto"
	"github.com/jinzhu/copier"
	"github.com/micro/go-micro/v2/errors"
	"time"
)

type AccountService struct {
	repo repository.UserRepository
}

func NewAccountService(repo repository.UserRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

// 登录方法
func (s AccountService) Login(username string, password string) (*schema.Account, error) {
	// 查询用户
	account, err := s.repo.FindByName(username)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errors.New(conf.AppId, "账号或者密码错误", 501)
	}
	isLogin := crypto.ComparePasswords(account.Password, password+account.Salt)
	if isLogin == false {
		return nil, errors.New(conf.AppId, "账号或者密码错误", 502)
	}
	item := new(schema.Account)
	_ = copier.Copy(item, account)
	return item, err
}

// 注册方法
func (s AccountService) Register(nickname, password, phone string) (*schema.Account, error) {
	account := &model.Account{}
	account.Nickname = nickname
	account.Phone = phone
	account.Salt = crypto.GetRandomString(8)
	account.Password = crypto.HashAndSalt(password, account.Salt)
	account.RegisterTime = time.Now().Unix()
	err := s.repo.Insert(account)
	if err != nil {
		return nil, err
	}
	item := new(schema.Account)
	_ = copier.Copy(item, account)
	return item, nil
}
