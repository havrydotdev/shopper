package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
	companyCtx = "companyId"
	adminCtx   = "isAdmin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect auth header")
		return
	}

	if headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect auth header")
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler) adminIdentity(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := h.services.GetUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if user.Email == "admin" && user.Username == "admin" {
		c.Set(adminCtx, true)
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

func getAdminCtx(c *gin.Context) (bool, error) {
	id, ok := c.Get(adminCtx)
	if !ok {
		return false, errors.New("admin id not found")
	}

	isAdmin, ok := id.(bool)
	if !ok {
		return false, errors.New("admin id not found")
	}

	return isAdmin, nil
}
