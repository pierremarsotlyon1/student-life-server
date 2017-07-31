package models

type Job struct {
	HeaderElasticsearch
	Source struct {
		Titre string `json:"titre" query:"titre" form:"titre"`
		Description string `json:"description" query:"description" form:"description"`
		Profil string `json:"profil" query:"profil" form:"profil"`
		Competences string `json:"competences" query:"competences" form:"competences"`
		TypeContrat string `json:"type_contrat" query:"type_contrat" form:"type_contrat"`
		DebutContrat string `json:"debut_contrat" query:"debut_contrat" form:"debut_contrat"`
		Remuneration int64 `json:"remuneration" query:"remuneration" form:"remuneration"`
		EmailContact string `json:"email_contact" query:"email_contact" form:"email_contact"`
		TelephoneContact string `json:"telephone_contact" query:"telephone_contact" form:"telephone_contact"`
		IdEntreprise string `json:"id_entreprise" query:"id_entreprise" form:"id_entreprise"`
		NomEntreprise string `json:"nom_entreprise" query:"nom_entreprise" form:"nom_entreprise"`
		LogoEntreprise string `json:"logo_entreprise" query:"logo_entreprise" form:"logo_entreprise"`
		Created string `json:"created" query:"created" form:"created"`
		IdTypeContrat string `json:"id_type_contrat" query:"id_type_contrat" form:"id_type_contrat"`
	} `json:"_source" form:"_source" query:"_source"`
}
