package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"local/wwc_cocksize_bot/pkg/server/handler/middleware"
	"local/wwc_cocksize_bot/pkg/service"
	"time"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/refresh", h.refresh)
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
