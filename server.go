package main

import (
	"github.com/labstack/echo"
	"tomuss_server/src/controllers"
	"github.com/labstack/echo/middleware"
	"tomuss_server/src/metiers"
	"golang.org/x/crypto/acme/autocert"
	"os"
)

func main() {
	e := echo.New()

	//CORS
	e.Use(middleware.CORS())

	//Association des routes
	//Définition des controllers
	etudiantController := new(controllers.EtudiantController)
	semestreController := new(controllers.SemestreController)

	//Gerant Controller sans JWT
	e.POST("/login", etudiantController.Login)
	e.POST("/register", etudiantController.Register)

	//Définition de l'api de base avec restriction JWT
	api := e.Group("/api")
	api.Use(middleware.JWT([]byte(metiers.GetSecretJwt())))

	//Api pour l'étudiant
	etudiantApi := api.Group("/etudiant")
	etudiantApi.PUT("/fcm", etudiantController.ChangeFcmToken)
	etudiantApi.PUT("/change/password", etudiantController.ChangePassword)
	etudiantApi.PUT("/change/informations", etudiantController.ChangeInformations)

	etudiantApi.GET("/semestres", semestreController.Find)
	etudiantApi.POST("/semestres", semestreController.Add)
	etudiantApi.PUT("/semestres/:id", semestreController.Update)
	etudiantApi.DELETE("/semestres/:id", semestreController.Remove)
	//semestreEtudiantApi.GET("", etudiantController.Profile)

	//Api pour les annonces
	/*annonceApi := api.Group("/annonce")
	annonceApi.GET("", annonceController.Find)
	annonceApi.GET("/:id", annonceController.Get)
	annonceApi.POST("", annonceController.Add)
	annonceApi.DELETE("/:id", annonceController.Delete)
	annonceApi.PUT("/:id", annonceController.Update)

	//Recherche des annonces par geolocation
	e.GET("/annonce/search/location", annonceController.SearchByLocation)*/

	go new(metiers.ScanRssMetier).Start()

	//On regarde comment on démarre le serveur
	env := os.Getenv("ENV")

	if env == "dev" {
		e.Logger.Fatal(e.Start(":1330"))
	} else{
		e.AutoTLSManager.Cache = autocert.DirCache("/var/www/Golang-Projects/src/tomuss_server.cache")
		e.Logger.Fatal(e.StartAutoTLS(":1330"))
	}
}