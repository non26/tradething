package curlcommon_test

import (
	"testing"
	"tradething/curl/curlcommon"

	"github.com/stretchr/testify/assert"
)

type TestStructTag struct {
	A string  `json:"a"`
	B string  `json:"b"`
	C float64 `json:"c"`
}

func TestGetJsonTagName(t *testing.T) {
	t.Run("get tag json from struct", func(t *testing.T) {
		tst := TestStructTag{}
		expected := []string{"a", "b", "c"}
		actual := curlcommon.GetJsonTag(tst)

		assert.Equal(t, expected[0], actual[0])
		assert.Equal(t, expected[1], actual[1])
		assert.Equal(t, expected[2], actual[2])
	})
}
