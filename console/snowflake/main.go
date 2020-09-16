package main

import (
	"context"
	conf "github.com/dashenwo/dashenwo/console/snowflake/config"
	"github.com/dashenwo/dashenwo/console/snowflake/handler"
	"github.com/dashenwo/dashenwo/console/snowflake/registry"
	tracer "github.com/dashenwo/dashenwo/pkg/opentracing"
	"github.com/dashenwo/dashenwo/pkg/plugins/wrapper/trace/opentracing"
	_ "github.com/dashenwo/dashenwo/profile"
	"github.com/dashenwo/plugins/logger/zap/v3"
	"github.com/micro/go-micro/v3/config"
	"github.com/micro/go-micro/v3/config/source/file"
	"github.com/micro/go-micro/v3/util/log"

	"github.com/micro/micro/v3/profile"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/urfave/cli/v2"
	zap2 "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func main() {

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "profile",
			Usage: "micro profile",
		},
		// TODO V3 命令行参数报错
		&cli.StringFlag{
			Name:  "conf_path",
			Value: "./conf/",
			Usage: "配置文件目录",
		},
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "starter kit debug.",
		},
	}
	app.Before = func(ctx *cli.Context) error {
		p := ctx.String("profile")
		// apply the profile
		profileD, err := profile.Load(p)
		if err != nil {
			logger.Fatal(err)
		}
		return profileD.Setup(ctx)
	}

	app.Action = func(ctx *cli.Context) error {
		confPath := ctx.String("conf_path")
		conf.ConfPath = confPath

		// TODO 配置加载，不在开箱即用
		_, err := config.NewConfig(
			config.WithSource(
				file.NewSource(
					file.WithPath(conf.ConfPath + "config.yaml"),
				),
			),
		)
		if err != nil {
			return err
		}
		return run()
	}

	app.Commands = cli.Commands{
		&cli.Command{
			Name:  "reload",
			Usage: "TODO",
			Action: func(ctx *cli.Context) error {
				return nil
			},
		},
	}

	ctx := context.TODO()
	if err := app.RunContext(ctx, os.Args); err != nil {
		logger.Fatal(err)
	}

}

func run() error {

	// 初始化日志库
	encodingConfig := zap2.NewProductionEncoderConfig()
	// 时间格式化
	encodingConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	l, err := zap.NewLogger(
		zap.WithCallerSkip(4),
		zap.WithConfig(zap2.NewProductionConfig()),
		zap.WithEncoderConfig(encodingConfig),
	)
	if err != nil {
		log.Fatal(err)
	}
	logger.DefaultLogger = l

	md := make(map[string]string)
	md["chain"] = "gray"

	srv := service.New(
		// 版本号
		service.Version("v3"),
		// 服务名
		service.Name(conf.AppId),
		// 启动端口
		service.Address(":8001"),
		// 附加信息
		service.Metadata(md),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer(conf.AppId, "203.195.200.40:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	srv.Init(
		// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
		service.WrapCall(opentracing.NewCallWrapper(t)),
		service.WrapHandler(opentracing.NewHandlerWrapper(t)),
	)

	c, err := registry.NewContainer(srv.Server())
	if err != nil {
		logger.Fatalf("failed to build container: %v", err)
	}

	// Register Handler
	handler.Apply(srv.Server(), c)

	// Run service
	return srv.Run()
}
