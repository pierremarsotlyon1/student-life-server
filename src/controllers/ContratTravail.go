package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/tools"
	"tomuss_server/src/metiers"
	"tomuss_server/src/models"
)

type ContratTravailController struct {}

func (*ContratTravailController) FindAll (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	contratsTravail, err := new(metiers.ContratTravailMetier).FindAll(client)

	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string][]*models.ContratTravail{
		"contrats_travail": contratsTravail,
	})
}
