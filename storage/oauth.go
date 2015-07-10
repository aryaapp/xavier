package storage

type NewTokenInput struct {
	Email     string   `json:"email" validate:"nonzero,email"`
	Password  string   `json:"password" validate:"nonzero"`
	GrantType string   `json:"grant_type" validate:"nonzero,grant_type"`
	Scopes    []string `json:"scopes"`
}
