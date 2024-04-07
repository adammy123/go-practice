package targetsum

import "fmt"

func changeSigns(nums []int, target int) int {
	fmt.Println("nums: ", nums)
	fmt.Println("targetSum: ", target)
	output := 0
	currSum := target * -1
	output = findSubsetCount(currSum, nums, []string{}, "")
	if output == 0 {
		output = -1
	}

	fmt.Println("output: ", output)
	return output
}

func findSubsetCount(currSum int, currList []int, history []string, action string) int {
	fmt.Println("currSum: ", currSum)
	fmt.Println("currList: ", currList)

	if action != "" {
		history = append(history, action)
	}
	fmt.Println("history: ", history)
	fmt.Println()
	if len(currList) == 0 {
		if currSum == 0 {
			fmt.Println("++")
			fmt.Println()
			return 1
		}
		return 0
	}
	toAdd := currList[0]
	currList = currList[1:]
	// fmt.Println("toAdd: ", toAdd)
	// fmt.Println("newList: ", currList)
	return findSubsetCount(currSum+toAdd, currList, history, "+") + findSubsetCount(currSum-toAdd, currList, history, "-")
}
