package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/celmysql-api/common"
	"github.com/celmysql-api/entity"
	"github.com/feiin/ploto"
	"github.com/google/uuid"
)

type ImageFileRepository struct {
}

func NewImageFileRepository() IImageFileRepository {
	return &ImageFileRepository{}
}

func (repository *ImageFileRepository) Save(ctx context.Context, tx *sql.Tx, imageFile entity.ImageFile) entity.ImageFile {

	id := uuid.New()
	SQL := "insert into imageFile(description, optimisticLockField, gCRecord, deleted, userInserted, insertedDate,oid,fileUrl,fileName) values (?,?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, imageFile.Description, imageFile.OptimisticLockField, imageFile.GCRecord, imageFile.Deleted, imageFile.UserInserted, imageFile.InsertedDate, id, imageFile.FileUrl, imageFile.FileName)
	common.PanicIfError(err)
	imageFile.Oid = id.String()
	return imageFile
}

func (repository *ImageFileRepository) Update(ctx context.Context, tx *sql.Tx, imageFile entity.ImageFile) entity.ImageFile {

	SQL := "update imageFile set description=?, optimisticLockField=?, gCRecord=?, deleted=?, lastUserId=?, lastUpdate=? ,fileUrl = ?,fileName = ? where oid = ?"
	_, err := tx.ExecContext(ctx, SQL, imageFile.Description, imageFile.OptimisticLockField, imageFile.GCRecord, imageFile.Deleted, imageFile.LastUserId, imageFile.LastUpdate, imageFile.FileUrl, imageFile.FileName, imageFile.Oid)
	common.PanicIfError(err)

	return imageFile
}

func (repository *ImageFileRepository) Delete(ctx context.Context, tx *sql.Tx, imageFile entity.ImageFile) {
	SQL := "delete from imageFile where oid = ?"
	_, err := tx.ExecContext(ctx, SQL, imageFile.Oid)
	common.PanicIfError(err)
}

func (repository *ImageFileRepository) FindById(ctx context.Context, tx *sql.Tx, oid string) (entity.ImageFile, error) {
	SQL := "select description, optimisticLockField, gCRecord, deleted, userInserted, insertedDate, lastUserId, lastUpdate, oid, fileUrl,fileName from imageFile where oid = ?"
	rows, err := tx.QueryContext(ctx, SQL, oid)
	common.PanicIfError(err)
	defer rows.Close()

	imageFile := entity.ImageFile{}
	if rows.Next() {
		err := ploto.Scan(rows, &imageFile)
		common.PanicIfError(err)
		return imageFile, nil
	} else {
		return imageFile, errors.New("imageFile is not found")
	}
}

func (repository *ImageFileRepository) Find(ctx context.Context, tx *sql.Tx, criteria string) []entity.ImageFile {
	SQL := "select description, optimisticLockField, gCRecord, deleted, userInserted, insertedDate, lastUserId, lastUpdate, oid,fileUrl,fileName from imageFile " + criteria
	rows, err := tx.QueryContext(ctx, SQL)
	common.PanicIfError(err)
	defer rows.Close()

	var imageFiles []entity.ImageFile
	for rows.Next() {
		imageFile := entity.ImageFile{}
		err := ploto.Scan(rows, &imageFile)
		common.PanicIfError(err)
		imageFiles = append(imageFiles, imageFile)
	}
	return imageFiles
}
