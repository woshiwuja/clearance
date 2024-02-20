package main

import (
	//"anarchy2036/queries"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"paths"
)

const SRV_PORT string = "8190"
const SRV_STRING string = "http://localhost:" + SRV_PORT
const DB_NAME string = "postgres"
const DB_PASS string = "postgres"
const DB_HOST string = "localhost"
const DB_PORT string = "5432"
const DB_USER string = "postgres"

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	http.Handle("/", http.FileServer(http.Dir("../static")))
	http.HandleFunc("/rule_creation", paths.GetRuleCreator)
	http.HandleFunc("/dashboard", paths.GetDashboard(db))
	http.HandleFunc("/regex_form", paths.GetRegexForm)
	http.HandleFunc("/device_form", paths.GetDeviceForm)
	http.HandleFunc("/devices", paths.GetDevices(db))
	http.HandleFunc("/device_name", paths.SearchDevices(db))
	http.HandleFunc("/search", paths.GetSearch)
	http.HandleFunc("/new_device", paths.PostNewDevice(db))
	http.HandleFunc("/delete", paths.Delete)
	fmt.Printf("server open on %s", SRV_PORT)

	err_a := http.ListenAndServe(":8190", nil)
	if err_a != nil {
		fmt.Printf("cant open server %v", err_a)
	}
}
