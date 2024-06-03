package main

import (
	"log"

	"my-pokedex-api/internal/model"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {
	pokemons := []model.Pokemon{
		{Name: "Pikachu", Abilities: []model.Ability{{Name: "Electric"}}},
		{Name: "Charizard", Abilities: []model.Ability{{Name: "Fire"}}},
	}

	for _, pokemon := range pokemons {
		if err := db.Create(&pokemon).Error; err != nil {
			log.Fatalf("Failed to insert the pokemon: %v", err)
		}
	}
}
