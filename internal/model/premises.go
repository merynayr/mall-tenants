package model

import "fmt"

// PremisesType определяет типы помещений
type PremisesType string

const (
	// Retail Торговое помещение
	Retail PremisesType = "retail"
	// Service Служебное помещение
	Service PremisesType = "service"
)

// Premises представляет собой модель данных для помещений
type Premises struct {
	Code            int          `json:"code"`
	Floor           int          `json:"floor"`
	Area            int          `json:"area"`
	Type            PremisesType `json:"type"`
	RentPerMonth    int          `json:"rent_per_month"`
	Status          string       `json:"status"`
	SecuritySystem  string       `json:"security_system"`
	AirConditioning string       `json:"air_conditioning"`
}

// Validate проверяет, что тип помещения допустимый
func (p *Premises) Validate() error {
	if p.Type != Retail && p.Type != Service {
		return fmt.Errorf("invalid premises type: %s", p.Type)
	}
	return nil
}
