package storage

type Client struct {
	ID     int    `json:"-" db:"id"`
	UUID   string `json:"uuid" db:"uuid"`
	Secret string `json:"-" db:"secret"`
	Name   string `json:"name" db:"name"`
	Url    string `json:"url" db:"url"`
	// GrantTypes GrantTypes `json:"grant_types" db:"grant_types"`
	// Scopes     Scopes     `json:"permitted_scopes" db:"permitted_scopes"`
}

type ClientStorage interface {
	All() ([]Client, error)
	Find(string) (*Client, error)
	FindByID(int) (*Client, error)
}

type ClientCache interface {
	Find(string) (*Client, error)
	Insert([]Client)
	// Insert(*Client) error
}
