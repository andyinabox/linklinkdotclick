package wsrouter

import (
	"fmt"
	"time"

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
	engine.GET("/ws", r.HandleWS)
}

func (r *Router) HandleWS(ctx *gin.Context) {

	conn, err := r.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("WS error: %v", err)
		return
	}
	defer conn.Close()

	conn.WriteMessage(websocket.TextMessage, []byte("First message!"))
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		time.Sleep(time.Second)
	}
}
