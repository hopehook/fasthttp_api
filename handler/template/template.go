package template

import (
	"github.com/hopehook/fasthttp_api/handler"

	"github.com/valyala/fasthttp"
)

func TemplateHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case "GET":
		switch string(ctx.UserValue("action").(string)) {
		case "home":
			home(ctx)
		case "login":
			login(ctx)
		case "ws":
			ws(ctx)
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

func home(ctx *fasthttp.RequestCtx) {
	handler.Render(ctx, "home.html")
}

func login(ctx *fasthttp.RequestCtx) {
	handler.Render(ctx, "login.html")
}

func ws(ctx *fasthttp.RequestCtx) {
	handler.Render(ctx, "websocket.html")
}
