package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"tomuss_server/src/daos"
)

type CategorieAnnonceMetier struct {}

func (*CategorieAnnonceMetier) FindAll (client *elastic.Client) ([]*models.CategorieAnnonce, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	categoriesAnnonce, err := new(daos.CategorieAnnonceDao).FindAll(client)

	if err != nil {
		return nil, err
	}

	return categoriesAnnonce, nil
}

func (*CategorieAnnonceMetier) GetById (client *elastic.Client, idCategorieAnnonce string) (*models.CategorieAnnonce, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idCategorieAnnonce) == 0 {
		return nil, errors.New("Erreur lors de la récupération de l'identifiant de la catégorie d'annonce")
	}

	categorieAnnonce, err := new(daos.CategorieAnnonceDao).GetById(client, idCategorieAnnonce)

	if err != nil {
		return nil, err
	}

	return categorieAnnonce, nil
}

func (*CategorieAnnonceMetier) Exist (client *elastic.Client, idCategorieAnnonce string) (bool, error) {
	if client == nil {
		return false, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idCategorieAnnonce) == 0 {
		return false, errors.New("Erreur lors de la récupération de l'identifiant de la catégorie d'annonce")
	}

	exist, err := new(daos.CategorieAnnonceDao).Exist(client, idCategorieAnnonce)

	if err != nil {
		return false, err
	}

	return exist, nil
}
