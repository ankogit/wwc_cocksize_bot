package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
	}

	api := router.Group("/api")
	{
		api.GET("/stats", h.getStats)
		//api.GET("/stats/:id")
	}

	return router
}
