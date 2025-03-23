package server

import (
	"fmt"
	"net/http"

	coreConfig "github.com/antorus-io/core/config"
)

func commonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		next.ServeHTTP(w, r)
	})
}

func handleCors(appConfig *coreConfig.ApplicationConfig, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")

		origin := r.Header.Get("Origin")

		if len(appConfig.ServerConfig.TrustedOrigins) != 0 {
			originTrusted := false

			for _, trustedOrigin := range appConfig.ServerConfig.TrustedOrigins {
				if trustedOrigin == "*" {
					w.Header().Set("Access-Control-Allow-Origin", "*")
					originTrusted = true

					break
				}

				if origin == trustedOrigin {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					originTrusted = true

					break
				}
			}

			if !originTrusted && origin != "" {
				http.Error(w, "Forbidden", http.StatusForbidden)

				return
			}
		}

		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.WriteHeader(http.StatusOK)

			return
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
