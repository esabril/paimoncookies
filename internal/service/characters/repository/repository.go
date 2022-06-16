package repository

import (
	"errors"
	"github.com/esabril/paimoncookies/internal/service/characters/model"
	"github.com/jmoiron/sqlx"
)

type ICharactersRepo interface {
	GetCharactersList() ([]model.Character, error)
	GetCharacterByName(name string) (model.Character, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) ICharactersRepo {
	return &repo{
		db: db,
	}
}

func (r *repo) GetCharactersList() ([]model.Character, error) {
	var characters []model.Character

	err := r.db.Select(&characters, `SELECT c.title AS title, e.title AS element 
		FROM character c JOIN element e ON c.element = e.name 
		ORDER BY c.id`,
	)

	return characters, err
}

func (r *repo) GetCharacterByName(name string) (model.Character, error) {
	var c []model.Character
	args := map[string]interface{}{
		"name": name,
	}

	stmt, err := r.db.PrepareNamed(
		`SELECT c.title AS title, e.title AS element, rarity, talent_book_type, r.title AS region 
		FROM character c 
		    JOIN element e ON c.element = e.name
			JOIN region r ON c.region = r.name
		WHERE c.title = :name
		LIMIT 1`)

	if err != nil {
		return model.Character{}, err
	}

	err = stmt.Select(&c, args)

	if len(c) == 0 {
		return model.Character{}, errors.New("character not found")
	}

	return c[0], err
}
