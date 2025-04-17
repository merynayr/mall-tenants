package payments

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
)

// Константы названий таблицы и столбцов
const (
	PaymentTable = "payments"
	IDColumn     = "payment_id"
	RentIDColumn = "rent_id"
	DateColumn   = "payment_date"
	AmountColumn = "amount"
)

// Структура репозитория с клиентом базы данных
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репозитория платежей
func NewRepository(db db.Client) repository.PaymentRepository {
	return &repo{db: db}
}

// CreatePayment создаёт новый платёж в базе данных
func (r *repo) CreatePayment(ctx context.Context, payment *model.Payment) error {
	query, args, err := sq.Insert(PaymentTable).
		Columns(
			RentIDColumn,
			DateColumn,
			AmountColumn,
		).
		Values(
			payment.RentID,
			payment.Date,
			payment.Amount,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "payment_repository.CreatePayment",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetPaymentByID получает платёж по его PaymentID
func (r *repo) GetPaymentByID(ctx context.Context, id int64) (*model.Payment, bool, error) {
	query, args, err := sq.Select(
		IDColumn,
		RentIDColumn,
		DateColumn,
		AmountColumn,
	).
		From(PaymentTable).
		Where(sq.Eq{IDColumn: id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, false, err
	}

	q := db.Query{
		Name:     "payment_repository.GetPaymentByID",
		QueryRaw: query,
	}

	var payment model.Payment
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&payment.PaymentID,
		&payment.RentID,
		&payment.Date,
		&payment.Amount,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &payment, true, nil
}

// GetLastPayment получает последний платёж по аренде
func (r *repo) GetLastPayment(ctx context.Context, rentID int64) (*model.Payment, bool, error) {
	query, args, err := sq.Select(
		IDColumn,
		RentIDColumn,
		DateColumn,
		AmountColumn,
	).
		From(PaymentTable).
		Where(sq.Eq{RentIDColumn: rentID}).
		OrderBy(DateColumn + " DESC").
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, false, err
	}

	q := db.Query{
		Name:     "payment_repository.GetLastPayment",
		QueryRaw: query,
	}

	var payment model.Payment
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&payment.PaymentID,
		&payment.RentID,
		&payment.Date,
		&payment.Amount,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &payment, true, nil
}

// GetOverduePayments получает список просроченных платежей по аренде
func (r *repo) GetOverduePayments(ctx context.Context, rentID int64) ([]model.Payment, error) {
	now := time.Now().UTC()
	query, args, err := sq.Select(
		IDColumn,
		RentIDColumn,
		DateColumn,
		AmountColumn,
	).
		From(PaymentTable).
		Where(sq.And{
			sq.Eq{RentIDColumn: rentID},
			sq.Lt{DateColumn: now},
		}).
		OrderBy(DateColumn + " ASC").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "payment_repository.GetOverduePayments",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []model.Payment
	for rows.Next() {
		var payment model.Payment
		err := rows.Scan(
			&payment.PaymentID,
			&payment.RentID,
			&payment.Date,
			&payment.Amount,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}

// GetPaymentByRentIDAndPeriod проверяет, существует ли платёж для аренды на данный период
func (r *repo) GetPaymentByRentIDAndPeriod(ctx context.Context, rentID int64, startDate, endDate time.Time) (bool, error) {
	query, args, err := sq.Select(
		"1",
	).
		PlaceholderFormat(sq.Dollar).
		From(PaymentTable).
		Where(sq.And{
			sq.Eq{RentIDColumn: rentID},
			sq.GtOrEq{DateColumn: startDate},
			sq.LtOrEq{DateColumn: endDate},
		}).ToSql()

	if err != nil {
		return true, err
	}

	q := db.Query{
		Name:     "payment_repository.GetPaymentByRentIDAndPeriod",
		QueryRaw: query,
	}

	var exists int32
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return exists == 1, nil
}
