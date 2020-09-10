package registry

import (
	"github.com/dashenwo/dashenwo/console/account/internal/repository/persistence/gorm"
	"github.com/dashenwo/dashenwo/console/account/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2/util/log"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()
	buildAccountUsecase(c)
	return c, nil
}

func buildAccountUsecase(c *dig.Container) {
	// DB初始化
	gorm.InitDb()
	// 初始化elasticsearch
	gorm.InitElasticsearch()

	err2 := c.Provide(gorm.NewAccountRepository)
	log.Info(err2)
	err3 := c.Provide(service.NewAccountService)
	log.Info(err3)
}
