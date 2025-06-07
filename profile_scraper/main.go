package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type BankIn struct {
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

type BankProfile struct {
	Name      string        `json:"name"`
	CEO       string        `json:"ceo"`
	Employees int           `json:"employees"`
	Lists     []ListProfile `json:"lists"`
}

type ListProfile struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	inputJSON := os.Args[1]
	var bank BankIn

	err := json.Unmarshal([]byte(inputJSON), &bank)
	if err != nil {
		fmt.Println("Erro ao fazer parse do JSON")
		return
	}

	req, err := http.Get(bank.Profile)
	if err != nil {
		fmt.Println("Erro ao fazer requisição HTTP")
		return
	}
	defer req.Body.Close()

	doc, err := goquery.NewDocumentFromReader(req.Body)
	if err != nil {
		fmt.Println("Erro ao carregar HTML")
		return
	}

	var ceo string
	doc.Find(".profile-stats__title").EachWithBreak(func(i int, s *goquery.Selection) bool {
		text := strings.TrimSpace(s.Text())
		if strings.Contains(strings.ToLower(text), "ceo") {
			ceo = s.Next().Text()
			return false
		}
		return true
	})

	var employees int
	doc.Find(".profile-stats__title").EachWithBreak(func(i int, s *goquery.Selection) bool {
		text := strings.TrimSpace(s.Text())
		if strings.Contains(strings.ToLower(text), "employees") {
			employeesStr := s.Next().Text()
			employeesStr = strings.ReplaceAll(employeesStr, ",", "")
			employees, _ = strconv.Atoi(employeesStr)
			return false
		}
		return true
	})

	lists := []ListProfile{}
	doc.Find("div.listuser-content__block.ranking a.listuser-item__list--title").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Text())
		url, exists := s.Attr("href")
		if exists {
			lists = append(lists, ListProfile{
				Name: name,
				URL:  url,
			})
		}
	})

	profile := BankProfile{
		Name:      bank.Name,
		CEO:       strings.TrimSpace(ceo),
		Employees: employees,
		Lists:     lists,
	}

	output, _ := json.MarshalIndent(profile, "", "  ")
	fmt.Println(string(output))
}
