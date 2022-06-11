package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"local/wwc_cocksize_bot/pkg/auth"
	"local/wwc_cocksize_bot/pkg/server/handler/response"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"

	userCtx = "userId"
)

func AuthUser(tokenManager auth.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseAuthHeader(c, tokenManager)
		if err != nil {
			response.NewResponse(c, http.StatusUnauthorized, err.Error())
		}

		c.Set(userCtx, id)
	}
}

func parseAuthHeader(c *gin.Context, tokenManager auth.TokenManager) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return tokenManager.ParseAccessToken(headerParts[1])
}
