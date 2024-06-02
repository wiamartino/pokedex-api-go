package repository

import (
	"my-pokedex-api/internal/model"

	"gorm.io/gorm"
)

type PokemonRepository struct {
	DB *gorm.DB
}

func NewPokemonRepository(db *gorm.DB) *PokemonRepository {
	return &PokemonRepository{DB: db}
}

func (r *PokemonRepository) GetAllPokemons() ([]model.Pokemon, error) {
	var pokemons []model.Pokemon
	if err := r.DB.Find(&pokemons).Error; err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (r *PokemonRepository) FindById(id uint64) (*model.Pokemon, error) {
	var pokemon model.Pokemon
	if err := r.DB.First(&pokemon, id).Error; err != nil {
		return nil, err
	}
	return &pokemon, nil
}

func (r *PokemonRepository) Save(pokemon *model.Pokemon) error {
	return r.DB.Save(pokemon).Error
}

func (r *PokemonRepository) Update(pokemon *model.Pokemon) error {
	return r.DB.Model(pokemon).Updates(pokemon).Error
}

func (r *PokemonRepository) Delete(pokemon *model.Pokemon) error {
	return r.DB.Delete(pokemon).Error
}

func (r *PokemonRepository) SearchByName(name string) (*model.Pokemon, error) {
	var pokemon model.Pokemon
	if err := r.DB.Where("name = ?", name).First(&pokemon).Error; err != nil {
		return nil, err
	}
	return &pokemon, nil
}

func (r *PokemonRepository) GetPokemonsByAbility(abilityName string) ([]model.Pokemon, error) {
	var pokemons []model.Pokemon
	if err := r.DB.Joins("Abilities").Where("abilities.name = ?", abilityName).Find(&pokemons).Error; err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (r *PokemonRepository) GetRandomPokemon() (*model.Pokemon, error) {
	var pokemon model.Pokemon
	if err := r.DB.Order("RANDOM()").First(&pokemon).Error; err != nil {
		return nil, err
	}
	return &pokemon, nil
}
