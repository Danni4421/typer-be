package models

import "time"

type TypingLog struct {
	ID             uint      `json:"id" db:"id" sql:"SERIAL PRIMARY KEY"`
	UserID         uint      `json:"user_id" db:"user_id" sql:"INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE"`
	LanguageID     uint      `json:"language_id" db:"language_id" sql:"INTEGER NOT NULL REFERENCES languages(id) ON DELETE CASCADE"`
	CharacterCount int16     `json:"character_count" db:"character_count" sql:"INTEGER NOT NULL CHECK (character_count >= 0)"`
	WordCount      int16     `json:"word_count" db:"word_count" sql:"INTEGER NOT NULL CHECK (word_count >= 0)"`
	Accuracy       float32   `json:"accuracy" db:"accuracy" sql:"REAL NOT NULL CHECK (accuracy >= 0 AND accuracy <= 100)"`
	WPM            int16     `json:"wpm" db:"wpm" sql:"REAL NOT NULL CHECK (wpm >= 0)"`
	CreatedAt      time.Time `json:"created_at" db:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at" sql:"TIMESTAMP"`
	DeletedAt      time.Time `json:"deleted_at" db:"deleted_at" sql:"TIMESTAMP"`
}
