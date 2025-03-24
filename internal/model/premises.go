package model

// PremisesType определяет типы помещений
type PremisesType string

const (
	// Retail Торговое помещение
	Retail PremisesType = "retail"
	// Service Служебное помещение
	Service PremisesType = "service"
)

// PremisesStatus определяет статус помещения
type PremisesStatus string

const (
	// Available помещение доступно
	Available PremisesStatus = "available"
	// Occupied помещение занято
	Occupied PremisesStatus = "occupied"
	// Maintenance помещение на обслуживании
	Maintenance PremisesStatus = "maintenance"
)

// Premises представляет собой модель данных для помещений
type Premises struct {
	Code            int64        `json:"code"`
	Floor           int64        `json:"floor"`
	Area            int64        `json:"area"`
	Type            PremisesType `json:"type"`
	RentPerMonth    int64        `json:"rent_per_month"`
	Status          string       `json:"status"`
	SecuritySystem  string       `json:"security_system"`
	AirConditioning string       `json:"air_conditioning"`
}
