package postgres

import (
	"xavier/storage"

	"github.com/jmoiron/sqlx"
)

type ClientDatabase struct {
	*sqlx.DB
}

func (db *ClientDatabase) All() ([]storage.Client, error) {
	c := []storage.Client{}
	err := db.Select(&c, "SELECT id, uuid, name, url, secret FROM clients c")
	return c, err
}

func (db *ClientDatabase) Find(uuid string) (*storage.Client, error) {
	c := &storage.Client{}
	err := db.Get(c, "SELECT id, uuid, name, url, secret, grant_types FROM clients c WHERE c.uuid = $1 LIMIT 1", uuid)
	return c, err
}

func (db *ClientDatabase) FindByID(id int) (*storage.Client, error) {
	c := &storage.Client{}
	err := db.Get(c, "SELECT id, uuid, name, url, secret FROM clients c WHERE c.id = $1 LIMIT 1", id)
	return c, err
}
