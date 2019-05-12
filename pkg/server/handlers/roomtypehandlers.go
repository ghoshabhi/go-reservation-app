package handlers

import (
	"fmt"
	"goreservationapp/pkg/storage"
	"goreservationapp/pkg/storage/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllRoomTypes is a handler for the GET /roomTypes route
func GetAllRoomTypes(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roomTypes []*models.RoomType
		var err error

		roomTypes, err = storage.GetAllRoomTypes()
		if err != nil {
			fmt.Println("errored getting all roomTypes")
		}
		c.JSON(http.StatusOK, gin.H{
			"roomTypes": roomTypes,
		})
	}
}

// AddNewRoomType is a handler for the POST /roomtypes route
func AddNewRoomType(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newRoomType *models.RoomType
		var newRoomTypeData *models.NewRoomTypeData
		var createdBy int
		var err error

		createdBy, err = strconv.Atoi(c.Request.Header.Get("created-by"))
		if err != nil || createdBy == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "created-by missing",
			})
			return
		}

		err = c.ShouldBindJSON(&newRoomTypeData)
		if err != nil {
			fmt.Printf("Error binding JSON %v", err)
		}

		newRoomType, err = models.NewRoomType(newRoomTypeData, createdBy)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		err = storage.AddNewRoomType(newRoomType)
		if err != nil {
			fmt.Printf("Error adding new room: %v", err)
		}

		c.Status(http.StatusNoContent)
	}
}
