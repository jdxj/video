package main

import (
	"github.com/jdxj/logger"
	"github.com/jdxj/video/config"
	"github.com/jdxj/video/model"
	"github.com/jdxj/video/service"
)

func main() {
	err := config.Init("./config.yaml")
	if err != nil {
		panic(err)
	}

	logger.Init(config.Log().Path, config.Mode())

	dbCfg := config.DB()
	err = model.InitDB(dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.DBName)
	if err != nil {
		panic(err)
	}

	service.InitRemoteService()
}
