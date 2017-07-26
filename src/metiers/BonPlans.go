package metiers

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"errors"
	"tomuss_server/src/daos"
	"github.com/go-shadow/moment"
	"time"
	"github.com/asaskevich/govalidator"
)

type BonPlansMetier struct {}

func (*BonPlansMetier) Find (client *elastic.Client, offset int) ([]*models.BonPlan, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	//On récup les bon plans
	bonplans, err := new(daos.BonPlansDao).Find(client, offset)

	if err != nil {
		return nil, err
	}

	return bonplans, nil
}

func (*BonPlansMetier) FindRecent (client *elastic.Client, date string) ([]*models.BonPlan, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(date) == 0 {
		return nil, errors.New("Erreur lors de la récupération de la date")
	}

	//On récup les bon plans
	bonplans, err := new(daos.BonPlansDao).FindRecent(client, date)

	if err != nil {
		return nil, err
	}

	return bonplans, nil
}

func (*BonPlansMetier) FindByEntreprise (client *elastic.Client, idEntreprise string, offset int) ([]*models.BonPlan, error) {
	if client == nil {
		return nil, errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEntreprise) == 0 {
		return nil, errors.New("Erreur lors de la récupération de l'identifiant de l'entreprise")
	}

	//On récup les bon plans
	bonplans, err := new(daos.BonPlansDao).FindByEntreprise(client, idEntreprise, offset)

	if err != nil {
		return nil, err
	}

	return bonplans, nil
}

func (*BonPlansMetier) Add (client *elastic.Client, idEntreprise string, bonplan *models.BonPlan) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idEntreprise) == 0 {
		return errors.New("Erreur lors de la récupération de l'identifiant de l'entreprise")
	}

	//On regarde si l'entreprise existe
	entreprise, err := new(daos.EntrepriseDao).GetById(client, idEntreprise)
	if err != nil {
		return err
	}

	if entreprise == nil {
		return errors.New("Erreur lors de la récupération du compte de l'entreprise")
	}

	if len(bonplan.Source.Title) == 0 {
		return errors.New("Vous devez saisir un titre")
	}

	if len(bonplan.Source.IdCategorie) == 0 {
		return errors.New("Erreur lors de la récupération de la catégorie d'annonce")
	}

	if len(bonplan.Source.Url) > 0 {
		if !govalidator.IsURL(bonplan.Source.Url) {
			return errors.New("L'url n'est pas au bon format")
		}
	}

	//On regarde si la catégorie d'annonce existe
	exist, err := new(CategorieAnnonceMetier).Exist(client, bonplan.Source.IdCategorie)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("Erreur lors de la récupération de la catégorie d'annonce")
	}

	//On affecte les propriétés
	m := moment.New()
	bonplan.Source.NomEnreprise = entreprise.Source.NomEntreprise
	bonplan.Source.LogoEntreprise = entreprise.Source.LogoEntreprise
	bonplan.Source.Created = m.UTC().FormatGo(time.RFC3339)
	bonplan.Source.DateDebut = m.Format("YYYY-MM-DD")
	bonplan.Source.DateFin = m.AddMonths(1).Format("YYYY-MM-DD")

	if err := new(daos.BonPlansDao).Add(client, idEntreprise, bonplan); err != nil {
		return err
	}

	return nil
}

func (*BonPlansMetier) Remove (client *elastic.Client, idEntreprise string, idBonPlan string) error {
	if client == nil {
		return errors.New("Erreur lors de la connexion à notre base de donnée")
	}

	if len(idBonPlan) == 0 {
		return errors.New("Erreur lors de la récupération de l'identifiant du bon plan")
	}

	if len(idEntreprise) == 0 {
		return errors.New("Erreur lors de la récupération de votre identifiant")
	}

	if err := new(daos.BonPlansDao).Remove(client, idEntreprise, idBonPlan); err != nil {
		return err
	}

	return nil
}
