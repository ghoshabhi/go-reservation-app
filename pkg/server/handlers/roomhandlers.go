package handlers

import (
	"fmt"
	"goreservationapp/pkg/storage"
	"goreservationapp/pkg/storage/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllRooms is a handler for the GET /rooms route
func GetAllRooms(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rooms []*models.Room
		var err error

		rooms, err = storage.GetAllRooms()
		if err != nil {
			fmt.Println("errored getting all rooms")
		}
		c.JSON(http.StatusOK, gin.H{
			"rooms": rooms,
		})
	}
}

// AddNewRoom is a handler for the POST /rooms route
func AddNewRoom(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newRoom *models.Room
		var newRoomData *models.NewRoomData
		var createdBy int
		var err error

		createdBy, err = strconv.Atoi(c.Request.Header.Get("created-by"))
		if err != nil || createdBy == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "created-by missing",
			})
			return
		}

		err = c.ShouldBindJSON(&newRoomData)
		if err != nil {
			fmt.Printf("Error binding JSON %v", err)
		}
		// fmt.Printf("newRoomData: %+v\n", newRoomData)

		newRoom, err = models.NewRoom(newRoomData, createdBy)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		err = storage.AddNewRoom(newRoom)
		if err != nil {
			fmt.Printf("Error adding new room: %v", err)
		}

		c.Status(http.StatusNoContent)
	}
}
