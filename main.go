package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"adam.com/toto/toto"
)

const (
	HB_FILE   = "hb.txt"
	ADAM_FILE = "adam.txt"
	WINNING_FILE = "winning.txt"
	INPUT_DIR = "./input"
	SEPERATOR = ","
)

var (
	winningNumbers   []string
	additionalNumber string
)

func main() {
	log.SetFlags(log.Lmicroseconds + log.Llongfile)

	log.Println("Starting program...")
	printSpaces(3)

	// getWinningNumbersFromFile()
	getWinningNumbersFromWeb()
	log.Printf("Winning Numbers: %v, Additional Number: %s", winningNumbers, additionalNumber)
	printSpaces(3)

	log.Println("Checking HB winnings...")
	checkWinnings(HB_FILE)
	printSpaces(3)

	log.Println("Checking Adam winnings...")
	checkWinnings(ADAM_FILE)
	printSpaces(3)
}

func checkWinnings(fileStr string) {
	hbData, err := os.ReadFile(INPUT_DIR + "/" + fileStr)
	if err != nil {
		log.Fatalln("error reading file: ", err)
	}
	betsRaw := string(hbData)
	betsString := strings.Split(betsRaw, "\n")
	lenBets := len(betsString)
	bets := make([][]string, lenBets)
	for i := 0; i < lenBets; i++ {
		bets[i] = strings.Split(betsString[i], SEPERATOR)
	}
	for i, bet := range bets {
		group := getGroup(bet)
		if group != 0 {
			log.Printf("Row: %d, Bet: %v, Group: %d", i+1, bet, group)
		}
	}
}

func getWinningNumbersFromWeb() {
	winningNumbers, additionalNumber = toto.GetTotoRaw()
}

func getWinningNumbersFromFile() {
	winningData, err := os.ReadFile(INPUT_DIR + "/" + WINNING_FILE)
	if err != nil {
		log.Fatalln("error reading file: ", err)
	}
	winningRaw := string(winningData)
	winningRows := strings.Split(winningRaw, "\n")
	winningNumbers = strings.Split(winningRows[0], SEPERATOR)
	additionalNumber = winningRows[1]

}

func getGroup(bet []string) int {
	winMatch, addMatch := 0, false
	for _, winningNum := range winningNumbers {
		if stringSliceContains(bet, winningNum) {
			winMatch += 1
		}
	}
	if stringSliceContains(bet, additionalNumber) {
		addMatch = true
	}
	return checkGroup(winMatch, addMatch)
}

// Group 1: 6 winning
// Group 2: 5 winning + additional
// Group 3: 5 winning
// Group 4: 4 winning + additional
// Group 5: 4 winning
// Group 6: 3 winning + additional
// Group 7: 3 winning
func checkGroup(winMatches int, addMatch bool) int {
	switch winMatches {
	case 6:
		return 1
	case 5:
		if addMatch {
			return 2
		}
		return 3
	case 4:
		if addMatch {
			return 4
		}
		return 5
	case 3:
		if addMatch {
			return 6
		}
		return 7
	default:
		return 0
	}
}

func stringSliceContains(strSlice []string, s string) bool {
	for _, str := range strSlice {
		if strings.EqualFold(str, s) {
			return true
		}
	}
	return false
}

func printSpaces(numSpaces int) {
	for i := 0; i < numSpaces; i++ {
		fmt.Println()
	}
}
