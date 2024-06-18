package curlcommon

import (
	"fmt"
	"reflect"
	"strings"
)

// func main() {
// 	w := Wham{
// 		Username: "maria",
// 		Password: "hunter2",
// 		ID:       42,
// 		Homebase: "2434 Main St",
// 	}
// 	fmt.Printf("%+v\n", w)
// 	SetField(&w, "username", "larry")
// 	fmt.Printf("%+v\n", w)
// }
/*
reference:
https://gist.github.com/lelandbatey/a5c957b537bed39d1d6fb202c3b8de06
*/
func SetField(item interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(item).Elem()
	if !v.CanAddr() {
		return fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}
	// It's possible we can cache this, which is why precompute all these ahead of time.
	findJsonName := func(t reflect.StructTag) (string, error) {
		if jt, ok := t.Lookup("json"); ok {
			return strings.Split(jt, ",")[0], nil
		}
		return "", fmt.Errorf("tag %v provided does not define a json tag", fieldName)
	}
	fieldNames := map[string]int{}
	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		jname, _ := findJsonName(tag)
		fieldNames[jname] = i
	}

	fieldNum, ok := fieldNames[fieldName]
	if !ok {
		return fmt.Errorf("field %s does not exist within the provided item", fieldName)
	}
	fieldVal := v.Field(fieldNum)
	fieldVal.Set(reflect.ValueOf(value))
	return nil
}
