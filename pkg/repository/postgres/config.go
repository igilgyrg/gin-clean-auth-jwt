package postgres

type postgresConfig struct {
	host        string
	port        string
	database    string
	username    string
	password    string
	maxAttempts int
}

func NewPostgresConfig(maxAttempts int, host, port, database, username, password string) *postgresConfig {
	return &postgresConfig{
		host:        host,
		port:        port,
		username:    username,
		password:    password,
		database:    database,
		maxAttempts: maxAttempts,
	}
}
