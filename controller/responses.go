package controller

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode uint   `json:"errorCode"`
}

type ListResponse struct {
	Data interface{} `json:"data"`
}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}
