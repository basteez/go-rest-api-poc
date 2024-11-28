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
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL primary key,
		"name" varchar NOT NULL,
		description varchar NOT NULL,
		"location" varchar NOT NULL,
		date_time timestamptz NOT NULL,
		user_id int NULL
	);
	`
	_, err := DB.Exec(createEventsTable)
	utils.PanicOnError(err, "Could not create events table")
}
