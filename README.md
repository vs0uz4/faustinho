# Faustinho

> [!IMPORTANT]
> Para poder executar o projeto contido neste repositório é necessário que se tenha o Docker instalado no computador. Para maiores informações siga o site <https://www.docker.com/>

- [Faustinho](#faustinho)
  - [Desafio GoLang Pós GoExpert - CLI Stress Test](#desafio-golang-pós-goexpert---cli-stress-test)
    - [Enunciado](#enunciado)
      - [Execução do Teste](#execução-do-teste)
      - [Geração de Relatório](#geração-de-relatório)
      - [Execução da Aplicação](#execução-da-aplicação)
  - [O que é o Faustinho](#o-que-é-o-faustinho)
  - [Funcionalidades](#funcionalidades)
  - [Requisitos](#requisitos)
    - [Tecnologias](#tecnologias)
  - [Executando o Projeto](#executando-o-projeto)
    - [Passos](#passos)
    - [Exemplo de Uso](#exemplo-de-uso)
  - [Relatório de Saída](#relatório-de-saída)
  - [Erros e Mensagens de Uso](#erros-e-mensagens-de-uso)

## Desafio GoLang Pós GoExpert - CLI Stress Test

Este projeto faz parte da Pós GoExpert como desafio, nele precisaríamos criar uma `CLI` que pudesse realizar testes de stress em serviços web, onde o usuário pudesse fornecer parâmetros como URL, o numero total de requests e quantidade total de chamadas simultaneas, o mesmo realizasse os devidos testes e ao final exibisse um relatório. Para isto temos o seguinte enunciado abaixo:

### Enunciado

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

- Entrada de Parâmetros via CLI:

```plain
--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.
```

#### Execução do Teste

- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.

#### Geração de Relatório

- Apresentar um relatório ao final dos testes contendo:
  - Tempo total gasto na execução
  - Quantidade total de requests realizados.
  - Quantidade de requests com status HTTP 200.
  - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

#### Execução da Aplicação

Poderemos utilizar essa aplicação fazendo uma chamada via Docker.

Ex:

```bash
docker run <sua imagem docker> —url=http://google.com —requests=1000 —concurrency=10
```

## O que é o Faustinho

Faustinho é o nome que dei a este projeto, onde implementa uma **CLI** robusta para execução de testes de stress em serviços web, contendo geração de relatórios, utilizando as seguintes tecnologias/pacotes/frameworks:

- net/http: Para gerenciar requisições HTTP;
- Flag: Para parsing de argumentos da CLI;
- Goroutines com Waitgroups: Para gerenciar concorrência.

## Funcionalidades

- Executa testes de stress em serviços web de forma simples e eficiente.
- Gera relatórios detalhados com informações como tempo total, requisições bem-sucedidas e distribuição de status HTTP.
- Limita o número de requisições simultâneas para otimizar o uso de recursos.

## Requisitos

### Tecnologias

- **Golang**: Linguagem principal;
- **Docker**: Containerização para facilitar a execução da aplicação.

## Executando o Projeto

O projeto foi desenvolvido utilizando Docker para facilitar sua execução. Basta seguir os passos descritos abaixo.

### Passos

1. Clone o repositório:

```bash
git clone https://github.com/vs0uz4/faustinho.git
cd faustinho
```

2. Construa a imagem Docker da aplicação:

```bash
docker build -t faustinho-cli .
```

3. Execute o container para realizar um teste de stress:

```bash
docker run --rm faustinho-cli --url=http://example.com --requests=1000 --concurrency=10
```

### Exemplo de Uso

Comando:

```bash
docker run --rm faustinho-cli --url=http://192.168.97.4:8000/health --requests=10000 --concurrency=50
```

> [!IMPORTANT]
> O uso de valores muito altos para `concurrency` poderá ocasionar em consumo elevado dos recursos de **CPU** e **MEMORIA** da máquina **HOST** executando o Docker, use com cautela. Caso queira prevenir esse tipo de comportamento, você pode inicializar o container, passando `flags` do Docker que limitam os recursos alocados pelo container, veja um exemplo abaixo:

```bash
docker run --rm --memory=256m --cpus=2 faustinho-cli --url=http://example.com:8000/health --requests=100000 --concurrency=50
```

> Ajuste as `flags` de limitação dos recursos, de acordo com a sua realidade.

## Relatório de Saída

Abaixo está um exemplo de saída gerada pelo sistema, incluindo todas as métricas relevantes.

```plain
Faustinho Stress Test Report
------------------------------
Total Requests: 10000
Successful Requests (200): 10000
Status Code Distribution:
  200: 10000
  Timeouts: 0
  Client Errors: 0
Total Time: 1.053109895s
Min Time per Request: 86.126µs
Max Time per Request: 1.049401582s
Avg Time per Request: 1.949941ms
Faustinho stress test completed!
```

## Erros e Mensagens de Uso

Caso algum parâmetro obrigatório não seja informado, a CLI exibe a seguinte mensagem de ajuda:

```plain
CLI Faustinho Stress Tester
---------------------------
Usage:
  ./faustinho [OPTIONS]

Options:
  --url         (string) URL to be tested. Required.
  --requests    (int) Number of total requests. Default: 100.
  --concurrency (int) Number of concurrent requests. Default: 10.

Example:
  ./faustinho --url=http://example.com --requests=1000 --concurrency=10

```
