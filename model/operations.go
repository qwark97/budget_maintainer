package model

import (
	"encoding/json"
	"io"
)

func NewOperation(body io.ReadCloser) (operation, error) {
	var data operation
	err := json.NewDecoder(body).Decode(&data)
	return data, err
}

func SaveOperation(data operation) error {
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
