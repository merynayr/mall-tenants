package model

// PaymentFilter - структура фильтрации для платежей
type PaymentFilter struct {
	IsPaid      *bool
	OverdueOnly bool
	Search      any
	SortBy      string
	SortOrder   string
	Limit       uint64
	Offset      uint64
}

// RentFilter - структура фильтрации для аренд
type RentFilter struct {
	Search    any
	SortBy    string
	SortOrder string
	Limit     uint64
	Offset    uint64
}
