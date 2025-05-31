package models

import "time"

type Word struct {
	ID         int       `json:"id" db:"id" sql:"SERIAL PRIMARY KEY"`
	LanguageID int       `json:"language_id" db:"language_id" sql:"INTEGER NOT NULL REFERENCES languages(id) ON DELETE CASCADE"`
	Word       string    `json:"word" db:"word" sql:"VARCHAR(50) NOT NULL"`
	Length     int       `json:"length" db:"length" sql:"INTEGER NOT NULL CHECK (length > 0)"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at" sql:"TIMESTAMP"`
	DeletedAt  time.Time `json:"deleted_at" db:"deleted_at" sql:"TIMESTAMP"`
}
