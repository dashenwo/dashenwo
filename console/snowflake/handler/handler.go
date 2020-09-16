package handler

import (
	"github.com/dashenwo/dashenwo/console/snowflake/proto"
	"github.com/micro/go-micro/v3/server"
	"go.uber.org/dig"
)

func Apply(server server.Server, c *dig.Container) {
	_ = c.Invoke(func() {
		_ = proto.RegisterSnowflakeHandler(server, NewSnowflake())
	})
}
