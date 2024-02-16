package startup

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func InitRedis(redisAddress string, redisDb int) *redis.Client {
	log.Info().Msg("connecting to redis")
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddress,
		DB:   redisDb,
	})

	err := rdb.Ping(context.TODO()).Err()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to ping redis")
	}

	return rdb
}
