package model

import "time"

// Payment представляет собой модель данных для платежей
type Payment struct {
	PaymentID int64     `json:"payment_id" db:"payment_id"`
	RentID    int64     `json:"rent_id" db:"rent_id"`
	Date      time.Time `json:"payment_date" db:"payment_date"`
	Amount    int64     `json:"amount" db:"amount"`
}
