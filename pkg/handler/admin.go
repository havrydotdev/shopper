package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type idJson struct {
	Id int `json:"id"`
}

func (h *Handler) moderateItem(c *gin.Context) {
	isAdmin, err := getAdminCtx(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !isAdmin {
		newErrorResponse(c, http.StatusUnauthorized, "you can not perform this action!")
		return
	}

	var id idJson

	if err = c.BindJSON(&id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.ModerateItem(id.Id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status{
		Ok: true,
	})
}

func (h *Handler) moderateCompany(c *gin.Context) {
	isAdmin, err := getAdminCtx(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !isAdmin {
		newErrorResponse(c, http.StatusUnauthorized, "you can not perform this action!")
		return
	}

	var id idJson
	if err = c.BindJSON(&id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.ModerateCompany(id.Id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status{
		Ok: true,
	})
}

func (h *Handler) getNotVerifiedItems(c *gin.Context) {
	items, err := h.services.GetItems(false)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getNotVerifiedCompanies(c *gin.Context) {
	items, err := h.services.GetCompanies(false)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
