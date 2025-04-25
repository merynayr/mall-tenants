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

// PremisePolygon представляет собой модель данных для полигонов на плане здания
type PremisePolygon struct {
	PremiseCode int64  `json:"premiseCode"`
	Floor       int64  `json:"floor"`
	Points      string `json:"points"`
	Label       string `json:"label"`
}

// Polygons представляет собой модель данных для полигонов на плане здания
type Polygons struct {
	PremiseCode int64  `json:"premiseCode" db:"premise_code"`
	Floor       int64  `json:"floor" db:"floor"`
	Points      string `json:"points" db:"points"`
	Label       string `json:"label" db:"label"`
	Status      string `json:"status" db:"status"`
}

// FloorPlan представляет собой модель данных плана здания
type FloorPlan struct {
	ID       int64  `db:"id"`
	Floor    int64  `db:"floor"`
	Filename string `db:"filename"`
}
