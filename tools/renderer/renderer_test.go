package renderer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderer_RenderSuccessful(t *testing.T) {
	params := struct {
		Username string
	}{
		Username: "Angel",
	}

	r := NewRenderer("../../test/_templates/")
	expected := `Hello, Angel

Nice to see you! Your test is green now! ✅
Have a good day. 💜`

	result, err := r.Render("simple_template.tpl", params)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestRenderer_RenderFail(t *testing.T) {
	params := struct {
		Username string
	}{
		Username: "Angel",
	}

	r := NewRenderer("wrong")
	result, err := r.Render("simple_template.tpl", params)

	assert.Error(t, err)
	assert.Equal(t, "", result)
}

func TestRenderer_RenderErrorSuccess(t *testing.T) {
	r := NewRenderer("../../test/_templates/")
	result := r.RenderError("Test Case")

	assert.Equal(t, "❌ I cannot show you what you want, only error in Test Case", result)
}

func TestRenderer_RenderErrorFail(t *testing.T) {
	r := NewRenderer("wrong")
	result := r.RenderError("Test Case")

	assert.Equal(t, "Ой, что-то пошло не так. Давай немного подождем, может позже восстановится?", result)
}

func TestRenderer_RenderErrorWrongRenderParams(t *testing.T) {
	r := NewRenderer("../../test/_templates/")
	params := struct {
		Animal string
	}{
		"Cat",
	}
	result, err := r.Render("wrong_template.tpl", params)

	assert.NoError(t, err)
	assert.Equal(t, "Ой, что-то пошло не так. Давай немного подождем, может позже восстановится?", result)
}

func TestRenderer_AddEmojiToElement(t *testing.T) {
	t.Parallel()

	r := NewRenderer("path")

	testCases := []struct {
		Element  string
		Expected string
	}{
		{"Гео", "🔶 Гео"},
		{"Гидро", "💧 Гидро"},
		{"Пиро", "🔥 Пиро"},
		{"Анемо", "🍃 Анемо"},
		{"Крио", "❄ Крио"},
		{"Электро", "⚡ Электро"},
		{"Дендро", "🌱 Дендро"},
		{"Неизвестный", "Неизвестный"},
	}

	for _, tt := range testCases {
		t.Run(tt.Element, func(t *testing.T) {
			assert.Equal(t, tt.Expected, r.AddEmojiToElement(tt.Element))
		})
	}
}

func TestRenderer_GetEmojiToElement(t *testing.T) {
	t.Parallel()

	r := NewRenderer("path")

	testCases := []struct {
		Element  string
		Expected string
	}{
		{"Гео", "🔶"},
		{"Гидро", "💧"},
		{"Пиро", "🔥"},
		{"Анемо", "🍃"},
		{"Крио", "❄"},
		{"Электро", "⚡"},
		{"Дендро", "🌱"},
		{"Неизвестный", "Неизвестный"},
	}

	for _, tt := range testCases {
		t.Run(tt.Element, func(t *testing.T) {
			assert.Equal(t, tt.Expected, r.GetEmojiToElement(tt.Element))
		})
	}
}

func TestRenderer_AddEmojiToGem(t *testing.T) {
	t.Parallel()

	r := NewRenderer("path")

	testCases := []struct {
		Gem      string
		Expected string
	}{
		{"Агат Агнидус", "🔴 Агат Агнидус"},
		{"Лазурит Варунада", "🔵 Лазурит Варунада"},
		{"Аметист Ваджрада", "\U0001F7E3 Аметист Ваджрада"},
		{"Бирюза Вайюда", "\U0001F7E2 Бирюза Вайюда"},
		{"Нефрит Шивада", "💎 Нефрит Шивада"},
		{"Топаз Притхива", "\U0001F7E1 Топаз Притхива"},
	}

	for _, tt := range testCases {
		t.Run(tt.Gem, func(t *testing.T) {
			assert.Equal(t, tt.Expected, r.AddEmojiToGem(tt.Gem))
		})
	}
}
