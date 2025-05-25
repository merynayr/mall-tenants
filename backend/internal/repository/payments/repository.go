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
	"github.com/merynayr/mall-tenants/internal/sys"
)

// Константы названий таблицы и столбцов
const (
	PaymentTable    = "payments"
	PaymentIDColumn = "payment_id"
	RentalIDColumn  = "rental_id"
	PeriodStartCol  = "period_start"
	PeriodEndCol    = "period_end"
	AmountColumn    = "amount"
	IsPaidColumn    = "is_paid"
	PaymentDateCol  = "payment_date"
	CreatedAtColumn = "created_at"
	UpdatedAtColumn = "updated_at"
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
func (r *repo) CreatePayment(ctx context.Context, p model.Payment) (int64, error) {
	query, args, err := sq.Insert(PaymentTable).
		Columns(RentalIDColumn, PeriodStartCol, PeriodEndCol, AmountColumn, IsPaidColumn, PaymentDateCol).
		Values(p.RentalID, p.PeriodStart, p.PeriodEnd, p.Amount, p.IsPaid, p.PaymentDate).
		Suffix("RETURNING " + PaymentIDColumn).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "payment_repository.CreatePayment",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, err

}

func (r *repo) MarkAsPaid(ctx context.Context, id int64, paidAt time.Time) error {
	query, args, err := sq.Update(PaymentTable).
		Set(IsPaidColumn, true).
		Set(PaymentDateCol, paidAt).
		Set(UpdatedAtColumn, time.Now()).
		Where(sq.Eq{PaymentIDColumn: id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "payment_repository.MarkAsPaid",
		QueryRaw: query,
	}

	ct, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return errors.New("no payment found to update")
	}
	return nil
}

func (r *repo) MarkAsPaidMany(ctx context.Context, ids []int64, paymentDate time.Time) error {
	query, args, err := sq.Update(PaymentTable).
		Set(IsPaidColumn, true).
		Set(PaymentDateCol, paymentDate).
		Set(UpdatedAtColumn, time.Now()).
		Where(sq.Eq{PaymentIDColumn: ids}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "payment_repository.MarkAsPaidMany",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) GetByID(ctx context.Context, id int64) (model.Payment, error) {
	query, args, err := sq.Select("*").
		From(PaymentTable).
		Where(sq.Eq{PaymentIDColumn: id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return model.Payment{}, err
	}

	q := db.Query{
		Name:     "payment_repository.GetByID",
		QueryRaw: query,
	}

	var p model.Payment
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&p.ID, &p.RentalID, &p.PeriodStart, &p.PeriodEnd,
		&p.Amount, &p.IsPaid, &p.PaymentDate, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Payment{}, sys.PaymentsNotFoundError
		}
		return model.Payment{}, err
	}
	return p, nil
}

func (r *repo) List(ctx context.Context, f model.PaymentFilter) ([]model.Payment, error) {
	qb := sq.Select("*").
		From(PaymentTable).
		OrderBy(PeriodStartCol + " DESC").
		Limit(f.Limit).
		Offset(f.Offset).
		PlaceholderFormat(sq.Dollar)

	if f.IsPaid != nil {
		qb = qb.Where(sq.Eq{IsPaidColumn: *f.IsPaid})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "payment_repository.List",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Payment
	for rows.Next() {
		var p model.Payment
		err := rows.Scan(
			&p.ID, &p.RentalID, &p.PeriodStart, &p.PeriodEnd,
			&p.Amount, &p.IsPaid, &p.PaymentDate, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, rows.Err()
}
