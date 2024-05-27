package types

import (
	"reflect"
	"strings"
)

func ReflectColumnFields(s interface{}) string {
	r := reflect.TypeOf(s)
	var columns []string

	for i := 0; i < r.NumField(); i++ {
		if tag, ok := r.Field(i).Tag.Lookup("db"); ok {
			if tag == "-" {
				continue
			}
			columns = append(columns, tag)
		} else {
			columns = append(columns, r.Field(i).Name)
		}
	}
	return strings.Join(columns, ",")
}
