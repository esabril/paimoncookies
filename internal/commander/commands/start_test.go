package commands

import (
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommander_GetStartSuccessful(t *testing.T) {
	s := &service.Service{}
	c := NewCommander(s, DefaultTemplatePath, nil)

	expected := `Привет, Путешественник! 🌸

Ты принес печеньки? 🍪 Паймон очень рада тебя видеть!
Надеюсь, Паймон поможет тебе в твоем путешествии, ведь она просто кладезь полезной информации! 💬

Например, я могу рассказать тебе о 🗓️ расписании на сегодня: /agenda

Кстати, тебе не обязательно запоминать команды: говорить со мной поможет клавиатура, расположенная внизу. 📑

Что ты хочешь узнать у Паймон? 🙃`

	assert.Equal(t, expected, c.GetStart())
}

func TestCommander_GetStartFail(t *testing.T) {
	s := &service.Service{}
	c := NewCommander(s, "wrongTemplatePath/", nil)

	expected := `Ой, что-то пошло не так. Давай немного подождем, может позже восстановится?`

	assert.Equal(t, expected, c.GetStart())
}
