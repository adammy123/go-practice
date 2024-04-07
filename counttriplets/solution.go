package solution

func countTriplets(arr []int64, r int64) int64 {
	inputMap := createInputMap(arr)
	

	return 0

}

func createInputMap(arr []int64) map[int64][]int {
	inputMap := map[int64][]int{}
	for idx, num := range arr {
		if _, ok := inputMap[num]; !ok {
			inputMap[num] = []int{}
		}
		inputMap[num] = append(inputMap[num], idx)
	}
	return inputMap
}