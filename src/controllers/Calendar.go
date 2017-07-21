package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/metiers"
	"tomuss_server/src/tools"
	"tomuss_server/src/models"
)

type CalendarController struct {}

func (*CalendarController) UpdateUrlIcs (c echo.Context) error {

	//Récupération du Token
	idEtudiant := new(metiers.JwtMetier).GetTokenByContext(c)

	//Création du client
	client := tools.CreateElasticsearchClient()

	urlIcs := new(models.UrlIcs)

	if err := c.Bind(urlIcs); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	if err := new(metiers.CalendarMetier).UpdateUrlIcs(client, idEtudiant, urlIcs.UrlIcs); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}

func (*CalendarController) FindEvents (c echo.Context) error {
	//Récupération du Token
	idEtudiant := new(metiers.JwtMetier).GetTokenByContext(c)

	//Création du client
	client := tools.CreateElasticsearchClient()

	events, err := new(metiers.CalendarMetier).FindEvents(client, idEtudiant)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string][]models.Event {
		"events": events,
	})
}

func (*CalendarController) RefreshEvents (c echo.Context) error {
	//Récupération du Token
	idEtudiant := new(metiers.JwtMetier).GetTokenByContext(c)

	//Création du client
	client := tools.CreateElasticsearchClient()

	if err := new(metiers.CalendarMetier).RefreshEvents(client, idEtudiant); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}