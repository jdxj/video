package router

import (
	v1 "github.com/jdxj/video/cmd/video/router/api/v1"
	"github.com/jdxj/video/config"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(config.Mod)
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	apiGroup := r.Group("api")
	v1.RegisterAPI(apiGroup)

	return r
}
