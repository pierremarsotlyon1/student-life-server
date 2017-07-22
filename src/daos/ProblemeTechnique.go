package daos

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"context"
	"errors"
)

type ProblemeTechniqueDao struct {}

func (*ProblemeTechniqueDao) Add(client *elastic.Client, problemeTechnique *models.ProblemeTechnique) error {
	result, err := client.Index().
		Index(index).
		Type("problemetechnique").
		BodyJson(problemeTechnique.Source).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return errors.New("Erreur lors de l'ajout du problème technique")
	}

	problemeTechnique.Id = result.Id
	problemeTechnique.Index = result.Index
	problemeTechnique.Type = result.Type
	problemeTechnique.Version = result.Version

	return nil
}
