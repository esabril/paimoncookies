package archive

import (
	"github.com/esabril/paimoncookies/internal/service/characters"
	cModel "github.com/esabril/paimoncookies/internal/service/characters/model"
	cRepo "github.com/esabril/paimoncookies/internal/service/characters/repository"
	"github.com/esabril/paimoncookies/internal/service/world"
	wRepo "github.com/esabril/paimoncookies/internal/service/world/repository"
)

// Archive service handle common information about all data in the game
type Archive struct {
	world      world.WorldInterface
	characters characters.CharacterInterface
}

func NewArchive(w world.WorldInterface, c characters.CharacterInterface) *Archive {
	return &Archive{
		world:      w,
		characters: c,
	}
}

func NewMock(wrepo wRepo.RepositoryInterface, crepo cRepo.RepositoryInterface) *Archive {
	return &Archive{
		world:      world.NewMock(wrepo),
		characters: characters.NewMock(crepo, nil, nil),
	}
}

func (a *Archive) GetCharacterInfo(name string) (cModel.Character, error) {
	character, err := a.characters.GetCharacterByName(name)
	if err != nil {
		return cModel.Character{}, err
	}

	amNames := []string{character.CommonAscensionMaterial, character.AscensionLocalSpeciality}
	commonAm, lsAm, err := a.world.GetCommonLocalSpecAscensionMaterialsByNames(amNames)
	if err != nil {
		return cModel.Character{}, err
	}

	ac, err := a.GetCharacterAscension(character)
	if err != nil {
		return cModel.Character{}, err
	}

	ac.LocalSpeciality = lsAm
	ac.CommonMaterial = commonAm

	tu, err := a.GetCharacterTalentUpgrade(character)
	if err != nil {
		return cModel.Character{}, err
	}

	tu.CommonMaterial = commonAm

	character.Materials.Ascension = ac
	character.Materials.TalentUpgrade = tu

	return character, nil
}

func (a *Archive) GetCharacterAscension(character cModel.Character) (cModel.CharacterAscension, error) {
	gem, err := a.world.GetGemByName(character.AscensionGem)
	if err != nil {
		return cModel.CharacterAscension{}, err
	}

	worldBd, err := a.world.GetWorldBossDropByName(character.AscensionBossDrop)
	if err != nil {
		return cModel.CharacterAscension{}, err
	}

	return cModel.CharacterAscension{
		Gem:      gem,
		BossDrop: worldBd,
	}, nil
}

func (a *Archive) GetCharacterTalentUpgrade(character cModel.Character) (cModel.CharacterTalentUpgrade, error) {
	talentBook, err := a.world.GetTalentBookByName(character.TalentBookType)
	if err != nil {
		return cModel.CharacterTalentUpgrade{}, err
	}

	weeklyBd, err := a.world.GetWeeklyBossDropByName(character.TalentBossDrop)
	if err != nil {
		return cModel.CharacterTalentUpgrade{}, err
	}

	return cModel.CharacterTalentUpgrade{
		TalentBook: talentBook,
		BossDrop:   weeklyBd,
	}, nil
}
