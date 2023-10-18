package server

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer      *http.Server
	shutdownTimeout time.Duration
}

type Opt func(*Server)

func NewServer(addr string, opts ...Opt) *Server {
	s := &Server{
		shutdownTimeout: 5 * time.Second, // default value, can be modified
		httpServer: &http.Server{
			Addr: addr,
		},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithShutdownTimeout(timeout time.Duration) Opt {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

func WithRoutes(r chi.Router) Opt {
	return func(s *Server) {
		s.httpServer.Handler = r
	}
}

func (s *Server) Start() error {
	// Create a channel to listen for OS signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine so it doesn't block
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server Started on %s\n", s.httpServer.Addr)

	<-done // Wait for an OS signal to be caught

	log.Print("Server Stopping...")

	// Set the context for Shutdown with the timeout from Server struct
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	// Call Shutdown on http.Server with the context
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	log.Print("Server Exited Properly")

	return nil
}
