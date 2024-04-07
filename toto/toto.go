package toto

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const (
	WINNING_NUMBERS_LEN = 6
	WIN_PREFIX          = "win"
)

func GetTotoRaw() (winningNumbers []string, additionalNumber string) {
	log.Println("Getting winning numbers from Singapore Pools")

	client := http.Client{}
	resp, err := client.Get("https://www.singaporepools.com.sg/en/product/sr/Pages/toto_results.aspx")
	if err != nil {
		log.Fatalf("Error calling toto result page: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Non-200 Status Code received: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Error creating goquery document from response body: %+v", err)
	}

	winningNumbers = make([]string, WINNING_NUMBERS_LEN)
	for i := 0; i < WINNING_NUMBERS_LEN; i++ {
		currWinNumStr := strconv.Itoa(i + 1)
		className := "." + WIN_PREFIX + currWinNumStr
		winningNumbers[i] = doc.Find(className).Text()
	}

	additionalNumber = doc.Find(".additional").Text()
	return
}
