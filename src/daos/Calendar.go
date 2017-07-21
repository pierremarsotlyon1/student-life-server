package daos

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"errors"
	"tomuss_server/src/models"
)

type CalendarDao struct{}

func (*CalendarDao) UpdateUrlIcs(client *elastic.Client, idUser string, urlIcs string) error {
	if _, err := client.Update().
		Index(index).
		Type("students").
		Id(idUser).
		Script(elastic.NewScriptInline("" +
		"ctx._source.url_ics = params.urlIcs").Param("urlIcs", urlIcs)).
		Do(context.Background()); err != nil {
		return errors.New("Erreur lors de la sauvegarde du lien de votre calendrier")
	}

	return nil
}

func (*CalendarDao) RefreshEvents(client *elastic.Client, idUser string, events []*models.Event) error {
	if _, err := client.Update().
		Index(index).
		Type("students").
		Id(idUser).
		Script(elastic.NewScriptInline("" +
		"ctx._source.calendar = params.calendar").Param("calendar", events)).
		Do(context.Background()); err != nil {
		return errors.New("Erreur lors de la sauvegarde des events de votre calendrier")
	}

	return nil
}

func (*CalendarDao) FindEvents(client *elastic.Client, idUser string) ([]models.Event, error) {
	etudiant, err := new(EtudiantDao).GetById(client, idUser)

	if err != nil || etudiant == nil {
		return nil, err
	}

	return etudiant.Source.Calendar, nil
}
