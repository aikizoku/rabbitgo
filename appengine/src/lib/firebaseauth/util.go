package firebaseauth

import (
	"strings"
)

const (
	headerPrefix      string = "BEARER"
	debugHeaderPrefix string = "user="
)

func getTokenByAuthHeader(ah string) string {
	pLen := len(headerPrefix)
	if len(ah) > pLen && strings.ToUpper(ah[0:pLen]) == headerPrefix {
		return ah[pLen+1:]
	}
	return ""
}

func getUserByAuthHeader(ah string) string {
	if strings.HasPrefix(ah, debugHeaderPrefix) {
		return ah[len(debugHeaderPrefix):]
	}
	return ""
}
