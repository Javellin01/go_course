package httpserver

import (
	"context"
	"fmt"
	"net/http"
)

type server struct {
	server http.Server
	notify chan error
}

// New returns the wrapper of the http server.
func New(handler http.Handler, opts ...Option) *server {
	srv := &server{
		server: http.Server{Handler: handler},
		notify: make(chan error),
	}
	for _, opt := range opts {
		opt(srv)
	}
	return srv
}

// Start starts listening on the TCP network port.
func (s *server) Start() {
	s.notify <- s.server.ListenAndServe()
}

// Notify returns channel with server runtime error.
// Will block until a non-nil error is written to the pipe.
func (s *server) Notify() <-chan error {
	return s.notify
}

// Shutdown gracefully shuts down the server without interrupting
// active connections.
func (s *server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %w", err)
	}
	return nil
}
