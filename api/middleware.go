package api

import (
	"net/http"

	"github.com/aryaapp/xavier/jwt"
	"github.com/labstack/echo"
)

func (a *AppContext) Bearer() echo.HandlerFunc {
	return func(c *echo.Context) error {
		accessToken := &jwt.AccessToken{}
		if err := a.JWTClient.DecodeFromRequest(accessToken, c.Request()); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if accessToken.UserID > 0 {
			c.Set("user.id", accessToken.UserID)
			return nil
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "token doesn't contain valid user credentials")
	}
}
