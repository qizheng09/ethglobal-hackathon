package config

import (
	"time"

	"github.com/spf13/viper"
)

var EnvConfig Config

type Config struct {
	Server ServerConfig
	Mysql  MysqlConfig
}

type ServerConfig struct {
	Port int
}

type MysqlConfig struct {
	Port             int
	Host             string
	Username         string
	Password         string
	Dbname           string
	MaxConn          int
	MaxOpen          int
	MaxLiftTimeInMin time.Duration
}

func InitConfig() error {
	viper.SetConfigType("toml")
	viper.SetConfigName("env")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&EnvConfig); err != nil {
		return err
	}

	return nil
}
