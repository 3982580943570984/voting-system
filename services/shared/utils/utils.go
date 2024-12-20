package utils

import (
	"context"
	"errors"

	"github.com/go-chi/jwtauth/v5"
)

func RetrieveIdFromToken(ctx context.Context) (int, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return -1, err
	}

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return -1, errors.New("invalid or missing 'id' in token claims")
	}

	return int(idFloat), nil
}
