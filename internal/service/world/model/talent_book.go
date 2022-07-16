package model

// TalentBook for game characters
type TalentBook struct {
	Id       int      `json:"id,omitempty"`
	Title    string   `json:"title"`
	Type     string   `json:"type,omitempty"`
	Location string   `json:"location,omitempty"`
	Weekdays []string `json:"-" db:"-"`
}

func (tb TalentBook) GetTotal() string {
	return "9/63/114"
}
