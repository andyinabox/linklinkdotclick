package wsrouter

import (
	"fmt"

	"github.com/andyinabox/linkydink/pkg/cssparser"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Router struct {
	upgrader *websocket.Upgrader
}

func New() *Router {
	return &Router{&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}}
}

func (r *Router) Register(engine *gin.Engine) {
	engine.GET("/ws/style-editor", r.HandleWS)
}

func (r *Router) HandleWS(ctx *gin.Context) {

	conn, err := r.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("WS error: %v", err)
		return
	}
	defer conn.Close()

	parseOptions := &cssparser.ParseOptions{}
	var result []byte
	var valid bool

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("WS error: %v", err)
			return
		}
		if mt == websocket.TextMessage {
			result, valid, err = cssparser.Parse(message, parseOptions)
			if err != nil {
				fmt.Printf("WS error: %v", err)
				return
			}
			if valid {
				err = conn.WriteMessage(websocket.TextMessage, result)
				if err != nil {
					fmt.Printf("WS error: %v", err)
					return
				}
			}
		}
	}
}
