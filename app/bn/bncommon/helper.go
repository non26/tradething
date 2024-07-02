package bncommon

import (
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
