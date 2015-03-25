package api

import "os"

type Environment struct {
	Name     string
	Secret   string
	Postgres string
	Redis    string
}

func DefaultEnvironment() *Environment {
	name := os.Getenv("ARYA_ENV")
	if len(name) == 0 {
		name = "development"
	}

	secret := os.Getenv("ARYA_SECRET")
	if len(secret) == 0 {
		secret = "#5K+･ｼミew{ｦ住ｳ(跼Tﾉ(ｩ┫ﾒP.ｿﾓ燾辻G�感ﾃwb=.!r.Oﾀﾍ奎gﾐ｣"
	}

	postgres := os.Getenv("ARYA_POSTGRES_URL")
	if len(postgres) == 0 {
		postgres = "postgres://@localhost/arya_development?sslmode=disable"
	}

	redis := os.Getenv("ARYA_REDIS_URL")
	if len(redis) == 0 {
		redis = "127.0.0.1:6379"
	}

	return &Environment{name, secret, postgres, redis}
}
