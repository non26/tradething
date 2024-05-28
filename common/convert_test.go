package common_test

import (
	"encoding/json"
	"testing"
	"tradetoolv2/common"

	"github.com/stretchr/testify/assert"
)

type TestingStruct struct {
	Sym string  `json:"sym"`
	Amt float64 `json:"amt"`
	Rat float64 `json:"rat"`
	Typ string  `json:"typ"`
}

type EmptyPayload struct{}

func TestJsonToString(t *testing.T) {
	t.Run("Convert json to string", func(t *testing.T) {
		body := TestingStruct{
			Sym: "BTC",
			Amt: 1000,
			Rat: 10,
			Typ: "limit",
		}
		jsonBody, _ := json.Marshal(body)
		expectedResult := `{"sym":"BTC","amt":1000,"rat":10,"typ":"limit"}`

		actual := common.JsonToString(jsonBody)

		assert.Equal(t, actual, expectedResult)
	})

	t.Run("Empty payload", func(t *testing.T) {
		body := EmptyPayload{}
		jsonBody, _ := json.Marshal(body)
		expectedResult := ""

		actual := common.JsonToString(jsonBody)

		assert.Equal(t, actual, expectedResult)
	})
}

func TestParamToQueryString(t *testing.T) {
	t.Run("Convert param to query string", func(t *testing.T) {
		param := map[string]string{}
		param["test1"] = "test1-value"
		param["test2"] = "test2-value"
		expected := []string{
			"?test1=test1-value&test2=test2-value",
			"?test2=test2-value&test1=test1-value",
		}

		actual := common.ParamToQueryString(param)

		assert.Contains(t,
			expected,
			actual,
		)

	})
}
