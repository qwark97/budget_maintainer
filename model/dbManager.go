package model

func SaveOperation(data Operation) error {
	return nil
}

func DeleteOperation(id int) error {
	return nil
}

func LoadAllOperations() (Operations, error) {
	return Operations{}, nil
}

func SaveAsset(assetName string) error {
	return nil
}

func AlterAsset(assetName string, amount int) error {
	return nil
}

func LoadAsset(assetName string) (Assets, error) {
	return Assets{}, nil
}

func LoadAssets() ([]Assets, error) {
	return []Assets{}, nil
}

func DeleteAsset(assetName string) error {
	return nil
}
