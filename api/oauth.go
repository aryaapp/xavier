package api

import (
	"net/http"

	"github.com/aryaapp/xavier/jwt"
	"github.com/aryaapp/xavier/storage"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func (a *AppContext) NewAccessToken(c *echo.Context) error {
	var input storage.NewTokenInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(422, Error{"Invalid input given."})
	}

	if input.GrantType == "password" {
		user, err := a.UserStorage.FindByEmail(input.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{"Database error"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			return c.JSON(http.StatusBadRequest, Error{"Could not be authenticate. Invalid email/password combination"})
		}

		token := jwt.NewAccessToken(user.ID)
		signedToken, err := a.JWTClient.Encode(token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Error{"Could not encode token. Internal error"})
		}
		return c.JSON(http.StatusCreated, Data{signedToken})
	}
	return c.JSON(422, Error{"Grant Type is not supported."})
}
