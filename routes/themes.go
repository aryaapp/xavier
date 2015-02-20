package routes

import (
	"fmt"
	"xavier/app"
)

func ThemesIndex(c *app.Context) *app.Error {
	t, err := c.ThemeStorage.All()
	if err != nil {
		c.LogError(err)
		return &app.Error{404, "Themes could not be found."}
	}
	return c.JSON(200, "themes", t)
}

func ThemesShow(c *app.Context) *app.Error {
	uuid := c.URLParams.ByName("theme")
	t, err := c.ThemeStorage.Find(uuid)
	if err != nil {
		c.LogError(err)
		return &app.Error{404, fmt.Sprintf("Theme could not be found for uuid %s", uuid)}
	}
	return c.JSON(200, "themes", t)
}
