package controller

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode uint   `json:"errorCode"`
}

func sendError(c *gin.Context, code int, msg string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}
