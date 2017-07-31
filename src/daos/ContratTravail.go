package daos

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"context"
	"errors"
	"encoding/json"
)

type ContratTravailDao struct{}

func (*ContratTravailDao) FindAll(client *elastic.Client) ([]*models.ContratTravail, error) {
	results, err := client.Search().
		Index(index).
		Type("contrat_travail").
		Pretty(true).
		Do(context.Background())

	if err != nil || results.Hits == nil {
		return nil, errors.New("Erreur lors de la récupération des contrats de travail")
	}

	var contratsTravail []*models.ContratTravail

	if results.TotalHits() == 0 {
		return contratsTravail, nil
	}

	for _, hit := range results.Hits.Hits {
		bytes, err := json.Marshal(hit)

		if err != nil {
			return nil, errors.New("Erreur lors de la récupération des contrats de travail")
		}

		contratTravail := new(models.ContratTravail)

		if err := json.Unmarshal(bytes, contratTravail); err != nil {
			return nil, errors.New("Erreur lors de la récupération des contrats de travail")
		}

		contratsTravail = append(contratsTravail, contratTravail)
	}

	return contratsTravail, nil
}

func (*ContratTravailDao) GetById(client *elastic.Client, idContratTravail string) (*models.ContratTravail, error) {
	result, err := client.Get().
		Index(index).
		Type("contrat_travail").
		Id(idContratTravail).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return nil, errors.New("Erreur lors de la récupération du contrat de travail")
	}

	bytes, err := json.Marshal(result)

	if err != nil {
		return nil, errors.New("Erreur lors de la récupération du contrat de travail")
	}

	contratTravail := new(models.ContratTravail)
	if err := json.Unmarshal(bytes, contratTravail); err != nil {
		return nil, errors.New("Erreur lors de la récupération du contrat de travail")
	}

	return contratTravail, nil
}
