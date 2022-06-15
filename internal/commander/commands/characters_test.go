package commands

import (
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/internal/service/characters"
	"github.com/esabril/paimoncookies/tools/renderer"
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
			Characters: characters.NewMock(nil, chs),
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
			Characters: characters.NewMock(elements, nil),
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
