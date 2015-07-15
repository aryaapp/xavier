package api

import (
	"encoding/json"

	"github.com/aryaapp/xavier/jwt"
	"github.com/aryaapp/xavier/storage"
	"github.com/labstack/echo"
)

type AppContext struct {
	Config              *Config
	JWTClient           *jwt.Client
	AppCache            storage.AppCache
	AppStorage          storage.AppStorage
	DeviceStorage       storage.DeviceStorage
	JournalStorage      storage.JournalStorage
	KeywordStorage      storage.KeywordStorage
	NoteStorage         storage.NoteStorage
	QuestionStorage     storage.QuestionStorage
	QuestionaireStorage storage.QuestionaireStorage
	ThemeStorage        storage.ThemeStorage
	UserStorage         storage.UserStorage
}

// Wrapper to properly marschal your JSON output
type Data struct {
	Object interface{}
}

func (d Data) MarshalJSON() ([]byte, error) {
	output := map[string]interface{}{
		"data": d.Object,
	}
	return json.Marshal(output)
}

// Wrapper to properly marschal your JSON errors
type Error struct {
	Message string
}

func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"error": e.Message,
	})
}

// Mux has all the routes needed for Xaxiver
func Mux(e *echo.Echo, a *AppContext) {
	// OAuth
	e.Post("/oauth/token", a.NewAccessToken)

	// Themes
	e.Get("/themes", a.FindAllThemes)
	e.Get("/themes/:uuid", a.FindThemeByUUID)

	// User
	e.Post("/user", a.NewUser)

	u := e.Group("/user")
	u.Get("", a.FindUserByID)
	u.Use(a.Bearer())
	{
		// Devices
		u.Get("/devices", a.FindAllDevices)

		// Journals
		u.Get("/journals", a.FindAllJournals)
		u.Get("/journals/:uuid", a.FindJournalByUUID)

		// Notes
		u.Get("/notes", a.FindAllNotes)
		u.Get("/notes/:uuid", a.FindNoteByUUID)
		u.Post("/notes", a.NewNote)

		u.Get("/questionaires", a.FindAllQuestionaires)
		u.Get("/questionaires/:uuid", a.FindQuestionaireByUUID)

		u.Get("/questions/:uuid", a.FindQuestionByUUID)
	}

	// u.Put("/devices/:token", a.)

	// ctx := NewContext()
	// router := api.NewRouter()
	// router.Get("/", api.Handler(ctx, routes.RootIndex))

	// router.Group("/", func(r *api.Router) {
	// 	r.Post("/oauth/tokens", api.Handler(ctx, routes.OAuthTokensCreate))

	// 	r.Get("/themes", api.Handler(ctx, routes.ThemesIndex))
	// 	r.Get("/themes/:theme", api.Handler(ctx, routes.ThemesShow))

	// 	r.Post("/user", api.Handler(ctx, routes.UserCreate))

	// }, api.Middleware(ctx, CurrentApp))

	// router.Group("/user", func(r *api.Router) {
	// r.Get("", api.Handler(ctx, routes.UserIndex))
	// // m.Put("", api.Handler(ctx, routes.UserCreate))

	// 	r.Get("/devices", api.Handler(ctx, routes.UserDevicesIndex))
	// 	r.Put("/devices/:device", api.Handler(ctx, routes.UserDevicesUpdate))

	// 	r.Get("/journals", api.Handler(ctx, routes.UserJournalsIndex))
	// 	r.Get("/journals/:journal", api.Handler(ctx, routes.UserJournalsShow))
	// 	r.Post("/journals", api.Handler(ctx, routes.UserJournalsCreate))

	// 	r.Get("/notes", api.Handler(ctx, routes.UserNotesIndex))
	// 	r.Get("/notes/:note", api.Handler(ctx, routes.UserNotesShow))
	// 	r.Post("/notes", api.Handler(ctx, routes.UserNotesCreate))

	// 	r.Get("/questionaires", api.Handler(ctx, routes.UserQuestionairesIndex))
	// 	r.Get("/questionaires/:questionaire", api.Handler(ctx, routes.UserQuestionairesShow))

	// 	r.Get("/questions/:question", api.Handler(ctx, routes.UserQuestionsShow))
	// 	r.Get("/questions/:question/keywords", api.Handler(ctx, routes.UserQuestionsKeywordsIndex))
	// 	r.Get("/questions/:question/keywords/:keyword", api.Handler(ctx, routes.UserQuestionsKeywordsShow))

	// }, api.Middleware(ctx, Bearer))

	// n := negroni.Classic()
	// n.Use(api.Middleware(ctx, ContentType))
	// n.UseHandler(router)
	// n.Run(":8000")
}
