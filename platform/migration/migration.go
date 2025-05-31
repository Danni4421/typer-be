package migration

import (
	"database/sql"
	"typer/app/models"
)

func AutoMigrate(db *sql.DB) {
	migrate(db, "users", models.User{})
	migrate(db, "sessions", models.Session{})
	migrate(db, "languages", models.Language{})
	migrate(db, "words", models.Word{})
}
