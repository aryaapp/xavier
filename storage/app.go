package storage

import "xavier/lib/util/pg"

type App struct {
	ID         int        `json:"-" db:"id"`
	UUID       string     `json:"uuid" db:"uuid"`
	Secret     string     `json:"-" db:"secret"`
	Name       string     `json:"name" db:"name"`
	Url        string     `json:"url" db:"url"`
	GrantTypes pg.Strings `json:"grant_types" db:"grant_types"`
	Scopes     pg.Strings `json:"permitted_scopes" db:"permitted_scopes"`
}

type AppStorage interface {
	All() ([]App, error)
	Find(string) (*App, error)
	FindByID(int) (*App, error)
}

type AppCache interface {
	Find(string) (*App, error)
	Insert([]App)
}
