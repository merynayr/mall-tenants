package rentals

import (
	"context"
	"time"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

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

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var rentalID int64
		var errTx error
		if rentalID, errTx = s.rentalRepository.CreateRental(ctx, &rental); errTx != nil {
			return errTx
		}

		updatePremise := model.Premises{
			Code:   rental.SpaceID,
			Status: string(model.Occupied),
		}
		if errTx := s.premiseRepository.UpdatePremise(ctx, &updatePremise); errTx != nil {
			return errTx
		}

		start := rental.StartDate
		end := rental.EndDate
		payments := make([]model.Payment, 0)

		currentStart := start

		for currentStart.Before(end) {
			nextMonth := currentStart.AddDate(0, 1, 0)

			var currentEnd time.Time
			if nextMonth.After(end) {
				currentEnd = end
			} else {
				currentEnd = nextMonth.AddDate(0, 0, -1)
			}

			payment := model.Payment{
				RentalID:    rentalID,
				PeriodStart: currentStart,
				PeriodEnd:   currentEnd,
				Amount:      int(premise.RentPerMonth),
				IsPaid:      false,
				CreatedAt:   time.Now().UTC(),
			}
			payments = append(payments, payment)

			currentStart = nextMonth
		}

		if errTx := s.paymentRepository.CreatePayment(ctx, payments); errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}
