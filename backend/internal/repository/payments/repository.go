package payments

import (
	"context"
	"errors"
	"strings"
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

	ClientIDColumn = "client_id"
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
func (r *repo) CreatePayment(ctx context.Context, payments []model.Payment) error {
	if len(payments) == 0 {
		return nil
	}

	builder := sq.Insert(PaymentTable).
		Columns(RentalIDColumn, PeriodStartCol, PeriodEndCol, AmountColumn, IsPaidColumn, PaymentDateCol).
		PlaceholderFormat(sq.Dollar)

	for _, p := range payments {
		builder = builder.Values(p.RentalID, p.PeriodStart, p.PeriodEnd, p.Amount, p.IsPaid, p.PaymentDate)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "payment_repository.CreatePaymentBatch",
		QueryRaw: query,
	}

	var ids []int64
	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		return err
	}

	return nil
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

func (r *repo) GetByClientID(ctx context.Context, id int64, f model.PaymentFilter) ([]model.PaymentList, error) {
	qb := sq.
		Select(
			"p.payment_id", "p.rental_id", "p.period_start", "p.period_end",
			"p.amount", "p.is_paid", "p.payment_date", "p.created_at", "p.updated_at",
			"pr.code", "c.organization_name",
		).
		From(PaymentTable + " AS p").
		Join("rentals r ON p.rental_id = r.rental_id").
		Join("premises pr ON r.space_code = pr.code").
		Join("clients c ON r.client_id = c.client_id").
		Where(sq.Eq{"c.client_id": id}).
		PlaceholderFormat(sq.Dollar)

	if f.IsPaid != nil {
		qb = qb.Where(sq.Eq{"p." + IsPaidColumn: *f.IsPaid})
	}

	validSortFields := map[string]string{
		"period_start": "p.period_start",
		"created_at":   "p.created_at",
		"updated_at":   "p.updated_at",
	}

	sortField, ok := validSortFields[f.SortBy]
	if !ok {
		sortField = "p.period_start"
	}

	order := "DESC"
	if strings.ToLower(f.SortOrder) == "asc" {
		order = "ASC"
	}

	qb = qb.OrderBy(sortField + " " + order).
		Limit(f.Limit).
		Offset(f.Offset)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "payment_repository.GetByClientID",
		QueryRaw: query,
	}

	var result []model.PaymentList
	err = r.db.DB().ScanAllContext(ctx, &result, q, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repo) List(ctx context.Context, f model.PaymentFilter) ([]model.PaymentList, error) {
	qb := sq.
		Select(
			"p.payment_id", "p.rental_id", "p.period_start", "p.period_end",
			"p.amount", "p.is_paid", "p.payment_date", "p.created_at", "p.updated_at",
			"pr.code", "c.organization_name",
		).
		From(PaymentTable + " AS p").
		Join("rentals r ON p.rental_id = r.rental_id").
		Join("premises pr ON r.space_code = pr.code").
		Join("clients c ON r.client_id = c.client_id").
		Where(sq.LtOrEq{"p." + PeriodStartCol: time.Now().AddDate(0, 0, -3)}).
		PlaceholderFormat(sq.Dollar)

	if f.IsPaid != nil {
		qb = qb.Where(sq.Eq{"p." + IsPaidColumn: *f.IsPaid})
	}

	if f.Client != "" {
		qb = qb.Where(sq.ILike{"c.organization_name": "%" + f.Client + "%"})
	}

	validSortFields := map[string]string{
		"period_start": "p.period_start",
		"created_at":   "p.created_at",
		"updated_at":   "p.updated_at",
	}

	sortField, ok := validSortFields[f.SortBy]
	if !ok {
		sortField = "p.period_start"
	}

	order := "DESC"
	if strings.ToLower(f.SortOrder) == "asc" {
		order = "ASC"
	}

	qb = qb.OrderBy(sortField + " " + order).
		Limit(f.Limit).
		Offset(f.Offset)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "payment_repository.List",
		QueryRaw: query,
	}

	var result []model.PaymentList
	err = r.db.DB().ScanAllContext(ctx, &result, q, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
