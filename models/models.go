package models

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"emil"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// type Customers struct {
// 	Customers []customer
// }
