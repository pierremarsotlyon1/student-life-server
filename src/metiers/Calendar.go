package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"errors"
	"github.com/asaskevich/govalidator"
	"tomuss_server/src/daos"
	"github.com/PuloV/ics-golang"
	"tomuss_server/src/models"
	"time"
)

type CalendarMetier struct{}

func (calendrierMetier *CalendarMetier) UpdateUrlIcs(client *elastic.Client, idUser string, urlIcs string) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idUser) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if len(urlIcs) == 0 {
		return errors.New("Erreur lors de la récupération de l'url du calendrier")
	}

	if !govalidator.IsURL(urlIcs) {
		return errors.New("L'url du calendrier est mal formatée")
	}

	//On update l'url de l'ics dans le document de l'étudiant
	if err := new(daos.CalendarDao).UpdateUrlIcs(client, idUser, urlIcs); err != nil {
		return err
	}

	//On démarre un routine pour parser les events
	go calendrierMetier.ParseIcs(client, idUser, urlIcs)

	return nil
}

/*func (*CalendarMetier) FindEvents(client *elastic.Client, idUser string) ([]models.Event, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idUser) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre identifiant")
	}

	events, err := new(daos.CalendarDao).FindEvents(client, idUser)
	if err != nil {
		return nil, err
	}

	return events, nil
}*/

func (calendrierMetier *CalendarMetier) Synchroniser (client *elastic.Client, idUser string) ([]*models.Event, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idUser) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre identifiant")
	}

	//On récupère l'étudiant
	etudiant, err := new(daos.EtudiantDao).GetById(client, idUser)
	if err != nil {
		return nil, err
	}

	if etudiant == nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	//On regarde si l'étudiant a déjà enregistré l'url de son calendrier
	if len(etudiant.Source.UrlIcs) == 0 {
		return nil, errors.New("Vous devez enregistrer l'url de votre calendrier avant de rafraichir les events")
	}

	//On lance une goroutine pour refresh le calendrier
	events, err := calendrierMetier.ParseIcs(client, idUser, etudiant.Source.UrlIcs)

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (*CalendarMetier) ParseIcs(client *elastic.Client, idUser string, urlIcs string) ([]*models.Event, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idUser) == 0 {
		return nil, errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if len(urlIcs) == 0 {
		return nil, errors.New("Erreur lors de la récupération de l'url du calendrier")
	}

	if !govalidator.IsURL(urlIcs) {
		return nil, errors.New("L'url du calendrier est mal formatée")
	}

	//Création du parser
	parser := ics.New()

	//On récup le chanel
	inputChan := parser.GetInputChan()

	//On lui demande de parser le fichier
	inputChan <- urlIcs

	//On attend que le parsing soit fini
	parser.Wait()

	//On récupère les calendrier
	cal, err := parser.GetCalendars()

	//On regarde si on a une erreur
	if err == nil {

		//On récup le calendrier
		calendar := cal[0]

		//Création du tableau des nouveaux events
		var events []*models.Event

		for _, e := range calendar.GetEvents() {
			event := new(models.Event)

			event.Location = e.GetLocation()
			event.Titre = e.GetSummary()
			event.Description = e.GetDescription()
			event.DateDebut = e.GetStart().Format(time.RFC3339)
			event.DateFin = e.GetEnd().Format(time.RFC3339)

			events = append(events, event)
		}

		//On update dans le document de l'étudiant
		if err := new(daos.CalendarDao).UpdateEvents(client, idUser, events); err != nil {
			return nil, err
		}

		return events, nil
	} else {
		return nil, errors.New("Erreur lors de la récupération des événements")
	}
}
