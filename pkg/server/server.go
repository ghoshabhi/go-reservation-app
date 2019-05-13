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
	// storage storage.Storage
	engine  *gin.Engine
	// logger ?
	// env variables?
}

// NewServer instantiates a new instance of server
func NewServer() Server {
	router := gin.Default()
	storage := storage.NewStorage()

	s := server{
		// storage: storage.NewStorage(),
		engine:  router,
	}

	apiMW := router.Group("/api")

	apiMW.GET("/rooms", handlers.GetAllRooms(storage))
	apiMW.GET("/rooms/:roomid", handlers.GetRoomByRoomID(storage))
	apiMW.POST("/rooms", handlers.AddNewRoom(storage))

	apiMW.GET("/roomtypes", handlers.GetAllRoomTypes(storage))
	apiMW.POST("/roomtypes", handlers.AddNewRoomType(storage))

	return &s
}

func (s *server) Start() {
	s.engine.Run()
}
