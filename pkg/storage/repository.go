package storage

import (
	"github.com/asdine/storm/v3"
	"local/wwc_cocksize_bot/pkg/storage/stormDB"
)

type Repositories struct {
	Users UserRepository
	Chats ChatRepository
}

func NewRepositories(db *storm.DB) *Repositories {
	return &Repositories{
		Users: stormDB.NewUserRepository(db),
		Chats: stormDB.NewChatRepository(db),
	}
}
