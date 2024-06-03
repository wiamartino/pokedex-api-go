package model

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(100);not null"`
	Type        PokeType `gorm:"many2many:pokemon_types;"`
	Description string   `gorm:"type:text;not null"`
	Height      float64
	Weight      float64
	HP          int
	Attack      int
	Defense     int
	Speed       int
	Evolution   string
	Abilities   []Ability `gorm:"many2many:pokemon_abilities;"`
	Moves       []Move    `gorm:"many2many:pokemon_moves;"`
	Generation  int
}
