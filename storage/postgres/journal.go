package postgres

import (
	"github.com/aryaapp/xavier/storage"
	"github.com/jmoiron/sqlx"
)

type JournalDatabase struct {
	*sqlx.DB
}

const (
	journalsAll      = `SELECT * FROM journals j WHERE j.user_id = $1`
	journalsFind     = `SELECT * FROM journals j WHERE j.uuid = $1 AND j.user_id = $2 LIMIT 1`
	journalsLastWeek = `SELECT * FROM journals j WHERE j.user_id = $1 AND created_at >= now()::date - 7`
	journalsAnswers  = `SELECT * FROM answers a WHERE a.journal_id = $1`
	journalsInsert   = `
		INSERT INTO journals (feeling, questions, created_at, updated_at, app_id, user_id) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING uuid, id, feeling, questions, created_at, updated_at, app_id, user_id`
	journalsAnswerInsert = `
		INSERT INTO answers (values, created_at, updated_at, journal_id, question_id) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING uuid, values, created_at, updated_at, journal_id, question_id`
)

func (db *JournalDatabase) FindAll(userID int) ([]storage.Journal, error) {
	j := []storage.Journal{}
	return j, db.Select(&j, journalsAll, userID)
}

func (db *JournalDatabase) FindForLastWeek(userID int) ([]storage.Journal, error) {
	j := []storage.Journal{}
	err := db.Select(&j, journalsLastWeek, userID)
	return j, err
}

func (db *JournalDatabase) FindByUUID(uuid string, userID int) (*storage.Journal, error) {
	j := &storage.Journal{}
	return j, db.Get(j, journalsFind, uuid, userID)
}

func (db *JournalDatabase) FindAllAnswers(journalID int) ([]storage.Answer, error) {
	a := []storage.Answer{}
	err := db.Select(&a, journalsAnswers, journalID)
	return a, err
}

func (db *JournalDatabase) New(journal *storage.Journal) error {
	tx := db.MustBegin()

	if err := tx.Get(journal, journalsInsert, journal.Feeling, journal.Questions, journal.CreatedAt, journal.UpdatedAt, journal.AppID, journal.UserID); err != nil {
		return err
	}

	answers := []storage.Answer{}
	for _, a := range journal.Answers {
		answer := storage.Answer{}
		if err := tx.Get(&answer, journalsAnswerInsert, a.Values, journal.CreatedAt, journal.UpdatedAt, journal.ID, a.QuestionID); err != nil {
			return err
		}
		answers = append(answers, answer)
	}
	journal.Answers = answers
	tx.Commit()

	return nil
}
