package server

import (
	"goreservationapp/pkg/server/handlers"
	"goreservationapp/pkg/storage"

	"github.com/gin-gonic/gin"
)

// Server defines the Server API interface
type Server interface {
	Start()
}

// server implements the Server interface
type server struct {
	storage storage.Storage
	engine  *gin.Engine
}

// NewServer instantiates a new instance of server
func NewServer() Server {
	router := gin.Default()

	s := server{
		storage: storage.NewStorage(),
		engine:  router,
	}

	apiMW := router.Group("/api")

	apiMW.GET("/rooms", handlers.GetAllRooms(s.storage))
	apiMW.POST("/rooms", handlers.AddNewRoom(s.storage))

	apiMW.GET("/roomtypes", handlers.GetAllRoomTypes(s.storage))
	apiMW.POST("/roomtypes", handlers.AddNewRoomType(s.storage))

	return &s
}

func (s *server) Start() {
	s.engine.Run()
}
