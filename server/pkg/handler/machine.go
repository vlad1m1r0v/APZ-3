package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

func(h *Handler) getAllMachines(c * gin.Context) {
	machines, _ := h.services.MachineService.GetAllMachines()
	c.JSON(http.StatusOK, machines)
}