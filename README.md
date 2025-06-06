# Birdie Challenge

## Parte 1 - Best Banks Scraper

### Pensamento e Abordagem

Para a primeira etapa do desafio, o objetivo era criar um código que consiga extrair os dados de bancos a partir de um URL, para o desafio em especifico, foi escolhido o site da Forbes.

A obtenção dos dados foi planejada a partir de uma análise do documento HTML, buscando por classes específicas na estrutura da lista dos bancos, conseguindo assim, identificar os dados de cada banco, como nome, cidade, país, ano de fundação, ranking e link para o perfil.

---

### Funcionamento

- O arquivo HTML é aberto e lido localmente.
- Em seguida, o `goquery` navega pelo arquivo HTML e seleciona os tipos de elementos desejados.
- Para cada elemento do banco identificado, são extraídos os dados.
- Por fim, os dados extraídos são organizados em um objeto e convertidos para JSON.

---
