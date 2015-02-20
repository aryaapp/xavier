package storage

import "time"

type Device struct {
	Token       string    `json:"token" db:"token"`
	Environment string    `json:"environment" db:"environment"`
	Name        string    `json:"name" db:"name"`
	Model       string    `json:"model" db:"model"`
	Os          string    `json:"os" db:"os"`
	OsVersion   string    `json:"os_version" db:"os_version"`
	AppVersion  string    `json:"app_version" db:"app_version"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	UserID      int       `json:"-" db:"user_id"`
}

type DeviceStorage interface {
	All(int) ([]Device, error)
	Find(string, int) (*Device, error)
	Exists(string, int) (bool, error)
	InsertOrUpdate(*Device) (bool, error)
}
