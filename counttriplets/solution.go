package solution

import (
	"sort"
)

func countTriplets(arr []int64, r int64) int64 {
	result := 0
	inputMap := createInputMap(arr)

	if r == 1 {
		for _, idxArr := range inputMap {
			result += binomial(len(idxArr), 3)
		}
		return int64(result)
	}

	// get keys, sort by ascending
	keys := createSortedKeys(inputMap)

	// map of lowest denominator : map of multiplier : arr of indices
	geoMap := createGeoMap(keys, r, inputMap)

	for _, subGeoMap := range geoMap {
		result += getTotalTriplets(subGeoMap)
	}

	return int64(result)
	
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

func createSortedKeys(inputMap map[int64][]int) []int64 {
	keys := make([]int64, 0, len(inputMap))
	for k := range inputMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func createGeoMap(keys []int64, r int64, inputMap map[int64][]int) map[int64]map[int][]int {
	geoMap := map[int64]map[int][]int{}
	for _, key := range keys {
		lowestDenominator, multiplier := getLowestDenominator(key, r, 1)
		if _, ok := geoMap[lowestDenominator]; !ok {
			geoMap[lowestDenominator] = map[int][]int{}
		}
		if _, ok := geoMap[lowestDenominator][multiplier]; !ok {
			geoMap[lowestDenominator][multiplier] = []int{}
		}
		geoMap[lowestDenominator][multiplier] = append(geoMap[lowestDenominator][multiplier], inputMap[key]...)
	}
	return geoMap
}

func getLowestDenominator(val int64, r int64, multiplier int) (int64, int) {
	if r==1 {
		return val, multiplier
	}
	if val % r > 0 {
		return val, multiplier
	}
	multiplier++
	return getLowestDenominator(val/r, r, multiplier)
}

func getTotalTriplets(subGeoMap map[int][]int) int {
	keys := make([]int, 0, len(subGeoMap))
	for k := range subGeoMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	
	result := 0
	for i:=0; i<len(keys)-2; i++ {
		if keys[i+1]-keys[i] != 1 || keys[i+2]-keys[i+1] != 1 {
			continue
		}
		iArr := subGeoMap[keys[i]]
		jArr := subGeoMap[keys[i+1]]
		kArr := subGeoMap[keys[i+2]]

		for _, iVal := range iArr {

			for _, jVal := range jArr {
				if jVal < iVal {
					continue
				}

				for _, kVal := range kArr {
					if kVal > jVal {
						result++ 
					}

				}
			}
		}
	}

	return result
}

func binomial(n, k int) int {
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}