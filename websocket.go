package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleChatConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		// 读取消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		// 处理消息(例如, 将其存储到数据库, 或广播到其他客户端)
		fmt.Printf("Received: %s\n", msg)

		// 向客户端发送消息(发送一个简单回复)
		err = conn.WriteMessage(websocket.TextMessage, []byte("Message received"))
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/chat", handleChatConnection)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
