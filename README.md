# 🚀 Backend - Go Todo App

Bem-vindo ao **Go todo App** desenvolvido em **Golang**! Este repositório contém o código-fonte para a API RESTful que alimenta a aplicação.

## 🛠️ Tecnologias

- ⚙️ **Golang**: Linguagem de programação utilizada para construir a aplicação.
- 🗄️ **MongoDB**: Banco de dados relacional para armazenar dados.
- 🔧 **Docker**: Para orquestração e execução do ambiente de desenvolvimento.
- 📦 **Fiber**: Framework de web leve e rápido.

## 📂 Estrutura do Projeto

```plaintext
project/
├── cmd/                            # Pontos de entrada da aplicação
│   └── app/                        # Main da API
│       └── main.go                 # Arquivo principal da aplicação
├── internal/                       # Código interno da aplicação
│   ├── config/                     # Configurações da aplicação
│   │   └── db/                     # Configurações relacionadas ao banco de dados
│   │       └── db_connection.go    # Arquivo de conexão com o banco de dados
│   ├── domain/                     # Regras de negócio e entidades       
│   │   └── todo.go                 # Entidade Todo
│   ├── handlers/                   # Controladores e rotas HTTP
│   │   └── todo_handler.go         # Manipuladores relacionados às todos
│   ├── repository/                 # Camada de acesso a dados (banco de dados)
│   │   └── todo_repo.go            # Repositório de todos
│   ├── routes/                     # Definição de rotas da aplicação
│   │   └── todo_routes.go          # Rotas de todos
│   ├── utils/                      # Funções utilitárias (helpers)
│   │   └── consts/                 # Constantes da aplicação
│   │       ├── error_consts.go     # Mensagems de erro
│   └──     └── messages.go         # Mensagens de sucesso
├── go.mod                          # Declaração de dependências do projeto
├── go.sum                          # Hashes de dependências
├── docker-compose.yml              # Arquivo de configuração do Docker Compose
├── Dockerfile                      # Arquivo de configuração do Docker
├── .gitignore                      # Pastas e diretórios a serem ignorados pelo Git
├── air.toml                        # Configuração do Air (hot reload)
└── README.md                       # Documentação do projeto

```

## 🚀 Como executar o projeto

### Pré-requisitos

- 🐳 **Docker** e **Docker Compose** instalados.
- 🔧 **Golang** (versão 1.20 ou superior).

### Passos

1. Clone a branch do repositório que contém o backend:
   ```bash
   git clone --single-branch --branch backend https://github.com/DevMatheusSilva/go-todo-app.git
   cd go-todo-app
   ```

2. Configure as variáveis de ambiente no arquivo `.env`:
   ```dotenv
   MONGODB_URI=<sua_string_de_conexão_com_o_mongodb>
   ```

3. Suba o ambiente com Docker:
   ```bash
   docker-compose up -d
   ```

4. Rode a aplicação:
   ```bash
   go run main.go
   ```

5. Acesse a API em: [http://localhost:3000](http://localhost:8080)


## 📋 Endpoints

- `GET /api/todos/`             - Retorna todas as todos cadastradas.
- `POST /api/todos/`            - Cria uma nova todo.
- `PATCH /api/todos/:id`        - Marca a todo como completa.
- `DELETE /api/todos/:id`       - Remove uma todo.

## 🗂️ Contribuição

1. Faça um fork do repositório.
2. Crie uma branch para sua feature:
   ```bash
   git checkout -b minha-feature
   ```
3. Envie sua feature:
   ```bash
   git push origin minha-feature
   ```
