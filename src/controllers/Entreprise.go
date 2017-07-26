package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/tools"
	"tomuss_server/src/models"
	"tomuss_server/src/metiers"
)

type EntrepriseController struct {}

func (*EntrepriseController) Register (c echo.Context) error {
	//Création du client
	client := tools.CreateElasticsearchClient()

	//Création et Bind du gérant
	registerEntreprise := new(models.RegisterEntreprise)
	if err := c.Bind(registerEntreprise); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	//On ajoute en BDD
	entreprise, err := new(metiers.EntrepriseMetier).Add(client, registerEntreprise)
	if err != nil {
		return c.JSON(404, models.JsonErrorResponse{Error: err.Error()})
	}

	//On regarde si on a bien l'entreprise
	if entreprise == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la création de votre compte"})
	}

	//Création du token
	token, err := new(metiers.JwtMetier).Encode(entreprise.Id)

	//On regarde si on a une erreur lors de la génération du token
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, models.Token{Token: token.Token})
}

func (*EntrepriseController) Login (c echo.Context) error {
	//Création du client
	client := tools.CreateElasticsearchClient()

	//Création et Bind du gérant
	login := new(models.Login)
	if err := c.Bind(login); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	//On récupère le compte entreprise en BDD
	entreprise, err := new(metiers.EntrepriseMetier).Login(client, login)
	if err != nil {
		return c.JSON(404, models.JsonErrorResponse{Error: err.Error()})
	}

	//On regarde si on a bien l'entreprise
	if entreprise == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la création de votre compte"})
	}

	//Création du token
	token, err := new(metiers.JwtMetier).Encode(entreprise.Id)

	//On regarde si on a une erreur lors de la génération du token
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, models.Token{Token: token.Token})
}

func (*EntrepriseController) Profile (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	entreprise, err := new(metiers.EntrepriseMetier).Profile(client, idEntreprise)

	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string]*models.Entreprise {
		"entreprise": entreprise,
	})
}

func (*EntrepriseController) UpdateInformations (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	entreprise := new(models.Entreprise)

	if err := c.Bind(entreprise); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	entrepriseNew, err := new(metiers.EntrepriseMetier).UpdateInformations(client, idEntreprise, entreprise)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string]*models.Entreprise {
		"entreprise": entrepriseNew,
	})
}

func (*EntrepriseController) UpdateUrlLogo (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	entreprise := new(models.Entreprise)

	if err := c.Bind(entreprise); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	if err := new(metiers.EntrepriseMetier).UpdateUrlLogo(client, idEntreprise, entreprise.Source.LogoEntreprise); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}