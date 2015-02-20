package oauth

const (
	AuthorizationCode = "authorization_code"
	Implicit          = "implicit"
	Password          = "password"
	RefreshToken      = "refresh_token"
)

func ValidGrantType(s string) bool {
	return s == AuthorizationCode || s == Implicit || s == Password || s == RefreshToken
}
