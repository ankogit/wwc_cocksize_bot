package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"time"
)

type TokenManager interface {
	NewAccessToken(userId string, ttl time.Duration) (string, error)
	NewRefreshToken() (string, error)
	ParseAccessToken(accessToken string) (string, error)
}

type Manager struct {
	secretKey string
}

func NewManager(secretKey string) (*Manager, error) {
	if secretKey == "" {
		return nil, errors.New("empty secretKey")
	}
	return &Manager{
		secretKey: secretKey,
	}, nil
}

func (m *Manager) NewAccessToken(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience: "",
		Id:       "123",
		//IssuedAt:  0,
		//Issuer:    "",
		//NotBefore: 0,

		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(m.secretKey))
}

func (m *Manager) NewRefreshToken() (string, error) {
	bytes := make([]byte, 32)

	randSource := rand.NewSource(time.Now().Unix())
	randNew := rand.New(randSource)

	if _, err := randNew.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}

func (m *Manager) ParseAccessToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secretKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
