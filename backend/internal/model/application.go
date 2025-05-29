package model

import "time"

// Application - модель заявок пользователей
type Application struct {
	ID               int64     `json:"id,omitempty" db:"id"`
	OrganizationName string    `json:"organizationName" db:"organization_name"`
	ContactPerson    string    `json:"contactPerson" db:"contact_person"`
	Address          string    `json:"address" db:"address"`
	Phone            string    `json:"phone" db:"phone"`
	Requisites       string    `json:"requisites" db:"requisites"`
	Email            string    `json:"email" db:"email"`
	PremiseNumber    int64     `json:"premiseNumber" db:"premise_number"`
	StartDate        time.Time `json:"startDate" db:"start_date"`
	EndDate          time.Time `json:"endDate" db:"end_date"`
	AdditionalInfo   *string   `json:"additionalInfo,omitempty" db:"additional_info"`
	IsProcessed      bool      `json:"isProcessed" db:"is_processed"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
}

// ApplicationFilter - структура фильтрации для заявок
type ApplicationFilter struct {
	IsProcessed *bool
	Limit       uint64
	Offset      uint64
}
