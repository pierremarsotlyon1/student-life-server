package daos

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"context"
	"errors"
	"encoding/json"
	"time"
)

type BonPlansDao struct{}

func (*BonPlansDao) Find(client *elastic.Client, offset int) ([]*models.BonPlan, error) {
	rangeQuery := elastic.NewRangeQuery("date_debut").Gte(time.Now().UTC().Format("2006-01-02"))
	results, err := client.Search().
		Index(index).
		Type("bonplans").
		From(offset).
		Query(rangeQuery).
		Pretty(true).
		Do(context.Background())

	if err != nil || results == nil {
		return nil, errors.New("Erreur lors de la récupération des bon plans")
	}

	var bonplans []*models.BonPlan

	if results.Hits.TotalHits == 0 {
		return bonplans, nil
	}

	for _, hit := range results.Hits.Hits {
		bytes, err := json.Marshal(hit)

		if err != nil {
			return nil, errors.New("Erreur lors de la récupération des bon plans")
		}

		bonplan := new(models.BonPlan)

		if err := json.Unmarshal(bytes, bonplan); err != nil {
			return nil, errors.New("Erreur lors de la récupération des bon plans")
		}

		bonplans = append(bonplans, bonplan)
	}

	return bonplans, nil
}

func (*BonPlansDao) FindByEntreprise(client *elastic.Client, idEntreprise string, offset int) ([]*models.BonPlan, error) {
	rangeQuery := elastic.NewRangeQuery("date_debut").Gte(time.Now().UTC().Format("2006-01-02"))
	matchQuery := elastic.NewMatchQuery("id_entreprise", idEntreprise)

	globalQuery := elastic.NewBoolQuery().Must(matchQuery, rangeQuery)
	results, err := client.Search().
		Index(index).
		Type("bonplans").
		From(offset).
		Query(globalQuery).
		Pretty(true).
		Do(context.Background())

	if err != nil || results == nil {
		return nil, errors.New("Erreur lors de la récupération des bon plans")
	}

	var bonplans []*models.BonPlan

	if results.Hits.TotalHits == 0 {
		return bonplans, nil
	}

	for _, hit := range results.Hits.Hits {
		bytes, err := json.Marshal(hit)

		if err != nil {
			return nil, errors.New("Erreur lors de la récupération des bon plans")
		}

		bonplan := new(models.BonPlan)

		if err := json.Unmarshal(bytes, bonplan); err != nil {
			return nil, errors.New("Erreur lors de la récupération des bon plans")
		}

		bonplans = append(bonplans, bonplan)
	}

	return bonplans, nil
}

func (*BonPlansDao) Add(client *elastic.Client, bonplan *models.BonPlan) error {
	result, err := client.Index().
		Index(index).
		Type("bonplans").
		BodyJson(bonplan.Source).
		Pretty(true).
		Do(context.Background())

	if result == nil || err != nil {
		return errors.New("Erreur lors de la création du bon plan")
	}

	bonplan.Id = result.Id
	bonplan.Type = result.Type
	bonplan.Version = result.Version

	return nil
}

func (*BonPlansDao) Remove(client *elastic.Client, idBonPlan string) error {
	_, err := client.Delete().
		Index(index).
		Type("bonplans").
		Id(idBonPlan).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return errors.New("Erreur lors de la suppression du bon plan")
	}

	return nil
}
