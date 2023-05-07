package controller

import (
	"net/http"

	"github.com/celmysql-api/common"
	"github.com/celmysql-api/services"
	"github.com/gin-gonic/gin"
)

type RajaOngkirController struct {
	RajaOngkirService services.IRajaOngkirService
}

func NewRajaOngkirController(rajaOngkirService services.IRajaOngkirService) IRajaOngkirController {
	return &RajaOngkirController{
		RajaOngkirService: rajaOngkirService,
	}
}

// // Create RajaOngkir
// // @Summary  create rajaOngkir
// // @Description create new rajaOngkir
// // @Tags rajaOngkirs
// // @Accept  json
// // @Produce  json
// // @Success 200 {string} string	"ok"
// // @Security ApiKeyAuth
// // @Param   rajaOngkir body  dto.CreateRajaOngkirDto true "rajaOngkir"
// // @Router /rajaOngkir/create [post]
// func (controller *RajaOngkirController) Create(c *gin.Context) {

// 	var createRajaOngkirDto dto.CreateRajaOngkirDto
// 	err := c.ShouldBindJSON(&createRajaOngkirDto)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
// 		return
// 	}

// 	rajaOngkirResponse := controller.RajaOngkirService.Create(c, createRajaOngkirDto)
// 	response := common.ResponseOk(rajaOngkirResponse, 1)

// 	c.JSON(http.StatusOK, response)
// }

// // Update RajaOngkir
// // @Summary  update rajaOngkir by id
// // @Description update existing rajaOngkir
// // @Tags rajaOngkirs
// // @Accept  json
// // @Produce  json
// // @Success 200 {string} string	"ok"
// // @Security ApiKeyAuth
// // @Param id path string true  "id"
// // @Param   rajaOngkir body  dto.UpdateRajaOngkirDto true "rajaOngkir"
// // @Router /rajaOngkir/{id} [put]
// func (controller *RajaOngkirController) Update(c *gin.Context) {
// 	id := c.Param("id")
// 	rajaOngkirUpdateRequest := dto.UpdateRajaOngkirDto{}
// 	err := c.ShouldBindJSON(&rajaOngkirUpdateRequest)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
// 		return
// 	}

// 	rajaOngkirUpdateRequest.Id = id

// 	rajaOngkirResponse := controller.RajaOngkirService.Update(c, rajaOngkirUpdateRequest, id)
// 	response := common.ResponseOk(rajaOngkirResponse, 1)

// 	c.JSON(http.StatusOK, response)
// }

// // Delete RajaOngkir
// // @Summary  delete rajaOngkir
// // @Description delete rajaOngkir by id
// // @Tags rajaOngkirs
// // @Accept  json
// // @Produce  json
// // @Success 200 {string} string	"ok"
// // @Security ApiKeyAuth
// // @Param id path string true  "id"
// // @Router /rajaOngkir/delete/{id} [delete]
// func (controller *RajaOngkirController) Delete(c *gin.Context) {
// 	id := c.Param("id")

// 	controller.RajaOngkirService.Delete(c, id)
// 	response := common.ResponseDeleteOk(id)

// 	c.JSON(http.StatusOK, response)
// }

// Get All Raja Ongkir
// @Summary List existing rajaOngkirs
// @Description Get all the existing rajaOngkirs
// @Tags rajaOngkirs
// @Accept  json
// @Produce  json
// @Success 200 {object} common.DefaultResponse	"ok"
// @Security ApiKeyAuth
// @Param   rajaOngkirDto body  dto.CreateRajaOngkirDto true "rajaOngkir"
// @Param   pageIndex    query     int  false  "page Index"  minimum(1)    maximum(10)
// @Param   pageSize    query     int  false  "page Size"  minimum(1)    maximum(10)
// @Router /raja-ongkir [post]
func (controller *RajaOngkirController) Find(c *gin.Context) {
	criteria := ""
	rajaOngkirs := controller.RajaOngkirService.Find(c, criteria)

	response := common.ResponseOk(rajaOngkirs, 1)
	c.JSON(http.StatusOK, response)
}

// // Find RajaOngkir By Id
// // @Summary  find rajaOngkir by id
// // @Description update new rajaOngkir
// // @Tags rajaOngkirs
// // @Accept  json
// // @Produce  json
// // @Success 200 {object} common.DefaultResponse	"ok"
// // @Security ApiKeyAuth
// // @Param id path string true  "id"
// // @Router /rajaOngkir/{id} [get]
// func (controller *RajaOngkirController) FindById(c *gin.Context) {
// 	id := c.Param("id")
// 	rajaOngkir := controller.RajaOngkirService.FindById(c, id)
// 	response := common.ResponseOk(rajaOngkir, 1)
// 	c.JSON(http.StatusOK, response)
// }
