package controller

import (
	"github.com/gin-gonic/gin"
)

type IRajaOngkirController interface {
	// Create(ctx *gin.Context)
	// Update(ctx *gin.Context)
	// Delete(ctx *gin.Context)
	// FindById(ctx *gin.Context)
	Find(ctx *gin.Context)
}
