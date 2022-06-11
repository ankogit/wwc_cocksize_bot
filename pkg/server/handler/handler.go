package handler

import (
	"github.com/gin-gonic/gin"
	"local/wwc_cocksize_bot/pkg/server/handler/middleware"
	"local/wwc_cocksize_bot/pkg/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/refresh", h.login)
	}

	api := router.Group("/api")
	{
		authenticated := api.Group("/", middleware.AuthUser(h.services.TokenManager))
		{
			authenticated.GET("/stats", h.getStats)
		}
		//api.GET("/stats", h.getStats)
		//api.GET("/stats/:id")
	}

	return router
}
