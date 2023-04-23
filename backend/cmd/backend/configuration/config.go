package configuration

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Server  struct {
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   int    `yaml:"port" env-default:"5555"`
	} `yaml:"server"`
	Storage StorageConfig `yaml:"storage"`
	Swagger SwaggerConfig `yaml:"swagger"`
}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SwaggerConfig struct {
	Title    string   `json:"title"`
	Version  string   `json:"version"`
	BasePath string   `json:"base-path"`
	Host     string   `json:"host"`
	Schemes  []string `json:"schemes"`
}

var instance *Config
var err error
var once sync.Once

func GetConfig(logger *logrus.Logger) (*Config, error) {
	once.Do(func() {
		logger.Info("read application configuration")
		instance = &Config{}
		err = cleanenv.ReadConfig("cmd/backend/config_local.yaml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	logger.Info(instance)
	return instance, err
}
