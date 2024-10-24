package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBUser string
	DBPass string
	DBPort string
}

func LoadConfig() (Config, error) {
	var config Config

	// 加载 .env 文件
	err := godotenv.Load("server.env")
	if err != nil {
		return config, err
	}

	// 从环境变量中获取配置
	config.DBHost = os.Getenv("DB_HOST")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPass = os.Getenv("DB_PASS")
	config.DBPort = os.Getenv("DB_PORT")

	return config, nil
}
