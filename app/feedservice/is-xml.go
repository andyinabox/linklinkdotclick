package feedservice

import (
	"strings"
)

func (s *Service) IsXml(body []byte) bool {
	docStart := strings.TrimSpace(string(body))[:5]
	return docStart == "<?xml"
}
