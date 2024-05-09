package utils

import (
	"errors"
	"strings"
)

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("Bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("Incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}
