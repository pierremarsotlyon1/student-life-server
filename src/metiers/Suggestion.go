package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"tomuss_server/src/daos"
)

type SuggestionMetier struct {}

func (*SuggestionMetier) Add (client *elastic.Client, idEtudiant string, suggestion *models.Suggestion) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEtudiant) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if len(suggestion.Source.Message) == 0 {
		return errors.New("Erreur lors de la récupération de votre suggestion")
	}

	//On ajoute la suggestion
	if err := new(daos.SuggestionDao).Add(client, idEtudiant, suggestion); err != nil {
		return err
	}

	return nil
}
