package model

type WeeklyBoss struct {
	Id              int      `db:"id"`
	Name            string   `db:"name"`
	Title           string   `db:"title"`
	Location        string   `db:"location"`
	Domain          string   `db:"domain"`
	TalentMaterials []string `db:"talent_materials"`
	Gems            []string `db:"gems"`
}

type WorldBoss struct {
	Id                int      `db:"id"`
	Name              string   `db:"name"`
	Title             string   `db:"title"`
	Location          string   `db:"location"`
	AscensionMaterial string   `db:"ascension_material"`
	Gems              []string `db:"gems"`
}

type BossDrop struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Title    string `db:"title"`
	Type     string `db:"type"`
	Boss     string `db:"boss"`
	Location string `db:"location"`
	Domain   string `db:"domain"`
}

func (wb BossDrop) GetTotal() string {
	if wb.Type == "world" {
		return "2/4/8/12/20 — 46"
	}

	return "18"
}
