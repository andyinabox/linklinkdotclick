package opmlparser

import "encoding/xml"

// https://siongui.github.io/2015/02/26/go-parse-opml-concisely/

type opmlDoc struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    opmlHead
	Body    opmlBody
}

type opmlHead struct {
	XMLName xml.Name `xml:"head"`
	Title   string   `xml:"title"`
}

type opmlBody struct {
	XMLName  xml.Name      `xml:"body"`
	Outlines []opmlOutline `xml:"outline"`
}

type opmlOutline struct {
	XMLName  xml.Name      `xml:"outline"`
	Text     string        `xml:"text,attr"`
	Title    string        `xml:"title,attr"`
	Type     string        `xml:"type,attr"`
	XmlUrl   string        `xml:"xmlUrl,attr"`
	HtmlUrl  string        `xml:"htmlUrl,attr"`
	Outlines []opmlOutline `xml:"outline"`
}

type Feed struct {
	Title   string
	XmlUrl  string
	HtmlUrl string
}

func ParseXml(xmlContent []byte) (feeds []Feed, err error) {
	feeds = []Feed{}
	opmlDoc := opmlDoc{}
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

func MarshallXml(feeds []Feed, title string) (xmlDoc []byte, err error) {
	doc := &opmlDoc{}
	doc.Version = "1.0"
	doc.Head.Title = title
	doc.Body.Outlines = make([]opmlOutline, len(feeds))
	for i, feed := range feeds {
		doc.Body.Outlines[i] = opmlOutline{
			Type:    "rss",
			Title:   feed.Title,
			Text:    feed.Title,
			XmlUrl:  feed.XmlUrl,
			HtmlUrl: feed.HtmlUrl,
		}
	}
	return xml.MarshalIndent(doc, "", "  ")
}
