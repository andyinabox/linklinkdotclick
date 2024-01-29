package feedreader

import "testing"

func TestParseRssUrl(t *testing.T) {
	reader := New()
	feed, err := reader.Parse("https://www.w3.org/blog/feed/")
	if err != nil {
		t.Fatalf("Error parsing feed url: %s", err)
	}
	if feed == nil {
		t.Fatal("Expected value for feed, got nil")
	}
	if len(feed.Items) == 0 {
		t.Fatal("Expected items in feed, got none")
	}
}

func TestParseBlogUrl(t *testing.T) {
	reader := New()
	feed, err := reader.Parse("https://www.w3.org/blog/")
	if err != nil {
		t.Fatalf("Error parsing feed url: %s", err)
	}
	if feed == nil {
		t.Fatal("Expected value for feed, got nil")
	}
	if len(feed.Items) == 0 {
		t.Fatal("Expected items in feed, got none")
	}
}
