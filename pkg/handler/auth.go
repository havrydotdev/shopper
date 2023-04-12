package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopper"
)

// @Summary SignUp
// @Tags auth
// @Description Sign-up endpoint
// @ID create-account
// @Accept json
// @Produce json
// @Param input body shopper.SignUpInput true "User`s info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input shopper.SignUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("invalid input body: %s", err.Error()))
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description Sign-in endpoint
// @ID login
// @Accept  json
// @Produce  json
// @Param input body shopper.SignInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input shopper.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
