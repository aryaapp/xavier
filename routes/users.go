package routes

import (
	"xavier/app"
	"xavier/storage"
)

type UserParams struct {
	Email                string   `json:"email" validate:"nonzero,email"`
	Password             string   `json:"password" validate:"nonzero,password"`
	PasswordConfirmation string   `json:"password_confirmation" validate:"nonzero"`
	Scopes               []string `json:"scopes"`
}

func UserIndex(c *app.Context) *app.Error {
	u, err := c.UserStorage.FindByID(c.GetUserID())
	if err != nil {
		c.LogError(err)
		return &app.Error{404, "User could not be found."}
	}

	j, err := c.JournalStorage.LastWeek(u.ID)
	if err != nil {
		c.LogError(err)
	}
	q, err := c.QuestionaireStorage.All(u.ID)
	if err != nil {
		c.LogError(err)
	}
	t, err := c.ThemeStorage.FindByID(u.ThemeID)
	if err != nil {
		c.LogError(err)
	}

	u.Journals = j
	u.Questionaires = q
	u.Theme = t

	return c.JSON(200, "user", u)
}

func UserCreate(c *app.Context) *app.Error {
	var params UserParams
	if err := c.BindParamsAndValidate(&params); err != nil {
		c.LogError(err)
		return &app.Error{422, "User could not be created. Invalid parameters:" + err.Error()}
	} else if params.Password != params.PasswordConfirmation {
		return &app.Error{422, "User could not be created. Invalid password confirmation"}
	}

	a := c.GetAppForCurrentRequest()
	user, err := c.UserStorage.Insert(&storage.UserRegistration{
		Email:    params.Email,
		Password: params.Password,
		AppID:    a.ID,
	})
	switch {
	case err == storage.UserConflictError:
		return &app.Error{409, "User could not be created. Already exists."}
	case err != nil:
		c.LogError(err)
		return &app.Error{500, "User could not be created. Internal error" + err.Error()}
	default:
		return c.JSON(201, "user", user)
	}
}
