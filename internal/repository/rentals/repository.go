package rental

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
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
			time.Now(),
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
func (r *repo) GetRentalByID(ctx context.Context, id int64) (*model.Rental, bool, error) {
	query, args, err := sq.Select(
		IDColumn,
		SpaceIDColumn,
		ClientIDColumn,
		StartDateColumn,
		EndDateColumn,
		PaidMonthsColumn,
	).
		From(RentalTable).
		Where(sq.Eq{IDColumn: id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, false, err
	}

	q := db.Query{
		Name:     "rental_repository.GetRentalByID",
		QueryRaw: query,
	}

	var rental model.Rental
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&rental.RentalID,
		&rental.SpaceID,
		&rental.ClientID,
		&rental.StartDate,
		&rental.EndDate,
		&rental.PaidMonths,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &rental, true, nil
}

// UpdateRental обновляет информацию о аренде
func (r *repo) UpdateRental(ctx context.Context, rental *model.Rental) error {
	builder := sq.Update(RentalTable)

	if rental.StartDate != 0 {
		builder = builder.Set(StartDateColumn, rental.StartDate)
	}
	if rental.EndDate != 0 {
		builder = builder.Set(EndDateColumn, rental.EndDate)
	}
	if rental.PaidMonths != 0 {
		builder = builder.Set(PaidMonthsColumn, rental.PaidMonths)
	}
	builder = builder.Set(UpdatedAtColumn, time.Now())

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
		return sys.NotFoundError
	}

	return nil
}

func (r *repo) CheckRentalOverlap(ctx context.Context, spaceCode int64, startDate int64, endDate int64) (bool, error) {
	query := `
	SELECT COUNT(*) 
	FROM rentals 
	WHERE space_code = $1 
	AND ((start_date BETWEEN $2 AND $3) 
	OR (end_date BETWEEN $2 AND $3) 
	OR ($2 BETWEEN start_date AND end_date))
	`

	q := db.Query{
		Name:     "rental_repository.CheckRentalOverlap",
		QueryRaw: query,
	}

	var count int
	err := r.db.DB().QueryRowContext(ctx, q, spaceCode, startDate, endDate).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
