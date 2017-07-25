package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/tools"
	"tomuss_server/src/metiers"
	"tomuss_server/src/models"
)

type CategorieAnnonceController struct {}

func (*CategorieAnnonceController) FindAll (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	categoriesAnnonce, err := new(metiers.CategorieAnnonceMetier).FindAll(client)

	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string][]*models.CategorieAnnonce {
		"categories_annonce": categoriesAnnonce,
	})
}
