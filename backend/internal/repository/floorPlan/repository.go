package floorplan

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	sq "github.com/Masterminds/squirrel"
	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
)

// Названия столбцов для таблицы polygons
const (
	PolygonsTable       = "polygons"
	PolygonCodeColumn   = "premise_code"
	PolygonPointsColumn = "points"
	PolygonLabelColumn  = "label"
)

// Названия таблиц и колонок
const (
	FloorPlansTable = "floor_plans"

	IDColumn       = "id"
	FloorColumn    = "floor"
	FilenameColumn = "filename"
)

// Структура репо с клиентом базы данных (интерфейсом)
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репо слоя
func NewRepository(db db.Client) repository.FloorPlanRepository {
	return &repo{db: db}
}

// CreateFloorPlan — создание записи о плане этажа
func (r *repo) CreateFloorPlan(ctx context.Context, plan *model.FloorPlan) error {
	query, args, err := sq.Insert(FloorPlansTable).
		Columns(FloorColumn, FilenameColumn).
		Values(plan.Floor, plan.Filename).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "floor_plan_repository.CreateFloorPlan",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

// GetFloorPlanByFloor — получить план по этажу
func (r *repo) GetFloorPlanByFloor(ctx context.Context, floor int64) (*model.FloorPlan, error) {
	query, args, err := sq.Select(IDColumn, FloorColumn, FilenameColumn).
		From(FloorPlansTable).
		Where(sq.Eq{FloorColumn: floor}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "floor_plan_repository.GetFloorPlanByFloor",
		QueryRaw: query,
	}

	var plan model.FloorPlan
	err = r.db.DB().ScanOneContext(ctx, &plan, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &plan, nil
}

// AddPolygon сохраняет полигон для помещения
func (r *repo) AddPolygon(ctx context.Context, poly *model.PremisePolygon) error {
	query, args, err := sq.Insert(PolygonsTable).
		Columns(
			PolygonCodeColumn,
			FloorColumn,
			PolygonPointsColumn,
			PolygonLabelColumn,
		).
		Values(
			poly.PremiseCode,
			poly.Floor,
			poly.Points,
			poly.Label,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "floor_plan_repository.AddPolygon",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) GetAllPolygons(ctx context.Context, floor int64) ([]*model.Polygons, error) {
	query, args, err := sq.
		Select(
			"p.premise_code",
			"pr.floor",
			"p.points",
			"p.label",
			"pr.status",
		).
		From(PolygonsTable + " AS p").
		Join("premises AS pr ON p.premise_code = pr.code").
		Where(
			sq.And{
				sq.Eq{"p.floor": floor},
				sq.Eq{"pr.floor": floor},
			}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "floor_plan_repository.GetAllPolygons",
		QueryRaw: query,
	}

	var polys []*model.Polygons
	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Polygons
		if err := rows.Scan(
			&p.PremiseCode,
			&p.Floor,
			&p.Points,
			&p.Label,
			&p.Status,
		); err != nil {
			return nil, err
		}
		polys = append(polys, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return polys, nil
}

// CheckPremiseCode уникальность кода помещения
func (r *repo) CheckPremiseCode(ctx context.Context, code int64) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM polygons WHERE premise_code = $1);"

	q := db.Query{
		Name:     "floor_plan_repository.CheckPremiseCode",
		QueryRaw: query,
	}

	var exists bool
	err := r.db.DB().QueryRowContext(ctx, q, code).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// DeletPolygon удаляет полигон по id помещения
func (r *repo) DeletPolygon(ctx context.Context, premiseCode int64) error {
	query, args, err := sq.Delete(PolygonsTable).
		Where(sq.Eq{PolygonCodeColumn: premiseCode}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "floor_plan_repository.DeletPolygon",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
