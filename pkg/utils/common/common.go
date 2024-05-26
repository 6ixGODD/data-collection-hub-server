package common

import (
	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return newUUID.String(), nil
}
