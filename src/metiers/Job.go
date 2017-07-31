package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"github.com/go-shadow/moment"
	"tomuss_server/src/daos"
)

type JobMetier struct {}

func (*JobMetier) Add (client *elastic.Client, idEntreprise string, job *models.Job) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEntreprise) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	//On vérifie que ce soit bien une entreprise
	if entreprise, err := new(EntrepriseMetier).GetById(client, idEntreprise); err != nil || entreprise == nil {
		return errors.New("Erreur lors de la récupération de votre compte")
	}

	//On vérifie le job
	if len(job.Source.Titre) == 0 {
		return errors.New("Vous devez saisir un titre")
	}

	if len(job.Source.EmailContact) == 0 && len(job.Source.TelephoneContact) == 0 {
		return errors.New("Vous devez saisir une coordonnée de contact")
	}

	if len(job.Source.IdTypeContrat) == 0 {
		return errors.New("Erreur lors de la récupération de l'identifiant du contrat de travail")
	}

	//On récupère le contrat de travail
	contratTravail, err := new(ContratTravailMetier).GetById(client, job.Source.IdTypeContrat)
	if err != nil {
		return err
	}

	if contratTravail == nil {
		return errors.New("Erreur lors de la récupération du contrat de travail")
	}

	//On affecte les valeurs
	job.Source.IdEntreprise = idEntreprise
	job.Source.Created = moment.New().Format("YYYY-MM-DD")
	job.Source.TypeContrat = contratTravail.Source.NomContratTravail

	if err := new(daos.JobDao).Add(client, job); err != nil {
		return err
	}

	return nil
}

func (*JobMetier) FindByDate (client *elastic.Client, offset int) ([]*models.Job, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	offresEmploi, err := new(daos.JobDao).FindByDate(client, offset)

	if err != nil {
		return nil, err
	}

	return offresEmploi, nil
}

func (*JobMetier) FindByEntreprise (client *elastic.Client, idEntreprise string, offset int) ([]*models.Job, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	//On vérifie que ce soit bien une entreprise
	if entreprise, err := new(EntrepriseMetier).GetById(client, idEntreprise); err != nil || entreprise == nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	offresEmploi, err := new(daos.JobDao).FindByEntreprise(client, idEntreprise, offset)

	if err != nil {
		return nil, err
	}

	return offresEmploi, nil
}

func (jobMetier *JobMetier) Update (client *elastic.Client, idEntreprise string, idJob string, job *models.Job) (*models.Job, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEntreprise) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre identifiant")
	}

	//On vérifie que ce soit bien une entreprise
	if entreprise, err := new(EntrepriseMetier).GetById(client, idEntreprise); err != nil || entreprise == nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	jobDao := new(daos.JobDao)

	//On regarde si l'entreprise possède l'offre d'emploi
	if job, err := jobDao.GetByIdEntreprise(client, idEntreprise, idJob); job == nil || err != nil {
		return nil, errors.New("Erreur lors de la récupération de l'offre d'emploi")
	}

	//On vérifie le job
	if len(job.Source.Titre) == 0 {
		return nil, errors.New("Vous devez saisir un titre")
	}

	if len(job.Source.EmailContact) == 0 && len(job.Source.TelephoneContact) == 0 {
		return nil, errors.New("Vous devez saisir une coordonnée de contact")
	}

	if len(job.Source.IdTypeContrat) == 0 {
		return nil, errors.New("Erreur lors de la récupération de l'identifiant du contrat de travail")
	}

	//On récupère le contrat de travail
	contratTravail, err := new(ContratTravailMetier).GetById(client, job.Source.IdTypeContrat)
	if err != nil {
		return nil, err
	}

	if contratTravail == nil {
		return nil, errors.New("Erreur lors de la récupération du contrat de travail")
	}

	if err := jobDao.Update(client, job); err != nil {
		return nil, err
	}

	newJob, err := jobDao.GetById(client, idJob)
	if newJob == nil || err != nil {
		return nil, errors.New("Erreur lors de la récupération de l'offre d'emploi")
	}

	return newJob, nil
}

func (*JobMetier) Remove (client *elastic.Client, idEntreprise string, idJob string) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEntreprise) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	//On vérifie que ce soit bien une entreprise
	if entreprise, err := new(EntrepriseMetier).GetById(client, idEntreprise); err != nil || entreprise == nil {
		return errors.New("Erreur lors de la récupération de votre compte")
	}

	//On vérifie que l'entreprise ait bien l'annonce
	jobDao := new(daos.JobDao)

	if offreEmploi, err := jobDao.GetByIdEntreprise(client, idEntreprise, idJob); offreEmploi == nil || err != nil {
		return errors.New("Erreur lors de la récupération de l'offre d'emploi")
	}

	if err := jobDao.Remove(client, idJob); err != nil {
		return err
	}

	return nil
}