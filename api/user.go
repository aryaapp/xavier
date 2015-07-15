package api

import (
	"net/http"

	"github.com/aryaapp/xavier/storage"
	"github.com/labstack/echo"
)

func (a *AppContext) FindUserByID(c *echo.Context) error {
	userID := c.Get("user.id").(int)
	u, err := a.UserStorage.FindByID(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"User could not be found."})
	}

	u.Journals, _ = a.JournalStorage.FindForLastWeek(u.ID)
	u.Questionaires, _ = a.QuestionaireStorage.FindAll(u.ID)
	// u.Theme, _ = a.ThemeStorage.FindByID(u.ThemeID)

	return c.JSON(http.StatusOK, Data{u})
}

func (a *AppContext) NewUser(c *echo.Context) error {
	var signup storage.UserSignup
	if err := c.Bind(&signup); err != nil {
		return c.JSON(422, Error{"Invalid input given."})
	}

	u, err := a.UserStorage.New(&signup, 1)
	switch {
	case err == storage.UserConflictError:
		return c.JSON(http.StatusConflict, Error{"User could not be created. Already exists"})
	case err != nil:
		return c.JSON(http.StatusBadRequest, Error{"User could not be created. Internal error"})
	default:
		return c.JSON(http.StatusCreated, Data{u})
	}
}
