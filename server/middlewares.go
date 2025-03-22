package server

import (
	"fmt"
	"net/http"

	coreConfig "github.com/antorus-io/core/config"
)

func commonHeaders(appConfig *coreConfig.ApplicationConfig, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		next.ServeHTTP(w, r)
	})
}

func handleCors(appConfig *coreConfig.ApplicationConfig, next http.Handler) http.Handler {
	trustedOrigins := make(map[string]struct{}, len(appConfig.ServerConfig.TrustedOrigins))

	for _, o := range appConfig.ServerConfig.TrustedOrigins {
		trustedOrigins[o] = struct{}{}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")

		origin := r.Header.Get("Origin")

		if origin != "" {
			if _, ok := trustedOrigins[origin]; ok || len(trustedOrigins) == 0 {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.Header().Set("Access-Control-Allow-Credentials", "true")

				if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
					w.WriteHeader(http.StatusOK)

					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")

				HandleHttpError(w, r, fmt.Errorf("%s", err), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
