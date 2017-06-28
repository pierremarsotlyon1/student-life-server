package models

type HeaderElasticsearch struct {
	Id string `json:"_id" form:"_id" query:"_id"`
	Type string `json:"_type" form:"_type" query:"_type"`
	Index string `json:"_index" form:"_index" query:"_index"`
	Score float32 `json:"_score" form:"_score" query:"_score"`
	Version int `json:"_version" form:"_version" query:"_version"`
	Parent string `json:"_parent" form:"_parent" query:"_parent"`
}
