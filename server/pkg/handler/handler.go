package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vlad1m1r0v/APZ-3/server/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	machines := router.Group("/machines")
	{
		machines.GET("/")
	}

	disks := router.Group("/disks")
	{
		disks.PUT("/:id")
	}

	return router
}
