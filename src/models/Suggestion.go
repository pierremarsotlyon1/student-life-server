package models

type Suggestion struct {
	HeaderElasticsearch
	Source struct {
		Message string `json:"message" query:"message" form:"message"`
		IdEtudiant string `json:"id_etudiant" query:"id_etudiant" form:"id_etudiant"`
	} `json:"_source" form:"_source" query:"_source"`
}
