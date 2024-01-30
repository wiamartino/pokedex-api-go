package handler

import (
	"database/sql"
	"fmt"
	"my-pokedex-api/internal/model"
	"my-pokedex-api/internal/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PokemonHandler struct {
	Repo *repository.PokemonRepository
}

func NewPokemonHandler(db *gorm.DB) *PokemonHandler {
	return &PokemonHandler{
		Repo: repository.NewPokemonRepository(db),
	}
}

func (h *PokemonHandler) GetAllPokemons(c *fiber.Ctx) error {
	pokemons, err := h.Repo.GetAllPokemons()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error", "detail": err.Error()})
	}
	return c.JSON(pokemons)
}

func (h *PokemonHandler) GetPokemon(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID", "detail": "The ID provided is not a valid unsigned integer"})
	}

	pokemon, err := h.Repo.FindById(idInt)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pokemon not found"})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error", "detail": err.Error()})
		}
	}
	return c.JSON(pokemon)
}

func (h *PokemonHandler) CreatePokemon(c *fiber.Ctx) error {
	p := new(model.Pokemon)
	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON", "details": err.Error()})
	}
	if err := h.Repo.Save(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot save Pokemon", "details": err.Error()})
	}
	return c.JSON(p)
}

func (h *PokemonHandler) UpdatePokemon(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	p := new(model.Pokemon)
	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Cannot parse JSON: %v", err)})
	}

	// Convert id to uint
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid ID: %v", err)})
	}
	p.ID = uint(idInt)

	updatedPokemon := h.Repo.Update(p)
	if updatedPokemon == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pokemon not found"})
	}
	return c.JSON(updatedPokemon)
}

func (h *PokemonHandler) DeletePokemon(c *fiber.Ctx) error {
	id := c.Params("id")
	p := new(model.Pokemon)

	// Convert id to uint
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	p.ID = uint(idInt)

	err = h.Repo.Delete(p)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pokemon not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
