package config

import (
	"strings"

	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
	"github.com/spf13/viper"
)

type Configuration struct {
	App AppConfiguration
	DB  DatabaseConfiguration
}

type AppConfiguration struct {
	Port string
}

type DatabaseConfiguration struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     uint
	DBName   string
	SSLMode  string
}

func LoadConfig(path string) Configuration {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.BindEnv("db.password", "DB_PASSWORD"); err != nil {
		logger.Fatal("failed to bind environment variable: ", err)
	}

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("failed to load configuration: ", err)
	}

	var configuration Configuration
	if err := viper.Unmarshal(&configuration); err != nil {
		logger.Fatal("failed to unmarshal configuration: ", err)
	}
	return configuration
}
