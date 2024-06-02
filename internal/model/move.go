package model

import (
	"gorm.io/gorm"
)

type Move struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Power    int
	Type     string    `gorm:"type:varchar(100);not null"`
	Pokemons []Pokemon `gorm:"many2many:pokemon_moves;"`
}
