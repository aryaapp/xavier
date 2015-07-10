package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (a *AppContext) FindAllThemes(c *echo.Context) error {
	t, err := a.ThemeStorage.FindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Themes could not be found."})
	}
	return c.JSON(http.StatusOK, Data{t})
}

func (a *AppContext) FindThemeByUUID(c *echo.Context) error {
	uuid := c.Param("uuid")
	t, err := a.ThemeStorage.FindByUUID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{fmt.Sprintf("Theme could not be found for id %s", uuid)})
	}
	return c.JSON(http.StatusOK, Data{t})
}
