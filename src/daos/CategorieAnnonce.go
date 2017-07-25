package daos

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"context"
	"errors"
	"encoding/json"
)

type CategorieAnnonceDao struct{}

func (*CategorieAnnonceDao) FindAll(client *elastic.Client) ([]*models.CategorieAnnonce, error) {
	results, err := client.Search().
		Index(index).
		Type("categorie_annonce").
		Size(9999).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return nil, errors.New("Erreur lors de la récupération des catégories d'annonce")
	}

	var categoriesAnnonce []*models.CategorieAnnonce

	if results.Hits.TotalHits == 0 {
		return categoriesAnnonce, nil
	}

	for _, hit := range results.Hits.Hits {
		bytes, err := json.Marshal(hit)

		if err != nil {
			return nil, errors.New("Erreur lors de la récupération des catégories d'annonce")
		}

		categorieAnnonce := new(models.CategorieAnnonce)

		if err := json.Unmarshal(bytes, categorieAnnonce); err != nil {
			return nil, errors.New("Erreur lors de la récupération des catégories d'annonce")
		}

		categoriesAnnonce = append(categoriesAnnonce, categorieAnnonce)
	}

	return categoriesAnnonce, nil
}

func (*CategorieAnnonceDao) GetById (client *elastic.Client, idCategorieAnnonce string) (*models.CategorieAnnonce, error) {
	result, err := client.Get().
		Index(index).
		Type("categorie_annonce").
		Id(idCategorieAnnonce).
		Pretty(true).
		Do(context.Background())

	if err != nil || result == nil {
		return nil, errors.New("Erreur lors de la récupération de la catégorie d'annonce")
	}

	categorieAnnonce := new(models.CategorieAnnonce)

	marshal, err := json.Marshal(result)
	if err != nil {
		return nil, errors.New("Erreur lors de la récupération de la catégorie d'annonce")
	}

	//On regarde si on peut deserialiser le hit en catégorie annonce
	if err := json.Unmarshal(marshal, categorieAnnonce); err != nil {
		return nil, errors.New("Erreur lors de la récupération de la catégorie d'annonce")
	}

	return categorieAnnonce, nil

}

func (*CategorieAnnonceDao) Exist (client *elastic.Client, idCategorieAnnonce string) (bool, error) {
	exists, err := client.Exists().
		Index(index).
		Type("categorie_annonce").
		Id(idCategorieAnnonce).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return false, errors.New("Erreur lors de la récupération de la catégorie d'annonce")
	}

	return exists, nil

}