package server

import (
	pb_user "github.com/jdxj/video/user/proto"
	"github.com/micro/micro/v3/service"
)

var (
	LoginService pb_user.LoginService
)

func InitRemoteService() {
	srv := service.New()
	srv.Init()

	LoginService = pb_user.NewLoginService("user", srv.Client())

	//srv.Run()
}
