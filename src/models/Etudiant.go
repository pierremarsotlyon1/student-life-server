package models

type Etudiant struct {
	HeaderElasticsearch
	Source struct {
		Nom string `json:"nom" query:"nom" form:"nom"`
		Prenom string `json:"prenom" query:"prenom" form:"prenom"`
		Email string `json:"email" query:"email" form:"email"`
		Password string `json:"password" query:"password" form:"password"`
		Semestres []Semestre `json:"semestres" form:"semestres" query:"semestres"`
	} `json:"_source" form:"_source" query:"_source"`
}
