package storage

type Question struct {
	ID            int       `json:"-" db:"id"`
	UUID          string    `json:"uuid" db:"uuid"`
	Title         string    `json:"title" db:"title"`
	Description   string    `json:"description" db:"description"`
	Processor     string    `json:"-" db:"processor"`
	View          string    `json:"view" db:"view"`
	Important     bool      `json:"important" db:"important"`
	Autocompletes bool      `json:"autocompletes" db:"autocompletes"`
	Keywords      []Keyword `json:"keywords,omitempty"`
	UserData      JSON      `json:"user_data" db:"user_data"`
}

type QuestionStorage interface {
	Find(string, int) (*Question, error)
	WhereIn([]string) ([]Question, error)
}
