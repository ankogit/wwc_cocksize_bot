package service

import (
	"context"
	"local/wwc_cocksize_bot/pkg/auth"
	"local/wwc_cocksize_bot/pkg/storage"
)

type Users interface {
	Login(ctx context.Context, input LoginInput) (Tokens, error)
}

type LoginInput struct {
	userId int64
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Services struct {
	Users Users
}
type Deps struct {
	UserRepository storage.UserRepository
	TokenManager   auth.TokenManager
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.UserRepository, deps.TokenManager)

	return &Services{
		Users: usersService,
	}
}
