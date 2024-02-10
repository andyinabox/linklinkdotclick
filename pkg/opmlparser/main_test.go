package opmlparser

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func Test_Parse(t *testing.T) {

	// non-nested opml
	{
		file, err := ioutil.ReadFile("../../test/fixtures/opml/feeds-simple.opml")
		if err != nil {
			t.Fatal(err.Error())
		}
		feeds, err := ParseXml(file)
		if err != nil {
			t.Fatal(err.Error())
		}
		expectedFeedCount := 3
		if len(feeds) != expectedFeedCount {
			t.Errorf("expected %d feeds, got %d", expectedFeedCount, len(feeds))
		}
	}

	// nested opml
	{
		file, err := ioutil.ReadFile("../../test/fixtures/opml/feeds-nested.opml")
		if err != nil {
			t.Fatal(err.Error())
		}
		feeds, err := ParseXml(file)
		if err != nil {
			t.Fatal(err.Error())
		}
		expectedFeedCount := 3
		if len(feeds) != expectedFeedCount {
			t.Errorf("expected %d feeds, got %d", expectedFeedCount, len(feeds))
		}
	}

}

func Test_Marshal(t *testing.T) {
	feeds := []Feed{
		{
			Title:   "CreativeApplications.Net",
			XmlUrl:  "https://www.creativeapplications.net/feed/",
			HtmlUrl: "https://www.creativeapplications.net",
		},
		{
			Title:   "Open Culture",
			XmlUrl:  "https://www.openculture.com/feed",
			HtmlUrl: "https://www.openculture.com/",
		},
		{
			Title:   "Chairblog.eu",
			XmlUrl:  "https://chairblog.eu/feed/",
			HtmlUrl: "https://chairblog.eu/",
		},
	}

	title := "My feeds"
	b, err := MarshallXml(feeds, "My feeds")
	if err != nil {
		t.Fatal(err.Error())
	}

	type testOutline struct{}
	type testDoc struct {
		Title    string        `xml:"head>title"`
		Outlines []testOutline `xml:"body>outline"`
	}

	doc := testDoc{}
	err = xml.Unmarshal(b, &doc)
	if err != nil {
		t.Fatal(err.Error())
	}
	if doc.Title != title {
		t.Errorf("expected title %s, got %s\n%s", title, doc.Title, string(b))
	}
	expectedFeedCount := 3
	if len(doc.Outlines) != expectedFeedCount {
		t.Errorf("expected %d feeds, got %d\n%s", expectedFeedCount, len(doc.Outlines), string(b))
	}
}
