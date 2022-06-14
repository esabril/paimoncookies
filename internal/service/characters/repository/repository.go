package repository

import (
	"github.com/esabril/paimoncookies/internal/service/characters/model"
	"github.com/jmoiron/sqlx"
)

type ICharactersRepo interface {
	GetCharactersList() ([]model.Character, error)
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
	if err != nil {
		return nil, err
	}

	return characters, nil
}
