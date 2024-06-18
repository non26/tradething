package curlcommon_test

import (
	"testing"
	"tradething/curl/curlcommon"

	"github.com/stretchr/testify/assert"
)

type Wham struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	ID       int64  `json:"_id"`
	Homebase string `json:"homebase"`
}

func TestSetFieldByTagName(t *testing.T) {
	t.Run("set field in struct", func(t *testing.T) {
		w := Wham{
			Username: "non",
			Password: "non2",
			ID:       1234,
			Homebase: "non3",
		}
		expected := Wham{
			Username: "chanon",
			Password: "non2",
			ID:       1234,
			Homebase: "non3",
		}

		err := curlcommon.SetField(&w, "username", "chanon")

		assert.Nil(t, err)
		assert.Equal(t, expected.Username, "chanon")
		assert.Equal(t, expected.Password, w.Password)
		assert.Equal(t, expected.ID, w.ID)
		assert.Equal(t, expected.Homebase, w.Homebase)
	})
}
