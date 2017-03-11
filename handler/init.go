package handler

import (
	"github.com/hopehook/fasthttp_api/config"
	"github.com/hopehook/fasthttp_api/lib"
)

// lib
var Logger = lib.Logger
var DB = lib.DB
var Cache = lib.Cache

// config
var TPL_PATH = config.TPL_PATH
var UPLOAD_PATH = config.UPLOAD_PATH
var DOWNLOAD_PATH = config.DOWNLOAD_PATH
