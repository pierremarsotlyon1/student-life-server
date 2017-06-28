package models

type Register struct {
	Nom string `json:"nom" query:"nom" form:"nom"`
	Prenom string `json:"prenom" query:"prenom" form:"prenom"`
	Email string `json:"email" query:"email" form:"email"`
	Password string `json:"password" query:"password" form:"password"`
	ConfirmPassword string `json:"confirmPassword" query:"confirmPassword" form:"confirmPassword"`
}
