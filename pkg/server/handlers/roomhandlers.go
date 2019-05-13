package handlers

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"

	"goreservationapp/pkg/storage"
	"goreservationapp/pkg/storage/models"
)

// GetAllRooms is a handler for the GET /rooms route
func GetAllRooms(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rooms []*models.Room
		var err error

		rooms, err = storage.GetAllRooms()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"rooms": rooms,
		})
	}
}

// GetRoomByRoomID fetches a room by roomid
func GetRoomByRoomID(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var room *models.Room
		var roomID string
		var err error

		if roomID = c.Param("roomid"); roomID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "roomid required",
			})
			return
		}

		if room, err = storage.GetRoomByRoomID(roomID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, room)
	}
}

// AddNewRoom is a handler for the POST /rooms route
func AddNewRoom(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newRoom *models.Room
		var newRoomData *models.NewRoomData
		var createdBy int
		var err error

		if createdBy, err = strconv.Atoi(c.Request.Header.Get("created-by")); err != nil || createdBy == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "created-by missing",
			})
			return
		}

		if err = c.ShouldBindJSON(&newRoomData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Printf("newRoomData: %+v\n", newRoomData)

		if newRoom, err = models.NewRoom(newRoomData, createdBy); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// TODO: Check if RoomType exists or not
		if err = storage.AddNewRoom(newRoom); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
