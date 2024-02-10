package opmlparser

import (
	"bytes"
	"encoding/xml"
)

// https://siongui.github.io/2015/02/26/go-parse-opml-concisely/

type OpmlDoc struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    OpmlHead
	Body    OpmlBody
}

type OpmlHead struct {
	XMLName xml.Name `xml:"head"`
	Title   string   `xml:"title"`
}

type OpmlBody struct {
	XMLName  xml.Name      `xml:"body"`
	Outlines []OpmlOutline `xml:"outline"`
}

type OpmlOutline struct {
	XMLName  xml.Name      `xml:"outline"`
	Text     string        `xml:"text,attr"`
	Title    string        `xml:"title,attr"`
	Type     string        `xml:"type,attr"`
	XmlUrl   string        `xml:"xmlUrl,attr"`
	HtmlUrl  string        `xml:"htmlUrl,attr"`
	Outlines []OpmlOutline `xml:"outline"`
}

type Feed struct {
	Title   string
	XmlUrl  string
	HtmlUrl string
}

func ParseXml(xmlContent []byte) (feeds []Feed, err error) {
	feeds = []Feed{}
	opmlDoc := OpmlDoc{}
	err = xml.Unmarshal(xmlContent, &opmlDoc)
	if err != nil {
		return
	}

	// top-level outlines
	for _, outline := range opmlDoc.Body.Outlines {
		if outline.Type == "rss" {
			feeds = append(feeds, Feed{
				Title:   outline.Title,
				XmlUrl:  outline.XmlUrl,
				HtmlUrl: outline.HtmlUrl,
			})
		}
		// sub-outlines
		for _, outline := range outline.Outlines {
			if outline.Type == "rss" {
				feeds = append(feeds, Feed{
					Title:   outline.Title,
					XmlUrl:  outline.XmlUrl,
					HtmlUrl: outline.HtmlUrl,
				})
			}
		}
	}

	return
}

func FeedsToOpml(feeds []Feed, title string) (doc OpmlDoc) {
	doc.Version = "1.0"
	doc.Head.Title = title
	doc.Body.Outlines = make([]OpmlOutline, len(feeds))
	for i, feed := range feeds {
		doc.Body.Outlines[i] = OpmlOutline{
			Type:    "rss",
			Title:   feed.Title,
			Text:    feed.Title,
			XmlUrl:  feed.XmlUrl,
			HtmlUrl: feed.HtmlUrl,
		}
	}
	return
}

func MarshallXml(feeds []Feed, title string) ([]byte, error) {
	doc := FeedsToOpml(feeds, title)

	b := &bytes.Buffer{}

	// write xml header
	b.Write([]byte(xml.Header))

	// create new encoder and encode doc
	enc := xml.NewEncoder(b)
	enc.Indent("", "  ")
	err := enc.Encode(doc)
	if err != nil {
		return []byte{}, err
	}

	return b.Bytes(), nil
}
