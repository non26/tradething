package okxcommon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

func GenerateKeyHeader(k string) (key string) {
	return k
}

/*
Sign the prehash string with the SecretKey using the HMAC SHA256.
*/
func GenerateSignHeader(tStp string, method string, requestPath string, body string, sk string) (signBase64 string) {
	queryString := fmt.Sprintf("%v%v", requestPath, body)
	s := fmt.Sprintf("%v%v%v", tStp, strings.ToUpper(method), queryString)
	mac := hmac.New(sha256.New, []byte(sk))
	_, _ = mac.Write([]byte(s))
	signBase64 = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return signBase64
}

/*
time show as UTC
2006-01-02T03:04:05.000Z
tISOWithMiliSec := tStp.UTC().Format("2006-01-02T03:04:05.000Z")
*/
func GenerateTimeHeader() (timeStamp string) {
	// tISOWithMiliSec := time.Now().UTC().Format("2006-01-02T03:04:05.000Z")
	// return tISOWithMiliSec
	utcTime := time.Now().UTC()
	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	return iso
}

// func GenerateTimeHeader() (t time.Time, timeStamp string) {
// 	tn := time.Now().UTC()
// 	tISOWithMiliSec := tn.Format("2006-01-02T03:04:05.000Z")
// 	return tn, tISOWithMiliSec
// }

/*
The passphrase you specified when creating the APIKey.
*/
func GeneratePassPhaseHeader(pp string) (passPhase string) {
	return pp
}

/*
`x-simulated-trading: 1` needs to be added to the header of the Demo Trading request.
*/
func GenerateDemoHeader() (v string) {
	return fmt.Sprintf("%v", 1)
}
