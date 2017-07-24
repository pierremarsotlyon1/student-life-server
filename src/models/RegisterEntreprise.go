package models

type RegisterEntreprise struct {
	Email string `json:"email" query:"email" form:"email"`
	Password string `json:"password" query:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" query:"confirm_password" form:"confirm_password"`
	NomEntreprise string `json:"nom_entreprise" query:"nom_entreprise" form:"nom_entreprise"`
	LogoEntreprise string `json:"logo_entreprise" query:"logo_entreprise" form:"logo_entreprise"`
}
