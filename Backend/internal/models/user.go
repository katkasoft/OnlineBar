package models

type User struct {
	ID       string  `json:"ID"`
	Name     string  `json:"Name"`
	Password string  `json:"Password"`
	Email    string  `json:"Email"`
	OS       string  `json:"OS"`
	Balance  float64 `json:"Balance"`
}
