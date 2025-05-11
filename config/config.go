package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App    AppConfig    `yaml:"app"`
		Logger LoggerConfig `yaml:"logger"`
		GRPC   GRPCConfig   `yaml:"grpc"`
		Token  TokenConfig  `yaml:"token"`
	}

	AppConfig struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	LoggerConfig struct {
		Level string `yaml:"level"`
	}

	GRPCConfig struct {
		Port    int `yaml:"port"`
		Timeout int `yaml:"timeout"`
	}

	TokenConfig struct {
		AuthorizationKey string `yaml:"authorization_key"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig("./config/config.yaml", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
