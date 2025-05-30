package rentals

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// ContractPath - путь, куда складываются документы с договорами
const ContractPath = "contracts/final"

func (s *srv) CreateRental(ctx context.Context, rental model.Rental) error {
	premise, exist, err := s.premiseRepository.GetPremisesByCode(ctx, rental.SpaceID)
	if err != nil {
		return err
	}
	if !exist {
		return sys.PremiseNotFoundError
	}

	if premise.Status == string(model.Occupied) {
		return sys.PremiseOccupiedError
	} else if premise.Status == string(model.Maintenance) {
		return sys.PremiseUnderRepairError
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) (err error) {
		rentalID, err := s.rentalRepository.CreateRental(ctx, &rental)
		if err != nil {
			return err
		}

		// Сохраняем загруженный договор
		if rental.TemplateFile == nil {
			return errors.New("договор не загружен")
		}

		src, err := rental.TemplateFile.Open()
		if err != nil {
			return fmt.Errorf("ошибка при открытии файла договора: %w", err)
		}
		defer func() {
			if cerr := src.Close(); cerr != nil && err == nil {
				err = cerr
			}
		}()

		// Создание пути, если не существует
		if err := os.MkdirAll(ContractPath, 0750); err != nil {
			return fmt.Errorf("ошибка при создании директории: %w", err)
		}

		// Сохраняем файл в /contracts/final/contract_rental_<id>.pdf
		filename := fmt.Sprintf("%s/contract_rental_%d.pdf", ContractPath, rentalID)
		dst, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("ошибка при создании файла договора: %w", err)
		}
		defer func() {
			if cerr := dst.Close(); cerr != nil && err == nil {
				err = cerr
			}
		}()

		if _, err := io.Copy(dst, src); err != nil {
			return fmt.Errorf("ошибка при сохранении договора: %w", err)
		}

		// Создаём запись о договоре
		contract := model.Contract{
			RentalID:  rentalID,
			FilePath:  filename,
			IsActive:  false,
			IsSigned:  false,
			CreatedAt: time.Now().UTC(),
		}

		if err := s.contractRepository.CreateContract(ctx, &contract); err != nil {
			return fmt.Errorf("ошибка при создании записи о договоре: %w", err)
		}

		return nil
	})

	return err
}
