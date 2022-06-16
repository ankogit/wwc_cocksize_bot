package storage

import (
	"github.com/asdine/storm/v3"
	"github.com/jmoiron/sqlx"
	"local/wwc_cocksize_bot/pkg/storage/postgresDB"
	"local/wwc_cocksize_bot/pkg/storage/stormDB"
)

type Repositories struct {
	Users         UserRepository
	Chats         ChatRepository
	RefreshTokens RefreshTokenRepository
}

func NewRepositories(db *storm.DB, dbPostgres *sqlx.DB) *Repositories {
	return &Repositories{
		Users:         stormDB.NewUserRepository(db),
		Chats:         stormDB.NewChatRepository(db),
		RefreshTokens: postgresDB.NewRefreshTokenRepository(dbPostgres),
	}
}
