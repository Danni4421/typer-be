package models

import "time"

type Session struct {
	ID        int       `json:"id" db:"id" sql:"SERIAL PRIMARY KEY"`
	UserID    int       `json:"user_id" db:"user_id" sql:"INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE"`
	Token     string    `json:"token" db:"token" sql:"TEXT NOT NULL"`
	CreatedAt time.Time `json:"created_at" db:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" sql:"TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at" sql:"TIMESTAMP"`
}
