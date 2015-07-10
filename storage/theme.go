package storage

type Theme struct {
	ID        int    `json:"-" db:"id"`
	UUID      string `json:"uuid" db:"uuid"`
	Color     string `json:"image" db:"color"`
	Wallpaper string `json:"wallpaper" db:"wallpaper"`
}

type ThemeStorage interface {
	FindAll() ([]Theme, error)
	FindByUUID(uuid string) (*Theme, error)
	FindByID(id int) (*Theme, error)
}
