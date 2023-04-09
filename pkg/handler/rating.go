package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopper"
	"strconv"
)

func (h *Handler) addNewRate(c *gin.Context) {
	var rate shopper.Rate
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&rate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rateId, err := h.services.CreateRate(id, rate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idJson{
		Id: rateId,
	})
}
