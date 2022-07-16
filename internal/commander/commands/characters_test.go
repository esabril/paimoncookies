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
				Title:                    "Венти",
				Region:                   "Мондштадт",
				Rarity:                   5,
				Element:                  "Анемо",
				TalentBookType:           "ballad",
				TalentBossDrop:           "tail_of_boreas",
				AscensionBossDrop:        "hurricane_seed",
				AscensionGem:             "vayuda_turquoise",
				AscensionLocalSpeciality: "cecilia",
				CommonAscensionMaterial:  "slime",
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
					Title: "Сесилия",
					Type:  "local_speciality",
				},
				{
					Title: "Слаймы",
					Type:  "common",
				},
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetGemByName("vayuda_turquoise").
		DoAndReturn(func(name string) (wModel.Gem, error) {
			return wModel.Gem{
				Name:  "vayuda_turquoise",
				Title: "Бирюза Вайюда",
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetGemDropInfoByName("vayuda_turquoise").
		DoAndReturn(func(name string) ([]wModel.BossDrop, error) {
			return []wModel.BossDrop{
				{
					Boss: "Анемо гипостазис",
					Type: "world",
				},
				{
					Boss: "Двалин",
					Type: "weekly",
				},
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetWorldBossDropByName("hurricane_seed").
		DoAndReturn(func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Title:    "Семя урагана",
				Boss:     "Анемо гипостазис",
				Location: "Мондштадт",
				Type:     "world",
			}, nil
		}).MinTimes(1).MaxTimes(1)

	w.EXPECT().
		GetWeeklyBossDropByName("tail_of_boreas").
		DoAndReturn(func(name string) (wModel.BossDrop, error) {
			return wModel.BossDrop{
				Title:    "Хвост Борея",
				Boss:     "Андриус",
				Location: "Мондштадт",
				Domain:   "Испытание Волка Севера",
				Type:     "weekly",
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

*Возвышение персонажа (1-90):*
*🟢 Бирюза Вайюда* — 1/9/9/6 шт.
💥 *Семя урагана* — 2/4/8/12/20 — 46 шт.
Можно получить: Анемо гипостазис, Мондштадт
🌺 *Сесилия* — 3/10/20/30/45/60 — 168 шт.
🦴 *Слаймы* — 18/30/36 шт.
🧠 *«Опыт героя»* — 432 шт.
💰 *Мора* — 420 000

*Возвышение талантов (1-10):*
📚 *Книги талантов:* «О Поэзии» — 9/63/114 шт.
Когда: среда, 📍 *суббота*, воскресенье
🦴 *Слаймы* — 18/30/36 шт.
⚜ *Хвост Борея* — 18 шт.
Можно получить: Андриус (Испытание Волка Севера), Мондштадт
👑 *Корона прозрения* — 3 шт.
💰 *Мора* — 4 950 000`

	result, element, gem := cmdr.GetCharacterInfo("Венти")
	assert.Equal(t, expected, result)
	assert.Equal(t, "Анемо", element)
	assert.Equal(t, "\U0001F7E2 Бирюза Вайюда", gem)
}
