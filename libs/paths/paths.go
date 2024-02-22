package paths

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

func GetDashboard(DB *sql.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Devices, err := DB.GetDevices(ctx)
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
		id := "%" + r.FormValue("id") + "%"
		name := "%" + r.FormValue("name") + "%"
		ip := "%" + r.FormValue("ip_addr") + "%"
		model := "%" + r.FormValue("model") + "%"
		mac := "%" + r.FormValue("mac_addr") + "%"
		search := sql.SearchDevicesParams{id, name, model, ip, mac}
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
		template.Execute(builder, Devices)
		s := builder.String()
		io.WriteString(w, s)
	}
}

func PostNewDevice(DB *sql.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		htmxFile, err := os.ReadFile("../static/modules/forms/device_added.html")
		if err != nil {
			fmt.Printf("error reading file %v", err)
		}
		Device := sql.AddDeviceParams{uuid.New().String(), r.FormValue("name"), r.FormValue("ip_addr"), r.FormValue("model"), r.FormValue("mac_addr")}
		id, err := DB.AddDevice(ctx, Device)
		if err != nil {
			io.WriteString(w, "<div>error with query</div>")
		}
		fmt.Println("added id:", id)
		htmlTemplate := string(htmxFile)
		template := template.Must(template.New("").Parse(htmlTemplate))
		builder := &strings.Builder{}
		template.Execute(builder, id)
		s := builder.String()
		io.WriteString(w, s)
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
