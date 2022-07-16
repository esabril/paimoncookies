package model

import "github.com/esabril/paimoncookies/internal/service/world/model"

type CharacterAscension struct {
	Gem             model.Gem
	BossDrop        model.BossDrop
	LocalSpeciality model.AscensionMaterial
	CommonMaterial  model.AscensionMaterial
	Mora            Mora
	HeroWit         HeroWit
}

type CharacterTalentUpgrade struct {
	TalentBook     model.TalentBook
	BossDrop       model.BossDrop
	CommonMaterial model.AscensionMaterial
	Crown          Crown
	Mora           Mora
}

type CharacterMaterials struct {
	Ascension     CharacterAscension
	TalentUpgrade CharacterTalentUpgrade
}

// Character structure
type Character struct {
	Id                       int    `json:"id"`
	Name                     string `json:"name"`
	Title                    string `json:"title"`
	Region                   string `json:"region,omitempty"`
	Rarity                   int8   `json:"rarity,omitempty"`
	Element                  string `json:"element,omitempty"`
	TalentBookType           string `json:"talent_book_type,omitempty" db:"talent_book_type"`
	TalentBossDrop           string `json:"talent_boss_drop" db:"talent_boss_drop"`
	AscensionBossDrop        string `json:"ascension_boss_drop" db:"ascension_boss_drop"`
	AscensionGem             string `json:"ascension_gem" db:"ascension_gem"`
	AscensionLocalSpeciality string `json:"ascension_local_speciality" db:"ascension_local_speciality"`
	CommonAscensionMaterial  string `json:"common_ascension_material" db:"common_ascension_material"`

	Materials CharacterMaterials `json:"materials,omitempty" db:"-"`
}

// Crown additional talent upgrade type
type Crown struct{}

func (c Crown) GetTitle() string {
	return "Корона прозрения"
}

func (c Crown) GetTotal() int {
	return 3
}

// Mora game currency
type Mora struct{}

func (m Mora) GetAscensionTotal() string {
	return "420 000"
}

func (m Mora) GetTalentUpgradeTotal() string {
	return "4 950 000"
}

// HeroWit Character increase of experience material
type HeroWit struct{}

func (hw HeroWit) GetTotal() int {
	return 432
}
