package gorm

import (
	"github.com/dashenwo/dashenwo/console/account/config"
	"github.com/dashenwo/dashenwo/console/account/internal/model"
	"github.com/dashenwo/dashenwo/console/account/internal/repository"
	"github.com/micro/go-micro/v2/errors"
)

type AccountRepository struct {
}

func NewAccountRepository() repository.UserRepository {
	return &AccountRepository{}
}

func (a *AccountRepository) FindByName(name string) (*model.Account, error) {

	return nil, nil
}

func (a *AccountRepository) Insert(account *model.Account) error {
	if err := db.Create(account).Error; err != nil {
		return errors.New(config.AppId, err.Error(), 201)
	}
	return nil
}
