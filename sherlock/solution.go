package solution

import "fmt"

const (
	NO  = "NO"
	YES = "YES"
)

func isValid(s string) string {

	charMap := map[string]int{}
	for _, charRune := range s {
		charStr := string(charRune)
		if _, ok := charMap[charStr]; !ok {
			charMap[charStr] = 0
		}
		charMap[charStr]++
	}

	countMap := map[int]int{}
	for _, v := range charMap {
		if _, ok := countMap[v]; !ok {
			countMap[v] = 0
		}
		countMap[v]++
	}
	fmt.Println("countMap: ", countMap)

	if len(countMap) > 2 {
		return NO
	}

	if len(countMap) == 1 {
		return YES
	}

	if isCountMapValid(countMap) {
		return YES
	}

	return NO
}

// countMap is map of count: num of occurences
// if there exist a 1:1 -> return true
// if the higher count value == 1 -> return true
func isCountMapValid(countMap map[int]int) bool {

	for k, v := range countMap {
		if k == 1 && v == 1 {
			return true
		}
	}

	keys := make([]int, 0, len(countMap))
	for k := range countMap {
		keys = append(keys, k)
	}

	if keys[0] > keys[1] {
		return keys[0]-keys[1] == 1 && countMap[keys[0]] == 1
	}
	return keys[1]-keys[0] == 1 && countMap[keys[1]] == 1
}
