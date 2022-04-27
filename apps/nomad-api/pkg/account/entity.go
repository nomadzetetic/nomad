package account

import "time"

type Entity struct {
	ID                 string    `db:"id"`
	Nickname           string    `db:"nickname"`
	Email              string    `db:"email"`
	AvatarUrl          *string   `db:"avatar_url"`
	CreatedAt          time.Time `db:"created_at"`
	ActivationToken    *string   `db:"activation_token"`
	Banned             bool      `db:"banned"`
	ObfuscatedPassword string    `db:"obfuscated_password"`
	Roles              []string  `db:"roles"`
}
