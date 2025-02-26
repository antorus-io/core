package server

import (
	"github.com/google/uuid"
)

func ValidatePathUUID(id string) (string, error) {
	uuid, err := uuid.Parse(id)

	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}
