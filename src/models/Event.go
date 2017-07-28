package models

type Event struct {
	Titre string `json:"titre" query:"titre" form:"titre"`
	DateDebut string `json:"date_debut" query:"date_debut" form:"date_debut"`
	DateFin string `json:"date_fin" query:"date_fin" form:"date_fin"`
	Description string `json:"description" query:"description" form:"description"`
	Location string `json:"location" query:"location" form:"location"`
}
