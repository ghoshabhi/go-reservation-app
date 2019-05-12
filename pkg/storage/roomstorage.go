package storage

import (
	"goreservationapp/pkg/storage/models"
)

func (s *storage) GetAllRooms() ([]*models.Room, error) {
	rooms := s.rooms
	return rooms, nil
}

func (s *storage) AddNewRoom(newRoom *models.Room) error {
	rooms := s.rooms
	rooms = append(rooms, newRoom)
	s.rooms = rooms
	return nil
}
