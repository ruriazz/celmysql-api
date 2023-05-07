package controller

import (
	"github.com/gin-gonic/gin"
)

type ISendEmailController interface {
	SendEmail(ctx *gin.Context)
}
