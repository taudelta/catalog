package config

import (
	"github.com/spf13/viper"
)

type LoggerConfig struct {
	Level            string                 `json:"level"`
	Encoding         string                 `json:"encoding"`
	OutputPaths      []string               `json:"outputPaths"`
	ErrorOutputPaths []string               `json:"errorOutputPaths"`
	InitialFields    map[string]interface{} `json:"initialFields"`
	EncoderConfig    struct {
		MessageKey   string `json:"messageKey"`
		LevelKey     string `json:"levelKey"`
		LevelEncoder string `json:"levelEncoder"`
	} `json:"encoderConfig"`
}

type DatabaseConfig struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Dbname string `json:"dbname"`
}

type Config struct {
	Logger   LoggerConfig   `json:"logger"`
	Database DatabaseConfig `json:"database"`
}

func Init() (*Config, error) {

	cfg := viper.New()

	cfg.AutomaticEnv()
	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath(".")
	cfg.AddConfigPath("./config")

	var fileConfig Config

	if err := cfg.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := cfg.Unmarshal(&fileConfig); err != nil {
		return nil, err
	}

	return &fileConfig, nil
}
