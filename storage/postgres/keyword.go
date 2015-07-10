package postgres

import (
	"github.com/aryaapp/xavier/storage"
	"github.com/jmoiron/sqlx"
)

type KeywordDatabase struct {
	*sqlx.DB
}

const (
	keywordsAll = `
		SELECT k.uuid, k.count, k.name, k.user_id FROM keywords k
		JOIN questions q ON q.id = k.question_id
		WHERE q.uuid = $1 AND k.user_id = $2
		ORDER BY k.count DESC
	`
	keywordsExists = `
		SELECT COUNT(*) FROM keywords k 
		WHERE k.name = $1 AND k.question_id = $2 AND k.user_id = $3
	`
	keywordsInsert = `
		INSERT INTO keywords (name, question_id, user_id) VALUES ($1, $2, $3) 
		RETURNING uuid, count, name, question_id, user_id
	`
	keywordsUpdate = `
		UPDATE keywords SET count = count + 1 
		WHERE name = $1 AND question_id = $2 AND user_id = $3
	`
)

func (db *KeywordDatabase) All(questionUUID string, userID int) ([]storage.Keyword, error) {
	k := []storage.Keyword{}
	return k, db.Select(&k, keywordsAll, questionUUID, userID)
}

func (db *KeywordDatabase) Exists(name string, questionID int, userID int) (bool, error) {
	var count int
	if err := db.Get(&count, keywordsExists, name, questionID, userID); err != nil {
		return false, err
	}
	return count == 1, nil
}

func (db *KeywordDatabase) InsertOrUpdate(k *storage.Keyword) error {
	exists, err := db.Exists(k.Name, k.QuestionID, k.UserID)
	if err != nil {
		return err
	}

	if exists {
		_, err := db.Query(keywordsUpdate, k.Name, k.QuestionID, k.UserID)
		return err
	}
	return db.Get(k, keywordsInsert, k.Name, k.QuestionID, k.UserID)
}
