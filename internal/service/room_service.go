package service

import (
	"chat/internal/models"
	"chat/internal/storage"
)

type RoomService struct {
	storage *storage.MongoDB
}

func (rs *RoomService) CreateRoom(room *models.Room) error {
	return rs.storage.CreateRoom(room)
}

// 更多业务逻辑如：添加成员、删除成员等
