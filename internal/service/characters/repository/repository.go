package repository

import (
	"errors"
	"fmt"
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

// NonTeyvatWorldChracters "Guest Stars" from another universes
var NonTeyvatWorldChracters = map[string]bool{
	"Элой": true,
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
		"name": fmt.Sprintf("%%%s%%", name),
	}

	_, NonTeyvatCharacter := NonTeyvatWorldChracters[name]
	regionColumn, joinRegion := "", ""

	if NonTeyvatCharacter {
		regionColumn = "c.region AS region"
	} else {
		regionColumn = "r.title AS region"
		joinRegion = "JOIN region r ON c.region = r.name"
	}

	query := fmt.Sprintf(
		`SELECT c.title AS title, e.title AS element, 
				rarity, talent_book_type, talent_boss_drop, ascension_boss_drop, ascension_gem,
				ascension_local_speciality, common_ascension_material,
				%s
			FROM character c 
			JOIN element e ON c.element = e.name %s
			WHERE c.title LIKE :name LIMIT 1`,
		regionColumn,
		joinRegion,
	)

	stmt, err := r.db.PrepareNamed(query)

	if err != nil {
		return model.Character{}, err
	}

	err = stmt.Select(&c, args)

	if len(c) == 0 {
		return model.Character{}, errors.New("character not found")
	}

	return c[0], err
}
