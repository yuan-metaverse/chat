package service

import (
	"chat/internal/storage"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"testing"
)

// 数据库交互
func TestSaveMessage(t *testing.T) {
	// 假设 MongoDB 已经连接并可以操作
	storage.ConnectToMongo()

	// 测试消息保存
	msg := []byte("Test message")
	err := SaveMessage(msg)
	assert.NoError(t, err)

	// 检查 MongoDB 中是否保存了该消息
	var result Message
	err = storage.GetCollection("messages").FindOne(nil, bson.M{"content": "Test message"}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "Test message", result.Content)
}
