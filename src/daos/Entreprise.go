package daos

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"context"
	"errors"
	"encoding/json"
)

type EntrepriseDao struct{}

func (*EntrepriseDao) Add(client *elastic.Client, entreprise *models.Entreprise) error {
	result, err := client.Index().
		Index(index).
		Type("entreprise").
		BodyJson(entreprise.Source).
		Pretty(true).
		Do(context.Background())

	if err != nil || result == nil {
		return errors.New("Erreur lors de la création de votre compte")
	}

	entreprise.Id = result.Id
	entreprise.Type = result.Type
	entreprise.Version = result.Version
	entreprise.Index = result.Index

	return nil
}

func (*EntrepriseDao) GetByEmail (client *elastic.Client, email string) (*models.Entreprise, error) {

	matchQuery := elastic.NewMatchQuery("email", email)

	results, err := client.Search().
	Index(index).
	Type("entreprise").
	Query(matchQuery).
	Pretty(true).
	Do(context.Background())

	if err != nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	if results.TotalHits() == 0 {
		return nil, nil
	}

	//On récup le premier compte
	first_entreprise := results.Hits.Hits[0]
	entreprise := new(models.Entreprise)

	bytes, err := json.Marshal(first_entreprise)

	//On parse le json en objet
	err_unmarshal := json.Unmarshal(bytes, &entreprise)
	if err_unmarshal != nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	return entreprise, nil
}

func (*EntrepriseDao) GetById(client *elastic.Client, idEntreprise string) (*models.Entreprise, error) {
	ctx := context.Background()

	result, err := client.Get().
		Index(index).
		Type("entreprise").
		Id(idEntreprise).
		Pretty(true).
		Do(ctx)

	if err != nil || result == nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	entreprise := new(models.Entreprise)

	marshal, err := json.Marshal(result)
	if err != nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	//On regarde si on peut deserialiser le hit en entreprise
	if err := json.Unmarshal(marshal, entreprise); err != nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	return entreprise, nil
}

func (*EntrepriseDao) UpdateInformations(client *elastic.Client, idEntreprise string, entreprise *models.Entreprise) error {
	ctx := context.Background()

	if updated, err := client.Update().
		Index(index).
		Type("entreprise").
		Id(idEntreprise).
		Script(elastic.NewScriptInline("" +
		"ctx._source.nom_entreprise = params.nomEntreprise; ").
		Lang("painless").
		Param("nomEntreprise", entreprise.Source.NomEntreprise)).
		Do(ctx); updated == nil || err != nil {
		return errors.New("Erreur lors de la modification de vos informations")
	}

	return nil
}
