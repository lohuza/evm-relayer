package webserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func NewWebServer() *gin.Engine {
	gin.SetMode(viper.GetString("app_mode"))
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}

func StartWebServer(ctx context.Context, wg *sync.WaitGroup, router *gin.Engine) {
	wg.Add(1)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", viper.GetInt("port")),
		Handler: router,
	}
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("HTTP server ListenAndServe error")
		}
	}()

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("HTTP server Shutdown error")
		}

		log.Debug().Msg("HTTP server has shut down gracefully.")
	}()
}
