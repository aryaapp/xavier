package storage

type Keyword struct {
	ID         int    `json:"-" db:"id"`
	UUID       string `json:"uuid" db:"uuid"`
	Name       string `json:"name" db:"name"`
	Count      int    `json:"count" db:"count"`
	QuestionID int    `json:"-" db:"question_id"`
	UserID     int    `json:"-" db:"user_id"`
}

type KeywordStorage interface {
	All(string, int) ([]Keyword, error)
	Exists(string, int, int) (bool, error)
	InsertOrUpdate(*Keyword) error
}
