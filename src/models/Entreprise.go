package models

type Entreprise struct {
	HeaderElasticsearch
	Source struct {
		Email string `json:"email" query:"email" form:"email"`
		Password string `json:"password" query:"password" form:"password"`
		NomEntreprise string `json:"nom_entreprise" query:"nom_entreprise" form:"nom_entreprise"`
		LogoEntreprise string `json:"logo_entreprise" query:"logo_entreprise" form:"logo_entreprise"`
	} `json:"_source" form:"_source" query:"_source"`
}
