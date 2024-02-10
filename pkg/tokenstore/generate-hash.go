package tokenstore

import (
	"github.com/google/uuid"
)

func generateHash(n int) (string, error) {
	id := uuid.New()
	return id.String(), nil
}
