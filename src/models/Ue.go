package models

type Ue struct {
	Name string `json:"name" form:"name" query:"name"`
	Abinj []Abinj `json:"abinj" form:"abinj" query:"abinj"`
	Abjus []Abjus `json:"abjus" form:"abjus" query:"abjus"`
	Notes []Note `json:"notes" form:"notes" query:"notes"`
	Prst []Prst `json:"prst" form:"prst" query:"prst"`
}
