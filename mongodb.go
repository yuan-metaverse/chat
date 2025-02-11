package chat

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
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

func main() {
	// MongoDB 链接 URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到 MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to MongoDB!")

	// 获取消息集合
	collection := client.Database("chat").Collection("messages")

	// 创建消息
	message := Message{
		SenderID:    "user123",
		ReceiverID:  "user456",
		GroupID:     "",
		MessageType: "text",
		Content:     "hello world",
		MediaURL:    "",
		Timestamp:   time.Now(),
	}

	// 插入消息到数据库
	_, err = collection.InsertOne(context.Background(), message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully inserted message!")
}
