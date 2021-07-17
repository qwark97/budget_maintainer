package model

import (
	"gorm.io/gorm"
	"time"
)

type baseMoneyStruct struct {
	gorm.Model
	Category string     `json:"category"`
	Amount   float64    `json:"amount"`
	Year     int        `json:"year"`
	Month    time.Month `json:"month"`
}

type Operation baseMoneyStruct

type BudgetPosition struct {
	baseMoneyStruct
	Group string `json:"group"`
}

type Assets struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique"`
	Amount int    `json:"amount"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"unique"`
}
