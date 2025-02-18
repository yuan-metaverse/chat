package main

import (
	"chat/internal/config"
	"chat/internal/handler"
	"chat/internal/storage"
	"log"
	"net/http"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 连接 MongoDB
	storage.ConnectToMongo()

	// 启动 WebSocket 服务器
	http.HandleFunc("/chat", handler.HandleChatConnection)

	// 启动 HTTP 服务
	log.Printf("Server running on port %s...", config.AppConfig.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.ServerPort, nil))
}
