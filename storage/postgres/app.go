package postgres

import (
	"xavier/storage"

	"github.com/jmoiron/sqlx"
)

type AppDatabase struct {
	*sqlx.DB
}

func (db *AppDatabase) All() ([]storage.App, error) {
	a := []storage.App{}
	err := db.Select(&a, "SELECT id, uuid, name, url, secret, grant_types, permitted_scopes FROM apps a")
	return a, err
}

func (db *AppDatabase) Find(uuid string) (*storage.App, error) {
	a := &storage.App{}
	err := db.Get(a, "SELECT id, uuid, name, url, secret, grant_types FROM apps a WHERE a.uuid = $1 LIMIT 1", uuid)
	return a, err
}

func (db *AppDatabase) FindByID(id int) (*storage.App, error) {
	a := &storage.App{}
	err := db.Get(a, "SELECT id, uuid, name, url, secret FROM apps a WHERE a.id = $1 LIMIT 1", id)
	return a, err
}
