package model

import "time"

// Rental представляет собой модель данных для аренды
type Rental struct {
	RentalID   int64      `json:"rental_id,omitempty" db:"rental_id"`
	SpaceID    int64      `json:"space_code" db:"space_code"`
	ClientID   int64      `json:"client_id" db:"client_id"`
	StartDate  time.Time  `json:"start_date" db:"start_date"`
	EndDate    time.Time  `json:"end_date" db:"end_date"`
	PaidMonths int64      `json:"paid_months" db:"paid_months"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Agreements представляет собой модель данных для аренды
type Agreements struct {
	RentalID   int64      `json:"rental_id,omitempty"`
	SpaceID    int64      `json:"space_code"`
	ClientName string     `json:"client_name"`
	StartDate  time.Time  `json:"start_date"`
	EndDate    time.Time  `json:"end_date"`
	PaidMonths int64      `json:"paid_months"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
