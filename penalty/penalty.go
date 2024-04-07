package penalty

import (
	"fmt"
	"math"
	"strings"
)

// get_best_closing_times("BEGIN Y Y END \nBEGIN N N END")
//   should return an array: [2, 0]
// get_best_closing_times("BEGIN BEGIN \nBEGIN N N BEGIN Y Y\n END N N END")
//   should return an array: [2]

// ASSUMPTIONS
// - remove all \n
// - input will only contains BEGIN END \n Y N
// - valid log will always have at least one opening hour
// - single spaces between each valid word
func Get_best_closing_times(input string) []int {
	result := []int{}

	// 1. clean input
	// - remove all \n
	input = strings.ReplaceAll(input, "\n", "")

	// 2. go through each string in cleaned input
	// - BEGIN -> reset TEMP = "BEGIN"
	// - END -> add to TEMP only if last string is Y/N -> find_best_closing_time & reset
	// - Y/N -> if last string is Y/N/BEGIN, add to TEMP and continue
	temp := ""
	for _, value := range strings.Fields(input) {
		switch value {
		case "BEGIN":
			temp = value
		case "Y", "N":
			if strings.HasSuffix(temp, "BEGIN") || strings.HasSuffix(temp, "Y") || strings.HasSuffix(temp, "N") {
				temp += " " + value
			}
		case "END":
			if strings.HasSuffix(temp, "Y") || strings.HasSuffix(temp, "N") {
				temp += " " + value
				fmt.Println(temp)
				temp = strings.ReplaceAll(temp, "BEGIN ", "")
        temp = strings.ReplaceAll(temp, " END", "")
        fmt.Println(temp)
				thisBestClosingTime := find_best_closing_time(temp)
				result = append(result, thisBestClosingTime)
				temp = ""
			}
		}
	}

	return result
}

// find_best_closing_time("Y Y N N") should return 2
// which is the closing time with the least penalty value

// ASSUMPTIONS:
// compare between same penalty values? - return either
// brute force: go through all values of `k` return smallest penalty
func find_best_closing_time(input string) int {
	var bestTime int
	leastPenalty := math.MaxInt32
	maxClosingTime := len(strings.Fields(input))+1

	for i:=0; i<maxClosingTime; i++ {
		thisPenalty := compute_penalty(input, i)
		if thisPenalty < leastPenalty {
			leastPenalty = thisPenalty
			bestTime = i
		}
	}

	return bestTime
}

func compute_penalty(input string, k int) int {
	penalty := 0
	isOpen := true

	// 1. iterate through string as an Array
	// 	- if open && N: penalty +=1
	//  - if closed && Y: penalty +=1
	for i, val := range strings.Fields(input) {

		// flips store from open to close
		if i == k {
			isOpen = !isOpen
		}

		if isOpen && strings.EqualFold(val, "N") {
			penalty += 1
		}

		if !isOpen && strings.EqualFold(val, "Y") {
      penalty += 1
    }
	}

	return penalty
}