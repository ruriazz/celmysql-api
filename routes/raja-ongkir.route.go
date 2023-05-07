package routes

import (
	"github.com/celmysql-api/controller"
	"github.com/gin-gonic/gin"
)

func RajaOngkirRouter(rajaOngkirController controller.IRajaOngkirController, group *gin.RouterGroup) *gin.RouterGroup {

	group.POST("/", rajaOngkirController.Find)
	// group.POST("/q", rajaOngkirController.Find)
	// group.GET("/:oid", rajaOngkirController.FindById)
	// group.POST("/create", rajaOngkirController.Create)
	// group.PUT("/:oid", rajaOngkirController.Update)
	// group.PUT("/delete/:oid", rajaOngkirController.Delete)

	return group
}
