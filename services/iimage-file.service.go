package services

import (
	"context"

	"github.com/celmysql-api/dto"
	"github.com/celmysql-api/mapping"
)

type IImageFileService interface {
	Create(ctx context.Context, request dto.CreateImageFileDto) mapping.ImageFileVm
	Update(ctx context.Context, request dto.UpdateImageFileDto, oid string) mapping.ImageFileVm
	Delete(ctx context.Context, oid string)
	FindById(ctx context.Context, oid string) mapping.ImageFileVm
	Find(ctx context.Context, criteria string) []mapping.ImageFileVm
}
