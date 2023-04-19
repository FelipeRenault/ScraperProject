# ScraperRepository

## Requisitos

- Versão do docker >= 20.10
- docker-compose >= 2.0.0

## Subindo o ambiente

- `make build`
- `make run`

## Chamando a API

- `http://localhost:8080/api/product?url=<URL>`

## Sites implementados

- Magalu:
  - http://localhost:8080/api/product?url=https://www.magazineluiza.com.br/cadeira-de-escritorio-giratoria-preta-rodinhas-estofada-importway/p/gacbf7c766/mo/moec/
  - http://localhost:8080/api/product?url=https://www.magazineluiza.com.br/gaveteiro-4-gavetas-preto-escritorio-home-office-treearbor/p/gj4dfdg115/mo/gaes/
  - http://localhost:8080/api/product?url=https://www.magazineluiza.com.br/smartphone-samsung-galaxy-m14-5g-128gb-4gb-ram-tela-infinita-de-6-6-dual-chip/p/aa1k3ckcg1/te/galx/

- Zattini:
  - http://localhost:8080/api/product?url=https://www.zattini.com.br/relogio-orient-feminino-analogico-fbss1193-d2sx-incolor-L20-3172-460
  - http://localhost:8080/api/product?url=https://www.zattini.com.br/bolsa-santa-lolla-mini-bag-matelasse-feminina-preto-H08-8087-006
  - http://localhost:8080/api/product?url=https://www.zattini.com.br/bota-cano-curto-comfortflex-feminina-preto-D82-2957-006

## To-do

- Testes unitários
- Tratar erros e informações faltantes na coleta de dados
- Possibilitar que o scrape diferencie entre produtos disponíveis e indisponíveis
- Implementar mais sites
