package paths

import (
	"fmt"
	"html/template"
	"io"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"models"
	"net/http"
	"os"
	"queries"
	"strings"
)

type Search struct {
	a string
	b string
	c string
	d string
}

func GetDashboard(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Devices []models.Device
		Devices = queries.GetDevices(DB)
		htmxFile, err := os.ReadFile("../static/modules/charts.html")
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
func SearchDevices(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Devices []models.Device
		name := r.FormValue("name")
		ip := r.FormValue("ip_addr")
		model := r.FormValue("model")
		mac := r.FormValue("mac_addr")
		search := models.Search{name, ip, model, mac}
		Devices = queries.SearchDevices(DB, search)
		htmxFile, err := os.ReadFile("../static/modules/devices.html")
		if err != nil {
			fmt.Printf("error reading file %v", err)
		}
		//num := rand.Int()
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
		htmxFile, err := os.ReadFile("../static/modules/forms/device_added.html")
		if err != nil {
			fmt.Printf("error reading file %v", err)
		}
		device := models.Device{uuid.New().String(), r.FormValue("name"), r.FormValue("ip_addr"), r.FormValue("model"), r.FormValue("mac_addr")}
		err_a := queries.CreateDevice(DB, device)
		if err_a != nil {
			io.WriteString(w, "<div>error with query</div>")
		}
		io.WriteString(w, string(htmxFile))
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "")
}
func GetRuleCreator(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/modules/rule_creation.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
func GetDeviceForm(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/modules/forms/device_form.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
func GetSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/modules/search.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
