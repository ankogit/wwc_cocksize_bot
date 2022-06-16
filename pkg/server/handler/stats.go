package handler

import (
	"github.com/gin-gonic/gin"
	"local/wwc_cocksize_bot/pkg/server/handler/response"
	"net/http"
)

func (h *Handler) getStats(c *gin.Context) {

	userStats, err := h.services.Repositories.Users.All()

	if err != nil {
		response.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.DataResponse{Data: userStats, Count: len(userStats)})
}
