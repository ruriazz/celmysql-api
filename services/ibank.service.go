package services

import (
	"context"

	"github.com/celmysql-api/dto"
	"github.com/celmysql-api/mapping"
)

type IBankService interface {
	Create(ctx context.Context, request dto.CreateBankDto) mapping.BankVm
	Update(ctx context.Context, request dto.UpdateBankDto, oid string) mapping.BankVm
	Delete(ctx context.Context, oid string)
	FindById(ctx context.Context, oid string) mapping.BankVm
	Find(ctx context.Context, criteria string) []mapping.BankVm
}
