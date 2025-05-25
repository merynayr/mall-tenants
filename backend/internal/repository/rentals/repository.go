package rental

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// Константы названий столбцов БД
const (
	RentalTable      = "rentals"
	IDColumn         = "rental_id"
	SpaceIDColumn    = "space_code"
	ClientIDColumn   = "client_id"
	StartDateColumn  = "start_date"
	EndDateColumn    = "end_date"
	PaidMonthsColumn = "paid_months"
	CreatedAtColumn  = "created_at"
	UpdatedAtColumn  = "updated_at"
)

// Структура репозитория с клиентом базы данных
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репозитория аренды
func NewRepository(db db.Client) repository.RentalRepository {
	return &repo{db: db}
}

// CreateRental создаёт новую аренду в базе данных
func (r *repo) CreateRental(ctx context.Context, rental *model.Rental) error {
	query, args, err := sq.Insert(RentalTable).
		Columns(
			SpaceIDColumn,
			ClientIDColumn,
			StartDateColumn,
			EndDateColumn,
			PaidMonthsColumn,
			CreatedAtColumn,
		).
		Values(
			rental.SpaceID,
			rental.ClientID,
			rental.StartDate,
			rental.EndDate,
			rental.PaidMonths,
			time.Now().UTC(),
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "rental_repository.CreateRental",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetRentalByID получает аренду по её RentalID
func (r *repo) GetRentalsByID(ctx context.Context, id int64) ([]model.Rental, error) {
	query, args, err := sq.Select(
		IDColumn,
		SpaceIDColumn,
		ClientIDColumn,
		StartDateColumn,
		EndDateColumn,
		PaidMonthsColumn,
		CreatedAtColumn,
	).
		From(RentalTable).
		Where(sq.Eq{ClientIDColumn: id}).
		OrderBy(CreatedAtColumn + " DESC").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "rental_repository.GetRentalsByID",
		QueryRaw: query,
	}

	var rentals []model.Rental
	err = r.db.DB().ScanAllContext(ctx, &rentals, q, args...)

	if err != nil {
		return nil, err
	}

	return rentals, nil
}

// UpdateRental обновляет информацию о аренде
func (r *repo) UpdateRental(ctx context.Context, rental *model.Rental) error {
	builder := sq.Update(RentalTable)

	if !rental.StartDate.IsZero() {
		builder = builder.Set(StartDateColumn, rental.StartDate)
	}
	if !rental.EndDate.IsZero() {
		builder = builder.Set(EndDateColumn, rental.EndDate)
	}
	if rental.PaidMonths != 0 {
		builder = builder.Set(PaidMonthsColumn, sq.Expr(PaidMonthsColumn+" + ?", rental.PaidMonths))
	}

	builder = builder.Set(UpdatedAtColumn, time.Now().UTC())

	query, args, err := builder.PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{IDColumn: rental.RentalID}).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "rental_repository.UpdateRental",
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return sys.RentalsNotFoundError
	}

	return nil
}

func (r *repo) GetAgreements(ctx context.Context, limit, offset uint64) ([]model.Agreements, error) {
	query, args, err := sq.Select(
		"r."+IDColumn,
		"r."+SpaceIDColumn,
		"c.organization_name",
		"r."+StartDateColumn,
		"r."+EndDateColumn,
		"r."+PaidMonthsColumn,
		"r."+CreatedAtColumn,
	).
		From(RentalTable + " r").
		Join("clients c ON c.client_id = r.client_id").
		PlaceholderFormat(sq.Dollar).
		Limit(limit).
		Offset(offset).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "rental_repository.GetAgreements",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rentals []model.Agreements

	for rows.Next() {
		var rental model.Agreements
		err := rows.Scan(
			&rental.RentalID,
			&rental.SpaceID,
			&rental.ClientName,
			&rental.StartDate,
			&rental.EndDate,
			&rental.PaidMonths,
			&rental.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		rentals = append(rentals, rental)
	}

	return rentals, nil
}
