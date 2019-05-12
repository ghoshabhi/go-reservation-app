package storage

import (
	"goreservationapp/pkg/storage/models"
)

func (s *storage) GetAllRoomTypes() ([]*models.RoomType, error) {
	roomTypes := s.roomTypes
	return roomTypes, nil
}

func (s *storage) AddNewRoomType(newRoomType *models.RoomType) error {
	roomTypes := s.roomTypes
	roomTypes = append(roomTypes, newRoomType)
	s.roomTypes = roomTypes
	return nil
}
