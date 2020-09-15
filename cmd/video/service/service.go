package service

import (
	"github.com/jdxj/video/cmd/video/other"
	"github.com/jdxj/video/cmd/video/router"
	"github.com/jdxj/video/config"
	"github.com/jdxj/video/database"
	"github.com/jdxj/video/logger"
	proto_user "github.com/jdxj/video/proto/user"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

const (
	Name = "video"
)

func StartService() {
	service := micro.NewService(
		micro.Flags(&cli.StringFlag{
			Name:  "c",
			Usage: "config file path",
			Value: "config.yaml",
		}),
		micro.Name(Name),
	)

	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			path := ctx.String("c")
			return InitBase(path)
		}),
		micro.AfterStart(router.StartServer),
		micro.BeforeStop(router.StopServer),
	)

	other.LoginService = proto_user.NewLoginService("user", service.Client())

	if err := service.Run(); err != nil {
		logger.Error("Run: %s", err)
	}
}

func InitBase(path string) error {
	err := config.Init(path)
	if err != nil {
		return err
	}

	logger.Init()

	return database.Init()
}
