package models

type Message struct {
	RoomID    string `json:"room_id" bson:"room_id"`     // 房间 ID
	Sender    string `json:"sender" bson:"sender"`       // 发送者
	Content   string `json:"content" bson:"content"`     // 消息内容
	Timestamp int64  `json:"timestamp" bson:"timestamp"` // 消息时间戳
}
