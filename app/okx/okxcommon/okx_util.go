package okxcommon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

func ToQueryParameter(s interface{}) string {
	structType := reflect.TypeOf(s)
	valueOfStruct := reflect.ValueOf(s)
	queryParam := ""
	numberOfField := structType.NumField()

	for idx := 0; idx < numberOfField; idx++ {
		tag := structType.Field(idx).Tag.Get("json")
		structFieldValue := valueOfStruct.Field(idx)
		if idx == numberOfField-1 {
			queryParam += fmt.Sprintf("%v=%v", tag, structFieldValue)
			break
		}
		queryParam += fmt.Sprintf("%v=%v&", tag, structFieldValue)

	}
	return queryParam
}

func StructToJson(s interface{}) ([]byte, error) {
	j, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func AddInstIdUSDTSWAPPostfix(instId string) string {
	return fmt.Sprintf("%v-USDT-SWAP", instId)
}

func OkxConditionResponseError(httpCode int, okxCode string, okxMsg string) error {
	msg := fmt.Sprintf("http=%v&code=%v&msg=%v", httpCode, okxCode, okxMsg)
	if httpCode != http.StatusOK {
		return errors.New(msg)
	}
	if okxCode == "1" {
		return errors.New(msg)
	}
	return nil
}

func ResponseToStruct[k comparable](s []k, mp []interface{}) ([]k, error) {
	for _, d := range mp {
		m := d.(map[string]interface{})
		e := new(k)
		jsonData, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(jsonData, e)
		if err != nil {
			return nil, err
		}
		s = append(s, *e)
	}

	return s, nil
}
