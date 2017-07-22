package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/metiers"
	"tomuss_server/src/tools"
	"tomuss_server/src/models"
)

type ProblemeTechniqueController struct {}

func (*ProblemeTechniqueController) Add (c echo.Context) error {
	//Récupération du Token
	idEtudiant := new(metiers.JwtMetier).GetTokenByContext(c)

	//Création du client
	client := tools.CreateElasticsearchClient()

	//On bind la suggestion
	problemeTechnique := new(models.ProblemeTechnique)

	if err := c.Bind(problemeTechnique); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de votre problème technique"})
	}

	if err := new(metiers.ProblemeTechniqueMetier).Add(client, idEtudiant, problemeTechnique); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}
