package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT" required:"true"`
	DBSource             string        `mapstructure:"DB_SOURCE" required:"true"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL" required:"true"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS" required:"true"`
	GrpcServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS" required:"true"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY" required:"true"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION" required:"true"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION" required:"true"`
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
