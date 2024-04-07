package vowel

import "strings"

var vowels = []string{"a", "e", "i", "o", "u"}

// A vowel substring is a substring that only consists of vowels
// ('a', 'e', 'i', 'o', and 'u') and has all five vowels present in it.
// Given a string word, return the number of vowel substrings in word.
func CountNumVowelSubstrings(input string) int {
	count := 0

	// start sliding window from the left
	// until fifth last letter
LeftLoop:
	for l := 0; l < len(input)-4; l++ {
		currSubstr := input[l : l+5]
		isCurrSubstrVowelSubstr := false

		// check that substr contains only vowels
		for _, byteChar := range currSubstr {
			char := string(byteChar)
			if !isCharVowel(char) {
				continue LeftLoop
			}
		}

		if isVowelSubstr(currSubstr) {
			count += 1
			isCurrSubstrVowelSubstr = true
		}

	RightLoop:
		for r := l + 5; r < len(input); r++ {
			// check new letter is vowel or not
			newChar := string(input[r])
			if !isCharVowel(newChar) {
				continue LeftLoop
			}

			if isCurrSubstrVowelSubstr {
				count += 1
				continue RightLoop
			}
			currSubstr += newChar
			if isVowelSubstr(currSubstr) {
				count += 1
				isCurrSubstrVowelSubstr = true
			}
		}
	}
	return count
}

func isVowelSubstr(input string) bool {
	if len(input) < 5 {
		return false
	}
	vowelsMap := map[string]bool{}
	for _, byteChar := range input {
		char := string(byteChar)
		if !isCharVowel(char) {
			return false
		}
		vowelsMap[char] = true
	}
	return len(vowelsMap) == 5
}

func isCharVowel(input string) bool {
	return len(input) == 1 && strSliceContains(vowels, input)
}

func strSliceContains(strSlice []string, char string) bool {
	for _, str := range strSlice {
		if strings.EqualFold(str, char) {
			return true
		}
	}
	return false
}
