package utils

import (
    "github.com/gin-gonic/gin"
)

type APIResponse struct {
    Error   bool        `json:"error"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}

func Respond(ctx *gin.Context, status int, error bool, message string, data interface{}) {
    ctx.JSON(status, APIResponse{
        Error:   error,
        Message: message,
        Data:    data,
    })
}

func RespondError(ctx *gin.Context, status int, message string) {
    Respond(ctx, status, true, message, nil)
}

func RespondSuccess(ctx *gin.Context, status int, message string, data interface{}) {
    Respond(ctx, status, false, message, data)
}
