package grpc

import (
	"fmt"
	"github.com/ankogit/wwc_cocksize_bot--api/api"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Deps struct {
	Logger *logrus.Logger

	StatsHandler api.StatsServer
}

type Server struct {
	Deps Deps
	srv  *grpc.Server
}

func NewServer(deps Deps) *Server {
	logger := logrus.NewEntry(deps.Logger)
	return &Server{
		srv: grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_logrus.StreamServerInterceptor(logger),
				grpc_recovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(logger),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		),
		Deps: deps,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	listner, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	api.RegisterStatsServer(s.srv, s.Deps.StatsHandler)

	if err := s.srv.Serve(listner); err != nil {
		return err
	}

	return nil
}
