package model

import "github.com/esabril/paimoncookies/internal/service/world/model"

type CharacterMaterials struct {
	TalentBook model.TalentBook
}

// Character structure
type Character struct {
	Id             int
	Name           string
	Title          string
	Region         string
	Rarity         int8
	Element        string
	TalentBookType string             `db:"talent_book_type"`
	Materials      CharacterMaterials `db:"-"`
}
