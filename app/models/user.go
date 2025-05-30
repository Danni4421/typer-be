package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id" sql:"SERIAL PRIMARY KEY"`
	Username  string    `json:"username" db:"username" sql:"VARCHAR(50) NOT NULL"`
	Name      string    `json:"name" db:"name" sql:"VARCHAR(120) NOT NULL"`
	Email     string    `json:"email" db:"email" sql:"VARCHAR(50) NOT NULL"`
	Password  string    `json:"-" db:"password" sql:"TEXT NOT NULL"`
	CreatedAt time.Time `json:"created_at" db:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" sql:"TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at" sql:"TIMESTAMP"`
}
