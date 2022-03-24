package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunHttp(port string, handler http.Handler) error {
	logger := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    1 << 20,
		//ConnState:         fConnState,
		ErrorLog: logger,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}

	return s.httpServer.ListenAndServe()
}

func fConnState(n net.Conn, state http.ConnState) {
	fmt.Println("ConnState")
	fmt.Println(n, state)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
