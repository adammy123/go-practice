package solution

func alternatingCharacters(s string) int32 {
	result := 0
	prevCharStr := string(s[0])

	for i:=1; i<len(s); i++ {
		currCharStr := string(s[i])
		if currCharStr == prevCharStr {
			result += 1
		} else {
			prevCharStr = currCharStr
		}
	}

	return int32(result)
}