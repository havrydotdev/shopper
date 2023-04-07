package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopper"
)

func (h *Handler) updateUserBalance(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var balance shopper.UpdateUserBalance
	if err := c.BindJSON(&balance); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateUserBalance(userId, balance.Value)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status{
		Ok: true,
	})
}

func (h *Handler) getUserHistory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	items, err := h.services.GetUserHistory(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) returnItem(c *gin.Context) {
	//userId, err := getUserId(c)
	//if err != nil {
	//	newErrorResponse(c, http.StatusUnauthorized, err.Error())
	//	return
	//}

}

func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.services.DeleteUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status{
		Ok: true,
	})
}

func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input shopper.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateUser(userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status{
		Ok: true,
	})
}

func (h *Handler) getUserNotifications(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	notifications, err := h.services.User.GetUserNotifications(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, notifications)
}
