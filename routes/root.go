package routes

import "xavier/app"

func RootIndex(c *app.Context) *app.Error {
	return c.JSON(200, "hello", "world")
}
