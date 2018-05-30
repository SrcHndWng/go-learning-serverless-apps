package utils

import (
	uuid "github.com/satori/go.uuid"
)

// GenerateID creates UUID.
func GenerateID() string {
	return uuid.NewV4().String()
}
