package curlcommon

import (
	"fmt"
	"reflect"
)

func GetFieldKind(m interface{}) ([]reflect.Kind, error) {
	kinds := []reflect.Kind{}
	v := reflect.ValueOf(m).Elem()
	if !v.CanAddr() {
		return nil, fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}
	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		kind := typeField.Type.Kind()
		kinds = append(kinds, kind)
	}
	return kinds, nil
}
