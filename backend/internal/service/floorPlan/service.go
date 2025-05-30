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

// SaveFloorPlan сохраняет SVG или PNG-файл и создаёт запись в БД
func (s *srv) SaveFloorPlan(ctx context.Context, floor int64, reader io.Reader, mimeType string) (err error) {
	if floor < 0 {
		return sys.FloorNumberInvalidError
	}

	var extension string
	switch mimeType {
	case "image/svg+xml":
		extension = ".svg"
	case "image/png":
		extension = ".png"
	default:
		return fmt.Errorf("unsupported file type: %s", mimeType)
	}

	filename := fmt.Sprintf("floor_%d%s", floor, extension)
	fullPath := filepath.Join(svgBasePath, filename)

	if err := os.MkdirAll(svgBasePath, 0750); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(fullPath) // #nosec G304
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	if _, err := io.Copy(f, reader); err != nil {
		return fmt.Errorf("failed to write file data: %w", err)
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
		return nil, sys.FloorNumberInvalidError
	}
	plan, err := s.floorPlanRepository.GetFloorPlanByFloor(ctx, floor)
	if err != nil {
		return nil, err
	}
	if plan == nil {
		return nil, sys.FloorPlanNotFoundError
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
		return sys.InvalidIDFormatError
	}
	nounique, err := s.floorPlanRepository.CheckPremiseCode(ctx, poly.PremiseCode)
	if err != nil {
		return err
	}
	if nounique {
		return sys.PremiseAlreadyExistsError
	}
	return s.floorPlanRepository.AddPolygon(ctx, poly)
}

// GetPolygons получает из репозитория все полигоны по коду помещения
func (s *srv) GetPolygons(ctx context.Context, floor int64) ([]*model.Polygons, error) {
	if floor < 1 {
		return nil, sys.FloorNumberInvalidError
	}
	return s.floorPlanRepository.GetAllPolygons(ctx, floor)
}

// DeletPolygon удаляет из репозитория полигон по коду помещения
func (s *srv) DeletPolygon(ctx context.Context, premiseCode int64) error {
	if premiseCode < 1 {
		return sys.FloorNumberInvalidError
	}
	return s.floorPlanRepository.DeletPolygon(ctx, premiseCode)
}
