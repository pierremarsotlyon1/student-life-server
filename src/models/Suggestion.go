package models

type Suggestion struct {
	HeaderElasticsearch
	Source struct {
		Message string `json:"message" query:"message" form:"message"`
	} `json:"_source" form:"_source" query:"_source"`
}
