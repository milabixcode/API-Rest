# Go-studies

# API-Rest

## Estrutura do Projeto

```sh
.
├── api-rest
│   ├── create-event.http
│   ├── create-user.http
│   ├── delete-event.http
│   ├── get-events.http
│   ├── get-single-event.http
│   ├── login.http
│   └── update-event.http
├── db
│   └── db.go
├── middlewares
│   └── auth.go
├── models
│   ├── event.go
│   └── user.go
├── routes
│   ├── events.go
│   ├── routes.go
│   └── users.go
├── utils
│   ├── hash.go
│   └── jwt.go
├── api.db
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Descrição dos Arquivos e Pastas

- `api-rest/`: Contém arquivos `.http` para testar as rotas da API diretamente.

  - `create-event.http`: Requisição para criar um evento.
  - `create-user.http`: Requisição para criar um usuário.
  - `delete-event.http`: Requisição para deletar um evento.
  - `get-events.http`: Requisição para listar todos os eventos.
  - `get-single-event.http`: Requisição para obter um evento específico.
  - `login.http`: Requisição para login do usuário.
  - `update-event.http`: Requisição para atualizar um evento.

- `db/`: Contém a lógica de inicialização e conexão com o banco de dados.

  - `db.go`: Arquivo responsável pela configuração e conexão ao banco de dados.

- `middlewares/`: Contém middlewares usados na aplicação.

  - `auth.go`: Middleware para autenticação de rotas.

- `models/`: Contém os modelos de dados usados pela aplicação.

  - `event.go`: Modelo do evento.
  - `user.go`: Modelo do usuário.

- `routes/`: Define as rotas da aplicação.

  - `events.go`: Rotas relacionadas a eventos.
  - `routes.go`: Arquivo que registra todas as rotas da aplicação.
  - `users.go`: Rotas relacionadas a usuários.

- `utils/`: Contém funções utilitárias usadas na aplicação.

  - `hash.go`: Funções para hashing de senhas.
  - `jwt.go`: Funções para geração e validação de tokens JWT.

- `api.db`: Arquivo de banco de dados SQLite.

- `go.mod`: Arquivo de configuração do módulo Go.

- `go.sum`: Registro das dependências do Go com suas respectivas versões.

- `main.go`: Ponto de entrada da aplicação.

`README.md`: Arquivo de documentação do projeto.

## Descrição

Este projeto é uma API RESTful construída com Gin que oferece recursos de autenticação e está integrada a um banco de dados SQL. A API permite que usuários se registrem, façam login e acessem recursos protegidos com base em permissões. A autenticação é baseada em tokens JWT (JSON Web Tokens).

## Funcionalidades

- Registro de usuários.
- Login e geração de tokens JWT.
- Autenticação baseada em JWT.
- Operações CRUD em [entidades principais do seu domínio].
- Integração com banco de dados SQL ([MySQL, PostgreSQL, SQLite, etc.]).
- Rotas protegidas que exigem autenticação.
- Validação de permissões para acessar recursos específicos.

## Tecnologias Utilizadas

- Linguagem: Go
- Framework: Gin
- Autenticação: JWT (JSON Web Token)
- Banco de Dados: MySQL
