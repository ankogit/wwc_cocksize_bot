package service

import (
	"context"
	"local/wwc_cocksize_bot/pkg/auth"
	"local/wwc_cocksize_bot/pkg/models"
	"local/wwc_cocksize_bot/pkg/storage"
	"log"
	"strconv"
	"time"
)

type UserService struct {
	UserRepository         storage.UserRepository
	RefreshTokenRepository storage.RefreshTokenRepository
	TokenManager           auth.TokenManager
}

func NewUsersService(userRep storage.UserRepository, refreshRep storage.RefreshTokenRepository, m auth.TokenManager) *UserService {
	return &UserService{
		UserRepository:         userRep,
		RefreshTokenRepository: refreshRep,
		TokenManager:           m,
	}
}

func (s *UserService) Login(ctx context.Context, input LoginInput) (models.Session, error) {
	user, err := s.UserRepository.Get(input.UserId)
	if err != nil {
		return models.Session{}, err
	}

	tokens, err := s.createTokens(ctx, user)
	if err != nil {
		return models.Session{}, err
	}

	return tokens, nil
}
func (s *UserService) RefreshToken(ctx context.Context, input RefreshInput) (models.Session, error) {
	refreshToken, err := s.RefreshTokenRepository.Find(input.Token)
	if err != nil {
		return models.Session{}, err
	}

	user, err := s.UserRepository.Get(refreshToken.UserID)
	if err != nil {
		return models.Session{}, err
	}

	tokens, err := s.createTokens(ctx, user)
	if err != nil {
		return models.Session{}, err
	}

	return tokens, nil
}

func (s *UserService) createTokens(ctx context.Context, user models.UserData) (token models.Session, err error) {
	accessToken, err := s.TokenManager.NewAccessToken(strconv.FormatInt(user.ID, 10), 15*time.Minute)
	refreshToken, err := s.TokenManager.NewRefreshToken()

	data, err := s.RefreshTokenRepository.Create(user.ID, refreshToken, time.Now().AddDate(0, 1, 0))
	log.Println(data)

	if err != nil {
		return models.Session{}, err
	}
	return models.Session{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
