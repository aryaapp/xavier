package routes

import "xavier/api"

func RootIndex(c *api.Context) *api.Error {
	return c.JSON(200, "hello", "world")
}
