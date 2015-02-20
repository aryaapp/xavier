package postgres

import (
	"xavier/storage"

	"github.com/jmoiron/sqlx"
)

type JournalDatabase struct {
	*sqlx.DB
}

func (db *JournalDatabase) All(userID int) ([]storage.Journal, error) {
	j := []storage.Journal{}
	err := db.Select(&j, "SELECT * FROM journals j WHERE j.user_id = $1", userID)
	return j, err
}

func (db *JournalDatabase) Find(uuid string, userID int) (*storage.Journal, error) {
	j := &storage.Journal{}
	err := db.Get(j, "SELECT * FROM journals j WHERE j.uuid = $1 && j.user_id = $2", uuid, userID)
	return j, err
}

func (db *JournalDatabase) LastWeek(userID int) ([]storage.Journal, error) {
	j := []storage.Journal{}
	err := db.Select(&j, "SELECT * FROM journals j WHERE j.user_id = $1 AND created_at >= now()::date - 7", userID)
	return j, err
}

func (db *JournalDatabase) Insert(journal *storage.Journal) error {
	tx := db.MustBegin()

	journalQuery := "INSERT INTO journals (feeling, created_at, updated_at, client_id, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING uuid, id, feeling, created_at, updated_at, client_id, user_id"
	if err := tx.Get(journal, journalQuery, journal.Feeling, journal.CreatedAt, journal.UpdatedAt, journal.ClientID, journal.UserID); err != nil {
		return err
	}

	answers := []storage.Answer{}
	answersQuery := "INSERT INTO answers (values, created_at, updated_at, journal_id, question_id) VALUES ($1, $2, $3, $4, $5) RETURNING uuid, values, created_at, updated_at, journal_id, question_id"
	for _, a := range journal.Answers {
		answer := storage.Answer{}
		if err := tx.Get(&answer, answersQuery, a.Values, journal.CreatedAt, journal.UpdatedAt, journal.ID, a.QuestionID); err != nil {
			return err
		}
		answers = append(answers, answer)
	}
	journal.Answers = answers
	tx.Commit()

	return nil
}
