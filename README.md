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
    - [Configurações](#configurações)
  - [Testes Automatizados](#testes-automatizados)
  - [Executando o Projeto](#executando-o-projeto)
    - [Passos](#passos)

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

- net/http;
- Flag;
- Goroutines com Waitgroups;

## Funcionalidades

WIP

## Requisitos

### Tecnologias

- **Golang**: Linguagem principal;
- **Docker**: Containerização da aplicação.

### Configurações

As configurações do sistema podem ser realizadas via **variáveis de ambiente** ou no arquivo `.env` na raiz do projeto:

| Variável          | Descrição                              | Valor Padrão     |
|-------------------|----------------------------------------|------------------|
| ENV               | DESCRICAO                              | VALOR            |

## Testes Automatizados

A aplicação possui **100%** de cobertura de testes nas seguintes áreas:

Caso queira visualizar as informações dos relatórios da execução e cobertura dos testes mais recentes, basta acesssar o link abaixo:

- [Relatórios](.doc/TEST_REPORT.md)

## Executando o Projeto

O projeto foi desenvolvido utilizando Docker/Docker Compose de modo a facilitar a execução do mesmo, entõ para poder executar o projeto basta seguir os passos descritos abaixo.

### Passos

1. Clone o repositório:

```bash
git clone REPO
cd DIRETORIO
```

2. Configure o `.env` na raiz com base no arquivo `.env.example`.

3. Suba o ambiente com Docker Compose:

```bash
docker-compose up --build
```

WIP...
