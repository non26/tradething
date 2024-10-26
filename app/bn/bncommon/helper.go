package bncommon

import (
	"fmt"
	"net/url"
	"reflect"
	"time"
)

func GetQueryStringFromStructType[T any](m *T) url.Values {
	st := reflect.TypeOf(m).Elem()
	v := reflect.ValueOf(m).Elem()
	q := url.Values{}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i).Tag.Get("json")
		value := v.FieldByIndex([]int{i}).String()
		q.Add(field, value)
	}
	return q
}

func GetTimeStamp() int64 {
	return time.Now().Unix() * 1000
}

func GetStructTagValueByIndex(st reflect.Type, tag string, index int) string {
	return st.FieldByIndex([]int{index}).Tag.Get(tag)
}

func GetStructTagValueByField(st reflect.Type, field string) (string, error) {
	_field, found := st.FieldByName(field)
	if !found {
		return "", fmt.Errorf("field not found under %s field", field)
	}
	return string(_field.Name), nil
}
