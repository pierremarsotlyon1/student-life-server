package models

type CategorieAnnonce struct {
	HeaderElasticsearch
	Source struct {
		NomCategorieAnnonce string `json:"nom_categorie_annonce" query:"nom_categorie_annonce" form:"nom_categorie_annonce"`
	} `json:"_source" form:"_source" query:"_source"`
}
