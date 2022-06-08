package model

// WeaponMaterial world material structure
type WeaponMaterial struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title" json:"title"`
	Alias    string `json:"alias,omitempty"`
	Type     string `json:"type,omitempty"`
	Location string `json:"location,omitempty"`
}
