package api

type Config struct {
	Port        int    `envconfig:"port"`
	Environment string `envconfig:"environment"`
	PrivateKey  string `envconfig:"private_key"`
	PostgresURL string `envconfig:"postgres_url"` // PostgreSQL URL
}

// NewConfig initializes a new default config
func NewConfig() *Config {
	c := &Config{
		Port:        3001,
		Environment: "development",
		PrivateKey: "#5K+･ｼミew{ｦ住ｳ(跼Tﾉ(ｩ┫ﾒP.ｿﾓ燾辻G�感ﾃwb=.!r.Oﾀﾍ奎gﾐ｣",
		PostgresURL: "postgres://@localhost/arya_development?sslmode=disable",
	}
	return c
}
