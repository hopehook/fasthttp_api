package handler

import (
	"strconv"

	"github.com/hopehook/fasthttp_api/define"

	"github.com/valyala/fasthttp"
)

// login
func LoginHandler(ctx *fasthttp.RequestCtx) {
	// upload param
	username := string(ctx.PostArgs().Peek("username"))
	password := string(ctx.PostArgs().Peek("password"))
	Logger.Info(username)
	Logger.Info(password)
	// auth
	sql := `SELECT id AS uid, password, phone, name FROM users WHERE phone = ? AND is_deleted = 0 AND app_type = 1`
	rows, _ := DB.Query(sql, username)
	if len(rows) != 1 {
		CommonWrite(ctx, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	userInfo := rows[0]
	if password != userInfo["password"] {
		CommonWrite(ctx, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	// session
	uid := userInfo["uid"].(int64)
	uidStr := strconv.Itoa(int(uid))
	sid := uidStr
	_, err := Cache.SetHashMap(sid, userInfo)
	if err != nil {
		CommonWrite(ctx, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	// cookie
	cookie := fasthttp.AcquireCookie()   // get cookie obj from pool
	defer fasthttp.ReleaseCookie(cookie) //release cookie to pool
	cookie.SetKey("uid")
	cookie.SetValue(uidStr)
	cookie.SetPath("/")
	ctx.Response.Header.SetCookie(cookie)
	cookie.SetKey("sid")
	cookie.SetValue(sid)
	cookie.SetPath("/")
	ctx.Response.Header.SetCookie(cookie)

	// login success
	CommonWriteSuccess(ctx)
}

// logout
func LogoutHandler(ctx *fasthttp.RequestCtx) {
	// 获取cookie的uid,sid
	var sid = string(ctx.Request.Header.Cookie("sid"))
	// session
	_, err := Cache.DelKey(sid)
	if err != nil {
		CommonWrite(ctx, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	// cookie
	ctx.Response.Header.DelClientCookie("uid")
	ctx.Response.Header.DelClientCookie("sid")

	// logout success
	CommonWriteSuccess(ctx)
}
