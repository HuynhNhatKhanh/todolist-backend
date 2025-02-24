# TodoList Backend

## Overview

This is the backend service for a TodoList application, built using Golang. The application provides a RESTful API for managing tasks.

## Features

- Create, read, update, and delete (CRUD) todos
- RESTful API design
- Lightweight and efficient Golang backend

## Technologies Used

- **Golang** (version 1.24.0)
- **Debian** (minimal runtime image)
- **Docker** (for containerization)

## Setup & Installation

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.24.0 or later)
- [Docker](https://www.docker.com/)

### Build and Run Locally

1. Clone the repository:
   ```sh
   git clone https://github.com/HuynhNhatKhanh/todolist-backend.git
   cd todolist-backend
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Build the application:
   ```sh
   go build -o todolist ./cmd/main.go
   ```
4. Run the application:
   ```sh
   ./todolist
   ```

## Docker Deployment

### Build Docker Image

To build the Docker image, run:

```sh
docker build -t todolist-backend .
```

### Run Docker Container

To run the application inside a Docker container:

```sh
docker run -p 8080:8080 todolist-backend
```

## API Endpoints

| Method | Endpoint    | Description             |
| ------ | ----------- | ----------------------- |
| GET    | /todos      | Get all todos           |
| POST   | /todos      | Create a new todo       |
| PUT    | /todos/{id} | Update an existing todo |
| DELETE | /todos/{id} | Delete a todo           |

## Environment Variables

| Variable | Description | Default |
| -------- | ----------- | ------- |
| PORT     | Server port | 8080    |

## License

This project is licensed under the MIT License.

## Author

[Huynh Nhat Khanh]
