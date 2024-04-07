package knapsack

import (
	"sort"
)

type item struct {
	weight int
	value  int
}

func (i item) getValueToWeightRatio() float32 {
	return float32(i.value) / float32(i.weight)
}

func GetMaxValueGivenMaxCapacity(capacity int, items []item) int {
	maxValue := 0
	currCapacity := 0

	if capacity < 1 || len(items) < 1 {
		return 0
	}

	sortItemsByValueToWeightRatio(items)

	for _, item := range items {
		if currCapacity+item.weight <= capacity {
			currCapacity = currCapacity + item.weight
			maxValue += item.value
		}
	}
	return maxValue
}

// higher value first
// lower weight first if same value
func sortItemsByValueToWeightRatio(items []item) {
	sort.Slice(items, func(i, j int) bool {
		// return if i should be before j
		iValue := items[i].getValueToWeightRatio()
		jValue := items[j].getValueToWeightRatio()
		if iValue == jValue {
			return items[i].weight < items[j].weight
		}
		return iValue > jValue
	})
}
