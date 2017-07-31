package controllers

import (
	"github.com/labstack/echo"
	"tomuss_server/src/models"
	"tomuss_server/src/tools"
	"tomuss_server/src/metiers"
	"strconv"
)

type JobController struct {}

func (*JobController) Add (c echo.Context) error {
	job := new(models.Job)

	if err := c.Bind(job); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	client := tools.CreateElasticsearchClient()

	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	if err := new(metiers.JobMetier).Add(client, idEntreprise, job); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string]*models.Job {
		"job": job,
	})
}

func (*JobController) Remove (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	idJob := c.Param("id")

	if err := new(metiers.JobMetier).Remove(client, idEntreprise, idJob); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.NoContent(200)
}

func (*JobController) FindByDate (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	offset := c.QueryParam("offset")
	if len(offset) == 0 {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de l'offset"})
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de l'offset"})
	}

	jobs, err := new(metiers.JobMetier).FindByDate(client, offsetInt)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string][]*models.Job{
		"jobs": jobs,
	})
}

func (*JobController) FindByIdEntreprise (c echo.Context) error {
	client := tools.CreateElasticsearchClient()

	offset := c.QueryParam("offset")
	if len(offset) == 0 {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de l'offset"})
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération de l'offset"})
	}

	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	jobs, err := new(metiers.JobMetier).FindByEntreprise(client, idEntreprise, offsetInt)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string][]*models.Job{
		"jobs": jobs,
	})
}

func (*JobController) Update (c echo.Context) error {
	job := new(models.Job)

	if err := c.Bind(job); err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération des informations"})
	}

	client := tools.CreateElasticsearchClient()

	idEntreprise := new(metiers.JwtMetier).GetTokenByContext(c)

	idJob := c.Param("id")

	newJob, err := new(metiers.JobMetier).Update(client, idEntreprise, idJob, job)
	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, map[string]*models.Job {
		"job": newJob,
	})
}