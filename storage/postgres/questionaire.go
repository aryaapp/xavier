package postgres

import (
	"github.com/aryaapp/xavier/storage"

	"github.com/jmoiron/sqlx"
)

type QuestionaireDatabase struct {
	*sqlx.DB
}

const (
	questionairesAll = `
		SELECT qs.id, qs.uuid, qs.title, json_agg(q.*) as questions FROM questionaires qs
		LEFT JOIN LATERAL (
			SELECT uuid, title, description, view, important, autocompletes, user_data FROM questions q
			WHERE q.id IN (SELECT(UNNEST(qs.questions)))
			LIMIT array_length(qs.questions, 1)
		) q ON TRUE
		JOIN questionaires_users qu ON qu.questionaire_id = qs.id AND qu.user_id = $1
		GROUP BY qs.id`
	questionairesFind = `
		SELECT qs.id, qs.uuid, qs.title, json_agg(q.*) as questions FROM questionaires qs
		LEFT JOIN LATERAL (
			SELECT uuid, title, description, view, important, autocompletes, user_data FROM questions q
			WHERE q.id IN (SELECT(UNNEST(qs.questions)))
			LIMIT array_length(qs.questions, 1)
		) q ON TRUE
		JOIN questionaires_users qu ON qs.uuid = $1 AND qu.user_id = $2
		GROUP BY qs.id
		LIMIT 1`
)

func (db *QuestionaireDatabase) FindAll(userID int) ([]storage.Questionaire, error) {
	questionaires := []storage.Questionaire{}
	err := db.Select(&questionaires, questionairesAll, userID)
	return questionaires, err
}

func (db *QuestionaireDatabase) FindByUUID(uuid string, userID int) (*storage.Questionaire, error) {
	q := &storage.Questionaire{}
	err := db.Get(q, questionairesFind, uuid, userID)
	return q, err
}
