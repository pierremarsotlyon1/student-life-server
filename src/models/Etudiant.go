package models

type Etudiant struct {
	HeaderElasticsearch
	Source struct {
		InformationStudent
		FcmToken
		Email string `json:"email" query:"email" form:"email"`
		Password string `json:"password" query:"password" form:"password"`
		Semestres []Semestre `json:"semestres" form:"semestres" query:"semestres"`
		UrlIcs string `json:"url_ics" query:"url_ics" form:"url_ics"`
		Calendar []Event `json:"calendar" query:"calendar" form:"calendar"`
	} `json:"_source" form:"_source" query:"_source"`
}
