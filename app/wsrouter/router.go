package wsrouter

import (
	"encoding/json"
	"fmt"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/cssparser"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type styleRequest struct {
	Styles string `json:"styles"`
}

type styleResponse struct {
	Styles   string   `json:"styles"`
	Valid    bool     `json:"valid"`
	Warnings []string `json:"warnings"`
	Errors   []error  `json:"errors"`
}

type Router struct {
	upgrader *websocket.Upgrader
	sc       app.ServiceContainer
}

func New(sc app.ServiceContainer) *Router {
	return &Router{&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}, sc}
}

func (r *Router) Register(engine *gin.Engine) {
	engine.GET("/ws/style-editor", r.HandleWS)
}

func (r *Router) HandleWS(ctx *gin.Context) {
	logger := r.sc.LogService()
	conn, err := r.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("WS error: %v", err)
		return
	}
	defer conn.Close()

	var result *cssparser.ParseResult

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("WS error: %v", err)
			return
		}

		// if it's a text message
		if mt == websocket.TextMessage {

			var request styleRequest
			err = json.Unmarshal(message, &request)

			if err != nil {
				logger.Error().Println(err.Error())
				return
			}

			result, err = cssparser.Parse([]byte(request.Styles), false)
			if err != nil {
				logger.Error().Println(err.Error())
				return
			}

			response := &styleResponse{
				Styles:   string(result.Output),
				Valid:    result.Valid,
				Warnings: result.Warnings,
				Errors:   result.Errors,
			}

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				logger.Error().Println(err.Error())
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, jsonResponse)
			if err != nil {
				logger.Error().Println(err.Error())
				return
			}

		}
	}
}
