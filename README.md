# Aplicativo ToDo com Go & React 🚀

Este é um simples aplicativo **ToDo** que utiliza Go no backend e React no frontend. O app é uma aplicação CRUD básica que permite aos usuários criar, ler, atualizar e deletar **todos**. Originalmente criado por Burak Orkmez, este projeto foi destaque no canal do freeCodeCamp no YouTube.

## 📋 Índice

- [Tecnologias](#tecnologias)
  - [Frontend](#frontend)
  - [Backend](#backend)
- [Começando 🚀](#começando-)
  - [Pré-requisitos](#pré-requisitos)
  - [Instalação](#instalação)
- [Funcionalidades 🌟](#funcionalidades-)
- [Uso 📖](#uso-)
- [Licença 📄](#licença-)

## Tecnologias

### Frontend

- **React**: Biblioteca JavaScript para construção de interfaces de usuário.
- **TypeScript**: Superset do JavaScript com tipagem estática opcional.
- **TanStack Query**: Biblioteca para gerenciamento de estado assíncrono e cache de dados.
- **Chakra UI**: Biblioteca de componentes modular e acessível para React.

### Backend

- **Go**: Linguagem de programação utilizada no backend.
- **Fiber**: Framework web inspirado no Express para Go.
- **MongoDB**: Banco de dados NoSQL para armazenar os **todos**.

## Começando 🚀

### Pré-requisitos

- **Go**
- **Node.js** e **npm** (incluído com o Node.js)
- **MongoDB**

### Instalação

Como o backend e o frontend estão em branches diferentes (`develop/backend` e `develop/frontend`), você precisará clonar cada branch separadamente.

#### Clonar o Repositório

##### Backend

Clone o branch do backend:

```bash
git clone -b develop/backend https://github.com/DevMatheusSilva/go-todo-app.git backend
```

##### Frontend

Clone o branch do frontend:

```bash
git clone -b develop/frontend https://github.com/DevMatheusSilva/go-todo-app.git frontend
```

#### Configuração do Backend

Navegue para o diretório do backend:

```bash
cd backend
```

Instale as dependências:

```bash
go mod download
```

Configure as variáveis de ambiente (crie um arquivo `.env` no diretório do backend):

```
MONGODB_URI=<sua_string_de_conexão_mongodb>
PORT=3000
```

Navegue para o diretório do arquivo principal:

```bash
cd cmd/app

```
Execute o backend:

```bash
go run main.go
```

#### Configuração do Frontend

Em uma nova janela de terminal, navegue para o diretório do frontend:

```bash
cd frontend
```

Instale as dependências:

```bash
npm install
```

Execute o frontend:

```bash
npm start
```

#### Acesse o App 🌐

Abra seu navegador e acesse: [http://localhost:3000](http://localhost:3000)

## Funcionalidades 🌟

- **Criar Todo**: Adicione um novo **todo** à lista.
- **Ler Todos**: Visualize a lista de **todos**.
- **Atualizar Todo**: Edite **todos** existentes.
- **Deletar Todo**: Remova **todos** da lista.
