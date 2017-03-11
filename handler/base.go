package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/hopehook/fasthttp_api/define"

	"github.com/valyala/fasthttp"
)

func loginAuth(ctx *fasthttp.RequestCtx) (ok bool) {
	// 获取cookie的uid,sid
	var uid = string(ctx.Request.Header.Cookie("uid"))
	var sid = string(ctx.Request.Header.Cookie("sid"))
	// 获取sid对应的session
	session_info, err := Cache.GetHashMapString(sid)
	if err != nil {
		return false
		Logger.Error(err.Error())
	}
	// cookie和session信息比对
	if session_info["uid"] != uid {
		return false
	}
	Logger.Info("auth success")
	return true
}

// login auth handler
func Auth(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		printBeforeLog(ctx)
		hasAuth := loginAuth(ctx)

		if hasAuth {
			// Delegate request to the given handle
			h(ctx)
			return
		}
		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
	})
}

// raw handler
func Raw(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		printBeforeLog(ctx)
		h(ctx)
		return
	})
}

// Write to response body
func WriteString(ctx *fasthttp.RequestCtx, result string) {
	printFinishLog(result)
	ctx.WriteString(result)
}

func WriteBytes(ctx *fasthttp.RequestCtx, result []byte) {
	printFinishLog(string(result))
	ctx.Write(result)
}

func CommonWriteSuccess(ctx *fasthttp.RequestCtx) {
	result := fmt.Sprintf(`{"ret": %d, "msg": %s, "data": {}}`, define.SUCCESS, define.SUCCESS_MSG)
	WriteString(ctx, result)
}

func CommonWriteError(ctx *fasthttp.RequestCtx) {
	result := fmt.Sprintf(`{"ret": %d, "msg": %s, "data": {}}`, define.ERR, define.ERR_MSG)
	WriteString(ctx, result)
}

func CommonWrite(ctx *fasthttp.RequestCtx, ret int64, msg string, data interface{}) {
	resultMap := map[string]interface{}{
		"ret":  ret,
		"msg":  msg,
		"data": data,
	}
	result, _ := json.Marshal(resultMap)
	WriteBytes(ctx, result)
}

func Render(ctx *fasthttp.RequestCtx, path string) {
	t, _ := template.ParseFiles(filepath.Join(TPL_PATH, path))
	t.Execute(ctx, nil)
	ctx.SetContentType("text/html; charset=UTF-8") // 渲染
}

// 打印handler执行之前的日志
func printBeforeLog(ctx *fasthttp.RequestCtx) {
	Logger.Info("================================================")
	Logger.Info("收到请求: method: %s, url: %s, from: %s", ctx.Method(), ctx.Path(), ctx.RemoteAddr())
	Logger.Info("上传参数: query: %v, post: %v", ctx.QueryArgs(), ctx.PostArgs())
}

// 打印handler执行结束的日志
func printFinishLog(result interface{}) {
	Logger.Info("请求返回: %v ", result)
	Logger.Info("================================================")
}
