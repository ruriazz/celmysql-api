package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/celmysql-api/entity"

	"github.com/celmysql-api/common"
	"github.com/feiin/ploto"
	"github.com/google/uuid"
)

type BankRepository struct {
}

func NewBankRepository() IBankRepository {
	return &BankRepository{}
}

func (repository *BankRepository) Save(ctx context.Context, tx *sql.Tx, bank entity.Bank) entity.Bank {

	id := uuid.New()
	SQL := `insert into bank( userInserted, insertedDate,
		oid,bankCode,bankName) values (?,?,?,?,?)`
	_, err := tx.ExecContext(ctx, SQL, bank.UserInserted, bank.InsertedDate,
		id, bank.BankCode, bank.BankName)
	common.PanicIfError(err)
	bank.Oid = id.String()
	return bank
}

func (repository *BankRepository) Update(ctx context.Context, tx *sql.Tx, bank entity.Bank) entity.Bank {

	SQL := "update bank set bankCode = ?,bankName = ? where oid = ?"
	_, err := tx.ExecContext(ctx, SQL, bank.BankCode, bank.BankName, bank.Oid)
	common.PanicIfError(err)

	return bank
}

func (repository *BankRepository) Delete(ctx context.Context, tx *sql.Tx, bank entity.Bank) {
	SQL := "delete from bank where oid = ?"
	_, err := tx.ExecContext(ctx, SQL, bank.Oid)
	common.PanicIfError(err)
}

func (repository *BankRepository) FindById(ctx context.Context, tx *sql.Tx, oid string) (entity.Bank, error) {
	SQL := "select  oid, bankCode,bankName from bank where oid = ?"
	rows, err := tx.QueryContext(ctx, SQL, oid)
	common.PanicIfError(err)
	defer rows.Close()

	bank := entity.Bank{}
	if rows.Next() {
		err := ploto.Scan(rows, &bank)
		common.PanicIfError(err)
		return bank, nil
	} else {
		return bank, errors.New("bank is not found")
	}
}

func (repository *BankRepository) Find(ctx context.Context, tx *sql.Tx, criteria string) []entity.Bank {
	SQL := "select oid,bankCode,bankName from bank " + criteria
	rows, err := tx.QueryContext(ctx, SQL)
	common.PanicIfError(err)
	defer rows.Close()

	var banks []entity.Bank
	for rows.Next() {
		bank := entity.Bank{}
		err := ploto.Scan(rows, &bank)
		common.PanicIfError(err)
		banks = append(banks, bank)
	}
	return banks
}
