package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"tomuss_server/src/daos"
	"golang.org/x/crypto/bcrypt"
)

type EtudiantMetier struct {}

func (*EtudiantMetier) Login (client *elastic.Client, login *models.Login) (*models.Etudiant, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if login == nil {
		return nil, errors.New("Erreur lors de la récupération des informations")
	}

	if len(login.Email) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre email")
	}

	if len(login.Password) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre mot de passe")
	}

	etudiant, err := new(daos.EtudiantDao).GetByEmail(client, login.Email)

	if err != nil {
		return nil, err
	}

	if etudiant == nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	if len(etudiant.Source.Password) == 0 {
		return nil, errors.New("Erreur lors de la comparaison des mots de passe")
	}

	//On regarde si le password du gerant = passwordHash => sinon problème
	if err := bcrypt.CompareHashAndPassword([]byte(etudiant.Source.Password), []byte(login.Password)); err != nil {
		return nil, errors.New("Les mots de passe ne sont pas identiques")
	}

	return etudiant, nil
}

func (*EtudiantMetier) Register (client *elastic.Client, register *models.Register) (*models.Etudiant, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(register.Email) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre email")
	}

	etudiantDao := new(daos.EtudiantDao)

	etudiant, err := etudiantDao.GetByEmail(client, register.Email)

	if err != nil {
		return nil, err
	}

	if etudiant != nil {
		return nil, errors.New("Un compte existe déjà avec cet email")
	}

	if len(register.Nom) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre nom")
	}

	if len(register.Prenom) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre prénom")
	}

	if len(register.Password) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre mot de passe")
	}

	if len(register.ConfirmPassword) == 0 {
		return nil, errors.New("Erreur lors de la récupération de la confirmation de votre mot de passe")
	}

	if register.Password != register.ConfirmPassword {
		return nil, errors.New("Vos mots de passe ne sont pas identique")
	}

	//On hash le mot de passe
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Erreur lors de la sécurisation de votre mot de passe")
	}

	etudiant = new(models.Etudiant)

	etudiant.Source.Nom = register.Nom
	etudiant.Source.Prenom = register.Prenom
	etudiant.Source.Password = string(passwordHash)
	etudiant.Source.Email = register.Email

	err = etudiantDao.Add(client, etudiant)

	if err != nil {
		return nil, err
	}

	return etudiant, nil
}

func (*EtudiantMetier) GetById (client *elastic.Client, idEtudiant string) (*models.Etudiant, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEtudiant) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre identifiant")
	}

	etudiant, err := new(daos.EtudiantDao).GetById(client, idEtudiant)

	if err != nil {
		return nil, err
	}

	return etudiant, nil
}

func (*EtudiantMetier) Find (client *elastic.Client) ([]*models.Etudiant, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	etudiants, err := new(daos.EtudiantDao).Find(client)

	if err != nil {
		return nil, err
	}

	return etudiants, nil
}

func (*EtudiantMetier) ChangePassword(client *elastic.Client, idEtudiant string, changePassword *models.ChangePassword) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEtudiant) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if len(changePassword.NewPassword) == 0 {
		return errors.New("Erreur lors de la récupération de votre nouveau mot de passe")
	}

	if len(changePassword.ConfirmNewPassword) == 0 {
		return errors.New("Erreur lors de la récupération de la confirmation de votre nouveau mot de passe")
	}

	if changePassword.ConfirmNewPassword != changePassword.NewPassword {
		return errors.New("Vos mot de passe doivent être identique")
	}

	//On hash le mot de passe
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(changePassword.ConfirmNewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Erreur lors de la sécurisation de votre mot de passe")
	}

	//On convertit en string le mot de passe hashé
	stringPasswordHash := string(passwordHash)
	if len(stringPasswordHash) == 0 {
		return errors.New("Erreur lors de la conversion de votre nouveau de mot de passe")
	}

	if err := new(daos.EtudiantDao).ChangePassword(client, idEtudiant, stringPasswordHash); err != nil {
		return err
	}

	return nil
}

func (*EtudiantMetier) ChangeInformations(client *elastic.Client, idEtudiant string, informationsStudent *models.InformationStudent) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEtudiant) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if len(informationsStudent.Nom) == 0 {
		return errors.New("Erreur lors de la récupération de votre nom")
	}

	if len(informationsStudent.Prenom) == 0 {
		return errors.New("Erreur lors de la récupération de votre prénom")
	}

	if err := new(daos.EtudiantDao).ChangeInformatios(client, idEtudiant, informationsStudent); err != nil {
		return err
	}

	return nil
}