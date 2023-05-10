package util

import (
	"strings"
)

func TrimStr(str string) string {
	s := strings.Trim(str, "\n")
	i := strings.Index(s, "]")
	s = strings.Trim(s[i+1:], " ")
	return s
}

func StrToArray(str string) []string {
	return strings.Split(TrimStr(str), "\n")
}
