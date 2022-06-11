package handler

import (
	"github.com/gin-gonic/gin"
	"local/wwc_cocksize_bot/pkg/server/handler/response"
	"net/http"
)

func (h *Handler) getStats(c *gin.Context) {
	//userStats := models.UserData{
	//	ID:        0,
	//	Username:  "1",
	//	FirstName: "2",
	//	LastName:  "3",
	//	CockSize:  10,
	//	Time:      time.Now(),
	//}

	userStats, err := h.services.Repositories.Users.All()

	if err != nil {
		response.NewResponse(c, http.StatusInternalServerError, err.Error())
	
		return
	}

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

	c.JSON(http.StatusOK, response.DataResponse{Data: userStats})
}
