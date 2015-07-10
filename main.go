package main

import (
	"fmt"
	"log"

	"github.com/aryaapp/xavier/api"
	"github.com/aryaapp/xavier/jwt"

	p "github.com/aryaapp/xavier/storage/postgres"
	r "github.com/aryaapp/xavier/storage/redis"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
)

func main() {
	c := api.NewConfig()
	err := envconfig.Process("hermes", c)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Postgres
	pg, err := sqlx.Connect("postgres", c.PostgresURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer pg.Close()

	a := &api.AppContext{
		Config:              c,
		JWTClient:           &jwt.Client{c.PrivateKey},
		AppCache:            &r.AppCache{},
		AppStorage:          &p.AppDatabase{pg},
		DeviceStorage:       &p.DeviceDatabase{pg},
		JournalStorage:      &p.JournalDatabase{pg},
		KeywordStorage:      &p.KeywordDatabase{pg},
		NoteStorage:         &p.NoteDatabase{pg},
		QuestionStorage:     &p.QuestionDatabase{pg},
		QuestionaireStorage: &p.QuestionaireDatabase{pg},
		ThemeStorage:        &p.ThemeDatabase{pg},
		UserStorage:         &p.UserDatabase{pg},
	}

	// Move all apps to memory cache
	apps, err := a.AppStorage.FindAll()
	if err != nil {
		log.Fatalln(err)
	}
	a.AppCache.New(apps)

	e := echo.New()
	e.Use(mw.Logger())

	api.Mux(e, a)

	addr := fmt.Sprintf(":%d", c.Port)
	log.Printf("Starting server on: %s", addr)
	e.Run(addr)
}
