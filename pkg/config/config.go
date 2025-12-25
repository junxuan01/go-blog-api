package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	DSN string
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")  // 配置文件名 (不带扩展名)
	viper.SetConfigType("yaml")    // 配置文件类型
	viper.AddConfigPath("configs") // 查找路径

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
