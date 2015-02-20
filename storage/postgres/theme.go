package postgres

import (
	"xavier/storage"

	"github.com/jmoiron/sqlx"
)

type ThemeDatabase struct {
	*sqlx.DB
}

func (db *ThemeDatabase) All() ([]storage.Theme, error) {
	t := []storage.Theme{}
	return t, db.Select(&t, "SELECT uuid, color, wallpaper FROM themes t")
}

func (db *ThemeDatabase) Find(uuid string) (*storage.Theme, error) {
	t := &storage.Theme{}
	return t, db.Get(t, "SELECT uuid, color, wallpaper FROM themes t WHERE uuid = $1", uuid)
}

func (db *ThemeDatabase) FindByID(id int) (*storage.Theme, error) {
	t := &storage.Theme{}
	return t, db.Get(t, "SELECT id, uuid, color, wallpaper FROM themes t WHERE id = $1", id)
}
