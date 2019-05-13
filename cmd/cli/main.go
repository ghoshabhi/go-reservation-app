package main

import (
	"fmt"
	"goreservationapp/pkg/helper/list"
	"goreservationapp/pkg/storage/models"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

func main() {
	rooms := make([]*models.Room, 0)

	for i := 0; i < 3; i++ {
		rooms = append(rooms, &models.Room{
			ID:         strconv.Itoa(i + 1),
			Name:       fmt.Sprintf("Foo Room - %v", strconv.Itoa(i+1)),
			NickName:   fmt.Sprintf("Foo-NickName - %v", strconv.Itoa(i+1)),
			RoomTypeID: uuid.NewV4().String(),
			CreatedAt:  time.Now(),
			CreatedBy:  i + 1,
		})
	}

	for _, v := range rooms {
		fmt.Printf("room: %+v\n\n", *v)
	}

	roomHelper := list.GetRoomsHelper(rooms)
	filtered := roomHelper.Filter(func(room *models.Room) bool {
		return room.ID == "1"
	})

	fmt.Printf("\n\nfiltered: %+v", *filtered[0])
}
