package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"tomuss_server/src/daos"
)

type ProblemeTechniqueMetier struct {}

func (*ProblemeTechniqueMetier) Add (client *elastic.Client, idEtudiant string, problemeTechnique *models.ProblemeTechnique) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEtudiant) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if len(problemeTechnique.Source.Message) == 0 {
		return errors.New("Erreur lors de la récupération de votre problème technique")
	}

	//On attribut l'id de l'étudiant au problème technique
	problemeTechnique.Source.IdEtudiant = idEtudiant

	//On ajoute le probleme technique
	if err := new(daos.ProblemeTechniqueDao).Add(client, problemeTechnique); err != nil {
		return err
	}

	return nil
}
