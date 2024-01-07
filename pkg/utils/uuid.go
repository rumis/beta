package utils

import (
	"github.com/google/uuid"
)

// UUID generate a uuid string
func UUID() string {
	return uuid.New().String()
}
