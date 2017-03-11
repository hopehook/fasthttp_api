package router

import (
	h "github.com/hopehook/fasthttp_api/handler"
	"github.com/hopehook/fasthttp_api/handler/template"
	"github.com/hopehook/fasthttp_api/handler/websocket"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var Router *fasthttprouter.Router = fasthttprouter.New()

func init() {
	// static目录
	Router.HandleMethodNotAllowed = false
	Router.NotFound = fasthttp.FSHandler("./static", 0)
	// login
	Router.POST("/login/", h.Raw(h.LoginHandler))
	Router.GET("/logout/", h.Raw(h.LogoutHandler))
	// file目录
	Router.GET("/tool/:action/", h.Raw(h.ToolHandler))
	Router.POST("/tool/:action/", h.Auth(h.ToolHandler))
	// template目录
	Router.GET("/t/:action/", h.Raw(template.TemplateHandler))
	Router.POST("/t/:action/", h.Raw(template.TemplateHandler))
	// websocket
	Router.GET("/websocket/:action/", h.Raw(websocket.WebsocketHandler)) // 建立websocket

}
