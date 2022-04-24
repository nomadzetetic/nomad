package db

import "time"

type AccountEntity struct {
	ID                 string    `db:"id"`
	Nickname           string    `db:"nickname"`
	Email              string    `db:"email"`
	AvatarUrl          *string   `db:"avatar_url"`
	CreatedAt          time.Time `db:"created_at"`
	ActivationToken    *string   `db:"activation_token"`
	Enabled            bool      `db:"enabled"`
	ObfuscatedPassword string    `db:"obfuscated_password"`
}
