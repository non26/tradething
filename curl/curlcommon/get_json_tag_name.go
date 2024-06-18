package curlcommon

import (
	"reflect"
)

func GetJsonTag[T any](m T) []string {
	st := reflect.TypeOf(m).Elem()
	fields := []string{}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i).Tag.Get("json")
		fields = append(fields, field)
	}
	return fields
}
