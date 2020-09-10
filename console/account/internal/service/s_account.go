package service

import (
	"github.com/dashenwo/dashenwo/console/account/internal/repository"
	"github.com/dashenwo/dashenwo/console/account/schema"
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
		return nil, nil
	}
	if password != account.Password {
		return nil, nil
	}
	return account, err
}
