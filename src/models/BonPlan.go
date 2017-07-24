package models

type BonPlan struct {
	HeaderElasticsearch
	Source struct {
		Created string `json:"created" query:"created" form:"created"`
		Title string `json:"title" query:"title" form:"title"`
		Description string `json:"description" query:"description" form:"description"`
		NomEnreprise string `json:"nom_entreprise" query:"nom_entreprise" form:"nom_entreprise"`
		IdEnreprise string `json:"id_entreprise" query:"id_entreprise" form:"id_entreprise"`
		LogoEntreprise string `json:"logo_entreprise" query:"logo_entreprise" form:"logo_entreprise"`
		Reduction int64 `json:"reduction" query:"reduction" form:"reduction"`
		DateDebut string `json:"date_debut" query:"date_debut" form:"date_debut"`
		DateFin string `json:"date_fin" query:"date_fin" form:"date_fin"`
		CodePromo string `json:"code_promo" query:"code_promo" form:"code_promo"`
	} `json:"_source" form:"_source" query:"_source"`
}
