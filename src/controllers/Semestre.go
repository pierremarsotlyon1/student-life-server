package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/metiers"
	"tomuss_server/src/tools"
	"tomuss_server/src/models"
)

type SemestreController struct {}

func (*SemestreController) Find(c echo.Context) error {
	//Récupération du Token
	jwtToken := &metiers.JwtMetier{}
	idEtudiant := jwtToken.GetTokenByContext(c)

	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	semestres, err := new(metiers.SemestreMetier).Find(client, idEtudiant)

	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, *semestres)
}

func (*SemestreController) Add (c echo.Context) error {
	//Récupération du Token
	jwtToken := &metiers.JwtMetier{}
	idEtudiant := jwtToken.GetTokenByContext(c)

	//On bind
	semestre := new(models.Semestre)
	if err := c.Bind(semestre); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	if err := new(metiers.SemestreMetier).Add(client, idEtudiant, semestre); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, semestre)
}

func (*SemestreController) Update (c echo.Context) error {
	//Récupération du Token
	jwtToken := &metiers.JwtMetier{}
	idEtudiant := jwtToken.GetTokenByContext(c)

	//On bind
	semestre := new(models.Semestre)
	if err := c.Bind(semestre); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	//On récup l'ancienne url du semestre
	id := c.Param("id")

	if len(id) == 0 {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de l'url du semestre"})
	}

	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	if err := new(metiers.SemestreMetier).Update(client, idEtudiant, id, semestre); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, semestre)
}

func (*SemestreController) Remove (c echo.Context) error {
	//Récupération du Token
	jwtToken := &metiers.JwtMetier{}
	idEtudiant := jwtToken.GetTokenByContext(c)

	//On récup l'id du semestre
	idSemestre := c.Param("id")

	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	if err := new(metiers.SemestreMetier).Remove(client, idEtudiant, idSemestre); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}