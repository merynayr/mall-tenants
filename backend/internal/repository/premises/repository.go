package premises

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// Константы названий столбцов БД
const (
	PremisesTable = "premises"

	CodeColumn            = "code"
	FloorColumn           = "floor"
	AreaColumn            = "area"
	TypeColumn            = "type"
	RentPerMonthColumn    = "rent_per_month"
	StatusColumn          = "status"
	SecuritySystemColumn  = "security_system"
	AirConditioningColumn = "air_conditioning"
)

// Структура репо с клиентом базы данных (интерфейсом)
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репо слоя
func NewRepository(db db.Client) repository.PremiseRepository {
	return &repo{db: db}
}

// CreatePremises создает новое помещение в базе данных
func (r *repo) CreatePremise(ctx context.Context, premise *model.Premises) error {
	query, args, err := sq.Insert(PremisesTable).
		Columns(
			CodeColumn,
			FloorColumn,
			AreaColumn,
			TypeColumn,
			RentPerMonthColumn,
			StatusColumn,
			SecuritySystemColumn,
			AirConditioningColumn,
		).
		Values(
			premise.Code,
			premise.Floor,
			premise.Area,
			premise.Type,
			premise.RentPerMonth,
			premise.Status,
			premise.SecuritySystem,
			premise.AirConditioning,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "premises_repository.CreatePremises",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePremises обновляет информацию о помещении в базе данных
func (r *repo) UpdatePremise(ctx context.Context, premise *model.Premises) error {
	builder := sq.Update(PremisesTable)

	if premise.Floor != 0 {
		builder = builder.Set(FloorColumn, premise.Floor)
	}
	if premise.Area != 0 {
		builder = builder.Set(AreaColumn, premise.Area)
	}
	if premise.Type != "" {
		builder = builder.Set(TypeColumn, premise.Type)
	}
	if premise.RentPerMonth != 0 {
		builder = builder.Set(RentPerMonthColumn, premise.RentPerMonth)
	}
	if premise.Status != "" {
		builder = builder.Set(StatusColumn, premise.Status)
	}
	if premise.SecuritySystem != "" {
		builder = builder.Set(SecuritySystemColumn, premise.SecuritySystem)
	}
	if premise.AirConditioning != "" {
		builder = builder.Set(AirConditioningColumn, premise.AirConditioning)
	}

	query, args, err := builder.PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{CodeColumn: premise.Code}).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "premises_repository.UpdatePremises",
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return sys.PremiseNotFoundError
	}

	return nil
}

// GetPremisesByCode получает помещение по его коду
func (r *repo) GetPremisesByCode(ctx context.Context, code int64) (*model.Premises, bool, error) {
	query, args, err := sq.Select(
		CodeColumn,
		FloorColumn,
		AreaColumn,
		TypeColumn,
		RentPerMonthColumn,
		StatusColumn,
		SecuritySystemColumn,
		AirConditioningColumn,
	).
		From(PremisesTable).
		Where(sq.Eq{CodeColumn: code}).
		PlaceholderFormat(sq.Dollar).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, false, err
	}

	q := db.Query{
		Name:     "premises_repository.GetPremisesByCode",
		QueryRaw: query,
	}

	var premise model.Premises
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&premise.Code,
		&premise.Floor,
		&premise.Area,
		&premise.Type,
		&premise.RentPerMonth,
		&premise.Status,
		&premise.SecuritySystem,
		&premise.AirConditioning,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &premise, true, nil
}

// GetAllPremises возвращает список всех помещений
func (r *repo) GetAllPremises(ctx context.Context) ([]model.Premises, error) {
	query, args, err := sq.Select(
		CodeColumn,
		FloorColumn,
		AreaColumn,
		TypeColumn,
		RentPerMonthColumn,
		StatusColumn,
		SecuritySystemColumn,
		AirConditioningColumn,
	).
		From(PremisesTable).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "premises_repository.GetAllPremises",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var premises []model.Premises
	for rows.Next() {
		var p model.Premises
		err := rows.Scan(
			&p.Code,
			&p.Floor,
			&p.Area,
			&p.Type,
			&p.RentPerMonth,
			&p.Status,
			&p.SecuritySystem,
			&p.AirConditioning,
		)
		if err != nil {
			return nil, err
		}
		premises = append(premises, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return premises, nil
}
