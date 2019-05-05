package utils

import (
	"strings"
)

func StringSliceContains(slice []string, str string) bool {
	for _, strInSlice := range slice {
		if strings.TrimSpace(strInSlice) == strings.TrimSpace(str) {
			return true
		}
	}

	return false
}

func StringSliceContainsIgnoreCase(slice []string, str string) bool {
	for _, strInSlice := range slice {
		if strings.ToLower(strInSlice) == strings.ToLower(str) {
			return true
		}
	}

	return false
}
