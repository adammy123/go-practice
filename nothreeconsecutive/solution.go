package solution

import "strings"

// assume there is a solution
// A, B = [0..100]
func Solution(A int, B int) string {
	if B > A {
		return makeString("b", B, "a", A)
	}
	return makeString("a", A, "b", B)
}

// assume A >= B
// divide A into groups of 2
// e.g. aa aa aa aa a
// slot Bs into each gap left to right until finish
func makeString(a string, A int, b string, B int) string {
	numA := A/2
	arrA := []string{}
	for i:=0; i<numA; i++ {
		arrA = append(arrA, a+a)
	}
	if A%2>0 {
		arrA = append(arrA, a)
	}

	arrB := []string{}
	countB := 0
	for i:=0; i<len(arrA); i++ {
		if countB == B {
			break
		}
		arrB = append(arrB, b)
		countB += 1
	}

	for i:=0; i<len(arrB); i++ {
		if countB == B {
			break
		}
		arrB[i] += b
		countB += 1
	}

	// form string, take from arrA then arrB
	result := ""
	for i, element := range arrA {
		result += element
		if i < len(arrB) {
			result += arrB[i]
		}
	}

	return result
}


func isStringValid(res string, A, B int) bool {
	// there are A 'a's
	if A != strings.Count(res, "a") {
		return false
	}

	// there are B 'b's
	if B != strings.Count(res, "b") {
		return false
	}

	// there are no occurrences of 'aaa's or 'bbb's
	if strings.Contains(res, "aaa") || strings.Contains(res, "bbb") {
		return false
	}

	return true
}

/*
assume A>=B
A:6, B:2 
aa b aa b aa
A: 7, B:2 [impossible]
A: 10, B: 4
aa b aa b aa b aa b aa 
A: 10, B: 9
aa bb aa bb aa bb aa b aa

if 2B <
*/
