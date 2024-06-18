package curlcommon

import (
	"reflect"
	"strconv"
)

func SetStructField[T any](m T, struct_field string, value string, value_type reflect.Kind) error {
	switch value_type.String() {
	case "string":
		err := SetField(m, struct_field, value)
		if err != nil {
			return err
		}
	case "int":
		vi, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		err = SetField(m, struct_field, vi)
		if err != nil {
			return err
		}
	case "float64":
		f64i, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		err = SetField(m, struct_field, f64i)
		if err != nil {
			return err
		}
	}
	return nil
}
