package main

import (
	"log"
	"xavier/app"
	"xavier/routes"

	"github.com/codegangsta/negroni"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	ctx := NewContext()
	router := app.NewRouter()
	router.Group("/", func(r *app.Router) {
		r.Post("/oauth/tokens", app.Handler(ctx, routes.OAuthTokensCreate))

		r.Get("/themes", app.Handler(ctx, routes.ThemesIndex))
		r.Get("/themes/:theme", app.Handler(ctx, routes.ThemesShow))

		r.Post("/user", app.Handler(ctx, routes.UserCreate))

	}, app.Middleware(ctx, Client))

	router.Group("/user", func(r *app.Router) {
		r.Get("", app.Handler(ctx, routes.UserIndex))
		// m.Put("", app.Handler(ctx, routes.UserCreate))

		r.Get("/devices", app.Handler(ctx, routes.UserDevicesIndex))
		r.Put("/devices/:device", app.Handler(ctx, routes.UserDevicesUpdate))

		r.Get("/journals", app.Handler(ctx, routes.UserJournalsIndex))
		r.Get("/journals/:journal", app.Handler(ctx, routes.UserJournalsShow))
		r.Post("/journals", app.Handler(ctx, routes.UserJournalsCreate))

		r.Get("/notes", app.Handler(ctx, routes.UserNotesIndex))
		r.Get("/notes/:note", app.Handler(ctx, routes.UserNotesShow))
		r.Post("/notes", app.Handler(ctx, routes.UserNotesCreate))

		r.Get("/questionaires", app.Handler(ctx, routes.UserQuestionairesIndex))
		r.Get("/questionaires/:questionaire", app.Handler(ctx, routes.UserQuestionairesShow))

		r.Get("/questions/:question", app.Handler(ctx, routes.UserQuestionsShow))
		r.Get("/questions/:question/keywords", app.Handler(ctx, routes.UserQuestionsKeywordsIndex))
		r.Get("/questions/:question/keywords/:keyword", app.Handler(ctx, routes.UserQuestionsKeywordsShow))

	}, app.Middleware(ctx, Bearer))

	n := negroni.Classic()
	n.Use(app.Middleware(ctx, ContentType))
	n.UseHandler(router)
	n.Run(":8000")
}

func NewContext() *app.Context {
	env := app.DefaultEnvironment()

	pg, err := sqlx.Connect("postgres", env.Postgres)
	if err != nil {
		log.Fatalln(err)
	}

	// red, err := redis.DialTimeout("tcp", env.Redis, time.Duration(10)*time.Second)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	ctx := app.NewContext(env, pg)
	clients, err := ctx.ClientStorage.All()
	if err != nil {
		log.Fatalln(err)
	}
	ctx.ClientCache.Insert(clients)

	// for i := 0; i < len(clients); i++ {
	// 	client := clients[i]
	// 	err := ctx.ClientCache.Insert(&client)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }

	return ctx
}
