package premises

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
	premiseRepository repository.PremiseRepository
}

// NewService возвращает новый объект сервисного слоя mall-tenants
func NewService(premiseRepo repository.PremiseRepository) service.PremiseService {
	return &srv{
		premiseRepository: premiseRepo,
	}
}

// SaveFloorPlan сохраняет SVG-файл и создаёт запись в БД
func (s *srv) SaveFloorPlan(ctx context.Context, floor int64, reader io.Reader) error {
	filename := fmt.Sprintf("floor_%d.svg", floor)
	fullPath := filepath.Join(svgBasePath, filename)

	if err := os.MkdirAll(svgBasePath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create svg file: %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, reader); err != nil {
		return fmt.Errorf("failed to write svg data: %w", err)
	}

	// сохраняем запись в БД
	plan := &model.FloorPlan{
		Floor:    floor,
		Filename: filename,
	}

	if err := s.premiseRepository.CreateFloorPlan(ctx, plan); err != nil {
		return fmt.Errorf("failed to create db entry: %w", err)
	}

	return nil
}

// GetFloorPlanContent возвращает SVG по этажу
func (s *srv) GetFloorPlanContent(ctx context.Context, floor int64) ([]byte, error) {
	plan, err := s.premiseRepository.GetFloorPlanByFloor(ctx, floor)
	if err != nil {
		return nil, fmt.Errorf("floor plan not found: %w", err)
	}

	fullPath := filepath.Join(svgBasePath, plan.Filename)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read svg file: %w", err)
	}

	return content, nil
}

// AddPolygon делегирует сохранение в репозиторий
func (s *srv) AddPolygon(ctx context.Context, poly *model.PremisePolygon) error {
	nounique, err := s.premiseRepository.CheckPremiseCode(ctx, poly.PremiseCode)
	if err != nil {
		return err
	}
	if nounique {
		return sys.NewCommonError("Premise code already exists", codes.Forbidden)
	}
	return s.premiseRepository.AddPolygon(ctx, poly)
}

// GetPolygons получает из репозитория все полигоны по коду помещения
func (s *srv) GetPolygons(ctx context.Context) ([]*model.PremisePolygon, error) {
	return s.premiseRepository.GetAllPolygons(ctx)
}
