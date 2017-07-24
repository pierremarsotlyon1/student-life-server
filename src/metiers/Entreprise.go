package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"tomuss_server/src/daos"
	"strings"
)

type EntrepriseMetier struct {}

func (*EntrepriseMetier) Add (client *elastic.Client, registerEntreprise *models.RegisterEntreprise) (*models.Entreprise, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(registerEntreprise.Email) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre email")
	}

	if !govalidator.IsEmail(registerEntreprise.Email) {
		return nil, errors.New("Votre email n'est pas au bon format")
	}

	//On regarde si on a pas déjà une entreprise avec le même mail
	entrepriseDao := new(daos.EntrepriseDao)

	entreprise, err := entrepriseDao.GetByEmail(client, registerEntreprise.Email)
	if err != nil {
		return nil, errors.New("Erreur lors de la vérification de l'existance d'un compte avec le même email")
	}

	if entreprise != nil {
		return nil, errors.New("Une entreprise à déjà un compte avec cet email")
	}

	if len(registerEntreprise.NomEntreprise) == 0 {
		return nil, errors.New("Erreur lors de la récupération du nom de votre entreprise")
	}

	if len(registerEntreprise.Password) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre mot de passe")
	}

	if len(registerEntreprise.ConfirmPassword) == 0 {
		return nil, errors.New("Erreur lors de la récupération de la confirmation de votre mot de passe")
	}

	if registerEntreprise.Password != registerEntreprise.ConfirmPassword {
		return nil, errors.New("Vos mot de passe ne sont pas identique")
	}

	//On hash le mot de passe
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerEntreprise.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Erreur lors de la sécurisation de votre mot de passe")
	}

	entreprise = new(models.Entreprise)
	entreprise.Source.Email = registerEntreprise.Email
	entreprise.Source.NomEntreprise = registerEntreprise.NomEntreprise
	entreprise.Source.Password = string(passwordHash)

	if err := entrepriseDao.Add(client, entreprise); err != nil {
		return nil, err
	}

	return entreprise, nil
}

func (entrepriseMetier *EntrepriseMetier) Login (client *elastic.Client, login *models.Login) (*models.Entreprise, error) {
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

	//Lowercase sur email
	emailLower := strings.ToLower(login.Email)

	//On regarde si c'est un email valide
	if isEmail := govalidator.IsEmail(emailLower); !isEmail {
		return nil, errors.New("L'email n'est pas au bon format")
	}

	entreprise, err := entrepriseMetier.GetByEmail(client, emailLower)

	if err != nil {
		return nil, err
	}

	if entreprise == nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	if len(entreprise.Source.Password) == 0 {
		return nil, errors.New("Erreur lors de la comparaison des mots de passe")
	}

	//On regarde si le password du gerant = passwordHash => sinon problème
	if err := bcrypt.CompareHashAndPassword([]byte(entreprise.Source.Password), []byte(login.Password)); err != nil {
		return nil, errors.New("Les mots de passe ne sont pas identiques")
	}

	return entreprise, nil
}

func (*EntrepriseMetier) GetById (client *elastic.Client, idEntreprise string) (*models.Entreprise, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEntreprise) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre identifiant")
	}

	entreprise, err := new(daos.EntrepriseDao).GetById(client, idEntreprise)
	if err != nil {
		return nil, err
	}

	if entreprise == nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	return entreprise, nil
}

func (*EntrepriseMetier) GetByEmail (client *elastic.Client, email string) (*models.Entreprise, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(email) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre email")
	}

	if !govalidator.IsEmail(email) {
		return nil, errors.New("L'email n'est pas au bon format")
	}

	entreprise, err := new(daos.EntrepriseDao).GetByEmail(client, email)
	if err != nil {
		return nil, err
	}

	if entreprise == nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	return entreprise, nil
}

func (entrepriseMetier *EntrepriseMetier) UpdateInformations (client *elastic.Client, idEntreprise string, entreprise *models.Entreprise) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEntreprise) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if err := new(daos.EntrepriseDao).UpdateInformations(client, idEntreprise, entreprise); err != nil {
		return err
	}

	return nil
}