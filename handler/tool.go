package handler

import (
	"path/filepath"

	"github.com/hopehook/fasthttp_api/define"
	"github.com/hopehook/fasthttp_api/util"

	"github.com/valyala/fasthttp"
)

func ToolHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case "GET":
		switch string(ctx.UserValue("action").(string)) {
		case "download":
			download(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}

	case "POST":
		switch string(ctx.UserValue("action").(string)) {
		case "upload":
			upload(ctx)
		case "batch_upload":
			batchUpload(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

// 下载"file/download/"目录下的文件
func download(ctx *fasthttp.RequestCtx) {
	filename := string(ctx.QueryArgs().Peek("filename"))
	filepath := filepath.Join(DOWNLOAD_PATH, filename)
	Logger.Info(filepath)
	ctx.SendFile(filepath)
}

// 上传文件到"file/upload/"目录
func upload(ctx *fasthttp.RequestCtx) {
	// 这里直接获取到 multipart.FileHeader, 需要手动打开文件句柄
	fileHeader, err := ctx.FormFile("file") // 默认是"file"
	if err != nil {
		Logger.Error("get upload file error: %s", err.Error())
		ctx.SetStatusCode(500)
		CommonWrite(ctx, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	// 文件名和路径
	filename := fileHeader.Filename
	filepath := filepath.Join(UPLOAD_PATH, filename)

	// 保存
	err = util.SaveFile(fileHeader, filepath)
	if err != nil {
		Logger.Error("upload file error: %s", err.Error())
		ctx.SetStatusCode(500)
		CommonWrite(ctx, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}

	CommonWriteSuccess(ctx)

}

// 批量上传文件到"file/upload/"目录
func batchUpload(ctx *fasthttp.RequestCtx) {
	mf, err := ctx.MultipartForm()
	if err != nil {
		Logger.Error("get upload file error: %s", err.Error())
		ctx.SetStatusCode(500)
		CommonWrite(ctx, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	for _, fileHeader := range mf.File["file"] {
		// 文件名和路径
		filename := fileHeader.Filename
		filepath := filepath.Join(UPLOAD_PATH, filename)

		// 保存
		err = util.SaveFile(fileHeader, filepath)
		if err != nil {
			Logger.Error("upload file error: %s", err.Error())
			ctx.SetStatusCode(500)
			CommonWrite(ctx, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
			return
		}
	}

	CommonWriteSuccess(ctx)
}
