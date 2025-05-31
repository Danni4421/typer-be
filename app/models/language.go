package models

import "time"

type Language struct {
	ID        uint      `json:"id" db:"id" sql:"SERIAL PRIMARY KEY"`
	Code      string    `json:"code" db:"code" sql:"VARCHAR(4) NOT NULL UNIQUE"`
	Name      string    `json:"name" db:"name" sql:"VARCHAR(50) NOT NULL UNIQUE"`
	CreatedAt time.Time `json:"-" db:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"-" db:"updated_at" sql:"TIMESTAMP"`
	DeletedAt time.Time `json:"-" db:"deleted_at" sql:"TIMESTAMP"`
}
