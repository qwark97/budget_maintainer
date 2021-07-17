package model

import "fmt"

func SaveAsset(assetName string) error {
	asset := assets{
		Name:   assetName,
		Amount: 0,
	}
	res := DBConn.Create(&asset)
	return res.Error
}

func AlterAsset(assetName string, amount int) error {
	var asset assets
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

func LoadAsset(assetName string) (assets, error) {
	var asset assets
	if fetchRes := DBConn.Where("name = ?", assetName).Find(&asset); fetchRes.Error != nil {
		return assets{}, fetchRes.Error
	}
	if asset.Name == "" {
		return assets{}, fmt.Errorf("did not find such asset")
	}
	return asset, nil
}

func LoadAssets() ([]assets, error) {
	var a []assets
	if fetchRes := DBConn.Find(&a); fetchRes.Error != nil {
		return []assets{}, fetchRes.Error
	}
	return a, nil
}

func DeleteAsset(assetName string) error {
	var asset assets
	if fetchRes := DBConn.Where("name = ?", assetName).Find(&asset); fetchRes.Error != nil {
		return fetchRes.Error
	}
	if asset.Name == "" {
		return fmt.Errorf("did not find such asset")
	}
	res := DBConn.Unscoped().Delete(&asset)
	return res.Error
}
