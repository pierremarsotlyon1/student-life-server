package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/tools"
	"tomuss_server/src/metiers"
	"tomuss_server/src/models"
	"strconv"
)

type BonPlansController struct {}

func (*BonPlansController) Find (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	offset := c.QueryParam("offset")
	if len(offset) == 0 {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de l'offset"})
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de l'offset"})
	}

	bonplans, err := new(metiers.BonPlansMetier).Find(client, offsetInt)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, bonplans)
}
