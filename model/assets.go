package model

import "fmt"

func SaveAsset(assetName string) error {
	asset := Assets{
		Name:   assetName,
		Amount: 0,
	}
	res := DBConn.Create(&asset)
	return res.Error
}

func AlterAsset(assetName string, amount int) error {
	var asset Assets
	if fetchRes := DBConn.Where("name = ?", assetName).Find(&asset); fetchRes.Error != nil {
		return fetchRes.Error
	}
	if asset.Name == "" {
		return fmt.Errorf("did not find such asset")
	}
	asset.Amount = asset.Amount + amount

	saveRes := DBConn.Save(&asset)
	return saveRes.Error
}

func LoadAsset(assetName string) (Assets, error) {
	var asset Assets
	if fetchRes := DBConn.Where("name = ?", assetName).Find(&asset); fetchRes.Error != nil {
		return Assets{}, fetchRes.Error
	}
	if asset.Name == "" {
		return Assets{}, fmt.Errorf("did not find such asset")
	}
	return asset, nil
}

func LoadAssets() ([]Assets, error) {
	var assets []Assets
	if fetchRes := DBConn.Find(&assets); fetchRes.Error != nil {
		return []Assets{}, fetchRes.Error
	}
	return assets, nil
}

func DeleteAsset(assetName string) error {
	var asset Assets
	if fetchRes := DBConn.Where("name = ?", assetName).Find(&asset); fetchRes.Error != nil {
		return fetchRes.Error
	}
	if asset.Name == "" {
		return fmt.Errorf("did not find such asset")
	}
	res := DBConn.Unscoped().Delete(&asset)
	return res.Error
}
