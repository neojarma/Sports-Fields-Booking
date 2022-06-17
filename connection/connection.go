package connection

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func LoadEnv(key string) string {
	return os.Getenv(key)
}

func GetConnection() (*sql.DB, error) {

	// for non heroku
	// dbHost := LoadEnv("DB_HOST")
	// dbUsername := LoadEnv("DB_USERNAME")
	// dbPassword := LoadEnv("DB_PASSWORD")
	// dbName := LoadEnv("DB_NAME")

	// dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUsername, dbPassword, dbHost, dbName)

	// for heroku
	dataSource := LoadEnv("DATABASE_URL")

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
