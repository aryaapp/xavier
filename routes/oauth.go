package routes

import (
	"time"
	"xavier/api"
	"xavier/lib/oauth"
	"xavier/lib/token"

	"golang.org/x/crypto/bcrypt"
)

type OAuthUserParams struct {
	Email     string   `json:"email" validate:"nonzero,email"`
	Password  string   `json:"password" validate:"nonzero"`
	GrantType string   `json:"grant_type" validate:"nonzero,grant_type"`
	Scopes    []string `json:"scopes"`
}

func OAuthTokensCreate(c *api.Context) *api.Error {
	var params OAuthUserParams
	if err := c.BindParamsAndValidate(&params); err != nil {
		c.LogError(err)
		return &api.Error{422, "Could not be authenticate. Invalid parameters, " + err.Error()}
	}

	if params.GrantType == oauth.Password {
		a := c.GetAppForCurrentRequest()
		user, err := c.UserStorage.FindByEmail(params.Email)
		if err != nil {
			c.LogError(err)
			return &api.Error{401, "Could not be authenticate. Invalid email/password combination"}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
			c.LogError(err)
			return &api.Error{401, "Could not be authenticate. Invalid email/password combination"}
		}

		expiration := time.Hour * 1000000
		if c.Production() {
			expiration = time.Hour * 3
		}

		token := token.New(a.UUID, expiration, user.ID, params.Scopes)
		encoded, err := token.Sign(c.Environment.Secret)
		if err != nil {
			c.LogError(err)
			return &api.Error{500, "Could not create token. Internal error"}
		}
		return c.JSON(201, "token", encoded)
	}
	return &api.Error{422, "Could not be authenticate. Grant Type is not supported"}
}
