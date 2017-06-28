package models

type Abjus struct {
	Created string `json:"created" form:"created" query:"created"`
	Name string `json:"name" form:"name" query:"name"`
}
