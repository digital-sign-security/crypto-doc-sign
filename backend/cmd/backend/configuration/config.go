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
	Storage   StorageConfig `yaml:"storage"`
	Swagger   SwaggerConfig `yaml:"swagger"`
	Generator KeyGenerator  `yaml:"generator"`
}

type StorageConfig struct {
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	Database           string `yaml:"database"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	ConnectionAttempts int    `yaml:"connection_attempts"`
}

type SwaggerConfig struct {
	Title    string   `yaml:"title"`
	Version  string   `yaml:"version"`
	BasePath string   `yaml:"base-path"`
	Host     string   `yaml:"host"`
	Schemes  []string `yaml:"schemes"`
}

type KeyGenerator struct {
	MasterKey string `yaml:"master_key"`
	BitSize   int    `yaml:"bit_size"`
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
