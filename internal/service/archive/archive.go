package archive

import (
	"github.com/esabril/paimoncookies/internal/service/characters"
	"github.com/esabril/paimoncookies/internal/service/characters/model"
	"github.com/esabril/paimoncookies/internal/service/world"
)

// Archive service handle common information about all data in the game
type Archive struct {
	world      *world.World
	characters *characters.Characters
}

func NewArchive(w *world.World, c *characters.Characters) *Archive {
	return &Archive{
		world:      w,
		characters: c,
	}
}

func (a *Archive) GetCharacterInfo(name string) (model.Character, error) {
	character, err := a.characters.GetCharacterByName(name)
	if err != nil {
		return model.Character{}, err
	}

	talentBook, err := a.world.GetTalentBookByName(character.TalentBookType)
	if err != nil {
		return model.Character{}, err
	}

	character.Materials.TalentBook = talentBook

	return character, nil
}
