package models

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Room represents
type Room struct {
	ID         string    `json:"roomId"`
	Name       string    `json:"roomName"`
	NickName   string    `json:"roomNickName"`
	RoomTypeID string    `json:"roomTypeId"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  int       `json:"createdBy"`
}

// NewRoomData defines the json payload expected to create a new Room
type NewRoomData struct {
	RoomName     string `json:"roomName"`
	RoomTypeID   string `json:"roomTypeID"`
	RoomNickName string `json:"roomNickName"`
}

func isValidNewRoomData(data *NewRoomData) bool {
	if data == nil {
		return false
	}
	if data.RoomName == "" || data.RoomNickName == "" || data.RoomTypeID == "" {
		return false
	}
	return true
}

// NewRoom validates and instantiates a Room type
func NewRoom(data *NewRoomData, createdBy int) (*Room, error) {
	isValidData := isValidNewRoomData(data)
	if !isValidData {
		return nil, fmt.Errorf("Invalid new RoomType data")
	}
	return &Room{
		ID:         uuid.NewV4().String(),
		Name:       data.RoomName,
		RoomTypeID: data.RoomTypeID,
		NickName:   data.RoomNickName,
		CreatedAt:  time.Now(),
		CreatedBy:  createdBy,
	}, nil
}
