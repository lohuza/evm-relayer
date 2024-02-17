package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"

	"github.com/lohuza/relayer/internal/infrastructure/cache"
	"github.com/lohuza/relayer/internal/infrastructure/database"
	"github.com/lohuza/relayer/internal/infrastructure/webserver"
	"github.com/lohuza/relayer/internal/scheduler"
	"github.com/lohuza/relayer/pkg/startup"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	startup.ReadConfig()
	bunDb := database.InitPg(viper.GetString("pg_connection_string"))
	redisDb := cache.InitRedis(viper.GetString("redis.connection_string"), viper.GetInt("redis.db"))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	scheduler := scheduler.NewScheduler(ctx, &wg)
	scheduler.Start()

	router := webserver.NewWebServer()

	webserver.StartWebServer(ctx, &wg, router)

	<-ctx.Done()
	stop()
	wg.Wait()

	log.Debug().Msg("graceful shutdown complete")
}
