package configuration

import "github.com/spf13/viper"

type Config struct {
}

func InitConfig() error {
	viper.SetConfigName("config_local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("cmd/backend/configuration/")
	return viper.ReadInConfig()
}
