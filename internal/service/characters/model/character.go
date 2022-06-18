package model

import "github.com/esabril/paimoncookies/internal/service/world/model"

type CharacterMaterials struct {
	TalentBook model.TalentBook
}

// Character structure
type Character struct {
	Id             int                `json:"id"`
	Name           string             `json:"name"`
	Title          string             `json:"title"`
	Region         string             `json:"region,omitempty"`
	Rarity         int8               `json:"rarity,omitempty"`
	Element        string             `json:"element,omitempty"`
	TalentBookType string             `json:"talent_book_type,omitempty" db:"talent_book_type"`
	Materials      CharacterMaterials `json:"materials,omitempty" db:"-"`
}
