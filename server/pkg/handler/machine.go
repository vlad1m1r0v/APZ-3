package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

func(h *Handler) getAllMachines(c * gin.Context) {
	machines, err := h.services.MachineService.GetAllMachines()
	if err != nil {
		c.String(http.StatusInternalServerError, "bad request: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, machines)
}