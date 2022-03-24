package storage

import "local/wwc_cocksize_bot/pkg/models"

type ChatRepository interface {
	Save(data models.Chat) error
	Get(chatId int64) (models.Chat, error)
	Delete(chat models.Chat) error
	All() ([]models.Chat, error)
}
