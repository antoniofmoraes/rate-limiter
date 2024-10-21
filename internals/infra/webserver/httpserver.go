package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antoniofmoraes/rate-limiter/internals/infra/webserver/middlewares"
	"github.com/antoniofmoraes/rate-limiter/internals/services"
)

type httpServer struct {
	rateLimiterService *services.RateLimiterService
}

func NewHttpServer(rateLimiterService *services.RateLimiterService) *httpServer {
	return &httpServer{
		rateLimiterService,
	}
}

func (s *httpServer) Start() {
	mux := http.NewServeMux()

	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	mux.Handle("/api", middlewares.RateLimiterMiddleware(s.rateLimiterService, helloHandler))

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
