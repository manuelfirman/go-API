package handler

import (
	"errors"
	"fmt"
)

var (
	// ErrHandlerMissingKey is the error returned when a required key is missing
	ErrHandlerMissingKey = errors.New("missing key")
	// ErrHandlerMissingField is the error returned when a required field is missing
	ErrHandlerMissingField = errors.New("missing field")
	// ErrHandlerIdInRequest is the error returned when the ID is in the request
	ErrHandlerIdInRequest = errors.New("id in request")
)

// Response is a struct that contains the response message and data
type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// validateKeyExistance validates if the key exists in the map
func validateKeyExistance(m map[string]any, keys ...string) error {
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			return fmt.Errorf("%w: %s not found", ErrHandlerMissingKey, k)
		}
	}

	return nil
}
