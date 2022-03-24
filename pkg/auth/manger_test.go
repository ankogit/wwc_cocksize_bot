package auth

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestManager_NewAccessToken(t *testing.T) {
	userId := "12313"
	secretKey := "test"
	manager, _ := NewManager(secretKey)

	token, err := manager.NewAccessToken(userId, time.Minute*15)
	if err != nil {
		errors.New(err.Error())
	}

	fmt.Println(token)
}

func TestManager_NewRefreshToken(t *testing.T) {
	secretKey := "test"
	manager, _ := NewManager(secretKey)

	refreshToken, err := manager.NewRefreshToken()
	if err != nil {
		errors.New(err.Error())
	}
	fmt.Println(refreshToken)
}
func TestManager_ParseAccessToken(t *testing.T) {
	secretKey := "test"
	manager, _ := NewManager(secretKey)
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDgwMzQ3MzEsImp0aSI6IjEyMyIsInN1YiI6IjEyMzEzIn0.Ly8PZ76FZ8x9ElVK0u5SKXUKBWAndVPTmDXZlbvVKCI"
	userId, err := manager.ParseAccessToken(accessToken)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userId)
}
