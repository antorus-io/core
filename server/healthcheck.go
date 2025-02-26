package server

import (
	"net/http"

	"github.com/antorus-io/core/utils"
)

type Healthcheck struct {
	Status string `json:"status"`
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	healthcheckResponse := Healthcheck{
		Status: "OK",
	}

	if err := utils.WriteJSON(w, http.StatusOK, healthcheckResponse, nil); err != nil {
		HandleHttpError(w, r, err, http.StatusInternalServerError)

		return
	}
}
