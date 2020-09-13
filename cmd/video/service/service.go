package service

import (
	"video/cmd/video/router"
	"video/config"
	"video/database"
	"video/logger"
	proto_user "video/proto/user"

	"github.com/micro/cli/v2"

	"github.com/micro/go-micro/v2"
)

const (
	Name = "video"
)

var (
	LoginService proto_user.LoginService
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

	LoginService = proto_user.NewLoginService("user", service.Client())

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
