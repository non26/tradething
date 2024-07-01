package curlcommon

import (
	"reflect"
	"strconv"
	"strings"
)

func SetStructField[T any](m T, struct_field string, value string, value_type reflect.Kind) error {
	switch value_type.String() {
	case "string":
		err := SetField(m, struct_field, strings.ToUpper(value))
		if err != nil {
			return err
		}
	case "int":
		var vi int
		var err error
		if value == "" {
			vi = 0
		} else {
			vi, err = strconv.Atoi(value)
			if err != nil {
				return err
			}
		}
		err = SetField(m, struct_field, vi)
		if err != nil {
			return err
		}
	case "float64":
		var f64i float64
		var err error
		if value == "" {
			f64i = 0
		} else {
			f64i, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
		}
		err = SetField(m, struct_field, f64i)
		if err != nil {
			return err
		}
	}
	return nil
}
