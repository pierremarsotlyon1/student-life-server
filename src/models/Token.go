package models

type Token struct {
	Token string `json:"token" query:"token" form:"token"`
}

