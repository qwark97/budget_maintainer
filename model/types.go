package model

type Operation struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`

	ID        int
	Timestamp int
}
type Operations []Operation

type Assets struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}
