package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/tools"
	"tomuss_server/src/metiers"
	"tomuss_server/src/models"
	"strconv"
	"fmt"
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

func (*BonPlansController) FindByEntreprise (c echo.Context) error {
	offset := c.QueryParam("offset")
	if len(offset) == 0 {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de l'offset"})
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de l'offset"})
	}

	client := tools.CreateElasticsearchClient()
	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	bonplans, err := new(metiers.BonPlansMetier).FindByEntreprise(client, idEntreprise, offsetInt)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string][]*models.BonPlan{
		"bons_plans": bonplans,
	})
}

func (*BonPlansController) Add (c echo.Context) error {
	bonplans := new(models.BonPlan)

	if err := c.Bind(bonplans); err != nil {
		fmt.Println(err.Error())
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	client := tools.CreateElasticsearchClient()

	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	//On ajout le bon plan
	if err := new(metiers.BonPlansMetier).Add(client, idEntreprise, bonplans); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string]*models.BonPlan {
		"bon_plan": bonplans,
	})
}

func (*BonPlansController) Remove (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	idBonPlan := c.Param("id")

	if err := new(metiers.BonPlansMetier).Remove(client, idEntreprise, idBonPlan); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}
