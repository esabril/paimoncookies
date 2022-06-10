package model

// TalentBook for game characters
type TalentBook struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title"`
	Type     string `json:"type,omitempty"`
	Location string `json:"location,omitempty"`
}
