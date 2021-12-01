package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
type idJSON struct {
	MachineId int `json:"machineId"`
}

func(h *Handler) connectToMachine(c * gin.Context) {
	var body idJSON
	idString := c.Param("id")
	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "internal error: %s", err.Error())
		return
	}
	machineId := body.MachineId
	id,_ := strconv.Atoi(idString)
	if err := h.services.DiskService.ConnectToMachine(id, machineId); err != nil {
		c.String(http.StatusInternalServerError, "bad request: %s", err.Error())
		return
	}
	c.String(http.StatusOK, "successful disk request")
}