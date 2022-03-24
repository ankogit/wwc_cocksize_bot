package service

import (
	"context"
	"local/wwc_cocksize_bot/pkg/auth"
	"local/wwc_cocksize_bot/pkg/models"
	"local/wwc_cocksize_bot/pkg/storage"
	"strconv"
	"time"
)

type UserService struct {
	UserRepository storage.UserRepository
	TokenManager   auth.TokenManager
}

func NewUsersService(r storage.UserRepository, m auth.TokenManager) *UserService {
	return &UserService{
		UserRepository: r,
		TokenManager:   m,
	}
}

func (s *UserService) Login(ctx context.Context, input LoginInput) (Tokens, error) {
	user, err := s.UserRepository.Get(input.userId)
	if err != nil {
		return Tokens{}, err
	}

	tokens, err := s.createTokens(ctx, user)
	if err != nil {
		return Tokens{}, err
	}

	return tokens, nil
}

func (s *UserService) createTokens(ctx context.Context, user models.UserData) (token Tokens, err error) {
	accessToken, err := s.TokenManager.NewAccessToken(strconv.FormatInt(user.ID, 10), 15*time.Minute)
	refreshToken, err := s.TokenManager.NewRefreshToken()

	if err != nil {
		return Tokens{}, err
	}
	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
