package handler

import (
	"chat/internal/storage"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

// 成功建立 WebSocket 连接
func TestWebSocketConnection_Success(t *testing.T) {
	// 连接 MongoDB
	storage.ConnectToMongo()

	// 启动 WebSocket 服务器
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/chat", HandleChatConnection)

	go func() { _ = server.ListenAndServe() }()
	defer func(server *http.Server) { _ = server.Close() }(&server)

	// 连接 WebSocket
	url := "ws://localhost:8080/chat"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err)
	defer func(conn *websocket.Conn) { _ = conn.Close() }(conn)

	// 发送消息
	err = conn.WriteMessage(websocket.TextMessage, []byte("hello world"))
	assert.NoError(t, err)

	// 接收响应
	_, message, err := conn.ReadMessage()
	assert.NoError(t, err)
	assert.Equal(t, []byte("Message received"), message)
}

// 广播消息给多个客户端
func TestWebSocket_Broadcast(t *testing.T) {
	// 启动 WebSocket 服务器
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/chat", HandleChatConnection)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	defer server.Close()

	// 连接多个 WebSocket 客户端
	url := "ws://localhost:8080/chat"
	conn1, _, err := websocket.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err)
	defer conn1.Close()

	conn2, _, err := websocket.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err)
	defer conn2.Close()

	// 发送广播消息
	err = conn1.WriteMessage(websocket.TextMessage, []byte("Broadcast Test"))
	assert.NoError(t, err)

	// 接收广播消息
	_, message1, err := conn1.ReadMessage()
	assert.NoError(t, err)
	assert.Equal(t, "Message received", string(message1))

	_, message2, err := conn2.ReadMessage()
	assert.NoError(t, err)
	assert.Equal(t, "Message received", string(message2))
}

// 连接时出错
func TestWebSocket_ConnectionError(t *testing.T) {
	// 启动 WebSocket 服务器
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/chat", HandleChatConnection)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	defer server.Close()

	// 错误的 URL
	url := "ws://localhost:8081/chat" // 错误的端口

	// 尝试连接失败
	_, _, err := websocket.DefaultDialer.Dial(url, nil)
	assert.Error(t, err)
}

// 消息持久化到 MongoDB
func TestMessagePersistence(t *testing.T) {
	// 连接到 MongoDB
	storage.ConnectToMongo()

	// 启动 WebSocket 服务器
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/chat", HandleChatConnection)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	defer server.Close()

	// 连接 WebSocket 客户端
	url := "ws://localhost:8080/chat"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	assert.NoError(t, err)
	defer conn.Close()

	// 发送消息
	err = conn.WriteMessage(websocket.TextMessage, []byte("Persisted message"))
	assert.NoError(t, err)

	// 确保消息已存入数据库
	collection := storage.GetCollection("messages")
	var result map[string]interface{}
	err = collection.FindOne(context.Background(), map[string]interface{}{"content": "Persisted message"}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "Persisted message", result["content"])
}

// 处理多个并发 WebSocket 连接
func TestWebSocket_Concurrency(t *testing.T) {
	// 连接到 MongoDB
	storage.ConnectToMongo()

	// 启动 WebSocket 服务器
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/chat", HandleChatConnection)

	go server.ListenAndServe()
	defer server.Close()

	// 创建多个并发 WebSocket 连接
	for i := 0; i < 100; i++ {
		go func(i int) {
			url := "ws://localhost:8080/chat"
			conn, _, err := websocket.DefaultDialer.Dial(url, nil)
			assert.NoError(t, err)
			defer conn.Close()

			// 向服务器发送消息
			err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Message %d", i)))
			assert.NoError(t, err)

			// 接收响应
			_, message, err := conn.ReadMessage()
			assert.NoError(t, err)
			assert.Equal(t, "Message received", string(message))
		}(i)
	}

	// 给并发连接一点时间完成
	time.Sleep(3 * time.Second)
}
