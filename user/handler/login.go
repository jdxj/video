package handler

import (
	"context"

	"github.com/jdxj/video/database"
	"github.com/jdxj/video/logger"
	pb_user "github.com/jdxj/video/user/proto"
)

type LoginService struct {
}

func (ls *LoginService) Login(ctx context.Context, req *pb_user.LoginRequest, resp *pb_user.LoginResponse) error {
	user, err := database.LoginCheck(req.Name, req.Pass)
	if err != nil {
		logger.Error("LoginCheck: %s", err)
		resp.Code = 123
		resp.Message = "err"
		return nil
	}

	resp.Message = "ok"
	resp.Id = int32(user.ID)
	return nil
}
