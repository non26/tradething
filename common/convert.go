package common

import (
	"encoding/json"
	"fmt"
	"strings"
)

func JsonToString(js []byte) string {
	jsonString := string(js)
	if jsonString == `{}` {
		return ""
	}
	return jsonString
}

func ParamToQueryString(params map[string]string) string {
	query := make([]string, 0)
	for param, value := range params {
		queryFormat := fmt.Sprintf("%v=%v", param, value)
		query = append(query, queryFormat)
	}
	return "?" + strings.Join(query, "&")
}

func ToJson[T any](body T) []byte {
	jsonBody, _ := json.Marshal(body)
	return jsonBody
}

func JsonToMapString(jsonObj []byte) map[string]string {
	m := make(map[string]string)
	json.Unmarshal(jsonObj, &m)
	return m
}
