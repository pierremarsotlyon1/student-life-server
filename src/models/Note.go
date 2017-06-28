package models

type Note struct {
	Created string `json:"created" form:"created" query:"created"`
	Denominateur float64 `json:"denominateur" form:"denominateur" query:"denominateur"`
	Guid string `json:"guid" form:"guid" query:"guid"`
	Name string `json:"name" form:"name" query:"name"`
	Note string `json:"note" form:"note" query:"note"`
}
