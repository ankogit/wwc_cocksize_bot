package grpc_handler

import (
	"local/wwc_cocksize_bot/pkg/service"
)

type GrpcHandler struct {
	StatsHandler
	services *service.Services
}

func NewGrpcHandler(services *service.Services) *GrpcHandler {

	return &GrpcHandler{
		services: services,
	}
}

func (h *GrpcHandler) InitHandlers() {
	h.StatsHandler = *NewStatsHandler(h.services)
}
