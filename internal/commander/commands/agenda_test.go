package commands

import (
	"errors"
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/internal/service/world"
	"github.com/esabril/paimoncookies/internal/service/world/model"
	repo "github.com/esabril/paimoncookies/test/world/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var templatePath = "../template/"

// TestCommander_GetAgendaSixDaysSuccessful Getting Agenda for all weekdays except Sunday
func TestCommander_GetAgendaSixDaysSuccessful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repo.NewMockIWorldRepo(ctrl)
	configureWorldMockRepo(m)

	s := service.Service{
		TodayWeekday: "monday",
		World:        world.NewMock(m),
	}
	c := NewCommander(
		&s,
		templatePath,
		nil,
	)

	expected := `🔔 Что, Путешественник, готов к приключениям?
Сегодня 🗓 *понедельник* и сегодня в Тейвате тебя ждут:

📚 *Книги на таланты*:
Мондштадт: «О Свободе»
Ли Юэ: «О Процветании»

🗡 *Материалы для улучшения оружия:*
Мондштадт: «Плитки Декарабиана» (плиточки)
Ли Юэ: «Столбы Гуюнь»

Запасись смолой и вперед! А Паймон всегда будет с тобой! 💫`

	assert.Equal(t, expected, c.GetAgenda())
}

// TestCommander_GetAgendaSundaySuccessful Getting Agenda for Sunday
func TestCommander_GetAgendaSundaySuccessful(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repo.NewMockIWorldRepo(ctrl)
	configureWorldMockRepo(m)

	s := service.Service{
		TodayWeekday: "sunday",
		World:        world.NewMock(m),
	}
	c := NewCommander(
		&s,
		templatePath,
		nil,
	)

	expected := `🔔 Что, Путешественник, готов к приключениям?
Сегодня 🗓 *воскресенье* и сегодня в Тейвате тебя ждут:

📚 *Книги на таланты*:
Все возможные книги во всех открытых тобой Подземельях! Ох и сложный у тебя сегодня выбор! Но Паймон здесь, чтобы помочь!

🗡 *Материалы для улучшения оружия:*
Сегодня мы можем получить все возможные материалы! Давай выбирать вместе, куда мы сегодня отправимся?

Запасись смолой и вперед! А Паймон всегда будет с тобой! 💫`

	assert.Equal(t, expected, c.GetAgenda())
}

func TestCommander_GetAgendaTemplateFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repo.NewMockIWorldRepo(ctrl)
	configureWorldMockRepo(m)

	s := service.Service{
		TodayWeekday: "monday",
		World:        world.NewMock(m),
	}

	c := NewCommander(
		&s,
		"wrongTemplatePath/",
		nil,
	)

	assert.Equal(t, "Ой, что-то пошло не так. Давай немного подождем, может позже восстановится?", c.GetAgenda())
}

func TestCommander_GetAgendaDataFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repo.NewMockIWorldRepo(ctrl)

	m.
		EXPECT().
		GetWeekdayTalentBooksWithLocation(gomock.Any()).
		DoAndReturn(func(w string) ([]model.TalentBook, error) {
			return nil, errors.New("something wrong with database")
		}).MaxTimes(1).MinTimes(1)

	s := service.Service{
		TodayWeekday: "monday",
		World:        world.NewMock(m),
	}
	c := NewCommander(
		&s,
		templatePath,
		nil,
	)

	expected := `🤔 Что? Ты спрашивала Паймон про «расписание дня»? Кажется, Паймон нечего тебе рассказать прямо сейчас.
Позволь мне немного передохнуть и мы снова поговорим. 🤗`

	assert.Equal(t, expected, c.GetAgenda())
}

func configureWorldMockRepo(m *repo.MockIWorldRepo) {
	m.
		EXPECT().
		GetWeekdayTalentBooksWithLocation(gomock.Any()).
		DoAndReturn(func(w string) ([]model.TalentBook, error) {
			if w == "sunday" {
				return []model.TalentBook{}, nil
			}

			return []model.TalentBook{
				{
					Title:    "О Свободе",
					Location: "Мондштадт",
				},
				{
					Title:    "О Процветании",
					Location: "Ли Юэ",
				},
			}, nil
		}).MaxTimes(1).MinTimes(0)

	m.
		EXPECT().
		GetWeekdayWeaponMaterialsWithLocation(gomock.Any()).
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
				{
					Title:    "Столбы Гуюнь",
					Location: "Ли Юэ",
				},
			}, nil
		}).MaxTimes(1).MinTimes(0)

	m.
		EXPECT().
		GetRegions().
		DoAndReturn(func() ([]model.Region, error) {
			return []model.Region{
				{
					Name:  "mondstadt",
					Title: "Мондштадт",
				},
				{
					Name:  "liyue",
					Title: "Ли Юэ",
				},
			}, nil
		}).MaxTimes(1).MinTimes(1)
}
