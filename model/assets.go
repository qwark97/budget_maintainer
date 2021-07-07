package model

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
