package routes

import (
	"github.com/celmysql-api/controller"
	"github.com/gin-gonic/gin"
)

func ImageFileRouter(imageFileController controller.IImageFileController, group *gin.RouterGroup) *gin.RouterGroup {

	group.GET("/", imageFileController.Find)
	group.POST("/q", imageFileController.Find)
	group.GET("/:oid", imageFileController.FindById)
	group.POST("/upload", imageFileController.Create)
	group.PUT("/:oid", imageFileController.Update)
	group.PUT("/delete/:oid", imageFileController.Delete)

	return group
}
