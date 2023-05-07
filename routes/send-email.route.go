package routes

import (
	"github.com/celmysql-api/controller"
	"github.com/gin-gonic/gin"
)

func SendEmailRouter(sendEmailController controller.ISendEmailController, group *gin.RouterGroup) *gin.RouterGroup {

	group.GET("/", sendEmailController.SendEmail)
	return group
}
