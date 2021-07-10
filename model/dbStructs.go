package model

import "gorm.io/gorm"

type Operation struct {
	gorm.Model
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}

type Assets struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique"`
	Amount int    `json:"amount"`
}
