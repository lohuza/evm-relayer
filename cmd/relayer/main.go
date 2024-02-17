package main

import (
	"github.com/lohuza/relayer/internal/infrastructure/cache"
	"github.com/lohuza/relayer/internal/infrastructure/database"
	"github.com/lohuza/relayer/internal/infrastructure/webserver"
	"github.com/lohuza/relayer/internal/scheduler"
	"github.com/lohuza/relayer/pkg/startup"
	"github.com/spf13/viper"
)

func main() {
	startup.ReadConfig()
	bunDb := database.InitPg(viper.GetString("pg_connection_string"))
	redisDb := cache.InitRedis(viper.GetString("redis.connection_string"), viper.GetInt("redis.db"))

	scheduler := scheduler.NewScheduler()
	scheduler.Start()

	router := webserver.NewWebServer()

	webserver.StartWebServer(router)
}
