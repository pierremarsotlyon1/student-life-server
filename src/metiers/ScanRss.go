package metiers

import (
	"tomuss_server/src/tools"
	"net/http"
	"encoding/xml"
	"tomuss_server/src/models"
	"io/ioutil"
	"strings"
	"strconv"
	"tomuss_server/src/daos"
	"gopkg.in/olivere/elastic.v5"
	"time"
	"github.com/jasonlvhit/gocron"
)

type ScanRssMetier struct{}

func (scanRssMetier *ScanRssMetier) Start() {
	gocron.Every(1).Minutes().Do(scanRssMetier.scanRss)
	gocron.Start()
}

func (scanRssMetier *ScanRssMetier) scanRss() {
	client := tools.CreateElasticsearchClient()
	if client == nil {
		return
	}

	//On récup tous les étudiants
	etudiants, err := new(EtudiantMetier).Find(client)

	if err != nil || len(etudiants) == 0 {
		return
	}

	const pas = 100

	for i := 0; i < len(etudiants); i += pas {
		if (i + pas) > len(etudiants) {
			go scanRssMetier.ThreadEtudiant(client, etudiants[i:])
		} else {
			go scanRssMetier.ThreadEtudiant(client, etudiants[i:pas])
		}
	}
}

func (scanRssMetier *ScanRssMetier) ThreadEtudiant(client *elastic.Client, etudiants []*models.Etudiant) {
	etudiantDao := new(daos.EtudiantDao)

	somethingChange := false

	for _, etudiant := range etudiants {
		somethingChange = false

		if etudiant == nil || len(etudiant.Source.Semestres) == 0 {
			continue
		}

		//On récup le semestre actif

		for idSemestre := range etudiant.Source.Semestres {
			if !etudiant.Source.Semestres[idSemestre].Actif {
				continue
			}

			//On regarde si on a une url
			if len(etudiant.Source.Semestres[idSemestre].Url) == 0 {
				break
			}

			//On télécharge le rss
			resp, err := http.Get(etudiant.Source.Semestres[idSemestre].Url)
			if err != nil {
				break
			}

			//On parse le body en byte[]
			body, err := ioutil.ReadAll(resp.Body)

			resp.Body.Close()

			if err != nil {
				break
			}

			//On parse en objet
			var xmlFeed models.Xml
			err = xml.Unmarshal(body, &xmlFeed.Rss)

			if err != nil {
				break
			}

			//On parcourt les items du rss
			for _, item := range xmlFeed.Rss.Channel.Items {
				if len(item.Title) == 0 {
					continue
				}

				//On split le titre car il contient => ue : sujet : note|content
				titles := strings.Split(item.Title, ":")
				if len(titles) != 3 {
					continue
				}


				//On regarde si l'étudiant a l'ue
				var ue *models.Ue

				for idUe := range etudiant.Source.Semestres[idSemestre].Ues {
					if etudiant.Source.Semestres[idSemestre].Ues[idUe].Name == titles[0] {
						ue = &etudiant.Source.Semestres[idSemestre].Ues[idUe]
						break
					}
				}

				//On regarde si on a trouvé l'ue
				if ue == nil {
					//On ajoute l'ue au semestre
					ue = new(models.Ue)
					ue.Name = titles[0]
					etudiant.Source.Semestres[idSemestre].Ues = append(etudiant.Source.Semestres[idSemestre].Ues, *ue)
					somethingChange = true
				}

				//Si on a pas d'erreur lors de la conversion, l'item est une note
				if strings.Contains(titles[2], "/") {

					notes := strings.Split(titles[2], "/")
					if len(notes) == 0 {
						break
					}

					//On trim la note
					notes[0] = strings.TrimSpace(notes[0])

					if _, err := strconv.ParseFloat(notes[0], 64); err == nil {

						//On parse le dénominateur
						if denominateur, err := strconv.ParseFloat(notes[1], 64); err == nil {

							//On ajoute la note si il ne l'a pas déjà
							find := false
							for _, note := range ue.Notes {
								if note.Guid == item.Guid {
									find = true
									break
								}
							}

							if !find {
								//On récupère le datetime courant
								t := time.Now().UTC().Format(time.RFC3339)

								ue.Notes = append(ue.Notes, models.Note{
									Guid:    item.Guid,
									Name:    titles[1],
									Created: t,//item.PubDate,
									Note:    notes[0],
									Denominateur: denominateur,
								})
								somethingChange = true
							}
						}
					}

				} else {
					switch titles[2] {
					case "PRST":

						break
					default:
						break
					}
				}
			}

			break
		}

		//On update le tableau des semetres de l'étudiant
		if somethingChange {
			etudiantDao.UpdateSemestresWithVersion(client, etudiant.Id, etudiant.Source.Semestres, etudiant.Version)
		}
	}
}
