package repositories

import "local/wwc_cocksize_bot/pkg/models"

type UserRepository interface {
	Save(data models.UserData) error
	Get(userId int64) (models.UserData, error)
}
