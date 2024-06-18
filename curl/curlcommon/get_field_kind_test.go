package curlcommon_test

import (
	"testing"
	"tradething/curl/curlcommon"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	a string
	b int
	c float64
	d bool
}

func TestGetFieldKind(t *testing.T) {
	t.Run("get struct field type", func(t *testing.T) {
		ts := testStruct{}
		expected := []string{"string", "int", "float64", "bool"}

		actual, err := curlcommon.GetFieldKind(&ts)
		if err != nil {
			t.Error(err.Error())
			return
		}
		assert.Equal(t, expected[0], actual[0].String())
		assert.Equal(t, expected[1], actual[1].String())
		assert.Equal(t, expected[2], actual[2].String())
		assert.Equal(t, expected[3], actual[3].String())

	})
}
