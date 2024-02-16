package queries

import (
	"fmt"
	"models"

	"gorm.io/gorm"
)

func CreateDevice(db *gorm.DB, device models.Device) error {
	result := db.Create(&device)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func GetDevices(db *gorm.DB) (Devices []models.Device) {

	db.Raw("SELECT * FROM devices").Scan(&Devices)
	fmt.Println(Devices)
	return Devices
}
func GetDevicesByName(db *gorm.DB, name string) (Devices []models.Device) {
	db.Where("name = ?", name).Scan(&Devices)
	fmt.Println(Devices)
	return Devices
}
