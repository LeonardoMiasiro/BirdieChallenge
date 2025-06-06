package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strconv"
	"strings"
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
	// Abre o arquivo HTML
	file, err := os.Open("banks.html")
	if err != nil {
		fmt.Println("Erro no arquivo")
		return
	}
	defer file.Close()

	// Carrega o HTML
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		fmt.Println("Erro ao carregar HTML")
		return
	}

	// Busca os dados de todos os bancos dentro da classe "a.table-row"
	doc.Find("a.table-row").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find(".nameField").Text())
		city := strings.TrimSpace(s.Find(".city .row-cell-value").Text())
		country := strings.TrimSpace(s.Find(".country .row-cell-value").Text())
		foundedStr := strings.TrimSpace(s.Find(".yearFounded .row-cell-value").Text())
		rankStr := strings.TrimSpace(s.Find(".searchIndustryRank .starRank").Text())

		// Transforma rank e data de fundação para inteiro
		founded, _ := strconv.Atoi(foundedStr)
		rank, _ := strconv.Atoi(rankStr)

		// Pegar a URL do perfil do banco
		uri, _ := s.Attr("uri")
		profile := ""
		if uri != "" {
			profile = fmt.Sprintf("https://www.forbes.com/companies/%s/?list=worlds-best-banks", uri)
		}

		// Cria um objeto bank com as informações extraidas
		bank := Bank{
			Name:    name,
			City:    city,
			Country: country,
			Founded: founded,
			Rank:    rank,
			Profile: profile,
		}

		// Converse o objeto para JSON e faz um print dele
		jsonBytes, _ := json.Marshal(bank)
		fmt.Println(string(jsonBytes))
	})
}
