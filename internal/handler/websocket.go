package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"sync"
)

// WebSocket 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源
	},
}

// WebSocketManager 连接管理器, 用于管理所有 WebSocket 连接
type WebSocketManager struct {
	clients  map[string]*websocket.Conn            // 用户ID -> WebSocket 连接
	sessions map[string]map[string]bool            // 用户ID -> 私聊用户ID/房间ID -> true
	rooms    map[string]map[string]*websocket.Conn // 房间ID -> 用户ID -> WebSocket 连接
	mu       sync.Mutex                            // 并发安全
}

var manager = &WebSocketManager{
	clients:  make(map[string]*websocket.Conn),
	sessions: make(map[string]map[string]bool),
	rooms:    make(map[string]map[string]*websocket.Conn),
}

// HandleChatConnection 处理 WebSocket 连接
func HandleChatConnection(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		http.Error(w, "UserID required", http.StatusBadRequest)
		return
	}

	// 升级 HTTP 请求为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	// 将新的连接添加到 WebSocket 管理器中
	manager.mu.Lock()
	manager.clients[userID] = conn
	manager.mu.Unlock()
	log.Printf("User %s connected", userID)

	// 监听并处理客户端发送的消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}

		manager.HandleMessage(userID, message)
		// 保存消息
		// err = message.SaveMessage(message)
	}

	// 断开连接, 从客户端列表中移除
	manager.mu.Lock()
	delete(manager.clients, userID)
	manager.mu.Unlock()
	log.Printf("User %s disconnected", userID)
}

// HandleMessage 处理消息并发送
func (manager *WebSocketManager) HandleMessage(userID string, message []byte) {
	// 私聊, message 格式为 "/to=userID:content"
	// 群聊, message 格式为 "/room=roomID:content"
	if strings.HasPrefix(string(message), "/to=") {
		var receiverID, content string
		n, err := fmt.Sscanf(string(message), "/to=%s:%s", &receiverID, &content)
		if err != nil || n != 2 {
			log.Printf("Invalid private message format: %s", message)
			return
		}
		manager.SendPrivateMessage(userID, receiverID, message)
	} else if strings.HasPrefix(string(message), "/room=") {
		var roomID, content string
		n, err := fmt.Sscanf(string(message), "/room=%s:%s", &roomID, &content)
		if err != nil || n != 2 {
			log.Printf("Invalid room message format: %s", message)
			return
		}
		if roomID != "default" {
			manager.JoinRoom(userID, roomID)
		}
		manager.SendRoomMessage(userID, roomID, message)
	}
}

// SendPrivateMessage 发送消息给指定用户
func (manager *WebSocketManager) SendPrivateMessage(senderID, receiverID string, message []byte) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	conn, exists := manager.clients[receiverID]
	if !exists {
		log.Printf("User %s not found for private message.", receiverID)
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("Error sending private message to", receiverID, err)
		conn.Close()
		delete(manager.clients, receiverID) // 移除无法发送消息的连接
	}
}

// SendRoomMessage 发送消息给群聊房间
func (manager *WebSocketManager) SendRoomMessage(senderID, roomID string, message []byte) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if _, exists := manager.rooms[roomID]; !exists {
		manager.rooms[roomID] = make(map[string]*websocket.Conn)
		manager.rooms[roomID][senderID] = manager.clients[senderID]
	}

	for userID, conn := range manager.rooms[roomID] {
		if userID == senderID {
			continue
		}
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error sending room message to:", userID, err)
			conn.Close()
			delete(manager.rooms[roomID], userID) // 移除无法发送消息的连接
		}

	}
}

// JoinRoom 加入群聊房间
func (manager *WebSocketManager) JoinRoom(userID, roomID string) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if _, exists := manager.rooms[roomID]; !exists {
		manager.rooms[roomID] = make(map[string]*websocket.Conn)
	}

	conn := manager.clients[userID]
	manager.rooms[roomID][userID] = conn
}

// LeaveRoom 离开群聊房间
func (manager *WebSocketManager) LeaveRoom(userID, roomID string) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if room, exists := manager.rooms[roomID]; exists {
		delete(room, userID)
	}
}
