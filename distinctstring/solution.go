package distinctstring

import (
	"fmt"
)

func Solution(P, Q string) int {
	return minInt(calculateDistinct(replace(P, Q)), calculateDistinct(replace(Q, P)))
}

func replace(P, Q string) string {
	var currDistinct int
	var newP string
	for i, char := range Q {
		// if same, also replace
		currDistinct = calculateDistinct(P)
		newP = P[:i] + string(char) + P[i+1:]
		newDistinct := calculateDistinct(newP)
		if newDistinct <= currDistinct {
			P = newP
		}
	}
	fmt.Println("P:", newP)
	return P
}

func calculateDistinct(s string) int {
	set := map[rune]bool{}
	for _, char := range s {
		set[char] = true
	}
	fmt.Printf("Num distinct in %s: %d\n", s, len(set))
	return len(set)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}