package solution

import "fmt"

// https://www.hackerrank.com/challenges/special-palindrome-again/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings
// incomplete solution.

func substrCount(n int32, s string) int64 {
	result := 0
	fmt.Println("n: ", n)
	fmt.Println("s: ", s)

	if n < 2 {
		return 1
	}

	currCount := 1
	targetByte := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == targetByte {
			currCount++
		} else {
			// check palindrome
			// true: ++ sum and shift i
			// false: triangleNum(currCount)
			if i+currCount < len(s) && stringContainsOnlyTarget(s[i+1:i+1+currCount], targetByte) {
				fmt.Println("result += ", currCount*2+1)
				result += (currCount*2 + 1)
				if currCount > 1 {
					targetByte = s[i]
					i += currCount + 1
				} else {
					i += currCount
				}
			} else {
				fmt.Println("result += ", trianglularNum(currCount))
				result += trianglularNum(currCount)
				targetByte = s[i]
			}
			currCount = 1
		}
	}

	return int64(result)
}

func stringContainsOnlyTarget(str string, target byte) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != target {
			return false
		}
	}
	return true
}

// returns number of substrings from aabaa where n length of string
func palindromeNum(n int) int {
	if n == 1 {
		return n
	}
	if n%2 == 0 {
		fmt.Errorf("even number %d was passed into palindromNum", n)
	}
	return (n/2 + 1) + trianglularNum(n/2)
}

func trianglularNum(n int) int {
	return n * (n + 1) / 2
}
