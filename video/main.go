package main

import (
	pb_user "github.com/jdxj/video/user/proto"
	"github.com/micro/micro/v3/service"
)

func main() {
	srv := service.New()
	srv.Init()

	pb_user.NewLoginService("user", srv.Client())
}
