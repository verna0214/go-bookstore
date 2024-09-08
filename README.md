# Go Bookstore

Go Bookstore is a simple side project in Go that demonstrates how to connect to a MySQL database and perform CRUD operations. The project follows the MVC architecture and uses the `gorilla/mux` router to manage HTTP requests.

## Features

- Connects to a MySQL database
- Implements CRUD operations (Create, Read, Update, Delete)
- Uses MVC architecture for code organization
- Handles routing with `gorilla/mux`

## Requirements

- Go 1.22 or higher
- MySQL database
- `gorilla/mux` and `gorm` packages

## Installation

1. Clone the repository:
  ```sh
  git clone https://github.com/verna0214/go-bookstore.git
  ```

2. Navigate to the project directory:
  ``` sh
  cd go-bookstore
  ```

3. Install the required Go packages:
  ``` sh
  go mod tidy
  ```

## Configuration
Create a .env file in the root directory with your MySQL database connection details. Refer to the .env.example file for the required format:
  ``` go
  MYSQL_USERNAME=
  MYSQL_PASSWORD=
  ```

## Running
Run the main program. The application will run on localhost:9010.
  ``` go
  go run main.go
  ```

## Routes
- GET /books - Retrieve all books
- GET /books/{id} - Retrieve a specific book
- POST /books - Create a new book
- PUT /books/{id} - Update book information
- DELETE /books/{id} - Delete a book