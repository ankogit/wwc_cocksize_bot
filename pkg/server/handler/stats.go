package handler

import (
	"github.com/gin-gonic/gin"
	"local/wwc_cocksize_bot/pkg/models"
	"net/http"
	"time"
)

func (h *Handler) getStats(c *gin.Context) {
	userStats := models.UserData{
		ID:        0,
		Username:  "1",
		FirstName: "2",
		LastName:  "3",
		CockSize:  10,
		Time:      time.Now(),
	}

	//if err != nil {
	//	newResponse(c, http.StatusInternalServerError, err.Error())
	//
	//	return
	//}

	//courses, err := h.services.Admins.GetCourses(c.Request.Context(), school.ID)
	//if err != nil {
	//	newResponse(c, http.StatusInternalServerError, err.Error())
	//
	//	return
	//}

	//response := make([]domain.Course, len(courses))
	//if courses != nil {
	//	response = courses
	//}

	c.JSON(http.StatusOK, dataResponse{Data: userStats})
}
