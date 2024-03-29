package characters

import (
	"log"

	"github.com/esabril/paimoncookies/internal/service/characters/model"
	"github.com/esabril/paimoncookies/internal/service/characters/repository"
	"github.com/jmoiron/sqlx"
)

type CharacterInterface interface {
	CheckCharacter(name string) bool
	CheckElement(element string) bool
	GetCharacterByName(name string) (model.Character, error)
	GetElementCharacters(element string, first int, last int) []string
	GetElements() map[string][]string
	GetInitialCharactersList() (elements map[string][]string, characters map[string]bool, err error)
	SimplifyCharacterName(name string) (string, bool)
}

type Characters struct {
	repo       repository.RepositoryInterface
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
	"Сиканоин Хэйдзо":   "Хэйдзо",
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

func NewMock(repo repository.RepositoryInterface, elements map[string][]string, characters map[string]bool) *Characters {
	return &Characters{
		repo:       repo,
		elements:   elements,
		characters: characters,
	}
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
		characterName, _ := c.SimplifyCharacterName(ch.Title)
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
func (c *Characters) SimplifyCharacterName(name string) (string, bool) {
	shortName, ok := simpleNames[name]

	if !ok {
		return name, false
	}

	return shortName, true
}

func (c *Characters) GetElements() map[string][]string {
	return c.elements
}

// GetElementCharacters little experiment: trying to get immutable characters data from memory
// without requests to DB
func (c *Characters) GetElementCharacters(element string, first, last int) []string {
	list, ok := c.elements[element]
	if !ok {
		log.Println("Attempt to get a list by an element that is not in the list:", element)

		return []string{}
	}

	if first == 0 && last == 0 {
		return list
	}

	if first < 0 {
		first = 0
	}

	count := len(list)

	if last > count {
		last = count
	}

	return list[first:last]
}

func (c *Characters) CheckCharacter(name string) bool {
	_, ok := c.characters[name]

	if !ok {
		_, ok := c.SimplifyCharacterName(name)

		return ok
	}

	return ok
}

func (c *Characters) CheckElement(element string) bool {
	_, ok := c.elements[element]

	return ok
}

func (c *Characters) GetCharacterByName(name string) (model.Character, error) {
	ch, err := c.repo.GetCharacterByName(name)
	if err != nil {
		log.Println("Error while getting character by name:", err.Error())
	}

	return ch, err
}
