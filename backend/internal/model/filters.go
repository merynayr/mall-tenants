package model

// PaymentFilter - структура фильтрации для платежей
type PaymentFilter struct {
	IsPaid    *bool
	Search    interface{}
	SortBy    string
	SortOrder string
	Limit     uint64
	Offset    uint64
}

// RentFilter - структура фильтрации для аренд
type RentFilter struct {
	Search    interface{}
	SortBy    string
	SortOrder string
	Limit     uint64
	Offset    uint64
}
