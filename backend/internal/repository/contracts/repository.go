package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
)

const (
	contractTable = "contracts"

	colContractID = "contract_id"
	colRentalID   = "rental_id"
	colFilePath   = "file_path"
	colSignature  = "signature"
	colPublicKey  = "public_key"
	colIsActive   = "is_active"
	colIsSigned   = "is_signed"
	colSignedAt   = "signed_at"
	colCreatedAt  = "created_at"
)

// Структура репозитория с клиентом базы данных
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репозитория договоров
func NewRepository(db db.Client) repository.ContractRepository {
	return &repo{db: db}
}

func (r *repo) CreateContract(ctx context.Context, c *model.Contract) error {
	query, args, err := sq.Insert(contractTable).
		Columns(
			colRentalID,
			colFilePath,
			colIsActive,
			colIsSigned,
			colCreatedAt,
		).
		Values(
			c.RentalID,
			c.FilePath,
			c.IsActive,
			c.IsSigned,
			c.CreatedAt,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "contract_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) UpdateSignature(ctx context.Context, contractID int64, signature []byte, publicKey string) error {
	query, args, err := sq.Update(contractTable).
		Set(colSignature, signature).
		Set(colPublicKey, publicKey).
		Set(colIsSigned, true).
		Set(colIsActive, true).
		Set(colSignedAt, time.Now().UTC()).
		Where(sq.Eq{colContractID: contractID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "contract_repository.UpdateSignature",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) GetByRentalID(ctx context.Context, rentalID int64) ([]model.Contracts, error) {
	queryBuilder := sq.Select(colContractID, colIsActive, colIsSigned, colCreatedAt).
		From(contractTable).
		Where(sq.Eq{colRentalID: rentalID}).
		OrderBy(colCreatedAt + " DESC").
		PlaceholderFormat(sq.Dollar)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "contract_repository.GetByRentalID",
		QueryRaw: query,
	}

	var contracts []model.Contracts
	err = r.db.DB().ScanAllContext(ctx, &contracts, q, args...)
	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (r *repo) GetByContractID(ctx context.Context, contractID int64) (*model.Contracts, error) {
	queryBuilder := sq.Select(colContractID, colRentalID, colFilePath, colIsActive, colSignedAt, colCreatedAt).
		From(contractTable).
		Where(sq.Eq{colContractID: contractID}).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "contract_repository.GetByRentalID",
		QueryRaw: query,
	}

	var contract model.Contracts
	err = r.db.DB().ScanOneContext(ctx, &contract, q, args...)
	if err != nil {
		return nil, err
	}

	return &contract, nil
}

func (r *repo) DeleteContract(ctx context.Context, contractID int64) error {
	query, args, err := sq.Delete(contractTable).
		Where(sq.Eq{colContractID: contractID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "contract_repository.GetByRentalID",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]model.Contract, error) {
	queryBuilder := sq.Select(colContractID, colRentalID, colFilePath, colIsActive, colSignedAt, colCreatedAt).
		From(contractTable).
		OrderBy(colCreatedAt + " DESC").
		PlaceholderFormat(sq.Dollar)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "contract_repository.GetAll",
		QueryRaw: query,
	}

	var contracts []model.Contract
	err = r.db.DB().ScanAllContext(ctx, &contracts, q, args...)
	if err != nil {
		return nil, err
	}

	return contracts, nil
}
