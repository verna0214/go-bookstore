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
Create a .env file in the root directory with your MySQL database connection details: