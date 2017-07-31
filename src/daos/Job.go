package daos

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"context"
	"errors"
	"encoding/json"
)

type JobDao struct{}

func (*JobDao) Add(client *elastic.Client, job *models.Job) error {
	result, err := client.Index().
		Index(index).
		Type("jobs").
		BodyJson(job.Source).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return errors.New("Erreur lors de l'ajout de l'offre d'emploi")
	}

	job.Index = result.Index
	job.Id = result.Id
	job.Type = result.Type
	job.Version = result.Version

	return nil
}

func (*JobDao) FindByDate(client *elastic.Client, offset int) ([]*models.Job, error) {

	results, err := client.Search().
		Index(index).
		Type("jobs").
		Sort("created", false).
		From(offset).
		Pretty(true).
		Do(context.Background())

	if err != nil || results.Hits == nil {
		return nil, errors.New("Erreur lors de la récupération des offres d'emploi")
	}

	var offresEmploi []*models.Job

	if results.TotalHits() == 0 {
		return offresEmploi, nil
	}

	for _, hit := range results.Hits.Hits {
		bytes, err := json.Marshal(hit)

		if err != nil {
			return nil, errors.New("Erreur lors de la récupération des offres d'emploi")
		}

		offreEmploi := new(models.Job)

		if err := json.Unmarshal(bytes, offreEmploi); err != nil {
			return nil, errors.New("Erreur lors de la récupération des offres d'emploi")
		}

		offresEmploi = append(offresEmploi, offreEmploi)
	}

	return offresEmploi, nil
}

func (*JobDao) FindByEntreprise(client *elastic.Client, idEntreprise string, offset int) ([]*models.Job, error) {

	matchQuery := elastic.NewMatchQuery("id_entreprise", idEntreprise)

	results, err := client.Search().
		Index(index).
		Type("jobs").
		Query(matchQuery).
		Sort("created", false).
		From(offset).
		Pretty(true).
		Do(context.Background())

	if err != nil || results.Hits == nil {
		return nil, errors.New("Erreur lors de la récupération des offres d'emploi")
	}

	var offresEmploi []*models.Job

	if results.TotalHits() == 0 {
		return offresEmploi, nil
	}

	for _, hit := range results.Hits.Hits {
		bytes, err := json.Marshal(hit)

		if err != nil {
			return nil, errors.New("Erreur lors de la récupération des offres d'emploi")
		}

		offreEmploi := new(models.Job)

		if err := json.Unmarshal(bytes, offreEmploi); err != nil {
			return nil, errors.New("Erreur lors de la récupération des offres d'emploi")
		}

		offresEmploi = append(offresEmploi, offreEmploi)
	}

	return offresEmploi, nil
}

func (*JobDao) Remove(client *elastic.Client, idJob string) error {
	deleted, err := client.Delete().
		Index(index).
		Type("jobs").
		Id(idJob).
		Do(context.Background())

	if err != nil || !deleted.Found {
		return errors.New("Erreur lors de la suppression de l'offre d'emploi")
	}

	return nil
}

func (*JobDao) Update(client *elastic.Client, job *models.Job) error {
	_, err := client.Update().
		Index(index).
		Type("jobs").
		Id(job.Id).
		Script(elastic.NewScriptInline("" +
		"ctx._source.competences = params.competences; " +
		"ctx._source.debut_contrat = params.debut_contrat; " +
		"ctx._source.description = params.description; " +
		"ctx._source.email_contact = params.email_contact; " +
		"ctx._source.profil = params.profil; " +
		"ctx._source.remuneration = params.remuneration; " +
		"ctx._source.telephone_contact = params.telephone_contact; " +
		"ctx._source.titre = params.titre; " +
		"ctx._source.type_contrat = params.type_contrat; " +
		"ctx._source.id_type_contrat = params.id_type_contrat;").
		Param("competences", job.Source.Competences).
		Param("debut_contrat", job.Source.DebutContrat).
		Param("description", job.Source.Description).
		Param("email_contact", job.Source.EmailContact).
		Param("profil", job.Source.Profil).
		Param("remuneration", job.Source.Remuneration).
		Param("telephone_contact", job.Source.TelephoneContact).
		Param("titre", job.Source.Titre).
		Param("type_contrat", job.Source.TypeContrat).
		Param("id_type_contrat", job.Source.IdTypeContrat)).
		Do(context.Background())

	if err != nil {
		return errors.New("Erreur lors de la modification de votre offre d'emploi")
	}

	return nil
}

func (*JobDao) GetByIdEntreprise(client *elastic.Client, idEntreprise string, idJob string) (*models.Job, error) {
	matchQuery := elastic.NewBoolQuery().Must(
		elastic.NewMatchQuery("_id", idJob),
		elastic.NewMatchQuery("id_entreprise", idEntreprise))

	result, err := client.Search().
		Index(index).
		Type("jobs").
		Query(matchQuery).
		Pretty(true).
		Do(context.Background())

	if err != nil || result.TotalHits() == 0 {
		return nil, errors.New("Erreur lors de la récupération de l'offre d'emploi")
	}

	offreEmploi := new(models.Job)

	bytes, err := json.Marshal(result.Hits.Hits[0])
	if err != nil {
		return nil, errors.New("Erreur lors de la récupération de l'offre d'emploi")
	}

	if err := json.Unmarshal(bytes, offreEmploi); err != nil {
		return nil, errors.New("Erreur lors de la récupération de l'offre d'emploi")
	}

	return offreEmploi, nil
}

func (*JobDao) GetById (client *elastic.Client, idJob string) (*models.Job, error) {
	result, err := client.Get().
		Index(index).
		Type("jobs").
		Id(idJob).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return nil, errors.New("Erreur lors de la récupération de l'offre d'emploi")
	}

	offreEmploi := new(models.Job)

	bytes, err := json.Marshal(result)
	if err != nil {
		return nil, errors.New("Erreur lors de la récupération de l'offre d'emploi")
	}

	if err := json.Unmarshal(bytes, offreEmploi); err != nil {
		return nil, errors.New("Erreur lors de la récupération de l'offre d'emploi")
	}

	return offreEmploi, nil
}