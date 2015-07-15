package storage

type Questionaire struct {
	ID        int    `json:"-" db:"id"`
	UUID      string `json:"uuid" db:"uuid"`
	Title     string `json:"title" db:"title"`
	Questions JSON   `json:"questions" db:"questions"`
}

type QuestionaireStorage interface {
	FindAll(int) ([]Questionaire, error)
	FindByUUID(string, int) (*Questionaire, error)
}
