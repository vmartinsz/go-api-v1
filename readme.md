# User API

Uma aplicação simples em Go para gerenciamento de usuários, seguindo um sistema de camadas profissional.

## Tecnologias Utilizadas

- **Go**
- **Gin** (framework web)
- **PostgreSQL** (ou outro banco de dados suportado)
- **Docker** (opcional, para containerização)

## Estrutura do Projeto

A aplicação segue uma estrutura de camadas bem definida:

```
/go-api-v1
│── cmd/                      # Ponto de entrada da aplicação.
│   ├── api/                  # Ponto para separar tipos de entrada como api, consumers, scripts.
│       ├─ server/            # Separando servidores dentro de api, podendo ter mais de um api dentro da pasta.
│          ├─ routes/         # A pasta de routes guarda todas as routas e cada arquivo é um tipo de rotas.
│              ── routes.go
│       ├── main.go
│── internal/                 # Core da aplicação
│   ├── contracts/            # Interfaces e contratos de infraestrutura e ferramentas externas como S3, por ex.
│   ├── persistence/          # Implementação das interfaces do Repository.
│   ├── repository/           # Interfaces e contratos do banco de dados.
│   ├── usecase/              # Regras de negócio de cada caso uso.
│   ├── controller/           # Controladores HTTP.
│   ├── entity/               # Definição de entidades.
│── config/                   # Configurações e variáveis de ambiente da aplicação.
```

## Endpoints

`POST` /users - Criação de usuários
**Body:**
```json
{
  "name": "John",
  "last_name": "Doe",
  "age": 23
}
```
---
`GET` /users/{id}: Recuperação de dados do usuário pelo ID.
**Body:**
```json
{
  "id": 1,
  "name": "John",
  "last_name": "Doe",
  "age": 23
}
```
---
`PATCH` /users/{id}: Alteração de campos do usuário pelo ID.
**Body:**
```json
{
  "age": 13
}
```
---
`DELETE` /users/{id}: Deleção de conta do usuário pelo ID.
**Body:**
```json
{
  "message": "usuário deletado com sucesso."
}
```

## Como Rodar

1. Clone o repositório:
```sh
git clone https://github.com/vmartinsz/go-api-v1.git
cd go-api-v1
```

2. Configure as variáveis de ambiente no arquivo `.env` ou `config.yaml`.

3. Inicie a aplicação:
   ```sh
   go run main.go
   ```

Ou utilizando Docker:
   ```sh
   docker-compose up --build
   ```
## Licença

Este projeto é licenciado sob a [MIT License](LICENSE).
