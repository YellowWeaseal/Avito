package handler

import (
	"awesomeProject1/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	segments := router.Group("/segments")
	{
		segments.POST("/", h.createSegment)
		segments.GET("/", h.getAllSegments)
		segments.GET("/:id", h.getSegmentById)
		segments.PUT("/:id", h.updateSegment)
		segments.DELETE("/:id", h.deleteSegment)

	}
	users := router.Group(":id/user")
	{
		users.POST("/", h.createUser)
		users.GET("/", h.getAllUsers)
		users.GET("/:id", h.getUserById)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
	}
	return router
}
