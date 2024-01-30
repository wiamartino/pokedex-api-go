package main

import (
	"fmt"
	"log"
	"my-pokedex-api/internal/handler"
	"my-pokedex-api/internal/model"
	"my-pokedex-api/pkg/config"
	"my-pokedex-api/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Convert cfg to DSN string
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Connect to the database
	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	db.AutoMigrate(model.Pokemon{})

	// Create a new Fiber instance
	app := fiber.New()

	// Initialize Pokemon handler with db connection
	pokemonHandler := handler.NewPokemonHandler(db)

	// Set up routes
	app.Get("/pokemon", pokemonHandler.GetAllPokemons)
	app.Get("/pokemon/:id", pokemonHandler.GetPokemon)
	app.Post("/pokemon", pokemonHandler.CreatePokemon)
	app.Put("/pokemon/:id", pokemonHandler.UpdatePokemon)
	app.Delete("/pokemon/:id", pokemonHandler.DeletePokemon)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
