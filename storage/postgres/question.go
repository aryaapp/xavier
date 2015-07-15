package postgres

import (
	"github.com/aryaapp/xavier/storage"
	"github.com/jmoiron/sqlx"
)

type QuestionDatabase struct {
	*sqlx.DB
}

const (
	questionFind = `
		SELECT q.* FROM questionaires qs
			JOIN (
				SELECT uuid, title, description, view, important, autocompletes, user_data FROM questions q
				WHERE q.uuid = $1
				LIMIT 1
			) q ON TRUE
		JOIN questionaires_users qu ON qu.questionaire_id = qs.id AND qu.user_id = $2
		LIMIT 1`
	questionWhereIn = `
		SELECT id, uuid, title, description, processor, view, important, autocompletes, user_data
		FROM questions q WHERE uuid IN (?)
	`
)

func (db *QuestionDatabase) FindByUUID(uuid string, userID int) (*storage.Question, error) {
	q := &storage.Question{}
	err := db.Get(q, questionFind, uuid, userID)
	return q, err
}

func (db *QuestionDatabase) WhereIn(uuids []string) ([]storage.Question, error) {
	q := []storage.Question{}
	query, _, err := sqlx.In(questionWhereIn, uuids)
	err = db.Select(&q, query)
	return q, err
}
