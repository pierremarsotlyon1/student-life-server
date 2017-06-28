package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"tomuss_server/src/daos"
	"github.com/satori/go.uuid"
)

type SemestreMetier struct {}

func (*SemestreMetier) Find (client *elastic.Client, idEtudiant string) (*[]models.Semestre, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEtudiant) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre identifiant")
	}

	etudiant, err := new(EtudiantMetier).GetById(client, idEtudiant)

	if err != nil {
		return nil, err
	}

	return &etudiant.Source.Semestres, nil
}

func (*SemestreMetier) Add (client *elastic.Client, idEtudiant string, semestre *models.Semestre) error {
	//Check des infos
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEtudiant) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if len(semestre.Url) == 0 {
		return errors.New("Vous devez spécifier une URL")
	}

	if len(semestre.Name) == 0 {
		return errors.New("Vous devez spécifier un nom")
	}

	//On récupère l'étudiant
	etudiant, err := new(EtudiantMetier).GetById(client, idEtudiant)

	//On regarde si on a bien l'étudiant
	if err != nil {
		return err
	}

	if etudiant == nil {
		return errors.New("Erreur lors de la récupération de votre compte")
	}

	//On regarde si l'étudiant n'a pas déjà ce semestre via l'url
	for _, s := range etudiant.Source.Semestres {
		if len(s.Url) != 0 && s.Url == semestre.Url {
			return errors.New("Vous avez déjà ajouté un semestre avec cet URL")
		}

		if semestre.Actif && s.Actif {
			return errors.New("Vous ne pouvez avoir qu'un semestre d'actif à la fois, merci de désactiver vos semestres actifs")
		}
	}

	//On génère l'id
	semestre.Id = uuid.NewV4().String()

	//On ajoute le semestre à l'étudiant
	if err := new(daos.EtudiantDao).AddSemestre(client, idEtudiant, semestre); err != nil {
		return err
	}

	return nil
}

func (*SemestreMetier) Update (client *elastic.Client, idEtudiant string, id string, newSemestre *models.Semestre) error {
	//Check des infos
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEtudiant) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if len(id) == 0 {
		return errors.New("Erreur lors de la récupération de l'identifiant du semestre")
	}

	if len(newSemestre.Url) == 0 {
		return errors.New("Vous devez spécifier une URL")
	}

	if len(newSemestre.Name) == 0 {
		return errors.New("Vous devez spécifier un nom")
	}

	//On récupère l'étudiant
	etudiant, err := new(EtudiantMetier).GetById(client, idEtudiant)

	//On regarde si on a bien l'étudiant
	if err != nil {
		return err
	}

	if etudiant == nil {
		return errors.New("Erreur lors de la récupération de votre compte")
	}

	//On récupère le semestre via oldUrl
	updated := false

	var newSemestres []models.Semestre


	for _, s := range etudiant.Source.Semestres {
		if s.Actif && newSemestre.Actif && s.Id != newSemestre.Id {
			return errors.New("Vous ne pouvez avoir qu'un semestre d'actif à la fois, merci de désactiver vos semestres actifs")
		}

		if id == s.Id {
			s.Url = newSemestre.Url
			s.Name = newSemestre.Name
			s.Actif = newSemestre.Actif
			updated = true
		}

		newSemestres = append(newSemestres, s)
	}

	if !updated {
		return errors.New("Erreur lors de la récupération de votre semestre")
	}

	//On ajoute le semestre à l'étudiant
	if err := new(daos.EtudiantDao).UpdateSemestres(client, idEtudiant, newSemestres); err != nil {
		return err
	}

	return nil
}

func (*SemestreMetier) Remove (client *elastic.Client, idEtudiant string, idSemestre string) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEtudiant) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if len(idSemestre) == 0 {
		return errors.New("Erreur lors de la récupération de l'identifiant du semestre")
	}

	etudiantDao := new(daos.EtudiantDao)

	//On récup l'étudiant
	etudiant, err := etudiantDao.GetById(client, idEtudiant)

	if err != nil {
		return err
	}

	if etudiant == nil {
		return errors.New("Erreur lors de la récupération de votre compte")
	}

	//On génère le nouveau tableau de semestre en enlevant le semestre voulu
	var newSemestres []models.Semestre

	for _, s := range etudiant.Source.Semestres {
		if s.Id != idSemestre {
			newSemestres = append(newSemestres, s)
		}
	}

	//On update le document etudiant
	if err := etudiantDao.UpdateSemestres(client, idEtudiant, newSemestres); err != nil {
		return err
	}

	return nil
}