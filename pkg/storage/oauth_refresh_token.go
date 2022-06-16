package storage

import (
	"local/wwc_cocksize_bot/pkg/models"
	"time"
)

type RefreshTokenRepository interface {
	Create(userId int64, refreshToken string, expiresIn time.Time) (models.RefreshToken, error)
	Find(refreshToken string) (models.RefreshToken, error)
}
