package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error string `json:"error" example:"something went wrong!"`
	Ok    bool   `json:"ok" example:"false"`
}

type status struct {
	Ok bool `json:"ok" example:"true"`
}

func newErrorResponse(c *gin.Context, statusCode int, error string) {
	logrus.Errorf(error)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{
		Error: error,
		Ok:    false,
	})
}
