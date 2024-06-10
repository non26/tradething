package bncommon

import (
	"net/url"
	"reflect"
)

func GetQueryStringFromStructType[T any](m *T) (q *url.Values) {
	st := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i).Tag.Get("json")
		value := v.FieldByIndex([]int{i}).String()
		q.Add(field, value)
	}
	return q
}