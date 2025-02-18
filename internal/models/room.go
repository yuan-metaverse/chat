package models

type Room struct {
	ID           string   `json:"id" bson:"id"`                     // 房间 ID
	Name         string   `json:"name" bson:"name"`                 // 房间名称
	Type         string   `json:"type" bson:"type"`                 // 房间类型(public/private)
	Creator      string   `json:"creator" bson:"creator"`           // 创建者
	Participants []string `json:"participants" bson:"participants"` // 房间成员
	Admins       []string `json:"admins" bson:"admins"`             // 房间管理员
	CreatedAt    int64    `json:"createdAt" bson:"createdAt"`       // 创建时间
}
