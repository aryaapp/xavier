package app

import (
	"encoding/json"
	"log"
	"net/http"
	"xavier/storage"
	p "xavier/storage/postgres"
	r "xavier/storage/redis"
	"xavier/validator"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/nbio/httpcontext"
)

type Context struct {
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

	Environment *Environment
	Request     *http.Request
	Scopes      []string
	URLParams   httprouter.Params
	Writer      http.ResponseWriter
	Validator   *validator.Validator
}

type Error struct {
	Code    int
	Message string
}

type M map[string]interface{}

func NewContext(env *Environment, pg *sqlx.DB) *Context {
	validator := validator.New()
	return &Context{
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
		Environment:         env,
		Validator:           validator,
	}
}

func ChildContext(parent *Context, r *http.Request, w http.ResponseWriter) *Context {
	ctx := &Context{
		AppCache:            parent.AppCache,
		AppStorage:          parent.AppStorage,
		DeviceStorage:       parent.DeviceStorage,
		JournalStorage:      parent.JournalStorage,
		KeywordStorage:      parent.KeywordStorage,
		NoteStorage:         parent.NoteStorage,
		QuestionStorage:     parent.QuestionStorage,
		QuestionaireStorage: parent.QuestionaireStorage,
		ThemeStorage:        parent.ThemeStorage,
		UserStorage:         parent.UserStorage,
		Environment:         parent.Environment,
		Validator:           parent.Validator,
		Request:             r,
		URLParams:           Params(r),
		Writer:              w,
	}
	return ctx
}

func (c *Context) Production() bool {
	return c.Environment.Name == "production"
}

/////////////////////////////////////////////////////////////////////////////
// Context Env
/////////////////////////////////////////////////////////////////////////////

func (c *Context) GetUserID() int {
	if userID, ok := httpcontext.Get(c.Request, "userID").(int); ok {
		return userID
	}
	return 1
}

func (c *Context) SetUserID(userID int) {
	httpcontext.Set(c.Request, "userID", userID)
}

func (c *Context) GetAppForCurrentRequest() *storage.App {
	app, ok := httpcontext.Get(c.Request, "app").(*storage.App)
	if !ok {
		return nil
	}
	return app
}

func (c *Context) SetAppForCurrentRequest(app *storage.App) {
	httpcontext.Set(c.Request, "app", app)
}

/////////////////////////////////////////////////////////////////////////////
// Render
/////////////////////////////////////////////////////////////////////////////

func (c *Context) JSON(code int, key string, obj interface{}) *Error {
	writeHeader(c.Writer, code, "application/json")
	if err := json.NewEncoder(c.Writer).Encode(M{key: obj}); err != nil {
		// TODO
		log.Fatalln(err)

		return &Error{500, "Internal JSON error."}
	}
	return nil
}

func (c *Context) JSONError(code int, title string, description string) {
	writeHeader(c.Writer, code, "application/json")
	if err := json.NewEncoder(c.Writer).Encode(M{"error": M{"code": code, "title": title, "description": description}}); err != nil {
		log.Fatal(err)
	}
}

func writeHeader(w http.ResponseWriter, code int, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)
}

/////////////////////////////////////////////////////////////////////////////
// Log
/////////////////////////////////////////////////////////////////////////////

func (c *Context) LogError(err error) {
	if c.Production() == false {
		log.Print("Error: " + err.Error())
	}
}

/////////////////////////////////////////////////////////////////////////////
// Parameter binding
/////////////////////////////////////////////////////////////////////////////

func (c *Context) BindParams(obj interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}

func (c *Context) BindParamsAndValidate(obj interface{}) error {
	if err := c.BindParams(obj); err != nil {
		return err
	}
	return c.Validator.Validate(obj)
}
