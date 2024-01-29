package feedreader

import "testing"

func TestParseRssUrl(t *testing.T) {
	url := "https://www.w3.org/blog/feed/"
	reader := New()
	feed, feedUrl, err := reader.Parse("https://www.w3.org/blog/feed/")
	if err != nil {
		t.Fatalf("Error parsing feed url: %s", err)
	}
	if feedUrl != url {
		t.Fatalf("Expected feedUrl to be %s, got %s", url, feedUrl)
	}
	if feed == nil {
		t.Fatal("Expected value for feed, got nil")
	}
	if len(feed.Items) == 0 {
		t.Fatal("Expected items in feed, got none")
	}
}

func TestParseBlogUrl(t *testing.T) {
	expectedFeedUrl := "https://www.w3.org/blog/feed/"
	reader := New()
	feed, feedUrl, err := reader.Parse("https://www.w3.org/blog/")
	if err != nil {
		t.Fatalf("Error parsing feed url: %s", err)
	}
	if feedUrl != expectedFeedUrl {
		t.Fatalf("Expected feedUrl to be %s, got %s", expectedFeedUrl, feedUrl)
	}
	if feed == nil {
		t.Fatal("Expected value for feed, got nil")
	}
	if len(feed.Items) == 0 {
		t.Fatal("Expected items in feed, got none")
	}
}
