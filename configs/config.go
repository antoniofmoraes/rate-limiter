package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	DBHost                string `mapstructure:"DB_HOST"`
	DBPort                string `mapstructure:"DB_PORT"`
	DBPassword            string `mapstructure:"DB_PASSWORD"`
	RateLimiterTimeout    int64  `mapstructure:"RATE_LIMITER_TIMEOUT"`
	RateLimiterTokenLimit int32  `mapstructure:"RATE_LIMITER_TOKEN_LIMIT"`
	RateLimiterIpLimit    int32  `mapstructure:"RATE_LIMITER_IP_LIMIT"`
}

func LoadConfig(path string, file string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("api_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(file)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
