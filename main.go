package main

import (
	"log"
	"xavier/api"
	"xavier/routes"

	"github.com/codegangsta/negroni"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	ctx := NewContext()
	router := api.NewRouter()
	router.Get("/", api.Handler(ctx, routes.RootIndex))

	router.Group("/", func(r *api.Router) {
		r.Post("/oauth/tokens", api.Handler(ctx, routes.OAuthTokensCreate))

		r.Get("/themes", api.Handler(ctx, routes.ThemesIndex))
		r.Get("/themes/:theme", api.Handler(ctx, routes.ThemesShow))

		r.Post("/user", api.Handler(ctx, routes.UserCreate))

	}, api.Middleware(ctx, CurrentApp))

	router.Group("/user", func(r *api.Router) {
		r.Get("", api.Handler(ctx, routes.UserIndex))
		// m.Put("", api.Handler(ctx, routes.UserCreate))

		r.Get("/devices", api.Handler(ctx, routes.UserDevicesIndex))
		r.Put("/devices/:device", api.Handler(ctx, routes.UserDevicesUpdate))

		r.Get("/journals", api.Handler(ctx, routes.UserJournalsIndex))
		r.Get("/journals/:journal", api.Handler(ctx, routes.UserJournalsShow))
		r.Post("/journals", api.Handler(ctx, routes.UserJournalsCreate))

		r.Get("/notes", api.Handler(ctx, routes.UserNotesIndex))
		r.Get("/notes/:note", api.Handler(ctx, routes.UserNotesShow))
		r.Post("/notes", api.Handler(ctx, routes.UserNotesCreate))

		r.Get("/questionaires", api.Handler(ctx, routes.UserQuestionairesIndex))
		r.Get("/questionaires/:questionaire", api.Handler(ctx, routes.UserQuestionairesShow))

		r.Get("/questions/:question", api.Handler(ctx, routes.UserQuestionsShow))
		r.Get("/questions/:question/keywords", api.Handler(ctx, routes.UserQuestionsKeywordsIndex))
		r.Get("/questions/:question/keywords/:keyword", api.Handler(ctx, routes.UserQuestionsKeywordsShow))

	}, api.Middleware(ctx, Bearer))

	n := negroni.Classic()
	n.Use(api.Middleware(ctx, ContentType))
	n.UseHandler(router)
	n.Run(":8000")
}

func NewContext() *api.Context {
	env := api.DefaultEnvironment()

	pg, err := sqlx.Connect("postgres", env.Postgres)
	if err != nil {
		log.Fatalln(err)
	}

	// red, err := redis.DialTimeout("tcp", env.Redis, time.Duration(10)*time.Second)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	ctx := api.NewContext(env, pg)

	apps, err := ctx.AppStorage.All()
	if err != nil {
		log.Fatalln(err)
	}
	ctx.AppCache.Insert(apps)

	// for i := 0; i < len(clients); i++ {
	// 	client := clients[i]
	// 	err := ctx.ClientCache.Insert(&client)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }

	return ctx
}
