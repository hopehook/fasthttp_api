// http: github.com/valyala/fasthttp
// router: github.com/buaazp/fasthttprouter
// mysql: github.com/go-sql-driver/mysql
// redis: github.com/garyburd/redigo/redis
// logs: github.com/astaxie/beego/logs
// websocket: github.com/gorilla/websocket ->  github.com/hopehook/websocket // modified for fasthttp
package main

import (
	"fmt"

	"github.com/hopehook/fasthttp_api/config"
	"github.com/hopehook/fasthttp_api/lib"
	"github.com/hopehook/fasthttp_api/router"

	"github.com/valyala/fasthttp"
)

func main() {
	host := fmt.Sprintf("%s:%s", config.DEFAULT_SVR["ip"], config.DEFAULT_SVR["port"])
	lib.Logger.Info("System is running on %s", host)
	if err := fasthttp.ListenAndServe(host, router.Router.Handler); err != nil {
		lib.Logger.Error("Start fasthttp failed:", err.Error())
	}
}
