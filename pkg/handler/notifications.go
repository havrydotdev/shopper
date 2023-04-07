package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopper"
	"time"
)

func (h *Handler) addNewNotification(c *gin.Context) {
	isAdmin, err := getAdminCtx(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if !isAdmin {
		newErrorResponse(c, http.StatusUnauthorized, "you don`t have permission to perform this action")
		return
	}

	var notification shopper.Notification
	if err := c.BindJSON(&notification); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	notification.UserId = userId
	notification.CreatedAt = time.Now()

	id, err := h.services.CreateNotification(notification)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
