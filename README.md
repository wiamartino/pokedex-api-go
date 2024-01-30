# My Pokedex API

This project is a RESTful API for a Pokedex, built with Go, Fiber, and GORM.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or later
- A running PostgreSQL database

### Installing

1. Clone the repository
```
git clone https://github.com/wiamartino/my-pokedex-api.git
```
2. Navigate to the project directory
```
cd my-pokedex-api
```
3. Install the dependencies
```
go mod download
```
4. Set the necessary environment variables (DB_HOST, DB_NAME, DB_USER, DB_PASSWORD)
5. Run the application
```
go run cmd/main.go
```

## API Endpoints

- `GET /pokemon` - Get all Pokemons
- `GET /pokemon/:id` - Get a Pokemon by ID
- `POST /pokemon` - Create a new Pokemon
- `PUT /pokemon/:id` - Update a Pokemon by ID
- `DELETE /pokemon/:id` - Delete a Pokemon by ID

## Built With

- [Go](https://golang.org/) - The programming language used
- [Fiber](https://gofiber.io/) - The web framework used
- [GORM](https://gorm.io/) - The ORM library used

