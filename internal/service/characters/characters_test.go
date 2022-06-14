package characters

import (
	"errors"
	"github.com/esabril/paimoncookies/internal/service/characters/model"
	characters_repo "github.com/esabril/paimoncookies/test/characters/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacters_GetInitialCharactersListSuccessful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := characters_repo.NewMockICharactersRepo(ctrl)

	repo.EXPECT().GetCharactersList().DoAndReturn(func() ([]model.Character, error) {
		return []model.Character{
			{
				Title:   "Кэйа",
				Element: "Крио",
			},
			{
				Title:   "Тарталья",
				Element: "Гидро",
			},
			{
				Title:   "Эола",
				Element: "Крио",
			},
		}, nil
	}).MinTimes(1).MaxTimes(1)

	s := Characters{
		repo: repo,
	}

	elements, characters, err := s.GetInitialCharactersList()

	assert.NoError(t, err)
	assert.Equal(t, map[string]bool{"Кэйа": true, "Тарталья": true, "Эола": true}, characters)
	assert.Equal(
		t,
		map[string][]string{"Крио": {"Кэйа", "Эола"}, "Гидро": {"Тарталья"}},
		elements,
	)
}

func TestCharacters_GetInitialCharactersListFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := characters_repo.NewMockICharactersRepo(ctrl)

	repo.EXPECT().GetCharactersList().DoAndReturn(func() ([]model.Character, error) {
		return nil, errors.New("something wrong with database")
	}).MinTimes(1).MaxTimes(1)

	s := Characters{
		repo: repo,
	}

	elements, characters, err := s.GetInitialCharactersList()

	assert.Error(t, err)
	assert.Equal(t, map[string]bool(nil), characters)
	assert.Equal(t, map[string][]string(nil), elements)
}

func TestCharacters_SimplifyCharacterName(t *testing.T) {
	s := Characters{}

	var TestCases = []struct {
		Name     string
		Expected string
	}{
		{"Кэйа", "Кэйа"},
		{"Дилюк", "Дилюк"},
		{"Камисато Аято", "Аято"},
		{"Аратаки Итто", "Итто"},
		{"Еще какой-то персонаж", "Еще какой-то персонаж"},
	}

	for _, tt := range TestCases {
		assert.Equal(t, tt.Expected, s.SimplifyCharacterName(tt.Name))
	}
}
