package db

import (
	"database/sql"
	"strconv"
	"strings"

	"bstz.it/rest-api/configuration"
	"bstz.it/rest-api/utils"
	_ "github.com/lib/pq"
)

var DB *sql.DB

var err error

func InitDB() {
	var connectionString = makeConnectionString()
	DB, err = sql.Open("postgres", connectionString)
	utils.PanicOnError(err, "Error connecting to the database")

	err = DB.Ping()
	utils.PanicOnError(err, "Error pinging the database")

	createTables()
}

func makeConnectionString() string {
	connectionString := "postgres://{username}:{password}@{host}:{port}/{dbName}?sslmode={sslMode}"
	connectionString = strings.ReplaceAll(connectionString, "{username}", configuration.Config.Database.Username)
	connectionString = strings.ReplaceAll(connectionString, "{password}", configuration.Config.Database.Password)
	connectionString = strings.ReplaceAll(connectionString, "{host}", configuration.Config.Database.Host)
	connectionString = strings.ReplaceAll(connectionString, "{port}", strconv.Itoa(configuration.Config.Database.Port))
	connectionString = strings.ReplaceAll(connectionString, "{dbName}", configuration.Config.Database.DatabaseName)
	connectionString = strings.ReplaceAll(connectionString, "{sslMode}", configuration.Config.Database.SslMode)

	return connectionString
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL primary key,
		email varchar NOT NULL UNIQUE,
		password TEXT NOT NULL)
	`

	_, err := DB.Exec(createUserTable)
	utils.PanicOnError(err, "Could not create users table")

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		"name" VARCHAR NOT NULL,
		description VARCHAR NOT NULL,
		"location" VARCHAR NOT NULL,
		date_time TIMESTAMPTZ NOT NULL,
		user_id INT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id)
	);
	`
	_, err = DB.Exec(createEventsTable)
	utils.PanicOnError(err, "Could not create events table")
}
