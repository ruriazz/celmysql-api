package controller

import (
	"net/http"
	"strings"

	"github.com/celmysql-api/common"
	"github.com/celmysql-api/dto"
	"github.com/celmysql-api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ImageFileController struct {
	ImageFileService services.IImageFileService
}

func NewImageFileController(imageFileService services.IImageFileService) IImageFileController {
	return &ImageFileController{
		ImageFileService: imageFileService,
	}
}

// Create ImageFile
// @Summary  create imageFile
// @Description create new imageFile
// @Tags imageFiles
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Param   imageFile body  dto.CreateImageFileDto true "imageFile"
// @Router /image-file/upload [post]
func (controller *ImageFileController) Create(c *gin.Context) {
	file, err := c.FormFile("file")
	mimeType := strings.Split(file.Filename, ".")

	if err != nil {
		panic(err)
	}
	user := c.PostForm("user")
	if err != nil {
		panic(err)
	}
	filePath := "public/" + uuid.New().String() + "." + mimeType[1] // file.Filename

	c.SaveUploadedFile(file, filePath)
	url := "http://localhost:3000/" + filePath
	createImageFileDto := dto.CreateImageFileDto{
		FileName:     file.Filename,
		FileUrl:      url,
		UserInserted: user,
	}

	imageFileResponse := controller.ImageFileService.Create(c, createImageFileDto)
	response := common.ResponseOk(imageFileResponse, 1)

	c.JSON(http.StatusOK, response)
}

// Update ImageFile
// @Summary  update imageFile by oid
// @Description update existing imageFile
// @Tags imageFiles
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Param oid path string true  "oid"
// @Param   imageFile body  dto.UpdateImageFileDto true "imageFile"
// @Router /image-file/{oid} [put]
func (controller *ImageFileController) Update(c *gin.Context) {
	oid := c.Param("oid")
	imageFileUpdateRequest := dto.UpdateImageFileDto{}
	err := c.ShouldBindJSON(&imageFileUpdateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
		return
	}

	imageFileUpdateRequest.Oid = oid

	imageFileResponse := controller.ImageFileService.Update(c, imageFileUpdateRequest, oid)
	response := common.ResponseOk(imageFileResponse, 1)

	c.JSON(http.StatusOK, response)
}

// Delete ImageFile
// @Summary  delete imageFile
// @Description delete imageFile by oid
// @Tags imageFiles
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Param oid path string true  "oid"
// @Router /image-file/delete/{oid} [delete]
func (controller *ImageFileController) Delete(c *gin.Context) {
	oid := c.Param("oid")

	controller.ImageFileService.Delete(c, oid)
	response := common.ResponseDeleteOk(oid)

	c.JSON(http.StatusOK, response)
}

// Find ImageFile By Oid
// @Summary  find imageFile by oid
// @Description update new imageFile
// @Tags imageFiles
// @Accept  json
// @Produce  json
// @Success 200 {object} common.DefaultResponse	"ok"
// @Security ApiKeyAuth
// @Param oid path string true  "oid"
// @Router /image-file/{oid} [get]
func (controller *ImageFileController) FindById(c *gin.Context) {
	oid := c.Param("oid")
	imageFile := controller.ImageFileService.FindById(c, oid)
	response := common.ResponseOk(imageFile, 1)
	c.JSON(http.StatusOK, response)
}

// Get All ImageFile
// @Summary List existing imageFiles
// @Description Get all the existing imageFiles
// @Tags imageFiles
// @Accept  json
// @Produce  json
// @Success 200 {object} common.DefaultResponse	"ok"
// @Security ApiKeyAuth
// @Param   imageFileDto body  dto.CreateImageFileDto true "imageFile"
// @Param   pageIndex    query     int  false  "page Index"  minimum(1)    maximum(10)
// @Param   pageSize    query     int  false  "page Size"  minimum(1)    maximum(10)
// @Router /image-file/q [post]
func (controller *ImageFileController) Find(c *gin.Context) {
	// @Param pageIndex  path int false "pageIndex"
	// @Param pageSize  path int true "pageSize"
	pageIndex := c.Query("pageIndex")
	pageSize := c.Query("pageSize")

	var imageFileDto dto.FilterImageFileDto
	err := c.ShouldBindJSON(&imageFileDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ResponseForm1Forbidden(err.Error()))
		return
	}
	criteria := ""
	criteriaQ := ""
	if len(pageIndex) == 0 {
		pageIndex = "0"
	}
	if len(pageSize) == 0 {
		pageSize = "10"
	}

	criteriaQ += "where fileName Like " + "'%" + imageFileDto.FileName + "%'" + " and " + "fileUrl Like " + "'%" + imageFileDto.FileUrl + "%'"
	criteria += criteriaQ + " order by fileName asc LIMIT " + pageIndex + "," + pageSize

	imageFiles := controller.ImageFileService.Find(c, criteria)
	imageFilesCount := controller.ImageFileService.Find(c, criteriaQ)

	response := common.ResponseOk(imageFiles, len(imageFilesCount))
	c.JSON(http.StatusOK, response)
}
