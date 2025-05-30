package database

import (
	"database/sql"
	"fmt"
	"log"
	"typer/package/utils"
	"typer/platform/migration"

	_ "github.com/lib/pq"
)

var dsn string
var DB *sql.DB

func init() {
	DatabaseHost := utils.GetEnv("DATABASE_HOST", "localhost")
	DatabasePort := utils.GetEnv("DB_PORT", "5432")
	DatabaseUsername := utils.GetEnv("DB_USERNAME", "postgres")
	DatabasePassword := utils.GetEnv("DB_PASSWORD", "")
	DatabaseName := utils.GetEnv("DB_DATABASE", "postgres")

	dsn = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		DatabaseUsername,
		DatabasePassword,
		DatabaseHost,
		DatabasePort,
		DatabaseName,
	)
}

func ConnectPostgres() {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	DB = db

	migration.AutoMigrate(db)
}
