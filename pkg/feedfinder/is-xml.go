package feedfinder

import (
	"strings"
)

func IsXml(body []byte) bool {
	docStart := strings.TrimSpace(string(body))[:5]
	return docStart == "<?xml"
}
