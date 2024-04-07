package solution

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}
	forwardMaxDepth := getMaxDepth(A)
	backwardMaxDepth := getMaxDepth(reverseSlice(A))
	return max(forwardMaxDepth, backwardMaxDepth)
}

func getMaxDepth(A []int) int {
	maxDepth, tmpDepth := 0, 0
	leftHeight := A[0]
	var currHeight int

	for i:=1; i<len(A); i++ {
		currHeight = A[i]
		if currHeight >= leftHeight {
			leftHeight = currHeight
			maxDepth = max(tmpDepth, maxDepth)
			tmpDepth = 0
			continue
		}
		tmpDepth = max(tmpDepth, leftHeight-currHeight)
	}
	
	return maxDepth
}
func reverseSlice(A []int) []int {
	newSlice := make([]int, len(A))
	for i, element := range A {
		newSlice[len(A)-i-1] = element
	}
	return newSlice
}