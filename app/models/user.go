package models

import "time"

type User struct {
	ID        int       `db:"id" sql:"SERIAL PRIMARY KEY"`
	Username  string    `db:"username" sql:"VARCHAR(50) NOT NULL"`
	Name      string    `db:"name" sql:"VARCHAR(120) NOT NULL"`
	Email     string    `db:"email" sql:"VARCHAR(50) NOT NULL"`
	Password  string    `db:"password" sql:"TEXT NOT NULL"`
	CreatedAt time.Time `db:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `db:"updated_at" sql:"TIMESTAMP"`
	DeletedAt time.Time `db:"deleted_at" sql:"TIMESTAMP"`
}
