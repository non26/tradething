package bncommon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
)

func Sign(payload []byte, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

func CreateBinanceSignature(data *url.Values, binanceSecretKey string) string {
	payload := data.Encode()
	encodeString := Sign([]byte(payload), []byte(binanceSecretKey))
	encodeData := fmt.Sprintf("%v&signature=%v", data.Encode(), encodeString)
	return encodeData
}
