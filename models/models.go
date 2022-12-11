package models

type Customer struct {
	ID        float64 `json:"id"`
	Name      string  `json:"name"`
	Role      string  `json:"role"`
	Email     string  `json:"emil"`
	Phone     string  `json:"phone"`
	Contacted bool    `json:"contacted"`
}
