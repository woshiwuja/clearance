package queries

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"models"
	"strconv"
)

func likeHelp(s string) (t string) {
	if s != "" {
		t = "%" + s + "%"
		return t
	}
	return ""
}

func randomize_event() (Events []models.Event) {
	for i := 0; i <= 10; i++ {
		Events[i].Source_addr = strconv.ParseInt(rand.Int()) + strconv.ParseInt(rand.Int())
	}
	return Events
}

func CreateDevice(db *gorm.DB, device models.Device) error {
	result := db.Create(&device)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func GetDevices(db *gorm.DB) (Devices []models.Device) {

	db.Find(&Devices).Limit(10).Scan(&Devices)
	fmt.Println(Devices)
	return Devices
}
func SearchDevices(db *gorm.DB, search models.Search) (Devices []models.Device) {
	db.Where("name ILIKE ? or mac_addr ILIKE ? or ip_addr ILIKE ? or model ILIKE ?", likeHelp(search.Name), likeHelp(search.MAC_addr), likeHelp(search.IP_addr), likeHelp(search.Model)).Limit(10).Find(&Devices)
	fmt.Println("devices:", Devices)
	return Devices
}

func InitRandomEvents(db *gorm.DB, events []models.Event) error {
	events = randomize_event()
	result := db.Create(&events)
	if err != nil {
		return err
	}
	return nil
}
