# Birdie Challenge

Explicação Inicial sobre o Desenvolvimento e abordagens sobre a parte 1 e 2 do Desafio, abaixo tem uma explicação mais detalhada do código.

## Parte 1 - Best Banks Scraper

### Modo de Uso
```bash
go run banks_scraper/main.go 
```
---

### Pensamento e Abordagem

Para a primeira etapa do desafio, o objetivo era criar um código que consiga extrair os dados de bancos a partir de uma URL, para o desafio em questão, foi escolhido o site da Forbes.

A obtenção dos dados foi planejada a partir de uma análise do documento HTML, buscando por classes específicas na estrutura da lista dos bancos, conseguindo assim, identificar em qual parte do HTML estavam localizado os dados de cada banco, como nome, cidade, país, ano de fundação, ranking e link para o perfil.

---

### Funcionamento

- O arquivo HTML é aberto e lido localmente.
- Em seguida, o `goquery` navega pelo arquivo HTML e seleciona os tipos de elementos desejados.
- Para cada elemento do banco identificado, são extraídos os dados.
- Por fim, os dados extraídos são organizados em um objeto e convertidos para JSON.

---

## Parte 2 - Best Profile Scraper

### Modo de Uso
```bash
go run profile_scraper/main.go 'JSON-do-Banco'
```
Exemplo:

```bash
go run profile_scraper/main.go '{"name":"Nubank","city":"Sao Paulo","country":"Brazil","founded":2013,"rank":11,"profile":"https://www.forbes.com/companies/nubank/?list=worlds-best-banks"}

```
---

### Pensamento e Abordagem
 
Na segunda etapa do desafio, o foco foi criar um novo código que buscasse dados específicos do banco, a partir da sua página de perfil da Forbes. As informações buscadas foram: nome do CEO, número de funcionários e as listas em que o banco está presente.

Para isso, similar à parte 1, foi feita uma inspeção no HTML das páginas de perfil, identificando as classes e elementos que continham essas informações.

---

### Funcionamento

- O programa recebe como entrada um JSON similar ao obtido na parte 1.
- Utiliza o pacote `net/http` para fazer uma requisição ao URL do perfil.
- Em seguida, o conteúdo da página é processado com `goquery`, que busca o nome do CEO, Número de funcionário e as Listas.
- Por fim, as informações são organizadas na estrutura BankProfile e imprimidas como JSON.

---

## Explicação do Código

### Parte 1

#### 1. Definição da estrutura
Define a estrutura para armazenar os dados extraídos de cada banco

```bash
type Bank struct {
    Name    string `json:"name"`
    City    string `json:"city"`
    Country string `json:"country"`
    Founded int    `json:"founded"`
    Rank    int    `json:"rank"`
    Profile string `json:"profile"`
}
```


#### 2. Abertura do Arquivo HTML
Abre o arquivo HTML local chamado 'banks.html', se não for possível abrir, exibe uma mensagem de erro e encerra o programa.

```bash
file, err := os.Open("banks.html")  
if err != nil {
    fmt.Println("Erro no arquivo")
    return
}
defer file.Close()
```

#### 3. Carrega o Arquivo HTML
Converte o conteúdo do arquivo HTML para um objeto goquery, se não for possível carregar o HTML, exibe uma mensagem de erro e encerra o programa.

```bash
doc, err := goquery.NewDocumentFromReader(file)
if err != nil {
    fmt.Println("Erro ao carregar HTML")
    return
}
```

#### 4. Busca os dados de todos os bancos 
Dentro da classe "table-row" presente no objeto goquerry, procura os dados (nome, cidade, país, ano de fundação e ranking) de cada banco

```bash
doc.Find("a.table-row").Each(func(i int, s *goquery.Selection) {
    name := strings.TrimSpace(s.Find(".nameField").Text())
    city := strings.TrimSpace(s.Find(".city .row-cell-value").Text())
    country := strings.TrimSpace(s.Find(".country .row-cell-value").Text())
    foundedStr := strings.TrimSpace(s.Find(".yearFounded .row-cell-value").Text())
    rankStr := strings.TrimSpace(s.Find(".searchIndustryRank .starRank").Text())
```

Além disso, realiza a conversão da Data de fundação e Rank para inteiro

```bash
founded, _ := strconv.Atoi(foundedStr)
rank, _ := strconv.Atoi(rankStr)
```

Por fim, cria a URL do perfil do banco presente na tag.
```bash
uri, _ := s.Attr("uri")
profile := ""
if uri != "" {
    profile = fmt.Sprintf("https://www.forbes.com/companies/%s/?list=worlds-best-banks", uri)
}
```

#### 5. Cria um objeto bank com as informações extraidas
Com todas as informações obtidas, cria um objeto bank

```bash
bank := Bank{
Name:    name,
    City:    city,
    Country: country,
    Founded: founded,
    Rank:    rank,
    Profile: profile,
}
```

#### 6. Converse o objeto para JSON e imprime no terminal
Converte o objeto bank criado anteriormente para JSON e mostra ele no terminal.

```bash
jsonBytes, _ := json.Marshal(bank)
fmt.Println(string(jsonBytes))
```

---

### Parte 2

#### 1. Definição das estruturas
Define as estruturas de dados utilizadas para entrada e saída.
- BankIn: Recebe o Nome e URL do perfil do JSON de entrada.
- BankProfile: Armazena as informações  extraídas da página do banco.
- ListProfile: Representa as listas  em que o banco aparece.

```bash
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
```

#### 2. Recebe o JSON de entrada
Recebe uma string JSON com os dados do banco e atribui os dados na estrutura BankIn, se não for possível atribuir a estruturar, exibe uma mensagem de erro e encerra o programa.

```bash
inputJSON := os.Args[1]
var bank BankIn

err := json.Unmarshal([]byte(inputJSON), &bank)
if err != nil {
    fmt.Println("Erro ao fazer parse do JSON")
    return
}
```

#### 3. Requisição HTTP para o perfil do banco
Faz uma requisição GET para a URL do perfil do banco, se não for possível fazer a requisição, exibe uma mensagem de erro e encerra o programa.

```bash
req, err := http.Get(bank.Profile)
if err != nil {
    fmt.Println("Erro ao fazer requisição HTTP")
    return
}
defer req.Body.Close()
```

#### 4. Carregamento do HTML
Carrega a resposta da requisição e transforma em um objeto goquery, se não for carregar o HTML, exibe uma mensagem de erro e encerra o programa.

```bash
doc, err := goquery.NewDocumentFromReader(req.Body)
if err != nil {
    fmt.Println("Erro ao carregar HTML")
    return
}
```

#### 5. Extração dos dados
Similar ao passo 4 da parte 1, realiza a busca do Nome do CEO, Número de funcionários dentro da página de perfil do banco e as listas. 

Busca o campo cujo texto contém a palavra "CEO" e pega o próximo valor.

```bash
var ceo string
doc.Find(".profile-stats__title").EachWithBreak(func(i int, s *goquery.Selection) bool {
    text := strings.TrimSpace(s.Text())
    if strings.Contains(strings.ToLower(text), "ceo") {
        ceo = s.Next().Text()
        return false
    }
    return true
})
```

Busca o campo com o texto "Employees", remove vírgulas e converte para inteiro.

```bash
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
```

Percorre os links de listas associadas ao banco e armazena seus nomes e URLs.

```bash
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
```

#### 6. Monta o objeto final com as informações extraídas
Cria o objeto BankProfile com todas as informações coletadas.

```bash
profile := BankProfile{
    Name:      bank.Name,
    CEO:       strings.TrimSpace(ceo),
    Employees: employees,
    Lists:     lists,
}
```

#### 7. Converte para JSON e imprime no terminal
Converte o objeto BankProfile criado anteriormente para JSON e mostra ele no terminal.

```bash
output, _ := json.MarshalIndent(profile, "", "  ")
fmt.Println(string(output))
```

