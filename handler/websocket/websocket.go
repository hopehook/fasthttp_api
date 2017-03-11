package websocket

import (
	"github.com/hopehook/websocket"
	"github.com/valyala/fasthttp"
)

func WebsocketHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case "GET":
		switch string(ctx.UserValue("action").(string)) {
		case "upgrade":
			upgrade(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	case "POST":
		switch string(ctx.UserValue("action").(string)) {

		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func upgrade(ctx *fasthttp.RequestCtx) {
	var upgrader = websocket.UpgraderFs{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	if err := upgrader.Upgrade(ctx, hijackHandler, fasthttp.ResponseHeader{}); err != nil {
		Logger.Error(err.Error())
		return
	}
	return
}

// hijackHandler is called on hijacked connection.
var hijackHandler = func(c *websocket.Conn) error {
	for {
		messageType, p, err := c.ReadMessage()
		if err != nil {
			return err
		}
		if err = c.WriteMessage(messageType, p); err != nil {
			return err
		}
	}

	return nil

}
