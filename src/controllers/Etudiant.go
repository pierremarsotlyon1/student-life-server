package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"tomuss_server/src/tools"
	"tomuss_server/src/models"
	"tomuss_server/src/metiers"
)

type EtudiantController struct {}

func (*EtudiantController) ChangePassword(c echo.Context) error {
	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	//Création et Bind du gérant
	changePassword := new(models.ChangePassword)
	if err := c.Bind(changePassword); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	//Récupération du Token
	jwtToken := &metiers.JwtMetier{}
	idEtudiant := jwtToken.GetTokenByContext(c)

	if err := new(metiers.EtudiantMetier).ChangePassword(client, idEtudiant, changePassword); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}

func (*EtudiantController) ChangeFcmToken(c echo.Context) error {
	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	//Création et Bind du gérant
	fcmToken := new(models.FcmToken)
	if err := c.Bind(fcmToken); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	//Récupération du Token
	jwtToken := &metiers.JwtMetier{}
	idEtudiant := jwtToken.GetTokenByContext(c)

	if err := new(metiers.EtudiantMetier).ChangeFcmToken(client, idEtudiant, fcmToken); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}

func (*EtudiantController) ChangeInformations(c echo.Context) error {
	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	//Création et Bind du gérant
	informationsStudent := new(models.InformationStudent)
	if err := c.Bind(informationsStudent); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	//Récupération du Token
	jwtToken := &metiers.JwtMetier{}
	idEtudiant := jwtToken.GetTokenByContext(c)

	if err := new(metiers.EtudiantMetier).ChangeInformations(client, idEtudiant, informationsStudent); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}

func (*EtudiantController) Register(c echo.Context) (err error) {

	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	//Création et Bind du gérant
	register := new(models.Register)
	if err = c.Bind(register); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	//Création de l'objet metier
	etudiantMetier := &metiers.EtudiantMetier{}

	//On ajoute en BDD
	etudiant, err := etudiantMetier.Register(client, register)
	if err != nil {
		return c.JSON(404, models.JsonErrorResponse{Error: err.Error()})
	}

	//On regarde si on a bien l'etudiant
	if etudiant == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la création de votre compte"})
	}

	//Création du token
	m := &metiers.JwtMetier{}
	token, err := m.Encode(etudiant)

	//On regarde si on a une erreur lors de la génération du token
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, models.Token{Token: token.Token})
}

func (*EtudiantController) Login(c echo.Context) error {
	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	//On récupère les informations via le BIND
	login := new(models.Login)

	if err := c.Bind(login); err != nil {
		return c.JSON(404, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	//Création de l'objet metier
	etudiantMetier := &metiers.EtudiantMetier{}

	//On récupère le compte
	etudiant, err := etudiantMetier.Login(client, login)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	//On regarde si on a bien un compte
	if etudiant == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de votre compte"})
	}

	//Création du token
	m := &metiers.JwtMetier{}
	token, err := m.Encode(etudiant)

	//On regarde si on a une erreur lors de la génération du token
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, models.Token{Token: token.Token})
}

