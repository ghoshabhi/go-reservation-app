package storage

import "goreservationapp/pkg/storage/models"

// Storage interface defines the methods available to interact with the storage system
type Storage interface {
	// Room
	AddNewRoom(newRoom *models.Room) error
	GetAllRooms() ([]*models.Room, error)
	GetRoomByRoomID(roomID string) (*models.Room, error)
	// RoomType
	AddNewRoomType(newRoomType *models.RoomType) error
	GetAllRoomTypes() ([]*models.RoomType, error)
}

// storage implements the Storage interface
type storage struct {
	rooms     []*models.Room
	roomTypes []*models.RoomType
}

// NewStorage configures new storage
func NewStorage() Storage {
	var rooms = make([]*models.Room, 0)
	var roomTypes = make([]*models.RoomType, 0)

	return &storage{
		rooms:     rooms,
		roomTypes: roomTypes,
	}
}
