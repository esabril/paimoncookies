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
				Title:          "Кэйа",
				Region:         "Мондштадт",
				Rarity:         4,
				Element:        "Крио",
				TalentBookType: "ballad",
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

	expectedCharacter := cModel.Character{
		Title:          "Кэйа",
		Region:         "Мондштадт",
		Rarity:         4,
		Element:        "Крио",
		TalentBookType: "ballad",
		Materials: cModel.CharacterMaterials{
			TalentBook: wModel.TalentBook{
				Title: "О Поэзии",
				Weekdays: []string{
					"среда",
					"суббота",
					"воскресенье",
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
				Title:          "Кэйа",
				TalentBookType: "ballad",
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

	a := NewMock(w, c)

	result, err := a.GetCharacterInfo("Кэйа")

	assert.Error(t, err)
	assert.Equal(t, "", result.Title)
}
