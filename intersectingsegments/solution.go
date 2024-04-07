package solution

import "fmt"

type pair struct {
	length, lastIdx int
}

func Solution(A []int) int {
	fmt.Println(A)
	// create a dictionary of {sum: at most 2 pairs}
	dict := map[int][]pair{}
	for i:=0; i<len(A)-1; i++ {
		currSum := A[i] + A[i+1]

		// create kv pair if not yet
		if _, ok := dict[currSum]; !ok {
			dict[currSum] = []pair{}
		}

		// first occurence of sum, create single pair
		if len(dict[currSum]) == 0 {
			dict[currSum] = append(dict[currSum], pair{length: 1, lastIdx: i})
		} else if len(dict[currSum]) == 1 { 
			// non-overlapping, increment pair length and set lastIdx
			if dict[currSum][0].lastIdx != i-1 { 
				dict[currSum][0].lastIdx = i
				dict[currSum][0].length += 1
			} else { // overlapping, create new pair
				newPair := pair{length: dict[currSum][0].length, lastIdx: i}
				dict[currSum] = append(dict[currSum], newPair)
			}
		} else { //len == 2
			// if adding to both pairs, just keep the one with bigger length
			addToBoth := true
			for j:=0; j<=1; j++ {
				if dict[currSum][j].lastIdx != i-1 {
					dict[currSum][j].lastIdx = i
					dict[currSum][j].length += 1
				} else {
					addToBoth = false
				}
			}
			if addToBoth { // keep pair with biggest length
				pairToKeep := 0
				if dict[currSum][1].length > dict[currSum][0].length {
					pairToKeep = 1
				}
				dict[currSum] = []pair{dict[currSum][pairToKeep]}
			} else if dict[currSum][0].length == dict[currSum][1].length { // just keep first
				dict[currSum] = []pair{dict[currSum][0]}
			}
			
		}
		fmt.Println(dict)
		
	}
	return getMaxSegmentLength(dict)
}

func getMaxSegmentLength(dict map[int][]pair) int {
	maxVal := 0
	for _, v := range dict {
		for _, pair := range v {
			if pair.length > maxVal {
				maxVal = pair.length
			}
		}
	}
	return maxVal
}