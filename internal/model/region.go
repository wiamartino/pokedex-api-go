package model

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);not null"`
	Pokemons []Pokemon `gorm:"many2many:region_pokemons;"`
}
