package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type status struct {
	Ok bool `json:"ok"`
}

func newErrorResponse(c *gin.Context, statusCode int, error string) {
	logrus.Errorf(error)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{
		Error: error,
	})
}
