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
	var validCategories categories
	budgetPositions := manyBudgetPositions{}
	year := data.Year
	month := data.Month
	if categories, err := LoadAllCategories(); err != nil {
		return err
	} else {
		validCategories = categories
	}

	for _, transitPosition := range data.BudgetPositions {
		if !validCategories.isKnownCategory(transitPosition.Category) {
			return UnknownCategoryError
		}
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
