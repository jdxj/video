package v1

import (
	"context"
	"net/http"
	"video/cmd/video/other"
	"video/cmd/video/router/api"
	"video/logger"
	proto_user "video/proto/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AddSession(c *gin.Context) {
	user := new(api.User)
	err := c.ShouldBind(user)
	if err != nil {
		logger.Error("ShouldBind: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	loginReq := &proto_user.LoginRequest{
		Name: user.Name,
		Pass: user.Password,
	}
	loginResp, err := other.LoginService.Login(context.TODO(), loginReq)
	if err != nil {
		logger.Error("Login: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if loginResp.Code != 0 {
		resp := api.NewResponse(int(loginResp.Code), loginResp.Message, nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	uc := api.NewUserClaims(int(loginResp.Id), user.Name)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	secret, _ := KeyFunc(nil)
	ss, err := token.SignedString(secret)
	if err != nil {
		logger.Error("SignedString: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := api.NewResponse(0, "token ok", ss)
	c.JSON(http.StatusOK, resp)
}
