package daos

import (
	"gopkg.in/olivere/elastic.v5"
	"tomuss_server/src/models"
	"context"
	"errors"
	"encoding/json"
)

const index = "tomuss"

type EtudiantDao struct{}

func (*EtudiantDao) ChangePassword(client *elastic.Client, idEtudiant string, newPassword string) error {
	ctx := context.Background()

	if updated, err := client.Update().
		Index(index).
		Type("students").
		Id(idEtudiant).
		Script(elastic.NewScriptInline("" +
		"ctx._source.password = params.newPassword ").
		Lang("painless").
		Param("newPassword", newPassword)).
		Do(ctx); updated == nil || err != nil {
		return errors.New("Erreur lors de la modification de votre mot de passe")
	}

	return nil
}

func (*EtudiantDao) ChangeInformatios(client *elastic.Client, idEtudiant string, informationsStudent *models.InformationStudent) error {
	ctx := context.Background()

	if updated, err := client.Update().
		Index(index).
		Type("students").
		Id(idEtudiant).
		Script(elastic.NewScriptInline("" +
		"ctx._source.nom = params.nom; " +
		"ctx._source.prenom = params.prenom;").
		Lang("painless").
		Param("prenom", informationsStudent.Prenom).
		Param("nom", informationsStudent.Nom)).
		Do(ctx); updated == nil || err != nil {
		return errors.New("Erreur lors de la modification de vos informations")
	}

	return nil
}

func (*EtudiantDao) AddSemestre(client *elastic.Client, idEtudiant string, semestre *models.Semestre) error {
	ctx := context.Background()

	if updated, err := client.Update().
		Index(index).
		Type("students").
		Id(idEtudiant).
		Script(elastic.NewScriptInline("" +
		"if (ctx._source.semestres == null) {ctx._source.semestres = new ArrayList();} " +
		"ctx._source.semestres.add(params.semestre); ").
		Lang("painless").
		Param("semestre", semestre)).
		Do(ctx); updated == nil || err != nil {
		return errors.New("Erreur lors de l'ajout du semestre")
	}

	return nil
}

func (*EtudiantDao) UpdateSemestres(client *elastic.Client, idEtudiant string, semestres []models.Semestre) error {
	ctx := context.Background()

	if updated, err := client.Update().
		Index(index).
		Type("students").
		Id(idEtudiant).
		Script(elastic.NewScriptInline("" +
		"ctx._source.semestres = params.semestres; ").
		Lang("painless").
		Param("semestres", semestres)).
		Do(ctx); updated == nil || err != nil {
		return errors.New("Erreur lors de la modification de votre semestre")
	}

	return nil
}

func (*EtudiantDao) UpdateSemestresWithVersion(client *elastic.Client, idEtudiant string, semestres []models.Semestre, version int) error {
	ctx := context.Background()

	if updated, err := client.Update().
		Index(index).
		Type("students").
		Version(int64(version)).
		Id(idEtudiant).
		Refresh("true").
		Script(elastic.NewScriptInline("" +
		"ctx._source.semestres = params.semestres; ").
		Lang("painless").
		Param("semestres", semestres)).
		Do(ctx); updated == nil || err != nil {
		return errors.New("Erreur lors de la modification de votre semestre")
	}

	return nil
}

func (*EtudiantDao) GetById(client *elastic.Client, idEtudiant string) (*models.Etudiant, error) {
	ctx := context.Background()

	result, err := client.Get().
		Index(index).
		Type("students").
		Id(idEtudiant).
		Pretty(true).
		Do(ctx)

	if err != nil || result == nil {
		return nil, errors.New("Erreur lors de la récupération de l'étudiant")
	}

	etudiant := new(models.Etudiant)

	marshal, err := json.Marshal(result)
	if err != nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	//On regarde si on peut deserialiser le hit en annonce
	if err := json.Unmarshal(marshal, etudiant); err != nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	return etudiant, nil
}

func (*EtudiantDao) GetByEmail(client *elastic.Client, email string) (*models.Etudiant, error) {
	ctx := context.Background()

	matchQuery := elastic.NewMatchQuery("email", email)

	results, err := client.Search().
		Index(index).
		Type("students").
		Query(matchQuery).
		Pretty(true).
		Do(ctx)

	if err != nil || results == nil || results.Hits == nil {
		return nil, errors.New("Erreur lors de la récupération de l'étudiant")
	}

	if results.Hits.TotalHits == 0 {
		return nil, errors.New("Aucun compte avec cet email")
	}

	//On récup le premier compte
	first_etudiant := results.Hits.Hits[0]
	etudiant := new(models.Etudiant)

	bytes, err := json.Marshal(first_etudiant)

	//On parse le json en objet
	err_unmarshal := json.Unmarshal(bytes, &etudiant)
	if err_unmarshal != nil {
		return nil, errors.New("Erreur lors de la récupération de votre compte")
	}

	return etudiant, nil
}

func (*EtudiantDao) Add(client *elastic.Client, etudiant *models.Etudiant) error {
	ctx := context.Background()

	result, err := client.Index().
		Index(index).
		Type("students").
		BodyJson(etudiant.Source).
		Pretty(true).
		Do(ctx)

	if result == nil || err != nil {
		return errors.New("Erreur lors de la création de votre compte")
	}

	etudiant.Id = result.Id
	etudiant.Type = result.Type
	etudiant.Version = result.Version

	return nil
}

func (*EtudiantDao) Find(client *elastic.Client) ([]*models.Etudiant, error) {
	//On récup le context
	ctx := context.Background()

	//On recherche tous les étudiants
	results, err := client.Search().
		Index(index).
		Type("students").
		Version(true).
		Pretty(true).
		Do(ctx)

	//On regarde si on a eu une erreur lors de la recherche
	if err != nil || results == nil || results.Hits == nil {
		return nil, errors.New("Erreur lors de la récupération des étudiants")
	}

	//Création du tableau de retour
	etudiants := make([]*models.Etudiant, 0)

	//Si on a aucun résultats, on retourne le tableau vide
	if results.Hits.TotalHits == 0 {
		return etudiants, nil
	}

	//On parcourt les résultats pour les ajouter au tableau d'étudiants
	for _, hit := range results.Hits.Hits {
		etudiant := new(models.Etudiant)

		marshal, err := json.Marshal(hit)
		if err != nil {
			return nil, errors.New("Erreur lors de la récupération des étudiants")
		}

		//On regarde si on peut deserialiser le hit en étudiant
		if err := json.Unmarshal(marshal, etudiant); err != nil {
			return nil, errors.New("Erreur lors de la récupération des étudiants")
		}

		etudiants = append(etudiants, etudiant)
	}

	return etudiants, nil
}
