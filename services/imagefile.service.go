package services

import (
	"context"
	"database/sql"

	"github.com/celmysql-api/common"
	"github.com/celmysql-api/dto"
	"github.com/celmysql-api/entity"
	"github.com/celmysql-api/mapping"
	"github.com/celmysql-api/repository"
	"github.com/go-playground/validator/v10"
)

type ImageFileService struct {
	ImageFileRepository repository.IImageFileRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewImageFileService(imageFileRepository repository.IImageFileRepository, DB *sql.DB, validate *validator.Validate) IImageFileService {
	return &ImageFileService{
		ImageFileRepository: imageFileRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *ImageFileService) Create(ctx context.Context, request dto.CreateImageFileDto) mapping.ImageFileVm {
	err := service.Validate.Struct(request)
	common.PanicIfError(err)

	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	imageFile := entity.ImageFile{
		FileName:     &request.FileName,
		FileUrl:      &request.FileUrl,
		UserInserted: &request.UserInserted,
	}

	imageFile = service.ImageFileRepository.Save(ctx, tx, imageFile)

	return mapping.ToImageFileResponse(imageFile)
}

func (service *ImageFileService) Update(ctx context.Context, request dto.UpdateImageFileDto, oid string) mapping.ImageFileVm {
	err := service.Validate.Struct(request)
	common.PanicIfError(err)

	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	imageFile, err := service.ImageFileRepository.FindById(ctx, tx, oid)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	imageFile.FileName = &request.FileName
	imageFile.FileUrl = &request.FileUrl
	imageFile.LastUserId = &request.LastUserId

	imageFile = service.ImageFileRepository.Update(ctx, tx, imageFile)

	return mapping.ToImageFileResponse(imageFile)
}

func (service *ImageFileService) Delete(ctx context.Context, oid string) {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	imageFile, err := service.ImageFileRepository.FindById(ctx, tx, oid)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	service.ImageFileRepository.Delete(ctx, tx, imageFile)
}

func (service *ImageFileService) FindById(ctx context.Context, oid string) mapping.ImageFileVm {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	imageFile, err := service.ImageFileRepository.FindById(ctx, tx, oid)
	if err != nil {
		panic(common.NewNotFoundError(err.Error()))
	}

	return mapping.ToImageFileResponse(imageFile)
}

func (service *ImageFileService) Find(ctx context.Context, criteria string) []mapping.ImageFileVm {
	tx, err := service.DB.Begin()
	common.PanicIfError(err)
	defer common.CommitOrRollback(tx)

	imageFiles := service.ImageFileRepository.Find(ctx, tx, criteria)

	return mapping.ToImageFileResponses(imageFiles)
}
