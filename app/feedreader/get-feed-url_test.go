package feedreader

import (
	"net/http"
	"testing"
)

func Test_getFeedUrlsFromResponse(t *testing.T) {

	res, err := http.Get("https://www.w3.org/blog/")
	if err != nil {
		t.Fatalf("Error fetching url: %e", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Bad status code: %d", res.StatusCode)
	}

	urls, err := getFeedUrlsFromResponse(res)
	if err != nil {
		t.Fatalf("Error checking feed urls: %e", err)
	}
	if len(urls) != 1 {
		t.Fatalf("Expected 1 url, got %d", len(urls))
	}

	if urls[0].String() != "https://www.w3.org/blog/feed/" {
		t.Fatalf("Undexpected first feed URL: %s", urls[0].String())
	}
}

func Test_GetFeedUrl(t *testing.T) {
	reader := New()
	urls, err := reader.GetFeedUrls("https://www.w3.org/blog/")
	if err != nil {
		t.Fatalf("Error checking feed urls: %e", err)
	}
	if len(urls) != 1 {
		t.Fatalf("Expected 1 url, got %d", len(urls))
	}

	if urls[0].String() != "https://www.w3.org/blog/feed/" {
		t.Fatalf("Undexpected first feed URL: %s", urls[0].String())
	}
}

func Test_GetFeedAtom(t *testing.T) {
	reader := New()
	urls, err := reader.GetFeedUrls("https://daringfireball.net/")
	if err != nil {
		t.Fatalf("Error checking feed urls: %e", err)
	}
	if len(urls) != 1 {
		t.Fatalf("Expected 1 url, got %d", len(urls))
	}

	if urls[0].String() != "https://daringfireball.net/feeds/main" {
		t.Fatalf("Undexpected first feed URL: %s", urls[0].String())
	}
}
