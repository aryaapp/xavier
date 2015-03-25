package routes

import (
	"fmt"
	"xavier/api"
)

func ThemesIndex(c *api.Context) *api.Error {
	t, err := c.ThemeStorage.All()
	if err != nil {
		c.LogError(err)
		return &api.Error{404, "Themes could not be found."}
	}
	return c.JSON(200, "themes", t)
}

func ThemesShow(c *api.Context) *api.Error {
	uuid := c.URLParams.ByName("theme")
	t, err := c.ThemeStorage.Find(uuid)
	if err != nil {
		c.LogError(err)
		return &api.Error{404, fmt.Sprintf("Theme could not be found for uuid %s", uuid)}
	}
	return c.JSON(200, "themes", t)
}
