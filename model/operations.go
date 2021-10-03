package model

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
)

func NewOperation(body io.ReadCloser) (operation, error) {
	var data operation
	err := json.NewDecoder(body).Decode(&data)
	return data, err
}

func SaveOperation(data operation) error {
	data.Amount = math.Round(data.Amount*100) / 100
	if data.Amount <= 0 {
		return fmt.Errorf("operation's amount must be higher or equal than 0.01")
	}
	res := DBConn.Create(&data)
	return res.Error
}

func DeleteOperation(id int) error {
	res := DBConn.Delete(&operation{}, id)
	return res.Error
}

func LoadAllOperations() (operations, error) {
	var operations operations
	res := DBConn.Find(&operations.Elements)
	return operations, res.Error
}
