package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopper"
	"strconv"
)

// @Summary AddCompany
// @Security ApiKeyAuth
// @Tags companies
// @Description Create comment endpoint
// @ID addNewCompany
// @Accept  json
// @Produce  json
// @Param input body shopper.Company true "Company body"
// @Success 200 {object} idJson
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/companies [post]
func (h *Handler) addNewCompany(c *gin.Context) {
	val, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.Authorization.GetUser(val)

	if user.CompanyId.Valid {
		newErrorResponse(c, http.StatusConflict, "user is already registered in company")
		return
	}

	var company shopper.Company
	if err := c.BindJSON(&company); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Company.CreateCompany(company, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Set(companyCtx, id)
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) updateCompany(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input shopper.UpdateCompanyInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = input.Validate()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateCompany(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status{
		Ok: true,
	})
}

func (h *Handler) getCompanyById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	company, err := h.services.Company.GetCompanyById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, company)
}
