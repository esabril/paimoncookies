package characters

import (
	"github.com/esabril/paimoncookies/internal/service/characters/repository"
	"github.com/jmoiron/sqlx"
	"log"
)

type Characters struct {
	repo       repository.ICharactersRepo
	elements   map[string][]string
	characters map[string]bool
}

// Short characters names for searching
var simpleNames = map[string]string{
	"Каэдэхара Кадзуха": "Кадзуха",
	"Камисато Аято":     "Аято",
	"Камисато Аяка":     "Аяка",
	"Аратаки Итто":      "Итто",
	"Куки Синобу":       "Куки",
	"Сангономия Кокоми": "Кокоми",
	"Кудзё Сара":        "Сара",
}

func NewService(db *sqlx.DB) *Characters {
	s := &Characters{
		repo: repository.New(db),
	}

	elements, chars, err := s.GetInitialCharactersList()
	if err != nil {
		log.Fatalln("Error while loading characters list:", err.Error())
	}

	s.elements = elements
	s.characters = chars

	return s
}

// GetInitialCharactersList perform structures with characters and elements
func (c *Characters) GetInitialCharactersList() (elements map[string][]string, characters map[string]bool, err error) {
	list, err := c.repo.GetCharactersList()
	if err != nil {
		return
	}

	elements = make(map[string][]string)
	characters = make(map[string]bool)

	for _, ch := range list {
		characterName := c.SimplifyCharacterName(ch.Title)
		characters[characterName] = true

		_, ok := elements[ch.Element]
		if !ok {
			elements[ch.Element] = make([]string, 0)
		}

		elements[ch.Element] = append(elements[ch.Element], characterName)
	}

	return
}

// SimplifyCharacterName example: Камисато Аято -> Аято
func (c *Characters) SimplifyCharacterName(name string) string {
	shortName, ok := simpleNames[name]

	if !ok {
		return name
	}

	return shortName
}
