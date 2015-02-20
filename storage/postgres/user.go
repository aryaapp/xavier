package postgres

import (
	"strings"
	"xavier/storage"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserDatabase struct {
	*sqlx.DB
}

func (db *UserDatabase) Find(uuid string) (*storage.User, error) {
	u := &storage.User{}
	err := db.Get(u, "SELECT id, uuid, email, created_at, updated_at, fullname, gender, theme_id FROM users u WHERE u.uuid = $1", uuid)
	return u, err
}

func (db *UserDatabase) FindByID(id int) (*storage.User, error) {
	u := &storage.User{}
	err := db.Get(u, "SELECT id, uuid, email, created_at, updated_at, fullname, gender, theme_id FROM users u WHERE u.id = $1", id)
	return u, err
}

func (db *UserDatabase) FindByEmail(email string) (*storage.User, error) {
	u := &storage.User{}
	err := db.Get(u, "SELECT id, uuid, email, password_digest FROM users u WHERE u.email = $1", strings.ToLower(email))
	return u, err
}

func (db *UserDatabase) Insert(user *storage.User) error {
	count := 0
	if err := db.Get(&count, "SELECT COUNT(*) FROM users u WHERE u.email = $1", strings.ToLower(user.Email)); err != nil {
		return err
	} else if count > 0 {
		return storage.UserConflictError
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (email, password_digest, created_at, updated_at, theme_id) VALUES ($1,$2,$3,$4,$5)"
	query += "RETURNING uuid, email, created_at, updated_at, theme_id"

	return db.Get(user, query, strings.ToLower(user.Email), password, user.CreatedAt, user.UpdatedAt, user.ThemeID)
	// fields := "email, created_at, updated_at, theme_id"
	// query := fmt.Sprintf("INSERT INTO users (%s, password) VALUES (?, ?, ?, ?) RETURNING uuid, %s", fields, fields)
	// return db.Get(user, query, strings.ToLower(user.Email), user.CreatedAt, user.UpdatedAt, user.ThemeID, user.Password)
}
