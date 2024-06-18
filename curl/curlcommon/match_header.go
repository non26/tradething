package curlcommon

import "fmt"

func MatchHeader(header string, value string) string {
	return fmt.Sprintf("%v:%v", header, value)
}
