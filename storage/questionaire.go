package storage

import "xavier/lib/util/pg"

type Questionaire struct {
	ID        int     `json:"-" db:"id"`
	UUID      string  `json:"uuid" db:"uuid"`
	Title     string  `json:"title" db:"title"`
	Questions pg.JSON `json:"questions" db:"questions"`
}

type QuestionaireStorage interface {
	All(int) ([]Questionaire, error)
	Find(string, int) (*Questionaire, error)
}
