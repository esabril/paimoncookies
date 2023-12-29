package archive

import (
	"errors"

	cModel "github.com/esabril/paimoncookies/internal/service/characters/model"
	characters_repo "github.com/esabril/paimoncookies/internal/service/characters/repository"
	wModel "github.com/esabril/paimoncookies/internal/service/world/model"
	world_repo "github.com/esabril/paimoncookies/internal/service/world/repository"

	// characters_repo "github.com/esabril/paimoncookies/test/characters/repository"
	// world_repo "github.com/esabril/paimoncookies/test/world/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArchive_GetCharacterInfoSuccessful(t *testing.T) {
	w := world_repo.Mock{
		GetTalentBookByTypeFunc: func(bookType string) (wModel.TalentBook, error) {
			return wModel.TalentBook{
				Title: "О Поэзии",
			}, nil
		},
		GetTalentBookWeekdaysFunc: func(bookType string) ([]string, error) {
			return []string{
				"wednesday",
				"saturday",
			}, nil
		},
		GetAscensionMaterialsByNamesFunc: func(names []string) ([]wModel.AscensionMaterial, error) {
			return []wModel.AscensionMaterial{
				{
					Title: "Лилия Калла",
					Type:  "local_speciality",
				},
				{
					Title: "Печати похитителей сокровищ",
					Type:  "common",
				},
			}, nil
		},
		GetGemByNameFunc: func(name string) (wModel.Gem, error) {
			return wModel.Gem{
				Name:  "shivada_jade",
				Title: "Нефрит Шивада",
			}, nil
		},
		GetGemDropInfoByNameFunc: func(name string) ([]wModel.BossDrop, error) {
			return []wModel.BossDrop{
				{
					Boss: "Крио папоротник",
					Type: "world",
				},
				{
					Boss: "Андриус",
					Type: "weekly",
				},
			}, nil
		},
		GetWorldBossDropByNameFunc: func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Boss: "Крио папоротник",
			}, nil
		},
		GetWeeklyBossDropByNameFunc: func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Boss: "Андриус",
			}, nil
		},
	}

	c := characters_repo.Mock{
		GetCharacterByNameFunc: func(name string) (cModel.Character, error) {
			return cModel.Character{
				Title:                    "Кэйа",
				Region:                   "Мондштадт",
				Rarity:                   4,
				Element:                  "Крио",
				TalentBookType:           "ballad",
				TalentBossDrop:           "spirit_locker_of_boreas",
				AscensionBossDrop:        "hoarfrost_core",
				AscensionGem:             "shivada_jade",
				AscensionLocalSpeciality: "calla_lily",
				CommonAscensionMaterial:  "treasure_hoarder_insignias",
			}, nil
		},
	}

	expectedCharacter := cModel.Character{
		Title:                    "Кэйа",
		Region:                   "Мондштадт",
		Rarity:                   4,
		Element:                  "Крио",
		TalentBookType:           "ballad",
		AscensionLocalSpeciality: "calla_lily",
		CommonAscensionMaterial:  "treasure_hoarder_insignias",
		AscensionGem:             "shivada_jade",
		TalentBossDrop:           "spirit_locker_of_boreas",
		AscensionBossDrop:        "hoarfrost_core",
		Materials: cModel.CharacterMaterials{
			Ascension: cModel.CharacterAscension{
				Gem: wModel.Gem{
					Name:  "shivada_jade",
					Title: "Нефрит Шивада",
					DropInfo: wModel.GemDropInfo{
						WorldBosses: []wModel.WorldBoss{
							{
								Title: "Крио папоротник",
							},
						},
						WeeklyBosses: []wModel.WeeklyBoss{
							{
								Title: "Андриус",
							},
						},
					},
				},
				BossDrop: wModel.BossDrop{
					Boss: "Крио папоротник",
				},
				LocalSpeciality: wModel.AscensionMaterial{
					Title: "Лилия Калла",
					Type:  "local_speciality",
				},
				CommonMaterial: wModel.AscensionMaterial{
					Title: "Печати похитителей сокровищ",
					Type:  "common",
				},
			},
			TalentUpgrade: cModel.CharacterTalentUpgrade{
				TalentBook: wModel.TalentBook{
					Title: "О Поэзии",
					Weekdays: []string{
						"среда",
						"суббота",
						"воскресенье",
					},
				},
				BossDrop: wModel.BossDrop{
					Boss: "Андриус",
				},
				CommonMaterial: wModel.AscensionMaterial{
					Title: "Печати похитителей сокровищ",
					Type:  "common",
				},
			},
		},
	}

	a := NewMock(w, c)

	result, err := a.GetCharacterInfo("Кэйа")

	assert.NoError(t, err)
	assert.Equal(t, expectedCharacter, result)
}

func TestArchive_GetCharacterInfoWorldFail(t *testing.T) {
	c := characters_repo.Mock{
		GetCharacterByNameFunc: func(name string) (cModel.Character, error) {
			return cModel.Character{
				Title:                    "Кэйа",
				TalentBookType:           "ballad",
				TalentBossDrop:           "spirit_locker_of_boreas",
				AscensionBossDrop:        "hoarfrost_core",
				AscensionGem:             "shivada_jade",
				AscensionLocalSpeciality: "calla_lily",
				CommonAscensionMaterial:  "treasure_hoarder_insignias",
			}, nil
		},
	}

	w := world_repo.Mock{
		GetAscensionMaterialsByNamesFunc: func(names []string) ([]wModel.AscensionMaterial, error) {
			return nil, errors.New("something wrong with database")
		},
	}

	a := NewMock(w, c)

	result, err := a.GetCharacterInfo("Кэйа")

	assert.Error(t, err)
	assert.Equal(t, "", result.Title)
}

func TestArchive_GetCharacterInfoCharactersFail(t *testing.T) {
	c := characters_repo.Mock{
		GetCharacterByNameFunc: func(name string) (cModel.Character, error) {
			return cModel.Character{}, errors.New("something wrong with database")
		},
	}

	w := world_repo.Mock{
		GetTalentBookByTypeFunc: func(bookType string) (wModel.TalentBook, error) {
			return wModel.TalentBook{}, errors.New("something wrong with database")
		},
	}

	a := NewMock(w, c)

	result, err := a.GetCharacterInfo("Кэйа")

	assert.Error(t, err)
	assert.Equal(t, "", result.Title)
}
