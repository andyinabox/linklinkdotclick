package feedhelper

import (
	"io"
	"net/http"
	"strings"
)

func (h *Helper) IsFeed(res *http.Response) bool {
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	docStart := strings.TrimSpace(string(b))[:5]
	return docStart == "<?xml"
}
