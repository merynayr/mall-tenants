package model

import "time"

// Contract модель данных таблицы договоров
type Contract struct {
	ContractID int64      `json:"contract_id" db:"contract_id"`
	RentalID   int64      `json:"rental_id" db:"rental_id"`
	FilePath   string     `json:"file_path" db:"file_path"`
	Signature  []byte     `json:"signature" db:"signature"`
	PublicKey  string     `json:"public_key" db:"public_key"`
	IsActive   bool       `json:"is_active" db:"is_active"`
	IsSigned   bool       `json:"is_signed" db:"is_signed"`
	SignedAt   *time.Time `json:"signed_at,omitempty" db:"signed_at"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
}

// Contracts - модель для получения данных о договорах из БД
type Contracts struct {
	ContractID int64      `json:"contract_id" db:"contract_id"`
	RentalID   int64      `json:"rental_id" db:"rental_id"`
	FilePath   string     `json:"file_path" db:"file_path"`
	IsActive   bool       `json:"is_active" db:"is_active"`
	IsSigned   bool       `json:"is_signed" db:"is_signed"`
	SignedAt   *time.Time `json:"signed_at,omitempty" db:"signed_at"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
}
