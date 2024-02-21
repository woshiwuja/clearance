package paths

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"html/template"
	"io"
	"net/http"
	"os"
	"sql"
	"strings"
)

type Search struct {
	a pgtype.UUID
	b string
	c string
	d string
	e string
}

/*
	func GetDashboard() http.HandlerFunc {
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
*/
func GetRegexForm(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/GetDashboard request received\n")
	htmxFile, err := os.ReadFile("../static/modules/forms/regex.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
func GetDevices(DB *sql.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Devices, err := DB.GetDevices(ctx)
		if err != nil {
			fmt.Println(err)
		}
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
func SearchDevices(DB *sql.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := "id"
		name := r.FormValue("name")
		ip := r.FormValue("ip_addr")
		model := r.FormValue("model")
		mac := "%" + r.FormValue("mac_addr") + "%"
		var uuid [16]byte
		copy(uuid[:], []byte(id))
		search := sql.SearchDevicesParams{name, model, ip, mac}
		Devices, err := DB.SearchDevices(ctx, search)
		if err != nil {
			fmt.Println(err)
		}
		for _ = range Devices {
			fmt.Println(Devices)
		}
		htmxFile, err := os.ReadFile("../static/modules/devices.html")
		if err != nil {
			fmt.Printf("error reading file %v", err)
		}
		//num := rand.Int()
		htmlTemplate := string(htmxFile)
		template := template.Must(template.New("").Parse(htmlTemplate))
		builder := &strings.Builder{}
		template.Execute(builder, &Devices)
		s := builder.String()
		io.WriteString(w, s)
	}
} /*
func PostNewDevice(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		htmxFile, err := os.ReadFile("../static/modules/forms/device_added.html")
		if err != nil {
			fmt.Printf("error reading file %v", err)
		}
		device := models.Device{uuid.New().String(), pgtype.Text{"name"), r.FormValue("ip_addr"), r.FormValue("model"), r.FormValue("mac_addr")}
		err_a := queries.CreateDevice(DB, device)
		if err_a != nil {
			io.WriteString(w, "<div>error with query</div>")
		}
		io.WriteString(w, string(htmxFile))
	}
}*/
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
