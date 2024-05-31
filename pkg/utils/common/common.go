package common

import (
	"regexp"

	"github.com/google/uuid"
)

func GenerateUUID4() (string, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return newUUID.String(), nil
}

func EscapeSpecialChars(input string) string {
	re := regexp.MustCompile(`[\\^$*+?.()|[\]{}]`)
	return re.ReplaceAllStringFunc(
		input, func(s string) string {
			return "\\" + s
		},
	)
}
