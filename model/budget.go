package model

import (
	"encoding/json"
	"io"
)

func NewTransitBudget(body io.ReadCloser) (transitBudget, error) {
	var data transitBudget
	err := json.NewDecoder(body).Decode(&data)
	return data, err
}

func SetBudget(data transitBudget) error {
	budgetPositions := manyBudgetPositions{}
	year := data.Year
	month := data.Month

	for _, transitPosition := range data.BudgetPositions {
		p := budgetPosition{
			Category: transitPosition.Category,
			Amount:   transitPosition.Amount,
			Year:     year,
			Month:    month,
			Group:    transitPosition.Group,
		}
		budgetPositions.elements = append(budgetPositions.elements, p)
	}
	budgetPositions.mergeCategories()
	if removeRes := DBConn.
		Where("year = ?", year).
		Where("month = ?", month).
		Unscoped().
		Delete(&budgetPosition{}); removeRes.Error != nil {
		return removeRes.Error
	}
	for _, position := range budgetPositions.elements {
		DBConn.Create(&position)
	}

	return nil
}

func LoadBudget(year, month int) (transitBudget, error) { return transitBudget{}, nil }
