package world

import (
	"errors"
	"testing"

	"github.com/esabril/paimoncookies/internal/service/world/model"
	repo "github.com/esabril/paimoncookies/internal/service/world/repository"

	"github.com/stretchr/testify/assert"
)

func TestWorld_CreateAgendaSuccess(t *testing.T) {
	m := repo.Mock{
		GetWeekdayTalentBooksWithLocationFunc: func(w string) ([]model.TalentBook, error) {
			if w == "sunday" {
				return []model.TalentBook{}, nil
			}

			return []model.TalentBook{
				{
					Title:    "О Свободе",
					Location: "Мондштадт",
				},
			}, nil
		},
		GetWeekdayWeaponMaterialsWithLocationFunc: func(w string) ([]model.WeaponMaterial, error) {
			if w == "sunday" {
				return []model.WeaponMaterial{}, nil
			}

			return []model.WeaponMaterial{
				{
					Title:    "Плитки Декарабиана",
					Location: "Мондштадт",
					Alias:    "плиточки",
				},
			}, nil
		},
		GetRegionsFunc: func() ([]model.Region, error) { return nil, nil },
	}

	t.Log("Test for everyday except Sunday")

	today := "monday"
	service := &World{
		repo: m,
	}

	result, err := service.CreateAgenda(today)

	assert.NoError(t, err)
	assert.Equal(t, "понедельник", result.Weekday)
	assert.Equal(t, result.Content.TalentBooks, map[string][]model.TalentBook{
		"Мондштадт": {
			{
				Title:    "О Свободе",
				Location: "Мондштадт",
			},
		},
	})
	assert.Equal(
		t,
		result.Content.WeaponMaterials,
		map[string][]model.WeaponMaterial{
			"Мондштадт": {
				{
					Title: "Плитки Декарабиана",
					Alias: "плиточки",
				},
			},
		})
	assert.False(t, result.SystemData.IsSunday)

	t.Log("Test for Sunday")

	today = "sunday"
	result, err = service.CreateAgenda(today)

	assert.NoError(t, err)
	assert.Equal(t, "воскресенье", result.Weekday)
	assert.Equal(t, result.Content.TalentBooks, map[string][]model.TalentBook(nil))
	assert.Equal(t, result.Content.WeaponMaterials, map[string][]model.WeaponMaterial(nil))
	assert.True(t, result.SystemData.IsSunday)

}

func TestWorld_CreateAgendaFailed(t *testing.T) {
	today := "monday"

	m := repo.Mock{
		GetWeekdayTalentBooksWithLocationFunc: func(w string) ([]model.TalentBook, error) {
			return nil, errors.New("something wrong with database")
		},
	}

	service := &World{
		repo: m,
	}

	result, err := service.CreateAgenda(today)

	assert.Error(t, err)
	assert.Nil(t, result)
}
