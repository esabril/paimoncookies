package model

type GemDropInfo struct {
	WeeklyBosses []WeeklyBoss
	WorldBosses  []WorldBoss
}

type Gem struct {
	Id       int         `db:"id"`
	Name     string      `db:"name"`
	Title    string      `db:"title"`
	DropInfo GemDropInfo `db:"-"`
}

func (g Gem) GetTotal() string {
	return "1/9/9/6"
}
