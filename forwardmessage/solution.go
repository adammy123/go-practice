package solution

func Solution(S string, A []int) string {
	// initialize result
	result := ""

	// first letter is always the first letter of S
	result += string(S[0])

	// the next person to pass the message to
	nextIdx := A[0]

	// pass the message until the 0th person receives the message
	for {
		if nextIdx == 0 {
			break
		}
		result += string(S[nextIdx])
		nextIdx = A[nextIdx]
	}

    return result
}
