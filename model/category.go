package model

import "fmt"

func SaveCategory(categoryName string) error {
	res := DBConn.Create(&Category{Name: categoryName})
	return res.Error
}

func DeleteCategory(categoryName string) error {
	var category Category
	if fetchRes := DBConn.Where("name = ?", categoryName).Find(&category); fetchRes.Error != nil {
		return fetchRes.Error
	}
	if category.Name == "" {
		return fmt.Errorf("did not find such category")
	}
	res := DBConn.Unscoped().Delete(&category)
	return res.Error
}

func LoadCategory(categoryName string) (Category, error) {
	var category Category
	if fetchRes := DBConn.Where("name = ?", categoryName).Find(&category); fetchRes.Error != nil {
		return Category{}, fetchRes.Error
	}
	if category.Name == "" {
		return Category{}, fmt.Errorf("did not find such category")
	}
	return category, nil
}

func LoadAllCategories() (Categories, error) {
	var operations Categories
	res := DBConn.Find(&operations)
	return operations, res.Error
}
