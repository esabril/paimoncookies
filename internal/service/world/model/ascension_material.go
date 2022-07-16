package model

type AscensionMaterial struct {
	Id       int      `db:"id"`
	Name     string   `db:"name"`
	Title    string   `db:"title"`
	Type     string   `db:"type"`
	Location []string `db:"location"`
}

func (am AscensionMaterial) GetLocalSpecialityTotal() string {
	return "3/10/20/30/45/60 — 168"
}

func (am AscensionMaterial) GetCommonTalentTotal() string {
	return "18/66/93"
}

func (am AscensionMaterial) GetCommonAscensionTotal() string {
	return "18/30/36"
}
