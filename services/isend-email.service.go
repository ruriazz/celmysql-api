package services

import (
	"context"
)

type ISendEmailService interface {
	SendEmail(ctx context.Context)
}
