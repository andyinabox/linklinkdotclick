// https://www.smashingmagazine.com/2017/04/guide-http2-server-push/
package pushit

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentType string
type LinkType string

type Resource struct {
	Path        string
	ContentType ContentType
}

const (
	ContentTypeStyle  ContentType = "style"
	ContentTypeScript ContentType = "script"
	ContentTypeJson   ContentType = "json"
	ContentTypeImage  ContentType = "image"
	ContentTypeVideo  ContentType = "video"
)
const (
	LinkTypePreload LinkType = "preload"
)

func Middleware(resources []Resource) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if pusher := ctx.Writer.Pusher(); pusher != nil {
			for _, r := range resources {

				header := http.Header{}
				header.Add("rel", string(LinkTypePreload))
				header.Add("as", string(r.ContentType))

				err := pusher.Push(r.Path, &http.PushOptions{
					Header: header,
				})

				if err != nil {
					log.Printf("Failed to push: %v", err)
				}
			}
		}

		ctx.Next()
	}
}
