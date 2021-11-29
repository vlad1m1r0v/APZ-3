package handler

import "github.com/gin-gonic/gin"

type Handler struct {}

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
