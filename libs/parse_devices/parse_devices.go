package parse_devices

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"models"
	"os"
	"queries"
	"reflect"
	"sort"
	"strings"
)

func NewDevice() {
	queries.NewDevice()
}

func InitDevices() {
	f := "devices.toml"

	var devices models.Device
	meta, err := toml.DecodeFile(f, &devices)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	indent := strings.Repeat(" ", 14)

	fmt.Print("Decoded")
	typ, val := reflect.TypeOf(devices), reflect.ValueOf(devices)
	for i := 0; i < typ.NumField(); i++ {
		indent := indent
		if i == 0 {
			indent = strings.Repeat(" ", 7)
		}
		fmt.Printf("%s%-11s → %v\n", indent, typ.Field(i).Name, val.Field(i).Interface())
	}

	fmt.Print("\nKeys")
	keys := meta.Keys()
	sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
	for i, k := range keys {
		indent := indent
		if i == 0 {
			indent = strings.Repeat(" ", 10)
		}
		fmt.Printf("%s%-10s %s\n", indent, meta.Type(k...), k)
	}

}
