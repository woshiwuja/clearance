package paths

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"html/template"
	"io"
	"models"
	"net/http"
	"os"
	"queries"
	"strings"
)

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/GetDashboard request received\n")
	htmxFile, err := os.ReadFile("../static/modules/charts.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
func GetRegexForm(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/GetDashboard request received\n")
	htmxFile, err := os.ReadFile("../static/modules/forms/regex.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
func GetDevices(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Devices []models.Device
		Devices = queries.GetDevices(DB)
		htmxFile, err := os.ReadFile("../static/modules/devices.html")
		if err != nil {
			fmt.Printf("error reading file %v", err)
		}
		htmlTemplate := string(htmxFile)
		template := template.Must(template.New("hello").Parse(htmlTemplate))
		builder := &strings.Builder{}
		template.Execute(builder, Devices)
		s := builder.String()
		io.WriteString(w, s)
	}
}
func GetDevicesByName(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Devices []models.Device
		name := r.FormValue("name")
		Devices = queries.GetDevicesByName(DB, name)
		htmxFile, err := os.ReadFile("../static/modules/devices.html")
		if err != nil {
			fmt.Printf("error reading file %v", err)
		}
		htmlTemplate := string(htmxFile)
		template := template.Must(template.New("hello").Parse(htmlTemplate))
		builder := &strings.Builder{}
		template.Execute(builder, Devices)
		s := builder.String()
		io.WriteString(w, s)
	}
}
func PostNewDevice(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		htmxFile, err := os.ReadFile("../static/modules/device.html")
		if err != nil {
			fmt.Printf("error reading file %v", err)
		}
		device := models.Device{uuid.New().String(), r.FormValue("name"), r.FormValue("ip_addr"), r.FormValue("model"), r.FormValue("mac_addr")}
		err_a := queries.CreateDevice(DB, device)
		if err_a != nil {
			io.WriteString(w, "<div>error with query")
		}
		io.WriteString(w, string(htmxFile))
	}
}

func GetRuleCreator(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/modules/rule_creation.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
