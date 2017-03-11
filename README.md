# fasthttp_api
fasthttp api demo = fasthttp + fasthttprouter + websocket

* It's fast
  * based on fasthttp and fasthttprouter
* It's simple
  * few main dependences
  ```
  http: github.com/valyala/fasthttp
  router: github.com/buaazp/fasthttprouter
  mysql: github.com/go-sql-driver/mysql
  redis: github.com/garyburd/redigo/redis
  logs: github.com/astaxie/beego/logs
  websocket: github.com/gorilla/websocket ->  github.com/hopehook/websocket // modified for fasthttp
  ```
  * just a demo scale, very clean
* It's enough
  * support http and websocket
  * support mysql connection pool
  * support redis connection pool