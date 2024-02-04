package magiclinkservice

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateToken generates a token
func (s *Service) generateToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), err
}
