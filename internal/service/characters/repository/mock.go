package repository

import "github.com/esabril/paimoncookies/internal/service/characters/model"

var _ RepositoryInterface = Mock{}

type (
	Mock struct {
		GetCharactersListFunc  GetCharactersListFunc
		GetCharacterByNameFunc GetCharacterByNameFunc
	}

	GetCharactersListFunc  func() ([]model.Character, error)
	GetCharacterByNameFunc func(name string) (model.Character, error)
)

func (r Mock) GetCharactersList() ([]model.Character, error) {
	return r.GetCharactersListFunc()
}

func (r Mock) GetCharacterByName(name string) (model.Character, error) {
	return r.GetCharacterByNameFunc(name)
}
