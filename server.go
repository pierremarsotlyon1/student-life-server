package main

import (
	"github.com/labstack/echo"
	"tomuss_server/src/controllers"
	"github.com/labstack/echo/middleware"
	"tomuss_server/src/metiers"
	"os"
	"tomuss_server/src/tools"
	"tomuss_server/src/models"
	"gopkg.in/olivere/elastic.v5"
	"time"
	"context"
	"encoding/json"
)

type Message struct {
	models.HeaderElasticsearch
	Source struct {
		Phrase string `json:"message" query:"message" form:"message"`
		Date string `json:"date" query:"date" form:"date"`
	} `json:"_source" form:"_source" query:"_source"`
}

func LoveMessage(c echo.Context) error {
	//Création du client
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la connexion à notre base de donnée"})
	}

	results, err := client.Search().
	Index("lovemessage").
	Type("messages").
	Query(elastic.NewMatchQuery("date", time.Now().UTC().Format("2006-01-02"))).
	Pretty(true).
	Do(context.Background())

	if err != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération du message du jour"})
	}

	//On récup le premier compte
	messageResult := results.Hits.Hits[0]
	message := new(Message)

	bytes, err := json.Marshal(messageResult)

	//On parse le json en objet
	err_unmarshal := json.Unmarshal(bytes, &message)
	if err_unmarshal != nil {
		return c.JSON(403, models.JsonErrorResponse{Error: "Erreur lors de la récupération du message du jour"})
	}

	return c.JSON(200, message)
}

func main() {
	e := echo.New()

	//CORS
	e.Use(middleware.CORS())

	e.GET("/lovemessage", LoveMessage)

	//Association des routes
	//Définition des controllers
	etudiantController := new(controllers.EtudiantController)
	semestreController := new(controllers.SemestreController)
	bonPlansController := new(controllers.BonPlansController)
	calendarController := new(controllers.CalendarController)
	suggestionController := new(controllers.SuggestionController)
	problemeTechniqueController := new(controllers.ProblemeTechniqueController)
	entrepriseController := new(controllers.EntrepriseController)
	categorieAnnonceController := new(controllers.CategorieAnnonceController)
	jobController := new(controllers.JobController)
	contratTravailController := new(controllers.ContratTravailController)

	//Etudiant auth sans JWT
	e.POST("/login", etudiantController.Login)
	e.POST("/register", etudiantController.Register)

	//Entreprise auth sans JWT
	e.POST("/entreprise/login", entrepriseController.Login)
	e.POST("/entreprise/register", entrepriseController.Register)

	//Contrat travail API
	contratTravailApi := e.Group("/contrat_travail")
	contratTravailApi.GET("", contratTravailController.FindAll)

	//Catégorie annonce API
	categorieAnnonceApi := e.Group("/categorie_annonce")
	categorieAnnonceApi.GET("", categorieAnnonceController.FindAll)

	//Définition de l'api de base avec restriction JWT
	api := e.Group("/api")
	api.Use(middleware.JWT([]byte(metiers.GetSecretJwt())))

	//Api pour l'étudiant
	etudiantApi := api.Group("/etudiant")
	etudiantApi.GET("", etudiantController.Profile)
	etudiantApi.PUT("/fcm", etudiantController.ChangeFcmToken)
	etudiantApi.PUT("/change/password", etudiantController.ChangePassword)
	etudiantApi.PUT("/change/informations", etudiantController.ChangeInformations)

	etudiantApi.GET("/semestres", semestreController.Find)
	etudiantApi.POST("/semestres", semestreController.Add)
	etudiantApi.PUT("/semestres/:id", semestreController.Update)
	etudiantApi.DELETE("/semestres/:id", semestreController.Remove)

	//Api bon plans
	bonPlansApi := api.Group("/bonplans")
	bonPlansApi.GET("", bonPlansController.Find)
	bonPlansApi.GET("/recent/:date", bonPlansController.FindRecent)

	//Api calendar
	calendarApi := api.Group("/calendar")
	calendarApi.POST("", calendarController.UpdateUrlIcs)
	calendarApi.GET("", calendarController.Synchroniser)

	//Suggestion api
	suggestionApi := api.Group("/suggestion")
	suggestionApi.POST("", suggestionController.Add)

	//Probleme technique api
	problemeTechniqueApi := api.Group("/probleme/technique")
	problemeTechniqueApi.POST("", problemeTechniqueController.Add)

	//Entreprise API
	entrepriseApi := api.Group("/entreprise")
	entrepriseApi.GET("", entrepriseController.Profile)
	entrepriseApi.PUT("", entrepriseController.UpdateInformations)
	entrepriseApi.PUT("/logo", entrepriseController.UpdateUrlLogo)

	//Bon plans entreprise API
	bonPlansEntreprise := entrepriseApi.Group("/bonsplans")
	bonPlansEntreprise.POST("", bonPlansController.Add)
	bonPlansEntreprise.GET("", bonPlansController.FindByEntreprise)
	bonPlansEntreprise.DELETE("/:id", bonPlansController.Remove)

	//Job API
	jobApi := api.Group("/jobs")
	jobApi.GET("/date", jobController.FindByDate)
	jobApi.GET("/entreprise", jobController.FindByIdEntreprise)
	jobApi.POST("", jobController.Add)
	jobApi.PUT("/:id", jobController.Update)
	jobApi.DELETE("/:id", jobController.Remove)

	//Go routine pour scanner les rSS
	go new(metiers.ScanRssMetier).Start()

	//On regarde comment on démarre le serveur
	env := os.Getenv("ENV")

	if env == "dev" {
		e.Logger.Fatal(e.Start(":1330"))
	} else{
		//e.AutoTLSManager.Cache = autocert.DirCache("/var/www/Golang-Projects/src/tomuss_server.cache")
		//e.Logger.Fatal(e.StartAutoTLS(":1330"))
		e.Logger.Fatal(e.Start(":1330"))
	}
}