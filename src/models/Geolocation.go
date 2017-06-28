package models

type Geolocation struct {
	Lat float64 `json:"lat" query:"lat" form:"lat"`
	Lon float64 `json:"lon" query:"lon" form:"lon"`
}
