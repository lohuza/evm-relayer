package startup

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func ReadConfig() {
	viper.SetConfigFile("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err)
	}
}
