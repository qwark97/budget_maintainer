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

type operation baseMoneyStruct

type budgetPosition struct {
	baseMoneyStruct
	Group string `json:"group"`
}

type assets struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique"`
	Amount int    `json:"amount"`
}

type category struct {
	gorm.Model
	Name string `gorm:"unique"`
}
