package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/metiers"
	"tomuss_server/src/tools"
	"tomuss_server/src/models"
	"fmt"
)

type SuggestionController struct {}

func (*SuggestionController) Add (c echo.Context) error {
	//Récupération du Token
	idEtudiant := new(metiers.JwtMetier).GetTokenByContext(c)

	//Création du client
	client := tools.CreateElasticsearchClient()

	//On bind la suggestion
	suggestion := new(models.Suggestion)

	fmt.Println("mdmsqd")
	if err := c.Bind(suggestion); err != nil {
		fmt.Println(err.Error())
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de votre suggestion"})
	}

	if err := new(metiers.SuggestionMetier).Add(client, idEtudiant, suggestion); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}
