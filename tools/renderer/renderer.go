package renderer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"
)

// CommonErrorMessage Little Trick: if user gives a message without any emoji — something wrong with renderer in common
const CommonErrorMessage = "Ой, что-то пошло не так. Давай немного подождем, может позже восстановится?"

type Renderer struct {
	TemplatePath      string
	ElementsToEmojis  map[string]string
	GemsToEmojis      map[string]string
	PreviousPageEmoji string
	NextPageEmoji     string
}

func NewRenderer(templatePath string) *Renderer {
	return &Renderer{
		TemplatePath: templatePath,
		ElementsToEmojis: map[string]string{
			"Крио":    "❄",
			"Пиро":    "🔥",
			"Анемо":   "🍃",
			"Электро": "⚡",
			"Гео":     "🔶",
			"Гидро":   "💧",
			"Дендро":  "🌱",
		},
		GemsToEmojis: map[string]string{
			"Агат Агнидус":     "🔴",
			"Лазурит Варунада": "🔵",
			"Аметист Ваджрада": "\U0001F7E3",
			"Бирюза Вайюда":    "\U0001F7E2",
			"Нефрит Шивада":    "💎",
			"Топаз Притхива":   "\U0001F7E1",
			"Изумруд Нагадус":  "\U0001F7E2",
		},
		PreviousPageEmoji: "⬅",
		NextPageEmoji:     "➡",
	}
}

func (r *Renderer) AddEmojiToElement(el string) string {
	emoji, ok := r.ElementsToEmojis[el]
	if !ok {
		return el
	}

	return fmt.Sprintf("%s %s", emoji, el)
}

func (r *Renderer) GetEmojiToElement(el string) string {
	emoji, ok := r.ElementsToEmojis[el]
	if !ok {
		return el
	}

	return emoji
}

func (r *Renderer) AddEmojiToGem(gem string) string {
	emoji, ok := r.GemsToEmojis[gem]
	if !ok {
		return gem
	}

	return fmt.Sprintf("%s %s", emoji, gem)
}

func (r *Renderer) Render(name string, params interface{}) (string, error) {
	funcMap := template.FuncMap{"join": strings.Join}
	t := template.New(name).Funcs(funcMap)
	t, err := t.ParseFiles(r.TemplatePath + name)
	if err != nil {
		log.Printf("Error ocurred while parsing template file: %s\n", err.Error())

		return "", err
	}

	return r.renderToString(t, params), err
}

// RenderError Returns error message from template with emoji
func (r *Renderer) RenderError(errorCase string) string {
	params := struct {
		ErrorCase string
	}{
		ErrorCase: errorCase,
	}

	t := template.New("error.tpl")
	t, err := t.ParseFiles(r.TemplatePath + "error.tpl")
	if err != nil {
		return CommonErrorMessage
	}

	return r.renderToString(t, params)
}

func (r *Renderer) renderToString(t *template.Template, params interface{}) string {
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, params); err != nil {
		log.Printf("Error ocurred while rendering: %s (params: %v)\n", err.Error(), params)
		return CommonErrorMessage
	}

	return tpl.String()
}
