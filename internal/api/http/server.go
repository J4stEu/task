package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/J4stEu/task/internal/service"
)

type HTTPServer struct {
	service    *service.Service
	httpServer *http.Server
}

func New(apiSecret []byte, service *service.Service) *HTTPServer {
	mux := InitMux(apiSecret, service)

	requestTimeout := 10 * time.Second

	return &HTTPServer{
		service: service,
		httpServer: &http.Server{
			Addr: "localhost:8085",

			ReadTimeout:  requestTimeout,
			WriteTimeout: requestTimeout,
			IdleTimeout:  requestTimeout * 12,

			Handler: http.TimeoutHandler(mux, requestTimeout, requestTimeoutMsg),
		},
	}
}

func (srv *HTTPServer) Run(
	// TODO: ctx can be used to stop services gracefully
	ctx context.Context,
) error {

	srv.service.Task.WatchAnalytics(ctx)

	fmt.Println("Running HTTP server: ", srv.httpServer.Addr)

	return srv.httpServer.ListenAndServe()
}
