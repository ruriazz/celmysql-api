package routes

import (
	"github.com/celmysql-api/controller"
	"github.com/gin-gonic/gin"
)

func AuthRouter(authController controller.IAuthController, group *gin.RouterGroup) *gin.RouterGroup {
	group.POST("/login", authController.Login)
	group.POST("/register", authController.Register)
	return group
}
