package router

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"video/config"
	"video/logger"
)

var (
	server *http.Server
)

func initServer() {
	server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "127.0.0.1", config.Server.Port),
		Handler: NewRouter(),
	}
}

func StartServer() (err error) {
	initServer()
	logger.Info("initServer success")

	go func() {
		logger.Info("ListenAndServe before")

		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			logger.Error("ListenAndServe: %s", err)
			return
		}

		logger.Info("ListenAndServe after")
	}()
	return
}

func StopServer() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}
