package api

// import (
// 	"reflect"
// )

// type IFieldType interface {
// 	string | int | int16 | int32 | int64 | float32 | float64
// }

// type RequestField[T IFieldType] struct {
// 	field         string
// 	is_user_input bool
// 	// description   string
// 	field_value T
// }

// func (r *RequestField[IFieldType]) GetUserInputField() string {
// 	if r.is_user_input {
// 		return r.field
// 	}
// 	return ""
// }

// func (r *RequestField[IFieldType]) SetField(field string) {
// 	r.field = field
// }

// func (r *RequestField[IFieldType]) IsUserInput(is bool) {
// 	r.is_user_input = is
// }

// func (r *RequestField[IFieldType]) SetValue(v IFieldType) {
// 	r.field_value = v
// }

// func (r *RequestField[IFieldType]) ConvertStringToNumber(source string) IFieldType {
// 	var x IFieldType
// 	switch reflect.TypeOf(x).Kind() {
// 	case reflect.Int:
// 	case reflect.Float64:
// 		return
// 	case reflect.String:
// 	}
// 	return x
// }
