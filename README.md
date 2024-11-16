# 🚀 Backend - Golang Application

Bem-vindo ao **Go todo App** desenvolvido em **Golang**! Este repositório contém o código-fonte para a API RESTful que alimenta a aplicação.

## 🛠️ Tecnologias

- ⚙️ **Golang**: Linguagem de programação utilizada para construir a aplicação.
- 🗄️ **MongoDB**: Banco de dados relacional para armazenar dados.
- 🔧 **Docker**: Para orquestração e execução do ambiente de desenvolvimento.
- 📦 **Fiber**: Framework de web leve e rápido.

## 📂 Estrutura do Projeto

```plaintext
project/
├── cmd/                    # Pontos de entrada da aplicação
│   └── app/                # Main da API
│       └── main.go         # Arquivo principal da aplicação
├── internal/               # Código interno da aplicação (não exportado)
│   ├── config/             # Configurações e variáveis de ambiente
│   │   └── config.go       # Arquivo de configuração
│   ├── domain/             # Regras de negócio e entidades
│   │   ├── user.go         # Entidade User
│   │   └── other.go        # Outras entidades
│   ├── handlers/           # Controladores e rotas HTTP
│   │   ├── user_handler.go # Manipuladores relacionados a usuários
│   │   └── other_handler.go
│   ├── middleware/         # Middlewares HTTP
│   │   └── auth.go         # Exemplo de middleware de autenticação
│   ├── repository/         # Camada de acesso a dados (banco de dados)
│   │   ├── user_repo.go    # Repositório de usuários
│   │   └── other_repo.go
│   ├── services/           # Lógica de aplicação (intermediária entre controllers e repositórios)
│   │   ├── user_service.go # Serviço de usuário
│   │   └── other_service.go
│   ├── utils/              # Funções utilitárias (helpers)
│   │   └── helpers.go
│   └── validators/         # Validações e schemas
│       └── user_validator.go
├── pkg/                    # Código reutilizável (público para outros projetos)
│   └── logger/             # Exemplo de biblioteca de logging
│       └── logger.go
├── migrations/             # Scripts de migração do banco de dados
│   └── 0001_initial.sql
├── docs/                   # Documentação da API (e.g., Swagger ou Postman)
│   └── openapi.yaml
├── tests/                  # Testes unitários e de integração
│   ├── integration/        # Testes de integração
│   └── unit/               # Testes unitários
├── go.mod                  # Declaração de dependências do projeto
├── go.sum                  # Hashes de dependências
└── README.md               # Documentação do projeto

```

## 🚀 Como executar o projeto

### Pré-requisitos

- 🐳 **Docker** e **Docker Compose** instalados.
- 🔧 **Golang** (versão 1.20 ou superior).
- 🛠️ **Make** (opcional, para facilitar os comandos).

### Passos

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/sua-aplicacao-backend.git
   cd sua-aplicacao-backend
   ```

2. Configure as variáveis de ambiente no arquivo `.env`:
   ```dotenv
   DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
   ```

3. Suba o ambiente com Docker:
   ```bash
   docker-compose up -d
   ```

4. Rode a aplicação:
   ```bash
   go run main.go
   ```

5. Acesse a API em: [http://localhost:8080](http://localhost:8080)

## 📋 Endpoints

- `GET /api/v1/resource` - Retorna todos os recursos.
- `POST /api/v1/resource` - Cria um novo recurso.
- `PUT /api/v1/resource/:id` - Atualiza um recurso existente.
- `DELETE /api/v1/resource/:id` - Remove um recurso.

## 🧪 Testes

Execute os testes com:
```bash
go test ./...
```

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

## 🛡️ Licença

Este projeto está sob a licença [MIT](LICENSE).

---

Se precisar de ajustes ou algo mais específico, é só avisar! 🚀