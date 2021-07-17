package model

import "time"

type Operations []Operation

type Categories []Category

type TransitBudget struct {
	Year            int                     `json:"year"`
	Month           time.Month              `json:"month"`
	BudgetPositions []transitBudgetPosition `json:"budget_positions"`
}

type transitBudgetPosition struct {
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Group    string  `json:"group"`
}
