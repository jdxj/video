package service

import (
	"video/cmd/user/handler"
	"video/config"
	"video/database"
	"video/logger"
	proto_user "video/proto/user"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

const (
	Name = "user"
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
		}))

	err := proto_user.RegisterLoginServiceHandler(
		service.Server(),
		new(handler.LoginService),
	)
	if err != nil {
		logger.Error("RegisterLoginServiceHandler: %s", err)
		return
	}

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