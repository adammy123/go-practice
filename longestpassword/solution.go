package longestpassword

import (
	"regexp"
	"strings"
	"unicode"
)

const ANTI_PATTERN = "[^A-Za-z0-9]+"

func Solution(S string) int {
	max := -1
	for _, word := range strings.Fields(S) {
		if isValidPassword(word) && len(word) > max {
			max = len(word)
		}
	}
	return max
}

func isValidPassword(pwd string) bool {
	if match, _ := regexp.MatchString(ANTI_PATTERN, pwd); match {
		return false
	}
	matchLetter := true
	matchDigit := false
	for _, char := range pwd {
		if unicode.IsDigit(char) {
			matchDigit = !matchDigit
		} else {
			matchLetter = !matchLetter
		}
	}
	return matchDigit && matchLetter
}
