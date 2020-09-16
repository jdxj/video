package server

import (
	"context"
	"fmt"
	"time"

	user "github.com/jdxj/user/proto"
	"github.com/jdxj/video/logger"
	"github.com/micro/micro/v3/service"
)

var (
	UserService user.UserService
)

func InitRemoteService() {
	srv := service.New(
		service.Name("video"))
	srv.Init()

	UserService = user.NewUserService("user", srv.Client())

	go func() {
		time.Sleep(3 * time.Second)
		Te()
	}()

	if err := srv.Run(); err != nil {
		panic(err)
	}

}

func Te() {
	req := &user.RequestLogin{
		Name:     "jdxj",
		Password: "jdxj",
	}
	resp, err := UserService.Login(context.TODO(), req)
	if err != nil {
		logger.Error("Login: %s", err)
	}
	fmt.Printf("%#v\n", resp)
}
