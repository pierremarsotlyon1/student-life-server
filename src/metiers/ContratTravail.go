package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"tomuss_server/src/daos"
)

type ContratTravailMetier struct{}

func (*ContratTravailMetier) FindAll(client *elastic.Client) ([]*models.ContratTravail, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	contratsTravail, err := new(daos.ContratTravailDao).FindAll(client)

	if err != nil {
		return nil, err
	}

	return contratsTravail, nil
}

func (*ContratTravailMetier) GetById(client *elastic.Client, idContratTravail string) (*models.ContratTravail, error) {
	//Check des informations
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idContratTravail) == 0 {
		return nil, errors.New("Erreur lors de la récupération de l'identifiant du contrat de travail")
	}

	//On récupère le contrat de travail
	contratTravail, err := new(daos.ContratTravailDao).GetById(client, idContratTravail)
	if err != nil {
		return nil, err
	}

	return contratTravail, nil
}
