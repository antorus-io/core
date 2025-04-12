package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antorus-io/core/config"
	"github.com/antorus-io/core/logs"
)

var ServerInstance *Server
var ServerInitialized = false

type Server struct {
	instance *http.Server
	service  string
}

func NewServer(appConfig *config.ApplicationConfig) {
	instance := &http.Server{
		Addr:         fmt.Sprintf(":%s", appConfig.ServerConfig.Port),
		ErrorLog:     log.New(os.Stdout, "", log.LstdFlags),
		Handler:      getRoutes(appConfig),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	ServerInstance = &Server{
		instance: instance,
		service:  appConfig.Service,
	}
	ServerInitialized = true
}

func (server *Server) Serve() error {
	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		logs.Logger.Info("Caught signal", "signal", s.String(), "operationName", "Serve")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		shutdownError <- server.instance.Shutdown(ctx)
	}()

	logs.Logger.Info("Starting server instance", "address", server.instance.Addr, "operationName", "Serve")

	if err := server.instance.ListenAndServe(); err != nil {
		return err
	}

	logs.Logger.Info("Server instance stopped", "address", server.instance.Addr, "operationName", "Serve")

	return nil
}
