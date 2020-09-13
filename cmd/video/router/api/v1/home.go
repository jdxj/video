package v1

import (
	"net/http"
	"video/cmd/video/router/api"

	"github.com/gin-gonic/gin"
)

func RegisterAPI(r *gin.RouterGroup) {
	r.GET("v1", Home)
}

func Home(c *gin.Context) {
	resp := &api.Response{
		Code:    0,
		Message: "ok",
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}
