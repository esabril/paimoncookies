package archive

import (
	"errors"
	cModel "github.com/esabril/paimoncookies/internal/service/characters/model"
	wModel "github.com/esabril/paimoncookies/internal/service/world/model"
	characters_repo "github.com/esabril/paimoncookies/test/characters/repository"
	world_repo "github.com/esabril/paimoncookies/test/world/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArchive_GetCharacterInfoSuccessful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := world_repo.NewMockIWorldRepo(ctrl)
	c := characters_repo.NewMockICharactersRepo(ctrl)

	c.EXPECT().
		GetCharacterByName("Кэйа").
		DoAndReturn(func(name string) (cModel.Character, error) {
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
		}).MaxTimes(1).MaxTimes(1)

	w.EXPECT().
		GetTalentBookByType("ballad").
		DoAndReturn(func(bookType string) (wModel.TalentBook, error) {
			return wModel.TalentBook{
				Title: "О Поэзии",
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetTalentBookWeekdays("ballad").
		DoAndReturn(func(bookType string) ([]string, error) {
			return []string{
				"wednesday",
				"saturday",
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetAscensionMaterialsByNames(gomock.Any()).
		DoAndReturn(func(names []string) ([]wModel.AscensionMaterial, error) {
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
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetGemByName("shivada_jade").
		DoAndReturn(func(name string) (wModel.Gem, error) {
			return wModel.Gem{
				Name:  "shivada_jade",
				Title: "Нефрит Шивада",
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetGemDropInfoByName("shivada_jade").
		DoAndReturn(func(name string) ([]wModel.BossDrop, error) {
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
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetWorldBossDropByName("hoarfrost_core").
		DoAndReturn(func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Boss: "Крио папоротник",
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetWeeklyBossDropByName("spirit_locker_of_boreas").
		DoAndReturn(func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Boss: "Андриус",
			}, nil
		}).MinTimes(1).MaxTimes(1)

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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := world_repo.NewMockIWorldRepo(ctrl)
	c := characters_repo.NewMockICharactersRepo(ctrl)

	c.EXPECT().
		GetCharacterByName("Кэйа").
		DoAndReturn(func(name string) (cModel.Character, error) {
			return cModel.Character{
				Title:                    "Кэйа",
				TalentBookType:           "ballad",
				TalentBossDrop:           "spirit_locker_of_boreas",
				AscensionBossDrop:        "hoarfrost_core",
				AscensionGem:             "shivada_jade",
				AscensionLocalSpeciality: "calla_lily",
				CommonAscensionMaterial:  "treasure_hoarder_insignias",
			}, nil
		}).MaxTimes(1).MaxTimes(1)

	w.EXPECT().
		GetTalentBookByType("ballad").
		DoAndReturn(func(bookType string) (wModel.TalentBook, error) {
			return wModel.TalentBook{}, errors.New("something wrong with database")
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetTalentBookWeekdays("ballad").
		DoAndReturn(func(bookType string) ([]string, error) {
			return []string{
				"wednesday",
				"saturday",
			}, nil
		}).MinTimes(0).MaxTimes(0)

	w.EXPECT().
		GetAscensionMaterialsByNames(gomock.Any()).
		DoAndReturn(func(names []string) ([]wModel.AscensionMaterial, error) {
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
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetGemByName("shivada_jade").
		DoAndReturn(func(name string) (wModel.Gem, error) {
			return wModel.Gem{
				Name:  "shivada_jade",
				Title: "Нефрит Шивада",
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetGemDropInfoByName("shivada_jade").
		DoAndReturn(func(name string) ([]wModel.BossDrop, error) {
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
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetWorldBossDropByName("hoarfrost_core").
		DoAndReturn(func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Boss: "Крио папоротник",
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetWeeklyBossDropByName("spirit_locker_of_boreas").
		DoAndReturn(func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Boss: "Андриус",
			}, nil
		}).MinTimes(0).MaxTimes(0)

	a := NewMock(w, c)

	result, err := a.GetCharacterInfo("Кэйа")

	assert.Error(t, err)
	assert.Equal(t, "", result.Title)
}

func TestArchive_GetCharacterInfoCharactersFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := world_repo.NewMockIWorldRepo(ctrl)
	c := characters_repo.NewMockICharactersRepo(ctrl)

	c.EXPECT().
		GetCharacterByName("Кэйа").
		DoAndReturn(func(name string) (cModel.Character, error) {
			return cModel.Character{}, errors.New("something wrong with database")
		}).MaxTimes(1).MaxTimes(1)

	w.EXPECT().
		GetTalentBookByType("ballad").
		DoAndReturn(func(bookType string) (wModel.TalentBook, error) {
			return wModel.TalentBook{}, errors.New("something wrong with database")
		}).MinTimes(0).MaxTimes(0)

	w.EXPECT().
		GetTalentBookWeekdays("ballad").
		DoAndReturn(func(bookType string) ([]string, error) {
			return []string{
				"wednesday",
				"saturday",
			}, nil
		}).MinTimes(0).MaxTimes(0)

	w.EXPECT().
		GetAscensionMaterialsByNames(gomock.Any()).
		DoAndReturn(func(names []string) ([]wModel.AscensionMaterial, error) {
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
		}).MinTimes(0).MaxTimes(0)

	w.EXPECT().
		GetGemByName("shivada_jade").
		DoAndReturn(func(name string) (wModel.Gem, error) {
			return wModel.Gem{
				Name:  "shivada_jade",
				Title: "Нефрит Шивада",
			}, nil
		}).MinTimes(0).MaxTimes(0)

	w.EXPECT().
		GetGemDropInfoByName("shivada_jade").
		DoAndReturn(func(name string) ([]wModel.BossDrop, error) {
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
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetWorldBossDropByName("hoarfrost_core").
		DoAndReturn(func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Boss: "Крио папоротник",
			}, nil
		}).MinTimes(0).MaxTimes(0)

	w.EXPECT().
		GetWeeklyBossDropByName("spirit_locker_of_boreas").
		DoAndReturn(func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Boss: "Андриус",
			}, nil
		}).MinTimes(0).MaxTimes(0)

	a := NewMock(w, c)

	result, err := a.GetCharacterInfo("Кэйа")

	assert.Error(t, err)
	assert.Equal(t, "", result.Title)
}
