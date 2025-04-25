package floorplan

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/sys/codes"
)

const svgBasePath = "./assets/floor_plans"

type srv struct {
	floorPlanRepository repository.FloorPlanRepository
}

// NewService возвращает новый объект сервисного слоя mall-tenants
func NewService(premiseRepo repository.FloorPlanRepository) service.FloorPlanService {
	return &srv{
		floorPlanRepository: premiseRepo,
	}
}

// SaveFloorPlan сохраняет SVG-файл и создаёт запись в БД
func (s *srv) SaveFloorPlan(ctx context.Context, floor int64, reader io.Reader) error {
	if floor < 0 {
		return fmt.Errorf("invalid floor number: %d", floor)
	}

	filename := fmt.Sprintf("floor_%d.svg", floor)
	fullPath := filepath.Join(svgBasePath, filename)

	if err := os.MkdirAll(svgBasePath, 0750); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(fullPath) // #nosec G304
	if err != nil {
		return fmt.Errorf("failed to create svg file: %w", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			return
		}
	}()

	if _, err := io.Copy(f, reader); err != nil {
		return fmt.Errorf("failed to write svg data: %w", err)
	}

	plan := &model.FloorPlan{
		Floor:    floor,
		Filename: filename,
	}

	if err := s.floorPlanRepository.CreateFloorPlan(ctx, plan); err != nil {
		return fmt.Errorf("failed to create db entry: %w", err)
	}

	return nil
}

// GetFloorPlanContent возвращает SVG по этажу
func (s *srv) GetFloorPlanContent(ctx context.Context, floor int64) ([]byte, error) {
	if floor < 1 {
		return nil, fmt.Errorf("%s", "номер этажа' должен быть положительным")
	}
	plan, err := s.floorPlanRepository.GetFloorPlanByFloor(ctx, floor)
	if err != nil {
		return nil, fmt.Errorf("floor plan not found: %w", err)
	}
	if plan == nil {
		return nil, sys.NotFoundError
	}
	fullPath := filepath.Join(svgBasePath, plan.Filename)
	content, err := os.ReadFile(fullPath) // #nosec G304
	if err != nil {
		return nil, fmt.Errorf("failed to read svg file: %w", err)
	}

	return content, nil
}

// AddPolygon делегирует сохранение в репозиторий
func (s *srv) AddPolygon(ctx context.Context, poly *model.PremisePolygon) error {
	if poly.PremiseCode < 1 {
		return fmt.Errorf("%s", "номер помещения должен быть положительным")
	}
	nounique, err := s.floorPlanRepository.CheckPremiseCode(ctx, poly.PremiseCode)
	if err != nil {
		return err
	}
	if nounique {
		return sys.NewCommonError("Premise code already exists", codes.Forbidden)
	}
	return s.floorPlanRepository.AddPolygon(ctx, poly)
}

// GetPolygons получает из репозитория все полигоны по коду помещения
func (s *srv) GetPolygons(ctx context.Context, floor int64) ([]*model.Polygons, error) {
	if floor < 1 {
		return nil, fmt.Errorf("%s", "номер этажа' должен быть положительным")
	}
	return s.floorPlanRepository.GetAllPolygons(ctx, floor)
}

// DeletPolygon удаляет из репозитория полигон по коду помещения
func (s *srv) DeletPolygon(ctx context.Context, premiseCode int64) error {
	if premiseCode < 1 {
		return fmt.Errorf("%s", "номер помещения' должен быть положительным")
	}
	return s.floorPlanRepository.DeletPolygon(ctx, premiseCode)
}
