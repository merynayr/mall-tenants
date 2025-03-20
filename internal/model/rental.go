package model

import "time"

// Rental представляет собой модель данных для аренды
type Rental struct {
	RentalID   int64     `json:"rental_id"`
	SpaceID    int64     `json:"space_code"`
	ClientID   int64     `json:"client_id"`
	StartDate  int64     `json:"start_date"`
	EndDate    int64     `json:"end_date"`
	PaidMonths int64     `json:"paid_months"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
