package mergeintervals

import (
	"sort"
)

// Merge Intervals. Given a collection of intervals, write a function to merge all overlapping intervals.

// For example, given [1,3],[2,6],[8,10],[15,18], the function should return [1,6],[8,10],[15,18].
// This is because intervals [1,3] and [2,6] overlap and should be merged into [1,6].

// 1. sort intervals by start, then end
// 2. merge subsequent intervals into prev interval
func mergeIntervals(intervals [][]int) [][]int {
	intervals = sortIntervals(intervals)
	result := [][]int{}
	used := make([]bool, len(intervals))

	for i, currInterval := range intervals {

		if used[i] {
			continue
		}

		for j := i; j < len(intervals); j++ {
			if !used[j] && intervalsOverlap(currInterval, intervals[j]) {
				currInterval = combineTwoIntervals(currInterval, intervals[j])
				used[j] = true
			} else {
				break
			}
		}
		result = append(result, currInterval)

	}
	return result
}

// assume i1 and i2 overlap
func combineTwoIntervals(i1, i2 []int) []int {
	return []int{min(i1[0], i2[0]), max(i1[1], i2[1])}
}

// return true if start/end of i1 and i2 overlap
func intervalsOverlap(i1, i2 []int) bool {
	if i1[0] < i2[0] {
		return i2[0] <= i1[1]
	}
	return i1[0] <= i2[1]
}

// sort intervals by start, then end
func sortIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return false
	})

	return intervals
}
