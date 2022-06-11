package model

// Agenda structure
type Agenda struct {
	Weekday    string           `json:"weekday"`
	Content    WorldContent     `json:"content"`
	SystemData AgendaSystemData `json:"system_data"`
}

// WorldContent with world resources
type WorldContent struct {
	TalentBooks     map[string][]TalentBook     `json:"talent_books"`
	WeaponMaterials map[string][]WeaponMaterial `json:"weapon_materials"`
}

type AgendaSystemData struct {
	IsSunday bool     `json:"is_sunday"`
	Regions  []Region `json:"regoins"`
}

// RussianWeekdays Russian translation for weekdays
var RussianWeekdays = map[string]string{
	"monday":    "понедельник",
	"tuesday":   "вторник",
	"wednesday": "среда",
	"thursday":  "четверг",
	"friday":    "пятница",
	"saturday":  "суббота",
	"sunday":    "воскресенье",
}
