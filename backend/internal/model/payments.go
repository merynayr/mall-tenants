package model

import "time"

// Payment представляет собой модель данных для платежей
type Payment struct {
	ID          int64      `json:"id"`
	RentalID    int64      `json:"rental_id"`
	PeriodStart time.Time  `json:"start_date"`
	PeriodEnd   time.Time  `json:"end_date"`
	Amount      int        `json:"amount"`
	IsPaid      bool       `json:"is_paid"`
	PaymentDate *time.Time `json:"payment_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// PaymentFilter - структура фильтрации для платежей
type PaymentFilter struct {
	IsPaid    *bool
	Client    string
	SortBy    string
	SortOrder string
	Limit     uint64
	Offset    uint64
}

// MarkPaymentsPaidRequest - структура
type MarkPaymentsPaidRequest struct {
	PaymentIDs []int64 `json:"payment_ids" swaggertype:"array,integer"`
}

// PaymentList структура, возвращаемая на фронт
type PaymentList struct {
	ID          int64      `json:"id"           db:"payment_id"`
	RentalID    int64      `json:"rental_id"    db:"rental_id"`
	PeriodStart time.Time  `json:"start_date"   db:"period_start"`
	PeriodEnd   time.Time  `json:"end_date"     db:"period_end"`
	Amount      int        `json:"amount"       db:"amount"`
	IsPaid      bool       `json:"status"       db:"is_paid"`
	PaymentDate *time.Time `json:"payment_date,omitempty" db:"payment_date"`
	CreatedAt   time.Time  `json:"created_at"   db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"   db:"updated_at"`
	SpaceCode   string     `json:"space_code"   db:"code"`
	ClientName  string     `json:"client_name"  db:"organization_name"`
}
