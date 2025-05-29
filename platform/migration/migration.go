package migration

import (
	"database/sql"
	"typer/app/models"
)

func AutoMigrate(db *sql.DB) {
	migrate(db, "users", models.User{})
}
