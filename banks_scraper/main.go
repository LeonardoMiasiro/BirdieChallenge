package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Bank struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
	Founded int    `json:"founded"`
	Rank    int    `json:"rank"`
	Profile string `json:"profile"`
}

func main() {
	file, err := os.Open("banks.html")
	if err != nil {
		fmt.Println("Erro no arquivo")
		return
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		fmt.Println("Erro ao carregar HTML")
		return
	}

	doc.Find("a.table-row").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find(".nameField").Text())
		city := strings.TrimSpace(s.Find(".city .row-cell-value").Text())
		country := strings.TrimSpace(s.Find(".country .row-cell-value").Text())
		foundedStr := strings.TrimSpace(s.Find(".yearFounded .row-cell-value").Text())
		rankStr := strings.TrimSpace(s.Find(".searchIndustryRank .starRank").Text())

		founded, _ := strconv.Atoi(foundedStr)
		rank, _ := strconv.Atoi(rankStr)

		uri, _ := s.Attr("uri")
		profile := ""
		if uri != "" {
			profile = fmt.Sprintf("https://www.forbes.com/companies/%s/?list=worlds-best-banks", uri)
		}

		bank := Bank{
			Name:    name,
			City:    city,
			Country: country,
			Founded: founded,
			Rank:    rank,
			Profile: profile,
		}

		jsonBytes, _ := json.Marshal(bank)
		fmt.Println(string(jsonBytes))
	})
}
