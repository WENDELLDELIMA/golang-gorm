# My First GORM Project Using Golang

This project is a simple REST API built with Golang and GORM, demonstrating basic CRUD (Create, Read, Update, Delete) operations. The API manages a collection of "Personalities" with attributes such as ID, Name, and History. It uses Gorilla Mux for routing and GORM as the ORM (Object Relational Mapping) tool to interact with a PostgreSQL database.

## Features

- **Create**: Add new personalities to the database.
- **Read**: Retrieve a list of all personalities or a single personality by ID.
- **Update**: Modify existing personalities' details.
- **Delete**: Remove personalities from the database.

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/WENDELLDELIMA/golang-gorm
   ```

2. **Navigate to the project directory:**

   ```sh
   cd repository-name
   ```

3. **Install dependencies:**

   Ensure you have Go installed, then run:

   ```sh
   go mod tidy
   ```

4. **Set up PostgreSQL database:**

   Ensure PostgreSQL is installed and running on your system. Create a database for your project.

5. **Configure database connection:**

   Update the database connection settings in `database/database.go` to match your PostgreSQL configuration.

6. **Run the application:**

   ```sh
   go run main.go
   ```

## Endpoints

- **GET /personalidades**: Retrieve all personalities.
- **GET /personalidades/{id}**: Retrieve a personality by ID.
- **POST /personalidades**: Create a new personality.
- **PUT /personalidades/{id}**: Update an existing personality.
- **DELETE /personalidades/{id}**: Delete a personality by ID.

## Example Data

```json
{
  "Id": 1,
  "Nome": "John Doe",
  "Historia": "A brief history about John Doe."
}
```

## Dependencies

This project relies on the following dependencies:

```go
require (
    github.com/gorilla/mux v1.8.1 // indirect
    github.com/jackc/pgpassfile v1.0.0 // indirect
    github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
    github.com/jackc/pgx/v5 v5.5.5 // indirect
    github.com/jackc/puddle/v2 v2.2.1 // indirect
    github.com/jinzhu/inflection v1.0.0 // indirect
    github.com/jinzhu/now v1.1.5 // indirect
    golang.org/x/crypto v0.17.0 // indirect
    golang.org/x/sync v0.1.0 // indirect
    golang.org/x/text v0.14.0 // indirect
    gorm.io/driver/postgres v1.5.9 // indirect
    gorm.io/gorm v1.25.10 // indirect
)
```

## Project Structure

```
/repository-name
├── main.go
├── models
│   └── personalidade.go
├── database
│   └── database.go
└── routes
    └── routes.go
```

- **main.go**: Entry point of the application. Initializes the database, sets up routes, and starts the server.
- **models/personalidade.go**: Defines the `Personalidade` struct and GORM model.
- **database/database.go**: Contains the database connection logic and migrations.
- **routes/routes.go**: Defines the HTTP routes and handlers.

## Contributing

Feel free to submit issues, fork the repository, and send pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License. See the LICENSE file for details.
