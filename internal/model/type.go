package model

import (
	"gorm.io/gorm"
)

type PokeType struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"`
}
