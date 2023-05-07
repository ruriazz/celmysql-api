package services

import (
	"context"

	"github.com/celmysql-api/dto"
	"github.com/celmysql-api/mapping"
)

type IRajaOngkirService interface {
	Create(ctx context.Context, request dto.CreateRajaOngkirDto) mapping.RajaOngkirVm
	// Update(ctx context.Context, request dto.UpdateRajaOngkirDto, id string) mapping.RajaOngkirVm
	// Delete(ctx context.Context, id string)
	// FindById(ctx context.Context, id string) mapping.RajaOngkirVm
	Find(ctx context.Context, criteria string) []mapping.RajaOngkirVm
}
