package storage

type App struct {
	ID         int     `json:"-" db:"id"`
	UUID       string  `json:"uuid" db:"uuid"`
	Secret     string  `json:"-" db:"secret"`
	Name       string  `json:"name" db:"name"`
	Url        string  `json:"url" db:"url"`
	GrantTypes Strings `json:"grant_types" db:"grant_types"`
	Scopes     Strings `json:"permitted_scopes" db:"permitted_scopes"`
}

type AppStorage interface {
	FindAll() ([]App, error)
	FindByUUID(uuid string) (*App, error)
	FindByID(id int) (*App, error)
}

type AppCache interface {
	FindByUUID(uuid string) (*App, error)
	New(apps []App)
}
