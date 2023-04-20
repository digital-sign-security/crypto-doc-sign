package configuration

import "github.com/spf13/viper"

type Config struct {
}

func InitConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config_local.yaml")
	return viper.ReadInConfig()
}
