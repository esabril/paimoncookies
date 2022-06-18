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
	t.Parallel()

	s := Characters{}

	var TestCases = []struct {
		Name     string
		Exists   bool
		Expected string
	}{
		{"Кэйа", false, "Кэйа"},
		{"Дилюк", false, "Дилюк"},
		{"Камисато Аято", true, "Аято"},
		{"Аратаки Итто", true, "Итто"},
		{"Еще какой-то персонаж", false, "Еще какой-то персонаж"},
	}

	for _, tt := range TestCases {
		t.Run(tt.Name, func(t *testing.T) {
			shortName, ok := s.SimplifyCharacterName(tt.Name)

			assert.Equal(t, tt.Exists, ok)
			assert.Equal(t, tt.Expected, shortName)
		})
	}
}

func TestCharacters_GetElementCharacters(t *testing.T) {
	t.Parallel()

	c := Characters{elements: map[string][]string{
		"Крио": {"Кэйя", "Эола", "Цици"},
	}}

	testCases := []struct {
		Element  string
		First    int
		Last     int
		Expected []string
	}{
		{"Крио", 0, 0, []string{"Кэйя", "Эола", "Цици"}},
		{"Крио", 1, 5, []string{"Эола", "Цици"}},
		{"Крио", 0, 5, []string{"Кэйя", "Эола", "Цици"}},
		{"Крио", 1, 4, []string{"Эола", "Цици"}},
		{"Крио", -1, 2, []string{"Кэйя", "Эола"}},
		{"Крио", 1, 3, []string{"Эола", "Цици"}},
		{"Пиро", 0, 1, []string{}},
	}

	for _, tt := range testCases {
		t.Run(tt.Element, func(t *testing.T) {
			assert.Equal(t, tt.Expected, c.GetElementCharacters(tt.Element, tt.First, tt.Last))
		})
	}
}

func TestCharacters_CheckElement(t *testing.T) {
	t.Parallel()

	c := Characters{elements: map[string][]string{
		"Крио":  {"Кэйя", "Эола", "Цици"},
		"Пиро":  {"Дилюк", "Йоимия"},
		"Гидро": {"Тарталья"},
	}}

	elements := []string{"Крио", "Пиро", "Дендро", "Крио", "Гидро", "Анемо"}

	trueCount, falseCount := 0, 0
	ch := make(chan bool)

	for _, el := range elements {
		go func(element string) {
			ch <- c.CheckElement(element)
		}(el)
	}

	for i := 0; i < len(elements); i++ {
		if <-ch {
			trueCount++
		} else {
			falseCount++
		}
	}

	close(ch)

	assert.Equal(t, 4, trueCount)
	assert.Equal(t, 2, falseCount)
}

func TestCharacters_CheckCharacter(t *testing.T) {
	c := Characters{
		characters: map[string]bool{
			"Кэйя":              true,
			"Эола":              true,
			"Цици":              true,
			"Дилюк":             true,
			"Йоимия":            true,
			"Тарталья":          true,
			"Камисато Аяка":     true,
			"Сангономия Кокоми": true,
		},
	}

	requests := []string{"Кэйя", "Эола", "Цици", "Сара", "Дилюк", "Аяка", "Йоимия", "Тарталья", "Лиза", "Камисато Аяка", "Сангономия Кокоми"}

	trueCount, falseCount := 0, 0
	ch := make(chan bool)

	for _, name := range requests {
		go func(name string) {
			ch <- c.CheckCharacter(name)
		}(name)
	}

	for i := 0; i < len(requests); i++ {
		if <-ch {
			trueCount++
		} else {
			falseCount++
		}
	}

	close(ch)

	assert.Equal(t, 8, trueCount)
	assert.Equal(t, 3, falseCount)
}
