package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

// Animal represents the structure of an animal entity.
type Animal struct {
	ID            int    `json:"id"`
	KindOfAnimal  string `json:"kindOfAnimal"`
	BreedOfAnimal string `json:"breedOfAnimal"`
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Description   string `json:"description"`
	AnimalPicture string `json:"animalPicture"`
}

// AnimalModel represents methods to interact with the "animals" table in the database.
type AnimalModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

// Insert inserts a new animal into the database.
func (a AnimalModel) Insert(animal *Animal) error {
	query := `
		INSERT INTO animals (kind_of_animal, breed_of_animal, name, age, description, animal_picture) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id
	`
	args := []interface{}{
		animal.KindOfAnimal,
		animal.BreedOfAnimal,
		animal.Name,
		animal.Age,
		animal.Description,
		animal.AnimalPicture,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := a.DB.QueryRowContext(ctx, query, args...).Scan(&animal.ID)
	return err
}

// Get retrieves an animal from the database based on the provided animal ID.
func (a AnimalModel) Get(animalID int) (*Animal, error) {
	query := `
		SELECT id, kind_of_animal, breed_of_animal, name, age, description, animal_picture
		FROM animals
		WHERE id = $1
	`
	var animal Animal
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := a.DB.QueryRowContext(ctx, query, animalID)
	err := row.Scan(
		&animal.ID,
		&animal.KindOfAnimal,
		&animal.BreedOfAnimal,
		&animal.Name,
		&animal.Age,
		&animal.Description,
		&animal.AnimalPicture,
	)
	if err != nil {
		return nil, err
	}
	return &animal, nil
}

// Similar methods for Update and Delete can be added based on your requirements.
