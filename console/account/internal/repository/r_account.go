package repository

import (
	"github.com/dashenwo/dashenwo/console/account/internal/model"
)

// 用户接口
type UserRepository interface {
	FindByName(name string) (*model.Account, error)
	Insert(account *model.Account) error
}
