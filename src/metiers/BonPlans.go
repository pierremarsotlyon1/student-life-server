package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"tomuss_server/src/daos"
)

type BonPlansMetier struct {}

func (*BonPlansMetier) Find (client *elastic.Client, offset int) ([]*models.BonPlan, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	//On récup les bon plans
	bonplans, err := new(daos.BonPlansDao).Find(client, offset)

	if err != nil {
		return nil, err
	}

	return bonplans, nil
}

func (*BonPlansMetier) Add (client *elastic.Client, bonplan *models.BonPlan) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if err := new(daos.BonPlansDao).Add(client, bonplan); err != nil {
		return err
	}

	return nil
}

func (*BonPlansMetier) Remove (client *elastic.Client, idBonPlan string) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idBonPlan) == 0 {
		return errors.New("Erreur lors de la récupération de l'identifiant du bon plan")
	}

	if err := new(daos.BonPlansDao).Remove(client, idBonPlan); err != nil {
		return err
	}

	return nil
}
