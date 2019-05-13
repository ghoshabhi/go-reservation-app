package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"goreservationapp/pkg/storage"
	"goreservationapp/pkg/storage/models"
)

// GetAllRoomTypes is a handler for the GET /roomTypes route
func GetAllRoomTypes(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roomTypes []*models.RoomType
		var err error

		if roomTypes, err = storage.GetAllRoomTypes(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
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

		if createdBy, err = strconv.Atoi(c.Request.Header.Get("created-by")); err != nil || createdBy == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "created-by missing",
			})
			return
		}

		err = c.ShouldBindJSON(&newRoomTypeData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if newRoomType, err = models.NewRoomType(newRoomTypeData, createdBy); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// TODO: Check if RoomType with same name exists already or not
		if err = storage.AddNewRoomType(newRoomType); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
