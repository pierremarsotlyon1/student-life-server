package models

type InformationStudent struct {
	Nom string `json:"nom" query:"nom" form:"nom"`
	Prenom string `json:"prenom" query:"prenom" form:"prenom"`
}