package main

func saveOperation(data Operation) error {
	return nil
}

func eraseOperation(id int) error {
	return nil
}

func loadAllOperations() (Operations, error) {
	return Operations{}, nil
}

func saveAsset(assetName string) error {
	return nil
}

func alterAsset(assetName string, amount int) error {
	return nil
}

func loadAsset(assetName string) (Assets, error) {
	return Assets{}, nil
}

func loadAssets() ([]Assets, error) {
	return []Assets{}, nil
}

func deleteAsset(assetName string) error {
	return nil
}
