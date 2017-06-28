package models

type Semestre struct {
	Id string `json:"id" form:"id" query:"id"`
	Actif bool `json:"actif" form:"actif" query:"actif"`
	Description string `json:"description" form:"description" query:"description"`
	Name string `json:"name" form:"name" query:"name"`
	Ues []Ue `json:"ues" form:"ues" query:"ues"`
	Url string `json:"url" form:"url" query:"url"`
}


