package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config 配置结构体, 用于从 .env 文件中读取配置
type Config struct {
	MongoURI string
	Port     string
}

// AppConfig 全局配置变量
var AppConfig *Config

// LoadConfig 加载 .env 文件并初始化配置
func LoadConfig() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化配置
	AppConfig = &Config{
		MongoURI: os.Getenv("MONGO_URI"),
		Port:     os.Getenv("SERVER_PORT"),
	}

	// 如果配置为空，打印警告
	if AppConfig.MongoURI == "" {
		log.Fatal("Mongo URI not provided in the .env file")
	}
	if AppConfig.Port == "" {
		log.Fatal("Server port not provided in the .env file")
	}
}
