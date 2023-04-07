package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopper"
)

func (h *Handler) addNewDiscount(c *gin.Context) {
	isAdmin, err := getAdminCtx(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if !isAdmin {
		newErrorResponse(c, http.StatusUnauthorized, "you don`t have permission to perform this action!")
		return
	}

	var discount shopper.Discount
	if err := c.BindJSON(&discount); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateDiscount(discount)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idJson{
		Id: id,
	})
}
