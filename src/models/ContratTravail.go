package models

type ContratTravail struct {
	HeaderElasticsearch
	Source struct {
		NomContratTravail string `json:"nom_contrat_travail" query:"nom_contrat_travail" form:"nom_contrat_travail"`
	} `json:"_source" form:"_source" query:"_source"`
}
