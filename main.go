package main

import (
	"github.com/jdxj/video/config"
	"github.com/jdxj/video/logger"
	"github.com/jdxj/video/model"
	"github.com/jdxj/video/server"
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

	server.InitRemoteService()
}