package grpc_handler

import (
	"context"
	"github.com/ankogit/wwc_cocksize_bot--api/api"

	"local/wwc_cocksize_bot/pkg/models"
	"local/wwc_cocksize_bot/pkg/service"
)

//type StatsService interface {
//	Get(ctx context.Context, req *api.SearchReq) (*api.GetStatsResp, error)
//}

type StatsHandler struct {
	api.UnimplementedStatsServer

	services *service.Services
}

func NewStatsHandler(s *service.Services) *StatsHandler {
	return &StatsHandler{services: s}
}

func (h *StatsHandler) Get(ctx context.Context, req *api.SearchReq) (*api.GetStatsResp, error) {
	userStats, err := h.services.Repositories.Users.All()

	if err != nil {
		return nil, err
	}

	return new(models.UserDataResponse).ArrayToProto(userStats), nil
}
