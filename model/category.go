package model

import "fmt"

func SaveCategory(categoryName string) error {
	res := DBConn.Create(&category{Name: categoryName})
	return res.Error
}

func DeleteCategory(categoryName string) error {
	var c category
	if fetchRes := DBConn.Where("name = ?", categoryName).Find(&c); fetchRes.Error != nil {
		return fetchRes.Error
	}
	if c.Name == "" {
		return fmt.Errorf("did not find such category")
	}
	res := DBConn.Unscoped().Delete(&c)
	return res.Error
}

func LoadCategory(categoryName string) (category, error) {
	var c category
	if fetchRes := DBConn.Where("name = ?", categoryName).Find(&c); fetchRes.Error != nil {
		return category{}, fetchRes.Error
	}
	if c.Name == "" {
		return category{}, fmt.Errorf("did not find such category")
	}
	return c, nil
}

func LoadAllCategories() (categories, error) {
	var operations categories
	res := DBConn.Find(&operations)
	return operations, res.Error
}
