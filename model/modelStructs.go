package model

import (
	"time"
)

type operations struct {
	Elements []operation
}

type categories struct {
	Elements []category
}

func (c *categories) isKnownCategory(catToCheck string) bool {
	for _, cat := range c.Elements {
		if cat.Name == catToCheck {
			return true
		}
	}
	return false
}

type transitBudget struct {
	Year            int                     `json:"year"`
	Month           time.Month              `json:"month"`
	BudgetPositions []transitBudgetPosition `json:"budget_positions"`
}

type transitBudgetPosition struct {
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Group    string  `json:"group"`
}

type manyBudgetPositions struct {
	elements []budgetPosition
}

func (mbp *manyBudgetPositions) mergeCategories() {
	if len(mbp.elements) < 1 {
		return
	}
	var tmpElements []*budgetPosition
	var mergedElements []budgetPosition
	var categories = make(map[string]func() func(amount float64) *budgetPosition)

	for i := range mbp.elements {
		if f, ok := categories[mbp.elements[i].Category]; ok {
			f()(mbp.elements[i].Amount)
		} else {
			bp := &mbp.elements[i]
			tmpElements = append(tmpElements, bp)
			categories[mbp.elements[i].Category] = func() func(amount float64) *budgetPosition {
				return func(amount float64) *budgetPosition {
					bp.Amount = bp.Amount + amount
					return bp
				}
			}
			categories[mbp.elements[i].Category]()(0)
		}
	}
	for _, e := range tmpElements {
		mergedElements = append(mergedElements, *e)
	}
	mbp.elements = mergedElements
}
