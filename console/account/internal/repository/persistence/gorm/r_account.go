package gorm

import (
	"github.com/dashenwo/dashenwo/console/account/internal/repository"
	"github.com/dashenwo/dashenwo/console/account/schema"
)

type AccountRepository struct {
}

func NewAccountRepository() repository.UserRepository {
	return &AccountRepository{}
}

func (a *AccountRepository) FindByName(name string) (*schema.Account, error) {

	return nil, nil
}

func (a *AccountRepository) Insert(account *schema.Account) error {
	return nil
}
