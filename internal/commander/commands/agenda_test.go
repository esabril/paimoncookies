package commands

import (
	"errors"

	"testing"

	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/internal/service/world"
	"github.com/esabril/paimoncookies/internal/service/world/model"
	repo "github.com/esabril/paimoncookies/internal/service/world/repository"

	"github.com/stretchr/testify/assert"
)

const DefaultTemplatePath = "../template/"

// TestCommander_GetAgendaSixDaysSuccessful Getting Agenda for all weekdays except Sunday
func TestCommander_GetAgendaSixDaysSuccessful(t *testing.T) {
	s := service.Service{
		TodayWeekday: "monday",
		World:        world.NewMock(getRepoMock()),
	}
	c := NewCommander(
		&s,
		DefaultTemplatePath,
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
	s := service.Service{
		TodayWeekday: "sunday",
		World:        world.NewMock(getRepoMock()),
	}
	c := NewCommander(
		&s,
		DefaultTemplatePath,
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
	s := service.Service{
		TodayWeekday: "monday",
		World:        world.NewMock(getRepoMock()),
	}

	c := NewCommander(
		&s,
		"wrongTemplatePath/",
		nil,
	)

	assert.Equal(t, "Ой, что-то пошло не так. Давай немного подождем, может позже восстановится?", c.GetAgenda())
}

func TestCommander_GetAgendaDataFail(t *testing.T) {
	m := getRepoMock()
	m.GetWeekdayTalentBooksWithLocationFunc = func(w string) ([]model.TalentBook, error) {
		return nil, errors.New("something wrong with database")
	}

	s := service.Service{
		TodayWeekday: "monday",
		World:        world.NewMock(m),
	}
	c := NewCommander(
		&s,
		DefaultTemplatePath,
		nil,
	)

	expected := `🤔 Что? Ты спрашивала Паймон про «расписание дня»? Кажется, Паймон нечего тебе рассказать прямо сейчас.
Позволь мне немного передохнуть и мы снова поговорим. 🤗`

	assert.Equal(t, expected, c.GetAgenda())
}

func getRepoMock() repo.Mock {
	return repo.Mock{
		GetWeekdayTalentBooksWithLocationFunc: func(w string) ([]model.TalentBook, error) {
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
				{
					Title:    "Столбы Гуюнь",
					Location: "Ли Юэ",
				},
			}, nil
		},
		GetRegionsFunc: func() ([]model.Region, error) {
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
		},
	}
}
