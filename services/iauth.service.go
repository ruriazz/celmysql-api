package services

import (
	"context"

	"github.com/celmysql-api/dto"
	"github.com/celmysql-api/entity"
	"github.com/celmysql-api/mapping"
)

type IAuthService interface {
	Login(ctx context.Context, criteria string) entity.PermissionPolicyUser
	Register(ctx context.Context, request dto.PayloadRegister) mapping.PermissionPolicyUserVm
}
