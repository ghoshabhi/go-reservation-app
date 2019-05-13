package list

import (
	"goreservationapp/pkg/storage/models"
)

// RoomsListHelper defines the interface to perform operations on a list
type RoomsListHelper interface {
	Filter(f func(*models.Room) bool) []*models.Room
	// .. Map, Any, All, Includes, Index
}

// RoomsListHelper implements the Helper interface
type roomsHelper struct {
	data []*models.Room
}

// GetRoomsHelper helper
func GetRoomsHelper(data []*models.Room) RoomsListHelper {
	return &roomsHelper{
		data: data,
	}
}

// Filter implements the Helper.Filter function
func (rh *roomsHelper) Filter(f func(room *models.Room) bool) []*models.Room {
	roomsSlice := make([]*models.Room, 0)
	for _, room := range rh.data {
		if f(room) {
			roomsSlice = append(roomsSlice, room)
		}
	}
	return roomsSlice
}
