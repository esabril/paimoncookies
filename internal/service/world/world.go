package world

import (
	"github.com/esabril/paimoncookies/internal/service/world/model"
	"github.com/esabril/paimoncookies/internal/service/world/repository"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
)

// World structure for resources
type World struct {
	repo repository.IRepo
}

// NewService creates new world service
func NewService(db *sqlx.DB) *World {
	return &World{
		repo: repository.New(db),
	}
}

// CreateAgenda with day's resources
func (w *World) CreateAgenda(weekday string) (*model.Agenda, error) {
	var books map[string][]string
	var materials map[string][]model.WeaponMaterial
	var err error

	isSunday := false

	// On Sunday we can receive all the talend books and materials, so we will not go to DB
	if weekday != "sunday" {
		books, err = w.GetAgendaTalentBooks(weekday)
		if err != nil {
			return nil, err
		}

		materials, err = w.GetAgendaWeaponMaterials(weekday)
		if err != nil {
			return nil, err
		}
	} else {
		isSunday = true
	}

	return &model.Agenda{
		Weekday: model.RussianWeekdays[weekday],
		Content: model.WorldContent{
			TalentBooks:     books,
			WeaponMaterials: materials,
		},
		SystemData: model.AgendaSystemData{
			IsSunday: isSunday,
		},
	}, nil
}

func (w *World) GetAgendaTalentBooks(weekday string) (map[string][]string, error) {
	booksList, err := w.repo.GetWeekdayTalentBooksWithLocation(weekday)
	if err != nil {
		log.Printf("Unable to get weekday's Talent Books list: %s\n", err.Error())

		return nil, errors.Wrap(err, "Unable to get weekday's Talent Books list")
	}

	books := make(map[string][]string)

	for _, b := range booksList {
		if _, ok := books[b.Location]; !ok {
			books[b.Location] = make([]string, 0)
		}

		books[b.Location] = append(books[b.Location], b.Title)
	}

	return books, nil
}

func (w *World) GetAgendaWeaponMaterials(weekday string) (map[string][]model.WeaponMaterial, error) {
	materialsList, err := w.repo.GetWeekdayWeaponMaterialsWithLocation(weekday)
	if err != nil {
		log.Printf("Unable to get weekday's Weapon Materials list: %s\n", err.Error())

		return nil, errors.Wrap(err, "Unable to get weekday's Weapon Materials list")
	}

	materials := make(map[string][]model.WeaponMaterial)

	for _, m := range materialsList {
		if _, ok := materials[m.Location]; !ok {
			materials[m.Location] = make([]model.WeaponMaterial, 0)
		}

		materials[m.Location] = append(materials[m.Location], model.WeaponMaterial{
			Title: m.Title,
			Alias: m.Alias,
		})
	}

	return materials, err
}
