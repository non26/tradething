package kccommon

import (
	"fmt"
	"time"
)

func CreateFutureClientId(
	product string,
	side string,
	symbol string,
	timefram string,
) string {
	var t time.Time
	var timeformat string
	if timefram == "" {
		t = time.Now()
		timeformat = t.Format("2006-01-02T15:04:05")
	}
	clientId := fmt.Sprintf("%v-%v%v%v", product, side, symbol, timeformat)
	return clientId
}
