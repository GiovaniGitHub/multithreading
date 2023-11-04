# multithreading

Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.
As duas requisições serão feitas simultaneamente para as seguintes APIs:
https://cdn.apicep.com/file/apicep/" + cep + ".json
http://viacep.com.br/ws/" + cep + "/json/
Os requisitos para este desafio são:
- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Iniciar o projeto
```bash
# Clonar o projeto
$ git clone https://github.com/GiovaniGitHub/multithreading

# Entrar no diretório do desafio
$ cd multithreading

# Executar o projeto
$ go run cmd/server/main.go
```

## Documentação da API de Consulta de CEP
A API de Consulta de CEP permite que você obtenha informações de endereço com base no CEP (Código de Endereçamento Postal) fornecido. Esta API suporta solicitações GET e retorna os dados em formato JSON.

### Teste via Swagger
 - Acessar a url http://localhost:8000/swagger/index.html


### Endpoint da API
O endpoint da API é:

```bash
http://localhost:8000/cep/{cep}
```
Substitua ``{cep}`` pelo CEP desejado. Por exemplo, se você deseja consultar o CEP "70070130", a URL seria:

```bash
http://localhost:8000/cep/70070130
```

### Requisição
Para fazer uma solicitação para a API, você pode usar o comando curl da seguinte maneira:

```bash
curl -X 'GET' \
 'http://localhost:8000/cep/{cep}' \
 -H 'accept: application/json'
```

Substitua {cep} pelo CEP que você deseja consultar.

#### Resposta

A API responderá com um objeto JSON que contém as informações do endereço associadas ao CEP fornecido. A estrutura do objeto de resposta pode variar, mas geralmente inclui campos como:

cep: O CEP consultado.
logradouro: O nome da rua ou avenida.
bairro: O bairro.
cidade: A cidade.
estado: O estado.
Outros campos específicos, dependendo da implementação da API.

Exemplo de resposta:

```json
{
  "Api":"http://viacep.com.br/ws/",
  "Cep":"70070-130",
  "Bairro":"Asa Sul",
  "Localidade":"Brasília",
  "UF":"DF",
  "Logradouro":"SBS Quadra 3"
}
```