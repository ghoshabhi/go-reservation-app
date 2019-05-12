package models

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// RoomType type of room
type RoomType struct {
	ID          string    `json:"roomTypeId"`
	Name        string    `json:"roomTypeName"`
	Description string    `json:"roomDescription"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   int       `json:"createdBy"`
}

// NewRoomTypeData defines the parameters to create a new RoomType
type NewRoomTypeData struct {
	RoomTypeName        string `json:"roomTypeName"`
	RoomTypeDescription string `json:"roomDescription"`
}

func isValidNewRoomTypeData(data *NewRoomTypeData) bool {
	if data == nil {
		return false
	}
	if data.RoomTypeName == "" || data.RoomTypeDescription == "" {
		return false
	}
	return true
}

// NewRoomType instantiates a Roomtype
func NewRoomType(data *NewRoomTypeData, createdBy int) (*RoomType, error) {
	isValidData := isValidNewRoomTypeData(data)
	if !isValidData {
		return nil, fmt.Errorf("Invalid new RoomType data")
	}
	return &RoomType{
		ID:          uuid.NewV4().String(),
		Name:        data.RoomTypeName,
		Description: data.RoomTypeDescription,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
	}, nil
}
