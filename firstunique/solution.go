package firstunique

func Solution(A []int) int {
	numCount := map[int]int{}
	for _, num := range A {
		if _, ok := numCount[num]; !ok {
			numCount[num] = 0
		}
		numCount[num] += 1
	}
	for _, num := range A {
		if numCount[num] == 1 {
			return num
		}
	}
	return -1
}