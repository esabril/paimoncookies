package world

import (
	"log"

	"github.com/esabril/paimoncookies/internal/service/world/model"
	"github.com/esabril/paimoncookies/internal/service/world/repository"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type WorldInterface interface {
	CreateAgenda(weekday string) (*model.Agenda, error)
	GetAgendaTalentBooks(weekday string) (map[string][]model.TalentBook, error)
	GetAgendaWeaponMaterials(weekday string) (map[string][]model.WeaponMaterial, error)
	GetTalentBookByName(bookType string) (model.TalentBook, error)
	GetGemByName(name string) (model.Gem, error)
	FindGemByTitle(title string) (model.Gem, error)
	GetWeeklyBossDropByName(name string) (model.BossDrop, error)
	GetWorldBossDropByName(name string) (model.BossDrop, error)
	GetCommonLocalSpecAscensionMaterialsByNames(names []string) (
		common model.AscensionMaterial,
		localSpec model.AscensionMaterial,
		err error,
	)
	GetGemDropInfo(name string) (model.GemDropInfo, error)
	GetWeekdayTranslation(wd string) string
}

// World structure for resources
type World struct {
	repo repository.RepositoryInterface
}

// NewService creates new world service
func NewService(db *sqlx.DB) *World {
	return &World{
		repo: repository.New(db),
	}
}

// NewMock Service with configured mock repository
func NewMock(mockRepo repository.RepositoryInterface) *World {
	return &World{
		repo: mockRepo,
	}
}

// CreateAgenda with day's resources
func (w *World) CreateAgenda(weekday string) (*model.Agenda, error) {
	var books map[string][]model.TalentBook
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

	regions, err := w.repo.GetRegions()
	if err != nil {
		log.Printf("Unable to get Regions list: %s\n", err.Error())

		return nil, err
	}

	return &model.Agenda{
		Weekday: model.RussianWeekdays[weekday],
		Content: model.WorldContent{
			TalentBooks:     books,
			WeaponMaterials: materials,
		},
		SystemData: model.AgendaSystemData{
			IsSunday: isSunday,
			Regions:  regions,
		},
	}, nil
}

func (w *World) GetAgendaTalentBooks(weekday string) (map[string][]model.TalentBook, error) {
	booksList, err := w.repo.GetWeekdayTalentBooksWithLocation(weekday)
	if err != nil {
		log.Printf("Unable to get weekday's Talent Books list: %s\n", err.Error())

		return nil, errors.Wrap(err, "Unable to get weekday's Talent Books list")
	}

	books := make(map[string][]model.TalentBook)

	for _, b := range booksList {
		if _, ok := books[b.Location]; !ok {
			books[b.Location] = make([]model.TalentBook, 0)
		}

		books[b.Location] = append(books[b.Location], b)
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

func (w *World) GetTalentBookByName(bookType string) (model.TalentBook, error) {
	book, err := w.repo.GetTalentBookByType(bookType)
	if err != nil {
		log.Println("Error while getting Talent Book by type:", err.Error())

		return model.TalentBook{}, err
	}

	weekdays, err := w.repo.GetTalentBookWeekdays(bookType)
	if err != nil {
		log.Println("Error while getting Talent Book Weekdays by type:", err.Error())

		return model.TalentBook{}, err
	}

	translate := make([]string, 0, len(weekdays))
	for _, wd := range weekdays {
		translate = append(translate, w.GetWeekdayTranslation(wd))
	}

	// We can get any resources on Sunday
	translate = append(translate, model.RussianSunday)
	book.Weekdays = translate

	return book, nil
}

func (w *World) GetGemByName(name string) (model.Gem, error) {
	gem, err := w.repo.GetGemByName(name)
	if err != nil {
		log.Println("Error while getting Gem by name:", err.Error())

		return model.Gem{}, err
	}

	di, err := w.GetGemDropInfo(name)
	if err != nil {
		log.Println("Error while getting Gem Drop Info:", err.Error())
	}

	gem.DropInfo = di

	return gem, nil
}

func (w *World) FindGemByTitle(title string) (model.Gem, error) {
	gem, err := w.repo.FindGemByTitle(title)
	if err != nil {
		log.Println("Error while getting Gem by title:", err.Error())

		return model.Gem{}, err
	}

	di, err := w.GetGemDropInfo(gem.Name)
	if err != nil {
		log.Println("Error while getting Gem Drop Info:", err.Error())
	}

	gem.DropInfo = di

	return gem, nil
}

func (w *World) GetWeeklyBossDropByName(name string) (model.BossDrop, error) {
	bd, err := w.repo.GetWeeklyBossDropByName(name)
	if err != nil {
		log.Println("Error while getting Weekly Boss Drop by name:", err.Error())

		return model.BossDrop{}, err
	}

	return bd, nil
}

func (w *World) GetWorldBossDropByName(name string) (model.BossDrop, error) {
	bd, err := w.repo.GetWorldBossDropByName(name)
	if err != nil {
		log.Println("Error while getting World Boss Drop by name:", err.Error())

		return model.BossDrop{}, err
	}

	return bd, nil
}

func (w *World) GetCommonLocalSpecAscensionMaterialsByNames(names []string) (
	common model.AscensionMaterial,
	localSpec model.AscensionMaterial,
	err error,
) {
	ams, err := w.repo.GetAscensionMaterialsByNames(names)
	if err != nil {
		log.Println("Error while getting Ascension Materials by name:", err.Error())

		return model.AscensionMaterial{}, model.AscensionMaterial{}, err
	}

	if len(ams) > 2 {
		ams = ams[:2]
	}

	if ams[0].Type == "common" {
		common, localSpec, err = ams[0], ams[1], nil
	} else {
		common, localSpec, err = ams[1], ams[0], nil
	}

	return
}

func (w *World) GetGemDropInfo(name string) (model.GemDropInfo, error) {
	bd, err := w.repo.GetGemDropInfoByName(name)
	if err != nil {
		log.Println("Error while getting Gem Drop Info:", err.Error())

		return model.GemDropInfo{}, err
	}

	weeklyBd, worldBd := make([]model.WeeklyBoss, 0), make([]model.WorldBoss, 0)

	for _, d := range bd {
		if d.Type == "weekly" {
			weeklyBd = append(weeklyBd, model.WeeklyBoss{
				Title:    d.Boss,
				Location: d.Location,
				Domain:   d.Domain,
			})
		} else {
			worldBd = append(worldBd, model.WorldBoss{
				Title:    d.Boss,
				Location: d.Location,
			})
		}
	}

	info := model.GemDropInfo{
		WeeklyBosses: weeklyBd,
		WorldBosses:  worldBd,
	}

	return info, nil
}

func (w *World) GetWeekdayTranslation(wd string) string {
	return model.RussianWeekdays[wd]
}
