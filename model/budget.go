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

func SetBudget(data transitBudget) error { return nil }

func LoadBudget(year, month string) (transitBudget, error) { return transitBudget{}, nil }
