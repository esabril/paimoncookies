package world

import (
	"errors"
	"github.com/esabril/paimoncookies/internal/service/world/model"
	repo "github.com/esabril/paimoncookies/test/world/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorld_CreateAgendaSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repo.NewMockIRepo(ctrl)

	t.Log("Test for everyday except Sunday")

	today := "monday"
	configureAgendaSuccessMock(m, today)

	service := &World{
		repo: m,
	}

	result, err := service.CreateAgenda(today)

	assert.NoError(t, err)
	assert.Equal(t, "Понедельник", result.Weekday)
	assert.Equal(t, result.Content.TalentBooks, map[string][]string{"Мондштадт": {"О Свободе"}})
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
	configureAgendaSuccessMock(m, today)
	result, err = service.CreateAgenda(today)

	assert.NoError(t, err)
	assert.Equal(t, "Воскресенье", result.Weekday)
	assert.Equal(t, result.Content.TalentBooks, map[string][]string(nil))
	assert.Equal(t, result.Content.WeaponMaterials, map[string][]model.WeaponMaterial(nil))
	assert.True(t, result.SystemData.IsSunday)

}

func TestWorld_CreateAgendaFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	today := "monday"
	m := repo.NewMockIRepo(ctrl)

	m.
		EXPECT().
		GetWeekdayTalentBooksWithLocation(today).
		DoAndReturn(func(w string) ([]model.TalentBook, error) {
			return nil, errors.New("something wrong with database")
		}).MaxTimes(1).MinTimes(1)

	service := &World{
		repo: m,
	}

	result, err := service.CreateAgenda(today)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func configureAgendaSuccessMock(m *repo.MockIRepo, today string) {
	m.
		EXPECT().
		GetWeekdayTalentBooksWithLocation(today).
		DoAndReturn(func(w string) ([]model.TalentBook, error) {
			if w == "sunday" {
				return []model.TalentBook{}, nil
			}

			return []model.TalentBook{
				{
					Title:    "О Свободе",
					Location: "Мондштадт",
				},
			}, nil
		}).MaxTimes(1).MinTimes(0)

	m.
		EXPECT().
		GetWeekdayWeaponMaterialsWithLocation(today).
		DoAndReturn(func(w string) ([]model.WeaponMaterial, error) {
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
		}).MaxTimes(1).MinTimes(0)
}
