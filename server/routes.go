package server

import (
	"expvar"
	"net/http"

	"github.com/antorus-io/core/config"
)

func getRoutes(appConfig *config.ApplicationConfig) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthcheck", healthcheckHandler)
	mux.Handle("GET /metrics", expvar.Handler())

	// Register dynamic routes
	for _, route := range appConfig.ServerConfig.Routes {
		mux.HandleFunc(route.Path, route.Handler)
	}

	mux.HandleFunc("/", wildcardHandler(appConfig.ServerConfig.Routes))

	return recoverPanic(handleCors(appConfig, logRequest(commonHeaders(appConfig, mux))))
}
