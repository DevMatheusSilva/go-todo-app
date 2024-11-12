# Go & React ToDo App

This is a simple **ToDo App** using **Go** for the backend and **React** for the frontend. The app is a basic **CRUD application** enabling users to **create**, **read**, **update**, and **delete** tasks. Originally created by [Burak Orkmez](https://github.com/burakorkmez), this project was featured on [freeCodeCamp's YouTube channel](https://www.youtube.com/watch?v=lNd7XlXwlho).

## Table of Contents
- [Technologies](#technologies)
    - [Frontend](#frontend)
    - [Backend](#backend)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Features](#features)
- [Usage](#usage)
- [License](#license)

## Technologies

### Frontend
- **React**: JavaScript library for building user interfaces
- **TypeScript**: Superset of JavaScript with optional static typing

### Backend
- **Go**: Programming language used for the backend
- **Fiber**: Express-inspired web framework for Go
- **MongoDB**: NoSQL database for storing tasks

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install)
- [Node.js](https://nodejs.org/) and npm (included with Node.js)
- [MongoDB](https://www.mongodb.com/try/download/community)

### Installation

#### Clone the Repository
```bash
git clone <repository-url>
cd todo-app
```

#### Backend Setup
1. **Navigate to the backend folder**:
   ```bash
   cd backend
   ```
2. **Install dependencies**:
   ```bash
   go mod download
   ```
3. **Set up environment variables** (create a `.env` file in the `backend` directory):
   ```
   MONGODB_URI=<your_mongodb_connection_string>
   PORT=8080
   ```
4. **Run the backend**:
   ```bash
   go run main.go
   ```

#### Frontend Setup
1. **Navigate to the frontend folder**:
   ```bash
   cd ../frontend
   ```
2. **Install dependencies**:
   ```bash
   npm install
   ```
3. **Run the frontend**:
   ```bash
   npm start
   ```

#### Access the App
- Open your browser and navigate to `http://localhost:3000`.

## Features
- **Create Task**: Add a new task to the list.
- **Read Tasks**: View a list of tasks.
- **Update Task**: Edit existing tasks.
- **Delete Task**: Remove tasks from the list.

## Usage
To use the app, start by adding tasks to your list. Each task can be edited or deleted as needed.

## License
This project is licensed under the MIT License.