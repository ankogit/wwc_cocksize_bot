package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"local/wwc_cocksize_bot/pkg/models"
	"local/wwc_cocksize_bot/pkg/server/handler/response"
	"local/wwc_cocksize_bot/pkg/service"
	"net/http"
	"strconv"
)

type refreshInput struct {
	RefreshToken string `json:"token" form:"token" binding:"required"`
}

type loginInput struct {
	UserId string `json:"user-id" form:"user-id" binding:"required"`
}

type tokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (h *Handler) login(c *gin.Context) {
	var input loginInput
	if ok := bindData(c, &input); !ok {
		return
	}

	user_id, err := strconv.ParseInt(input.UserId, 10, 64)
	if err != nil {
		response.NewResponse(c, http.StatusBadRequest, "Uncorrected user-id")
		return
	}

	resp, err := h.services.Users.Login(c, service.LoginInput{
		UserId: user_id,
	})

	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			response.NewResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		if errors.Is(err, models.ErrStudentBlocked) {
			response.NewResponse(c, http.StatusForbidden, err.Error())

			return
		}

		response.NewResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	response.JsonResponse(c, tokenResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, http.StatusOK)

}

func (h *Handler) refresh(c *gin.Context) {
	var input refreshInput
	if ok := bindData(c, &input); !ok {
		return
	}

	resp, err := h.services.Users.RefreshToken(c, service.RefreshInput{
		Token: input.RefreshToken,
	})

	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			response.NewResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		if errors.Is(err, models.ErrStudentBlocked) {
			response.NewResponse(c, http.StatusForbidden, err.Error())

			return
		}

		response.NewResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	response.JsonResponse(c, tokenResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, http.StatusOK)
}
