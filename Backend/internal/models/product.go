package models

import "time"

type Product struct {
	Data     time.Time
	Name     string  `json:"name"`
	Cost     float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

type ProductList struct {
	Products []Product `json:"products"`
}
