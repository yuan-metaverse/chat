package models

type User struct {
	ID        string   `json:"id" bson:"id"`                 // 用户 ID
	Name      string   `json:"name" bson:"name"`             // 用户名
	Status    string   `json:"status" bson:"status"`         // 当前状态(如在线、离线等)
	RoomIDs   []string `json:"room_ids" bson:"room_ids"`     // 用户参与的房间 ID 列表
	CreatedAt int64    `json:"created_at" bson:"created_at"` // 创建时间
}
