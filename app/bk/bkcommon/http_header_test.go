package bkcommon_test

import (
	"net/http"
	"testing"
	"tradething/app/bk/bkcommon"

	"github.com/stretchr/testify/assert"
)

func TestCreateSignUrlMethod(t *testing.T) {
	t.Run("Create sign url for POST method", func(t *testing.T) {
		method := http.MethodPost
		var tStp int64 = 1699376552354
		urlPath := "/api/v3/market/place-bid"
		bodyJson := `{"sym":"thb_btc","amt": 1000,"rat": 10,"typ": "limit"}`
		sk := "test"
		param := make(map[string]string)

		actual := bkcommon.CreateSignUrlMethod(
			method,
			tStp,
			urlPath,
			[]byte(bodyJson),
			sk,
			param,
		)

		assert.NotEmpty(t, actual)
	})
}
