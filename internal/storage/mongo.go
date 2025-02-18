package storage

import (
	"chat/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDB struct {
	client *mongo.Client
}

func (db *MongoDB) CreateRoom(room *models.Room) error {
	collection := db.client.Database("chat").Collection("rooms")
	_, err := collection.InsertOne(context.TODO(), room)
	return err
}

func (db *MongoDB) GetRoomByID(roomID string) (*models.Room, error) {
	var room models.Room
	collection := db.client.Database("chat").Collection("rooms")
	if err := collection.FindOne(context.TODO(), bson.M{"id": roomID}).Decode(&room); err != nil {
		return nil, err
	}
	return &room, nil
}

// // ConnectToMongo 连接 MongoDB
// func ConnectToMongo() {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	var err error
// 	client, err = mongo.Connect(clientOptions)
// 	if err != nil {
// 		log.Fatal("Error connecting to MongoDB:", err)
// 	}
//
// 	// 检查连接
// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal("Error pinging MongoDB:", err)
// 	}
//
// 	// 选择数据库
// 	db = client.Database("chat")
// 	log.Println("Connected to MongoDB!")
// }
//
// // GetCollection 获取指定的集合
// func GetCollection(collectionName string) *mongo.Collection {
// 	return db.Collection(collectionName)
// }
