package model

// Agenda structure
type Agenda struct {
	Weekday    string           `json:"weekday"`
	Content    WorldContent     `json:"content"`
	SystemData AgendaSystemData `json:"system_data"`
}

// WorldContent with world resources
type WorldContent struct {
	TalentBooks     map[string][]string         `json:"talent_books"`
	WeaponMaterials map[string][]WeaponMaterial `json:"weapon_materials"`
}

type AgendaSystemData struct {
	IsSunday bool `json:"is_sunday"`
}

// RussianWeekdays Russian translation for weekdays
var RussianWeekdays = map[string]string{
	"monday":    "Понедельник",
	"tuesday":   "Вторник",
	"wednesday": "Среда",
	"thursday":  "Четверг",
	"friday":    "Пятница",
	"saturday":  "Суббота",
	"sunday":    "Воскресенье",
}
