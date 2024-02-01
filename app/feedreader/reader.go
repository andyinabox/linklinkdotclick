package feedreader

import "github.com/mmcdole/gofeed"

type Reader struct {
}

type Result struct {
	*gofeed.Feed
}

func New() *Reader {
	return &Reader{}
}
