package storage

import "time"

type Journal struct {
	ID        int       `json:"-" db:"id"`
	UUID      string    `json:"uuid" db:"uuid"`
	Feeling   int       `json:"feeling" db:"feeling"`
	Questions JSON      `json:"questions" db:"questions"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	AppID     int       `json:"-" db:"app_id"`
	UserID    int       `json:"-" db:"user_id"`
	Answers   []Answer  `json:"answers,omitempty"`
	Metadata  JSON      `json:"metadata,omitempty" db:"metadata"`
}

type Answer struct {
	ID         int       `json:"-" db:"id"`
	UUID       string    `json:"uuid" db:"uuid"`
	Values     JSON      `json:"values" db:"values"`
	Answered   bool      `json:"answered" db:"answered"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	JournalID  int       `json:"-" db:"journal_id"`
	QuestionID int       `json:"-" db:"question_id"`
}

type JournalStorage interface {
	FindAll(userID int) ([]Journal, error)
	FindForLastWeek(userID int) ([]Journal, error)
	FindByUUID(uuid string, userID int) (*Journal, error)
	FindAllAnswers(journalID int) ([]Answer, error)
	New(*Journal) error
}
