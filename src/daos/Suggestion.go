package daos

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"context"
	"errors"
)

type SuggestionDao struct{}

func (*SuggestionDao) Add(client *elastic.Client, suggestion *models.Suggestion) error {
	result, err := client.Index().
		Index(index).
		Type("suggestion").
		BodyJson(suggestion.Source).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return errors.New("Erreur lors de l'ajout de la suggestion")
	}

	suggestion.Id = result.Id
	suggestion.Index = result.Index
	suggestion.Type = result.Type
	suggestion.Version = result.Version

	return nil
}
