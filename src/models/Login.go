package models

type Login struct {
	Email string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}