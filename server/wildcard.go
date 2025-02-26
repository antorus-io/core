package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/antorus-io/core/config"
)

func wildcardHandler(routes map[string]config.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path
		pathExists := false

		// Default routes for healthcheck and metrics
		routesList := []string{"GET /healthcheck", "GET /metrics"}

		// Append dynamic routes as well
		for _, route := range routes {
			routesList = append(routesList, route.Path)
		}

		for _, r := range routesList {
			if strings.Split(r, " ")[1] == requestPath {
				pathExists = true

				break
			}
		}

		fmt.Println(requestPath)

		if pathExists {
			handleHttpError(w, r, MethodNotAllowed, http.StatusMethodNotAllowed)
		} else {
			handleHttpError(w, r, ResourceNotFound, http.StatusNotFound)
		}
	}
}
