package applications

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
)

// Константы названий столбцов БД
const (
	applicationsTable = "applications"

	idColumn               = "id"
	PremiseNumberColumn    = "premise_number"
	EmailColumn            = "email"
	OrganizationNameColumn = "organization_name"
	ContactPersonColumn    = "contact_person"
	AddressColumn          = "address"
	PhoneColumn            = "phone"
	RequisitesColumn       = "requisites"
	AdditionalInfoColumn   = "additional_info"
	IsProcessedColumn      = "is_processed"
	CreatedAtColumn        = "created_at"
)

// Структура репо с клиентом базы данных (интерфейсом)
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репо слоя
func NewRepository(db db.Client) repository.ApplicationRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, a *model.Application) error {
	query, args, err := squirrel.
		Insert(applicationsTable).
		Columns(OrganizationNameColumn, ContactPersonColumn, AddressColumn, PhoneColumn, RequisitesColumn, EmailColumn, PremiseNumberColumn, AdditionalInfoColumn, IsProcessedColumn, CreatedAtColumn).
		Values(a.OrganizationName, a.ContactPerson, a.Address, a.Phone, a.Requisites, a.Email, a.PremiseNumber, a.AdditionalInfo, a.IsProcessed, a.CreatedAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "application_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) GetAll(ctx context.Context, filter model.ApplicationFilter) ([]*model.Application, error) {
	sb := squirrel.
		Select(
			idColumn,
			OrganizationNameColumn,
			ContactPersonColumn,
			AddressColumn,
			PhoneColumn,
			RequisitesColumn,
			EmailColumn,
			PremiseNumberColumn,
			AdditionalInfoColumn,
			CreatedAtColumn,
			IsProcessedColumn,
		).
		From(applicationsTable).
		OrderBy("created_at DESC").
		Limit(filter.Limit).
		Offset(filter.Offset).
		PlaceholderFormat(squirrel.Dollar)

	if filter.IsProcessed != nil {
		sb = sb.Where(squirrel.Eq{IsProcessedColumn: *filter.IsProcessed})
	}

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "application_repository.GetAll",
		QueryRaw: query,
	}

	var apps []*model.Application
	err = r.db.DB().ScanAllContext(ctx, &apps, q, args...)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (r *repo) SetProcessedStatus(ctx context.Context, id int64, isProcessed bool) error {
	query, args, err := squirrel.
		Update(applicationsTable).
		Set(IsProcessedColumn, isProcessed).
		Where(squirrel.Eq{idColumn: id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "application_repository.SetProcessedStatus",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}
