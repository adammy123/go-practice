package duplicate

import "strings"

type element struct {
	char  string
	count int
}

// given string input, remove k adjacent duplicates recursively
// e.g. k=2, assassinate -> aassinate -> ssinate -> inate
// e.g. k=2, mississippi -> m
func RemoveDuplicate(input string, k int) string {
	stack := []element{}
	for _, currRune := range input {
		currChar := string(currRune)
		if len(stack) > 0 && strings.EqualFold(stack[len(stack)-1].char, currChar) {
			stack[len(stack)-1].count += 1
			if stack[len(stack)-1].count == k {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, element{char: currChar, count: 1})
		}
	}

	return convertStackToString(stack)
}

func convertStackToString(stack []element) string {
	res := ""
	for _, e := range stack {
		for i := 0; i < e.count; i++ {
			res += e.char
		}
	}
	return res
}
