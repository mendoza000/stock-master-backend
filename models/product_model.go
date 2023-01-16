package models

type Product struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Brand   string  `json:"brand"`
	Stock   bool    `json:"stock"`
	Price   float32 `json:"price"`
	Details string  `json:"details"`
	Amount  int     `json:"amount"`
}