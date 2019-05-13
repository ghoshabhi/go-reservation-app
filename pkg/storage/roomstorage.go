package storage

import (
	"goreservationapp/pkg/helper/list"
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

func (s *storage) GetRoomByRoomID(roomID string) (*models.Room, error) {
	var room *models.Room

	roomsHelper := list.GetRoomsHelper(s.rooms)

	room = roomsHelper.Filter(func(r *models.Room) bool {
		return r.ID == roomID
	})[0]

	return room, nil
}
