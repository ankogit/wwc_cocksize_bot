package stormDB

import (
	"github.com/asdine/storm/v3"
	"local/wwc_cocksize_bot/pkg/models"
)

type UserRepository struct {
	db *storm.DB
}

func NewUserRepository(db *storm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(data models.UserData) error {
	return r.db.Save(&data)
}

func (r *UserRepository) Get(userId int64) (models.UserData, error) {
	var user models.UserData
	err := r.db.One("ID", userId, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) All() ([]models.UserData, error) {
	var users []models.UserData
	err := r.db.All(&users)
	if err != nil {
		return users, err
	}
	return users, nil
}
