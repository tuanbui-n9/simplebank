package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource            string        `mapstructure:"DB_SOURCE" required:"true"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS" required:"true"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMENTRIC_KEY" required:"true"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION" required:"true"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
