package utils

import (
	"strings"
)

func LastSplit(str, sep string) string {
	if str == "" || sep == "" {
		return ""
	}

	splits := strings.Split(str, sep)

	return splits[len(splits)-1]
}
