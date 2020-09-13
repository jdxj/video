package router

import (
	v1 "video/cmd/video/router/api/v1"
	"video/config"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(config.Mod)
	r := gin.Default()

	apiGroup := r.Group("api")
	v1.RegisterAPI(apiGroup)

	return r
}
