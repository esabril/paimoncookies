package repository

import "github.com/esabril/paimoncookies/internal/service/world/model"

var _ RepositoryInterface = Mock{}

type (
	Mock struct {
		GetWeekdayTalentBooksWithLocationFunc     GetWeekdayTalentBooksWithLocationFunc
		GetWeekdayWeaponMaterialsWithLocationFunc GetWeekdayWeaponMaterialsWithLocationFunc
		GetRegionsFunc                            GetRegionsFunc
		GetTalentBookByTypeFunc                   GetTalentBookByTypeFunc
		GetTalentBookWeekdaysFunc                 GetTalentBookWeekdaysFunc
		GetGemByNameFunc                          GetGemByNameFunc
		GetWeeklyBossDropByNameFunc               GetWeeklyBossDropByNameFunc
		GetWorldBossDropByNameFunc                GetWorldBossDropByNameFunc
		GetAscensionMaterialsByNamesFunc          GetAscensionMaterialsByNamesFunc
		FindGemByTitleFunc                        FindGemByTitleFunc
		GetGemDropInfoByNameFunc                  GetGemDropInfoByNameFunc
	}

	GetWeekdayTalentBooksWithLocationFunc     func(weekday string) ([]model.TalentBook, error)
	GetWeekdayWeaponMaterialsWithLocationFunc func(weekday string) ([]model.WeaponMaterial, error)
	GetRegionsFunc                            func() ([]model.Region, error)
	GetTalentBookByTypeFunc                   func(bookType string) (model.TalentBook, error)
	GetTalentBookWeekdaysFunc                 func(bookType string) ([]string, error)
	GetGemByNameFunc                          func(name string) (model.Gem, error)
	GetWeeklyBossDropByNameFunc               func(name string) (model.BossDrop, error)
	GetWorldBossDropByNameFunc                func(name string) (model.BossDrop, error)
	GetAscensionMaterialsByNamesFunc          func(names []string) ([]model.AscensionMaterial, error)
	FindGemByTitleFunc                        func(title string) (model.Gem, error)
	GetGemDropInfoByNameFunc                  func(name string) ([]model.BossDrop, error)
)

func (r Mock) GetWeekdayTalentBooksWithLocation(weekday string) ([]model.TalentBook, error) {
	return r.GetWeekdayTalentBooksWithLocationFunc(weekday)
}

func (r Mock) GetWeekdayWeaponMaterialsWithLocation(weekday string) ([]model.WeaponMaterial, error) {
	return r.GetWeekdayWeaponMaterialsWithLocationFunc(weekday)
}

func (r Mock) GetRegions() ([]model.Region, error) {
	return r.GetRegionsFunc()
}

func (r Mock) GetTalentBookByType(bookType string) (model.TalentBook, error) {
	return r.GetTalentBookByTypeFunc(bookType)
}

func (r Mock) GetTalentBookWeekdays(bookType string) ([]string, error) {
	return r.GetTalentBookWeekdaysFunc(bookType)
}

func (r Mock) GetGemByName(name string) (model.Gem, error) {
	return r.GetGemByNameFunc(name)
}

func (r Mock) GetWeeklyBossDropByName(name string) (model.BossDrop, error) {
	return r.GetWeeklyBossDropByNameFunc(name)
}

func (r Mock) GetWorldBossDropByName(name string) (model.BossDrop, error) {
	return r.GetWorldBossDropByNameFunc(name)
}

func (r Mock) GetAscensionMaterialsByNames(names []string) ([]model.AscensionMaterial, error) {
	return r.GetAscensionMaterialsByNamesFunc(names)
}

func (r Mock) FindGemByTitle(title string) (model.Gem, error) {
	return r.FindGemByTitleFunc(title)
}

func (r Mock) GetGemDropInfoByName(name string) ([]model.BossDrop, error) {
	return r.GetGemDropInfoByNameFunc(name)
}
