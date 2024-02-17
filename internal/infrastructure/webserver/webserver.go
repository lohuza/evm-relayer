package webserver

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewWebServer() *gin.Engine {
	gin.SetMode(viper.GetString("app_mode"))
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}

func StartWebServer(router *gin.Engine) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", viper.GetInt("microservice_port")),
		Handler: router,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return server.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
