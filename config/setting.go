package config

import (
	"path/filepath"
)

const DEBUG = true

var LOG map[string]string = map[string]string{
	"path":     "hopehook.com.log",
	"level":    "7",
	"maxdays":  "10",
	"separate": `["error"]`,
}

var DEFAULT_SVR map[string]string = map[string]string{
	"ip":   "127.0.0.1",
	"port": "8000",
}

var MYSQL map[string]string = map[string]string{
	"host":         "127.0.0.1:3306",
	"database":     "db",
	"user":         "user",
	"password":     "pwd",
	"maxOpenConns": "10",
	"maxIdleConns": "10",
}

var REDIS map[string]string = map[string]string{
	"host":         "127.0.0.1:6379",
	"database":     "0",
	"password":     "pwd",
	"maxOpenConns": "10",
	"maxIdleConns": "10",
}

// paths
var TPL_PATH, UPLOAD_PATH, DOWNLOAD_PATH = func() (string, string, string) {
	// TODO: type your project path "pwd". 
	// WHY: we can run golang project in many environment, but the project path is different !
	// So, fix project path, you can run this web project via "go run main.go" or "go build" & run or debug in any ide/tools.
	pwd := ""
	TemplatePath := filepath.Join(pwd, "template/")
	UploadPath := filepath.Join(pwd, "file/upload/")
	DownloadPath := filepath.Join(pwd, "file/download/")
	return TemplatePath, UploadPath, DownloadPath
}()
