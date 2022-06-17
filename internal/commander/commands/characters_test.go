package commands

import (
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/internal/service/archive"
	"github.com/esabril/paimoncookies/internal/service/characters"
	cModel "github.com/esabril/paimoncookies/internal/service/characters/model"
	wModel "github.com/esabril/paimoncookies/internal/service/world/model"
	characters_repo "github.com/esabril/paimoncookies/test/characters/repository"
	world_repo "github.com/esabril/paimoncookies/test/world/repository"
	"github.com/esabril/paimoncookies/tools/renderer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommander_isCharacter(t *testing.T) {
	chs := map[string]bool{
		"Кэйя":     true,
		"Эола":     true,
		"Цици":     true,
		"Дилюк":    true,
		"Йоимия":   true,
		"Тарталья": true,
	}

	c := Commander{
		service: &service.Service{
			Characters: characters.NewMock(nil, nil, chs),
		},
	}

	testCases := []struct {
		Reply    string
		Expected bool
	}{
		{"Сара", false},
		{"Кэйя", true},
		{"Йоимия", true},
		{"Рэйзор", false},
		{"Тарталья", true},
		{"Кэ Цин", false},
	}

	for _, tt := range testCases {
		t.Run(tt.Reply, func(t *testing.T) {
			assert.Equal(t, tt.Expected, c.isCharacter(tt.Reply))
		})
	}
}

func TestCommander_isElement(t *testing.T) {
	t.Parallel()

	elements := map[string][]string{
		"Крио":  {"Кэйя", "Эола", "Цици"},
		"Пиро":  {"Дилюк", "Йоимия"},
		"Гидро": {"Тарталья"},
	}

	c := Commander{
		service: &service.Service{
			Characters: characters.NewMock(nil, elements, nil),
		},
		renderer: renderer.NewRenderer("path"),
	}

	testCases := []struct {
		Reply    string
		Expected bool
	}{
		{"🔥 Пиро ➡", true},
		{"⬅ 💧 Гидро", true},
		{"❄ Крио", true},
		{"Гидро", true},
		{"🌏 К стихиям", false},
		{"Ответ боту", false},
	}

	for _, tt := range testCases {
		t.Run(tt.Reply, func(t *testing.T) {
			assert.Equal(t, tt.Expected, c.isElement(tt.Reply))
		})
	}
}

func TestCommander_getElementFromReply(t *testing.T) {
	c := Commander{
		renderer: renderer.NewRenderer("path"),
	}

	testCases := []struct {
		Reply    string
		Expected string
	}{
		{"🔥 Пиро ➡", "Пиро"},
		{"⬅ 🔥 Пиро", "Пиро"},
		{"🔥 Пиро", "Пиро"},
		{"Пиро", "Пиро"},
		{"Любой другой текст с пробелами", "Любой другой текст с пробелами"},
	}

	for _, tt := range testCases {
		t.Run(tt.Reply, func(t *testing.T) {
			assert.Equal(t, tt.Expected, c.getElementFromReply(tt.Reply))
		})
	}
}

func TestCommander_GetCharacterMenuRulesSuccessful(t *testing.T) {
	t.Parallel()

	c := Commander{
		renderer: renderer.NewRenderer(DefaultTemplatePath),
	}

	testCases := []struct {
		Element  string
		Expected string
	}{
		{
			Element: "",
			Expected: `Ты можешь двигаться по меню этого раздела внизу, а можешь просто ввести имя персонажа *в любой момент*, в поле отправки (даже не в этом разделе), и тебе тут же покажется вся информация. Удобно, да? 🌸

*Пара правил:*
- имя персонажа нужно вводить таким же, какое оно указано в игре. Я не знаю «Чичу», а вот про милашку Ци Ци расскажу с удовольствием;
- для наших друзей из Инадзумы ты можешь просто ввести его имя. Аратаки Итто, думаю, не обидится, если мы буем искать его просто как «Итто».

Итак, о ком ты хочешь узнать?`,
		},
		{
			Element: "Гидро",
			Expected: `Паймон может рассказать тебе вот об этих персонажах стихии 💧 Гидро

*Пара правил:*
- имя персонажа нужно вводить таким же, какое оно указано в игре. Я не знаю «Чичу», а вот про милашку Ци Ци расскажу с удовольствием;
- для наших друзей из Инадзумы ты можешь просто ввести его имя. Аратаки Итто, думаю, не обидится, если мы буем искать его просто как «Итто».

Итак, о ком ты хочешь узнать?`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Element, func(t *testing.T) {
			assert.Equal(t, tt.Expected, c.GetCharacterMenuRules(tt.Element))
		})
	}
}

func TestCommander_GetCharacterMenuRulesFail(t *testing.T) {
	c := Commander{
		renderer: renderer.NewRenderer("path"),
	}

	assert.Equal(
		t,
		"Я не могу напомнить тебе правила поиска персонажей... но ты ведь и так их помнишь, правда?",
		c.GetCharacterMenuRules("Пиро"),
	)
}

func TestCommander_GetCharacterInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := world_repo.NewMockIWorldRepo(ctrl)
	c := characters_repo.NewMockICharactersRepo(ctrl)

	c.EXPECT().
		GetCharacterByName("Венти").
		DoAndReturn(func(name string) (cModel.Character, error) {
			return cModel.Character{
				Title:          "Венти",
				Region:         "Мондштадт",
				Rarity:         5,
				Element:        "Анемо",
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

	cmdr := Commander{
		service: &service.Service{
			TodayWeekday: "saturday",
			Archive:      archive.NewMock(w, c),
		},
		renderer: renderer.NewRenderer(DefaultTemplatePath),
	}

	expected := `*Венти* 🍃 *5*★
Регион: Мондштадт

📚 *Книги талантов:* «О Поэзии»
Можно получить: среда, 🗓 *суббота*, воскресенье`

	result, element := cmdr.GetCharacterInfo("Венти")
	assert.Equal(t, expected, result)
	assert.Equal(t, "Анемо", element)
}
