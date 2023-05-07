package repository

import (
	"context"
	"database/sql"

	"github.com/celmysql-api/entity"
)

type IImageFileRepository interface {
	Save(ctx context.Context, tx *sql.Tx, imageFile entity.ImageFile) entity.ImageFile
	Update(ctx context.Context, tx *sql.Tx, imageFile entity.ImageFile) entity.ImageFile
	Delete(ctx context.Context, tx *sql.Tx, imageFile entity.ImageFile)
	FindById(ctx context.Context, tx *sql.Tx, oid string) (entity.ImageFile, error)
	Find(ctx context.Context, tx *sql.Tx, criteria string) []entity.ImageFile
}
