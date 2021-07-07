package model

func SaveOperation(data Operation) error {
	res := DBConn.Create(&data)
	return res.Error
}

func DeleteOperation(id int) error {
	res := DBConn.Delete(&Operation{}, id)
	return res.Error
}

func LoadAllOperations() (Operations, error) {
	var operations Operations
	res := DBConn.Find(&operations)
	return operations, res.Error
}
