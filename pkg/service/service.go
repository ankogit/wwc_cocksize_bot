package service

import (
	"context"
	"local/wwc_cocksize_bot/pkg/auth"
	"local/wwc_cocksize_bot/pkg/models"
	"local/wwc_cocksize_bot/pkg/storage"
)

type Users interface {
	Login(ctx context.Context, input LoginInput) (models.Session, error)
	RefreshToken(ctx context.Context, input RefreshInput) (models.Session, error)
}

type LoginInput struct {
	UserId int64
}

type RefreshInput struct {
	Token string
}

type Services struct {
	Users        Users
	TokenManager auth.TokenManager
	Repositories *storage.Repositories
}
type Deps struct {
	Repositories *storage.Repositories
	TokenManager auth.TokenManager
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repositories.Users, deps.Repositories.RefreshTokens, deps.TokenManager)

	return &Services{
		Users:        usersService,
		TokenManager: deps.TokenManager,
		Repositories: deps.Repositories,
	}
}
