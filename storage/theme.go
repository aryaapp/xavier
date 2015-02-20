package storage

type Theme struct {
	ID        int    `json:"-" db:"id"`
	UUID      string `json:"uuid" db:"uuid"`
	Color     string `json:"image" db:"color"`
	Wallpaper string `json:"wallpaper" db:"wallpaper"`
}

type ThemeStorage interface {
	All() ([]Theme, error)
	Find(string) (*Theme, error)
	FindByID(int) (*Theme, error)
}
