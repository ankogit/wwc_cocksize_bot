package models

import "time"

type RefreshToken struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	ExpiresIn    time.Time `db:"expires_in"`
	CreatedAt    time.Time `db:"created_at"`
	Revoked      bool      `db:"revoked"`
}
