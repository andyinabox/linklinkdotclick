package feedservice

import (
	"io"
	"net/http"
	"strings"
)

func (s *Service) IsXml(res *http.Response) bool {
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return false
	}
	docStart := strings.TrimSpace(string(b))[:5]
	return docStart == "<?xml"
}
