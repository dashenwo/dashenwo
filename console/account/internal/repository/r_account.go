package repository

import "github.com/dashenwo/dashenwo/console/account/schema"

// 用户接口
type UserRepository interface {
	FindByName(name string) (*schema.Account, error)
	Insert(account *schema.Account) error
}
