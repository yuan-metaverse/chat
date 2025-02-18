package service

import (
	"chat/internal/storage"
	"context"
	"time"
)

type Message struct {
	SenderID    string    `bson:"sender_id"`
	ReceiverID  string    `bson:"receiver_id"`
	GroupID     string    `bson:"group_id"`
	MessageType string    `bson:"message_type"`
	Content     string    `bson:"content"`
	MediaURL    string    `bson:"media_url"`
	Timestamp   time.Time `bson:"timestamp"`
}

// SaveMessage 保存消息到 MongoDB
func SaveMessage(msg []byte) error {
	message := Message{
		SenderID:    "user123", // 假设一个默认用户ID
		ReceiverID:  "user456",
		GroupID:     "",
		MessageType: "text",
		Content:     string(msg),
		MediaURL:    "",
		Timestamp:   time.Now(),
	}

	_, err := storage.GetCollection("messages").InsertOne(context.Background(), message)
	return err
}
