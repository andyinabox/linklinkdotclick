package tokenstore

import (
	"github.com/google/uuid"
)

func generateHash(n int) (string, error) {
	// b := make([]byte, n)
	// _, err := rand.Read(b)
	// if err != nil {
	// 	return "", err
	// }
	id := uuid.New()
	return id.String(), nil
}
