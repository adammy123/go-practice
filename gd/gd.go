package gd

import "strings"

const (
	HEADERS_SEP = ","
)

func Pal(headers string, supported []string) []string {
	result := []string{}
	headerSlice := strings.Split(headers, HEADERS_SEP)
	for _, header := range headerSlice {
		if sliceContains(supported, header) {
			result = append(result, header)
		}
	}

	return result
}

func sliceContains(slice []string, target string) bool {
	for _, element := range slice {
		if strings.EqualFold(element, target) {
			return true
		}
	}
	return false
}