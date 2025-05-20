package model

import "time"

// Payment представляет собой модель данных для платежей
type Payment struct {
	ID          int64      `json:"id"`
	RentalID    int64      `json:"rental_id"`
	PeriodStart time.Time  `json:"period_start"`
	PeriodEnd   time.Time  `json:"period_end"`
	Amount      int        `json:"amount"`
	IsPaid      bool       `json:"is_paid"`
	PaymentDate *time.Time `json:"payment_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// PaymentFilter - структура фильтрации для платежей
type PaymentFilter struct {
	IsPaid *bool
	Limit  uint64
	Offset uint64
}

// MarkPaymentsPaidRequest - структура
type MarkPaymentsPaidRequest struct {
	PaymentIDs []int64 `json:"payment_ids" swaggertype:"array,integer"`
}
