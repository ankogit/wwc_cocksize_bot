package stormDB

import (
	"github.com/asdine/storm/v3"
	"local/wwc_cocksize_bot/pkg/models"
)

type ChatRepository struct {
	db *storm.DB
}

func NewChatRepository(db *storm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) Save(data models.Chat) error {
	return r.db.Save(&data)
}

func (r *ChatRepository) Get(chatId int64) (models.Chat, error) {
	var chat models.Chat
	err := r.db.One("ID", chatId, &chat)
	if err != nil {
		return chat, err
	}
	return chat, nil
}
func (r *ChatRepository) Delete(chat models.Chat) error {
	return r.db.DeleteStruct(&chat)
}

func (r *ChatRepository) All() ([]models.Chat, error) {
	var chats []models.Chat
	err := r.db.All(&chats)
	if err != nil {
		return chats, err
	}
	return chats, nil
}
